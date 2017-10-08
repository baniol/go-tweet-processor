// All material is licensed under the Apache License Version 2.0, January 2004
// http://www.apache.org/licenses/LICENSE-2.0

package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/baniol/go-tweet-processor/web"
	log "github.com/sirupsen/logrus"
)

// RequestLogger writes some information about the request to the logs in
// the format: TraceID : (200) GET /foo -> IP ADDR (latency)
func RequestLogger(next web.Handler) web.Handler {

	// Wrap this handler around the next one provided.
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, params map[string]string) error {
		v := ctx.Value(web.KeyValues).(*web.Values)

		next(ctx, w, r, params)
		// TODO: rework the access log fomar - to JSON
		// TODO: now its info level
		log.Printf("%s : (%d) : %s %s -> %s (%s)",
			v.TraceID,
			v.StatusCode,
			r.Method, r.URL.Path,
			r.RemoteAddr, time.Since(v.Now),
		)

		// This is the top of the food chain. At this point all error
		// handling has been done including logging.
		return nil
	}

}
