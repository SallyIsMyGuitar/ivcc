package provider

import (
	"context"
	"errors"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/andig/evcc/api"
	"github.com/kballard/go-shellquote"
)

var log = api.NewLogger("exec")

func truish(s string) bool {
	return s == "1" || strings.ToLower(s) == "true" || strings.ToLower(s) == "on"
}

// Script implements shell script-based providers and setters
type Script struct {
	timeout time.Duration
}

// NewScriptProvider creates a script provider.
// Script execution is aborted after given timeout.
func NewScriptProvider(timeout time.Duration) *Script {
	return &Script{
		timeout: timeout,
	}
}

// StringGetter returns string from exec result. Only STDOUT is considered.
func (e *Script) StringGetter(script string) StringGetter {
	args, err := shellquote.Split(script)
	if err != nil {
		panic(err)
	} else if len(args) < 1 {
		panic("exec: missing script")
	}

	// return func to access cached value
	return func() (string, error) {
		ctx, cancel := context.WithTimeout(context.Background(), e.timeout)
		defer cancel()

		cmd := exec.CommandContext(ctx, args[0], args[1:]...)
		b, err := cmd.Output()

		s := strings.TrimSpace(string(b))

		if err != nil {
			// use STDOUT if available
			var ee *exec.ExitError
			if errors.As(err, &ee) {
				s = strings.TrimSpace(string(ee.Stderr))
			}

			log.ERROR.Printf("%s: %s", strings.Join(args, " "), s)
			return "", err
		}

		log.TRACE.Printf("%s: %s", strings.Join(args, " "), s)
		return s, nil
	}
}

// IntGetter parses int64 from exec result
func (e *Script) IntGetter(script string) IntGetter {
	exec := e.StringGetter(script)

	// return func to access cached value
	return func() (int64, error) {
		s, err := exec()
		if err != nil {
			return 0, err
		}

		return strconv.ParseInt(s, 10, 64)
	}
}

// FloatGetter parses float from exec result
func (e *Script) FloatGetter(script string) FloatGetter {
	exec := e.StringGetter(script)

	// return func to access cached value
	return func() (float64, error) {
		s, err := exec()
		if err != nil {
			return 0, err
		}

		return strconv.ParseFloat(s, 64)
	}
}

// BoolGetter parses bool from exec result. "on", "true" and 1 are considerd truish.
func (e *Script) BoolGetter(script string) BoolGetter {
	exec := e.StringGetter(script)

	// return func to access cached value
	return func() (bool, error) {
		s, err := exec()
		if err != nil {
			return false, err
		}

		return truish(s), nil
	}
}

// IntSetter invokes script with parameter replaced by int value
func (e *Script) IntSetter(param, script string) IntSetter {
	// return func to access cached value
	return func(i int64) error {
		cmd, err := replaceFormatted(script, map[string]interface{}{
			param: i,
		})
		if err != nil {
			return err
		}

		exec := e.StringGetter(cmd)
		if _, err := exec(); err != nil {
			return err
		}

		return nil
	}
}

// BoolSetter invokes script with parameter replaced by bool value
func (e *Script) BoolSetter(param, script string) BoolSetter {
	// return func to access cached value
	return func(b bool) error {
		cmd, err := replaceFormatted(script, map[string]interface{}{
			param: b,
		})
		if err != nil {
			return err
		}

		exec := e.StringGetter(cmd)
		if _, err := exec(); err != nil {
			return err
		}

		return nil
	}
}
