package handlers

import (
	"net/http"
	"gopkg.in/yaml.v2"
)

//MapHandler func
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

//YAMLHandler func
func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	pathUrls, err := parseYaml(yamlBytes)
	if err != nil {
		return nil, err
	}

	pathsToUrls := buildMap(pathUrls)

	pathsToUrls := map[string] string{}
	return MapHandler(pathsToUrls, fallback), nil
}

func buildMap(pathUrls []pathUrl) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, pu := range pathsToUrls {
		pathsToUrls[pu.Path] = pu.URL
	}
	return pathUrls
}

func parseYaml(data []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(yamlBytes, &pathUrls) 
	if err != nil {
		return nil, err
	}
	return pathUrls, nil
}

type pathUrl struct {
	Path string `yaml:"path"`
	URL string `yaml:"url"`
}