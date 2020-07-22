package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/andig/evcc/provider"
	"github.com/andig/evcc/server"
	"github.com/andig/evcc/server/updater"
	"github.com/andig/evcc/util"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	log     = util.NewLogger("main")
	cfgFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "evcc",
	Short:   "EV Charge Controller",
	Version: fmt.Sprintf("%s (%s)", server.Version, server.Commit),
	Run:     run,
}

func bind(cmd *cobra.Command, flag string) {
	if err := viper.BindPFlag(flag, cmd.PersistentFlags().Lookup(flag)); err != nil {
		panic(err)
	}
}

func configureCommand(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP(
		"log", "l",
		"error",
		"Log level (fatal, error, warn, info, debug, trace)",
	)
	bind(cmd, "log")

	cmd.PersistentFlags().StringVarP(&cfgFile,
		"config", "c",
		"",
		"Config file (default \"~/evcc.yaml\" or \"/etc/evcc.yaml\")",
	)
	cmd.PersistentFlags().BoolP(
		"help", "h",
		false,
		"Help for "+cmd.Name(),
	)
}

func init() {
	cobra.OnInitialize(initConfig)
	configureCommand(rootCmd)

	rootCmd.PersistentFlags().StringP(
		"uri", "u",
		"0.0.0.0:7070",
		"Listen address",
	)
	bind(rootCmd, "uri")

	rootCmd.PersistentFlags().DurationP(
		"interval", "i",
		10*time.Second,
		"Update interval",
	)
	bind(rootCmd, "interval")
}

// initConfig reads in config file and ENV variables if set
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {
		// Search for config in home directory if available
		if home, err := os.UserHomeDir(); err == nil {
			viper.AddConfigPath(home)
		}

		// Search config in home directory with name "mbmd" (without extension).
		viper.AddConfigPath(".")    // optionally look for config in the working directory
		viper.AddConfigPath("/etc") // path to look for the config file in

		viper.SetConfigName("evcc")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in
	if err := viper.ReadInConfig(); err == nil {
		// using config file
		cfgFile = viper.ConfigFileUsed()
	} else if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
		// parsing failed - exit
		fmt.Println(err)
		os.Exit(1)
	} else {
		// not using config file
		cfgFile = ""
	}
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) {
	util.LogLevel(viper.GetString("log"), viper.GetStringMapString("levels"))
	log.INFO.Printf("evcc %s (%s)", server.Version, server.Commit)

	// load config and re-configure logging after reading config file
	conf := loadConfigFile(cfgFile)
	util.LogLevel(viper.GetString("log"), viper.GetStringMapString("levels"))

	uri := viper.GetString("uri")
	log.INFO.Println("listening at", uri)

	// setup mqtt
	if viper.Get("mqtt") != nil {
		provider.MQTT = provider.NewMqttClient(conf.Mqtt.Broker, conf.Mqtt.User, conf.Mqtt.Password, clientID(), 1)
	}

	// setup loadpoints
	site := loadConfig(conf)

	// start broadcasting values
	tee := &Tee{}

	// value cache
	cache := util.NewCache()
	go cache.Run(tee.Attach())

	// setup database
	if conf.Influx.URL != "" {
		configureDatabase(conf.Influx, site.LoadPoints(), tee.Attach())
	}

	// create webserver
	socketHub := server.NewSocketHub()
	httpd := server.NewHTTPd(uri, conf.Menu, site, socketHub, cache)

	// publish to UI
	go socketHub.Run(tee.Attach(), cache)

	// setup values channel
	valueChan := make(chan util.Param)
	go tee.Run(valueChan)

	// version check
	go updater.Run(log, valueChan)

	// capture log messages for UI
	util.CaptureLogs(valueChan)

	// setup messaging
	pushChan := configureMessengers(conf.Messaging, cache)

	// set channels
	site.Prepare(valueChan, pushChan)

	site.DumpConfig()
	go site.Run(conf.Interval)

	log.FATAL.Println(httpd.ListenAndServe())
}
