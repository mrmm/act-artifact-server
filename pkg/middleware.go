package pkg

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func BearerAuth(handler httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, reqParams httprouter.Params) {
		const bearerAuthPrefix string = "Bearer"

		// Get the Basic Authentication credentials
		auth := request.Header.Get("Authorization")
		log.Trace(fmt.Sprintf("Authorization: %s", auth))
		if strings.HasPrefix(auth, bearerAuthPrefix) {
			// Check credentials
			if auth == fmt.Sprintf("%s %s", bearerAuthPrefix, CLI.Server.Token) {

				// Delegate request to the given handle
				handler(writer, request, reqParams)
				return
			}
		}
		http.Error(writer, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	}
}

func WithLogging(handler httprouter.Handle) httprouter.Handle {
	return func(writer http.ResponseWriter, request *http.Request, reqParams httprouter.Params) {
		start := time.Now()

		uri := request.RequestURI
		method := request.Method
		handler(writer, request, reqParams) // serve the original request

		duration := time.Since(start)

		// log request details
		log.WithFields(log.Fields{
			"uri":      uri,
			"method":   method,
			"duration": duration,
		}).Info("Request")
	}
}
