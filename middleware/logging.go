package middleware

import (
	"github.com/rs/zerolog"
	"net/http"
	"os"
)

// Logs - customized logger by Zerolog
var Logs zerolog.Logger

func init() {
	Logs = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "01-02-2006 15:04:05"}).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Logger()
}

// LockDebug function control the debug level logs
// 'true' - debug don't write
// 'false' - debug write
func LockDebug(debug bool) {
	if debug {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		Logs.Info().Msgf("Debug level is OFF")
	} else {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		Logs.Info().Msgf("Debug level is ON")
	}
}

// Logging for write logs from http client
func Logging(wrapHandle http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		wrapHandle.ServeHTTP(w, req)
		Logs.Info().Msgf("%s: %s %s", req.Host, req.Method, req.URL)
		//req.UserAgent()
	})
}
