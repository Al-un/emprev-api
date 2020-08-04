package core

import (
	"net/http"
)

// DoIfLogged is the authentication guard middleware which only allows
// access to authenticated users
func DoIfLogged(next func(w http.ResponseWriter, r *http.Request, claims JwtClaims)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, httpStatus := DecodeJWT(r)
		if httpStatus != http.StatusOK {
			w.WriteHeader(httpStatus)
			return
		}

		next(w, r, *claims)
	})
}

// DoIfAdmin is the authorization guard middleware which allows an handler
// to be executed only if the JWT contains a valid admin profile
func DoIfAdmin(next func(w http.ResponseWriter, r *http.Request, claims JwtClaims)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims, httpStatus := DecodeJWT(r)
		if httpStatus != http.StatusOK {
			w.WriteHeader(httpStatus)
			return
		}

		if !claims.IsAdmin {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		next(w, r, *claims)
	})
}

// AddCorsHeaders provides the necessary CORS headers for each request and
// also handle OPTIONS requests.
//
// Method field is handled by Gorilla
func AddCorsHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Default configuration is quite loose...
		corsAllowedHosts := "*"
		corsAllowedHeaders := "*"
		corsAllowedMethods := "*"

		// CORS
		w.Header().Set("Access-Control-Allow-Origin", corsAllowedHosts)
		w.Header().Set("Access-Control-Allow-Methods", corsAllowedMethods)
		w.Header().Set("Access-Control-Allow-Headers", corsAllowedHeaders)

		// Proceed for non-preflight requests only
		if r.Method != http.MethodOptions {
			next.ServeHTTP(w, r)
		} else {
			// Handle OPTIONS requests here
		}
	})
}

// AddJSONHeaders add the required header for accepting and providing JSON
func AddJSONHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// JSON
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Accept", "application/json")

		// Next
		next.ServeHTTP(w, r)
	})
}
