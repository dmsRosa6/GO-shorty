package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/dmsRosa6/go-shorty/internal/service"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	service service.URLShortenerService
}

func NewURLShortenerServer() *Server {
	return &Server{service: *service.NewURLShortenerService()}
}

func (s *Server) Routes() http.Handler {
	r := chi.NewRouter()

	r.Post("/shorten", s.shortenHandler)
	r.Get("/{code}", s.redirectHandler)

	return r
}

func (s *Server) shortenHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	url := r.FormValue("url")

	shorten, err := s.service.Shorten(ctx, url)
	if err != nil {
		http.Error(w, "failed to store url", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Shortened URL code: %s", shorten)
}

func (s *Server) redirectHandler(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	ctx := context.Background()

	url, err := s.service.Resolve(ctx, code)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	http.Redirect(w, r, url, http.StatusFound)
}
