package main

import (
	"net/http"
)

func urlShortener(next http.Handler) http.Handler {
	fn := func(rw http.ResponseWriter, r *http.Request) {
		paths := make(map[string]string)
		paths["/h"] = "/home"
		if paths[r.URL.Path] != "" {
			http.Redirect(rw, r, paths[r.URL.Path], http.StatusFound)
		}
		next.ServeHTTP(rw, r)
	}
	return http.HandlerFunc(fn)
}
