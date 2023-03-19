package middlewares

import "net/http"

func ContentTypeJson(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=UTF-8")

		next.ServeHTTP(w, req)
	})
}
