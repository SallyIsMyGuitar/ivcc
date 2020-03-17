package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/andig/evcc/api"
	"github.com/andig/evcc/core"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

//go:generate esc -o assets.go -pkg server -modtime 1566640112 -ignore .DS_Store -prefix ../assets ../assets

const (
	liveAssets = false
)

type MenuConfig struct {
	Title    string
	Subtitle string
	Img      string
	Iframe   string
	Link     string
}

type chargeModeJSON struct {
	Mode string `json:"mode"`
}

type route struct {
	Methods     []string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// loadPoint is the well-definied minimal interface for accessing loadpoint methods
type loadPoint interface {
	GetMode() api.ChargeMode
	SetMode(api.ChargeMode)
	Configuration() core.Configuration
}

// routeLogger traces matched routes including their executing time
func routeLogger(inner http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)
		log.TRACE.Printf(
			"%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			time.Since(start),
		)
	}
}

func indexHandler(links []MenuConfig, liveAssets bool) http.HandlerFunc {
	indexTemplate, err := FSString(liveAssets, "/index.html")
	if err != nil {
		log.FATAL.Fatal("httpd: failed to load embedded template: " + err.Error())
	}

	t, err := template.New("evcc").Delims("<<", ">>").Parse(indexTemplate)
	if err != nil {
		log.FATAL.Fatal("httpd: failed to create main page template: ", err.Error())
	}

	_, debug := _escData["/js/debug.js"]

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=UTF-8")
		w.WriteHeader(http.StatusOK)

		// _, err := fmt.Fprint(w, indexTemplate)
		if err := t.Execute(w, map[string]interface{}{
			"Debug": debug,
			"Links": links,
		}); err != nil {
			log.ERROR.Println("httpd: failed to render main page: ", err.Error())
		}
	})
}

// jsonHandler is a middleware that decorates responses with JSON and CORS headers
func jsonHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		h.ServeHTTP(w, r)
	})
}

// HealthHandler returns current charge mode
func HealthHandler(lp loadPoint) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := struct{ OK bool }{OK: true}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(res); err != nil {
			log.ERROR.Printf("httpd: failed to encode JSON: %v", err)
		}
	}
}

// ConfigHandler returns current charge mode
func ConfigHandler(lp loadPoint) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := lp.Configuration()

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(res); err != nil {
			log.ERROR.Printf("httpd: failed to encode JSON: %v", err)
		}
	}
}

// CurrentChargeModeHandler returns current charge mode
func CurrentChargeModeHandler(lp loadPoint) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res := chargeModeJSON{
			Mode: string(lp.GetMode()),
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(res); err != nil {
			log.ERROR.Printf("httpd: failed to encode JSON: %v", err)
		}
	}
}

// ChargeModeHandler updates charge mode
func ChargeModeHandler(lp loadPoint) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		mode, ok := vars["mode"]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		lp.SetMode(api.ChargeMode(mode))

		res := chargeModeJSON{
			Mode: string(lp.GetMode()),
		}

		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(res); err != nil {
			log.ERROR.Printf("httpd: failed to encode JSON: %v", err)
		}
	}
}

// SocketHandler attaches websocket handler to uri
func SocketHandler(hub *SocketHub) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ServeWebsocket(hub, w, r)
	}
}

// NewHttpd creates HTTP server with configured routes for loadpoint
func NewHttpd(url string, links []MenuConfig, lp loadPoint, hub *SocketHub) *http.Server {
	var routes = []route{
		route{
			[]string{"GET"},
			"/health",
			HealthHandler(lp),
		},
		route{
			[]string{"GET"},
			"/config",
			ConfigHandler(lp),
		},
		route{
			[]string{"GET"},
			"/mode",
			CurrentChargeModeHandler(lp),
		},
		route{
			[]string{"PUT", "POST", "OPTIONS"},
			"/mode/{mode:[a-z]+}",
			ChargeModeHandler(lp),
		},
	}

	router := mux.NewRouter().StrictSlash(true)

	// static
	router.HandleFunc("/", indexHandler(links, liveAssets))

	// individual handlers per folder
	for _, folder := range []string{"js", "css", "webfonts", "ico"} {
		prefix := fmt.Sprintf("/%s/", folder)
		router.PathPrefix(prefix).Handler(http.StripPrefix(prefix, http.FileServer(Dir(liveAssets, prefix))))
	}

	// api
	api := router.PathPrefix("/api").Subrouter()
	api.Use(jsonHandler)
	for _, r := range routes {
		api.
			Methods(r.Methods...).
			Path(r.Pattern).
			Handler(routeLogger(r.HandlerFunc))
	}

	// websocket
	router.HandleFunc("/ws", SocketHandler(hub))

	// add handlers
	handler := handlers.CompressHandler(router)
	handler = handlers.CORS(
		handlers.AllowedHeaders([]string{
			"Accept", "Accept-Language", "Content-Language", "Content-Type", "Origin",
		}),
	)(handler)

	srv := &http.Server{
		Addr:         url,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		ErrorLog:     log.ERROR,
	}
	srv.SetKeepAlivesEnabled(true)

	return srv
}
