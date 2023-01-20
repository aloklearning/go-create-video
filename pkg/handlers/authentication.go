// The file is different as it is a helper and it should be in a seperate file
// For clean coding principals
package handlers

import (
	"fmt"
	"net/http"
)

// Adding basic API Key here as a global variable
// Could have been done via api_key table as well
const (
	projectAPIKey = "3251744f-125c-458b-b80d-6f623d2a34bc"
)

// HandleAuthentication handles the request and pass it to
// the real handler and doing the authentication
func AuthenticationMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Doing things before the handler
		if r.Header.Get("api_key") != projectAPIKey {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Printf("401 STATUS: User UnAuthorised\n")

			return
		}

		handler.ServeHTTP(w, r)
	})
}
