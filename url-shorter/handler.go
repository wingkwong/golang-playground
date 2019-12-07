package urlshorter

import (
	"log"
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if target, ok := pathsToUrls[path]; ok {
			log.Printf("Redirecting %s to %s\n", path, target)
			http.Redirect(w, r, target, http.StatusMovedPermanently)
		}
		fallback.ServeHTTP(w, r)
	}
}
