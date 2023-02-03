package urlshort

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v3"
)

// declare global types for YAML KeyValue pairs &
// list of pairs
type Pair struct {
	Path string
	URL  string
}

type Pairs struct {
	Pairs []Pair
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	//	TODO: Implement this...
	return func(w http.ResponseWriter, r *http.Request) {
		// if this is not a GET Request
		if r.Method != http.MethodGet {
			// Fallback and redirect
			fallback.ServeHTTP(w, r)
			return
		}

		// verify the URL and handle any errors
		url, ok := pathsToUrls[r.URL.Path]

		if !ok {
			fallback.ServeHTTP(w, r)
			return
		}

		w.Header().Add("Location", url)
		w.WriteHeader(301)
		fmt.Printf("%s %s %d: %s\n", r.Method, r.URL.Path, 301, url)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	var pairs Pairs
	err := yaml.Unmarshal(yml, &pairs)

	// create a path to existing URLS
	pathsToUrls := make(map[string]string, len(pairs.Pairs))

	// iterate over the URLS
	for _, entry := range pairs.Pairs {
		pathsToUrls[entry.Path] = entry.URL
	}
	return MapHandler(pathsToUrls, fallback), err
}
