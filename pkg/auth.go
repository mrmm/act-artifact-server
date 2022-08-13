package pkg

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
)

func BearerAuth(token string, h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		const bearerAuthPrefix string = "Bearer"

		// Get the Basic Authentication credentials
		auth := r.Header.Get("Authorization")
		log.Info(fmt.Sprintf("Authorization: %s", auth))
		if strings.HasPrefix(auth, bearerAuthPrefix) {
			// Check credentials
			if auth == fmt.Sprintf("%s %s", bearerAuthPrefix, token) {

				// Delegate request to the given handle
				h(w, r, ps)
				return
			}
		}
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
	}
}
