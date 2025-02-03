package main

import (
    "context"
    "net/http"
    "log"
)

// Register the plugin
func registerHandlers(registerFunc func(name string, handler func(context.Context, map[string]interface{}, http.Handler) (http.Handler, error))) {
    registerFunc("header_overwrite_plugin", myHandler)
}

// myHandler modifies the request headers
func myHandler(ctx context.Context, extra map[string]interface{}, h http.Handler) (http.Handler, error) {
    return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
        // Log the original request headers
        log.Printf("Original Headers: %v", req.Header)

        // Overwrite a specific header
        req.Header.Set("X-Custom-Header", "MyValue")

        // Log the modified request headers
        log.Printf("Modified Headers: %v", req.Header)

        // Call the next handler in the chain
        h.ServeHTTP(w, req)
    }), nil
}

func main() {}