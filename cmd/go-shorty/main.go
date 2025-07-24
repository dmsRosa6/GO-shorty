package main

import (
	"net/http"

	"github.com/dmsRosa6/go-shorty/internal/handler"
)

func main() {
	srv := handler.NewURLShortenerServer()
	http.ListenAndServe(":8080", srv.Routes())
}
