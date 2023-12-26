package handlers

import (
	"context"
	"net/http"
	"regexp"
	"strings"
)

type routeCtxKey struct{}

func newRoute(method, pattern string, handler http.HandlerFunc ) route {
	return route{method, regexp.MustCompile("^"+pattern+"$"), handler}
}

type route struct {
	method string
	regex *regexp.Regexp
	handler http.HandlerFunc
}


func Serve(w http.ResponseWriter, r *http.Request) {
	var allow []string
	for _, route := range routes {
		matches := route.regex.FindStringSubmatch(r.URL.Path)

		if len(matches) > 0 {
			if r.Method != route.method {
				allow = append(allow, route.method)
				continue
			}

			ctx := context.WithValue(r.Context(), routeCtxKey{}, matches[1:] )
			route.handler(w, r.WithContext(ctx))
			return 
		}
	}

	if len(allow) > 0 {
		w.Header().Set("Allow", strings.Join(allow, ", "))
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.NotFound(w, r)
}

func getField(r *http.Request, index int) string {
	fields := r.Context().Value(routeCtxKey{}).([]string)
	return fields[index]
}