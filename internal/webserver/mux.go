package webserver

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func NewMux() *chi.Mux {
	mux := chi.NewRouter()

	// middleware
	mux.Use(middleware.Logger)
	mux.Use(middleware.Recoverer)
	mux.Use(middleware.CleanPath)
	mux.Use(middleware.Heartbeat("/"))
	mux.Use(middleware.AllowContentType("application/json"))
	mux.Use(middleware.AllowContentEncoding("deflate", "gzip"))

	// profiler & metrics
	mux.Mount("/debug", middleware.Profiler())
	mux.Mount("/metrics", promhttp.Handler())

	// routes
	bundleRoutes := chi.NewRouter()
	bundleRoutes.Post("/", CreateBundle)
	bundleRoutes.Get("/{bundle_id}", GetBundle)
	bundleRoutes.Patch("/{bundle_id}", UpdateBundle)
	bundleRoutes.Delete("/{bundle_id}", DeleteBundle)
	mux.Mount("/bundle", bundleRoutes)

	return mux
}
