package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/gorilla/mux"
)

func benchmark(b *testing.B, handler http.Handler, path string) {
	res := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, path, nil)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for range b.N {
		handler.ServeHTTP(res, req)
		if res.Code != http.StatusOK {
			b.Fatalf("failed to execute. status: %d", res.Code)
		}
	}
}

func BenchmarkServeMux0(b *testing.B) {
	benchmark(b, createServeMux(), "/")
}

func BenchmarkServeMux1(b *testing.B) {
	benchmark(b, createServeMux(), "/foo")
}

func BenchmarkServeMux5(b *testing.B) {
	benchmark(b, createServeMux(), "/foo/bar/baz/qux/quux")
}

func BenchmarkServeMux10(b *testing.B) {
	benchmark(b, createServeMux(), "/foo/bar/baz/qux/quux/corge/grault/garply/waldo/fred")
}

func BenchmarkChi0(b *testing.B) {
	benchmark(b, createChiMux(), "/")
}

func BenchmarkChi1(b *testing.B) {
	benchmark(b, createChiMux(), "/foo")
}

func BenchmarkChi5(b *testing.B) {
	benchmark(b, createChiMux(), "/foo/bar/baz/qux/quux")
}

func BenchmarkChi10(b *testing.B) {
	benchmark(b, createChiMux(), "/foo/bar/baz/qux/quux/corge/grault/garply/waldo/fred")
}

func BenchmarkGorilla0(b *testing.B) {
	benchmark(b, createGorillaMux(), "/")
}

func BenchmarkGorilla1(b *testing.B) {
	benchmark(b, createGorillaMux(), "/foo")
}

func BenchmarkGorilla5(b *testing.B) {
	benchmark(b, createGorillaMux(), "/foo/bar/baz/qux/quux")
}

func BenchmarkGorilla10(b *testing.B) {
	benchmark(b, createGorillaMux(), "/foo/bar/baz/qux/quux/corge/grault/garply/waldo/fred")
}

func createServeMux() http.Handler {
	mux := http.NewServeMux()
	// No path parametes.
	mux.HandleFunc("GET /{$}", func(w http.ResponseWriter, r *http.Request) {
	})
	// Only one path parameter.
	mux.HandleFunc("GET /{foo}", func(w http.ResponseWriter, r *http.Request) {
		_ = r.PathValue("foo")
	})
	// 5 path parameters.
	mux.HandleFunc("GET /{foo}/{bar}/{baz}/{qux}/{quux}", func(w http.ResponseWriter, r *http.Request) {
		_ = r.PathValue("foo")
		_ = r.PathValue("bar")
		_ = r.PathValue("baz")
		_ = r.PathValue("qux")
		_ = r.PathValue("quux")
	})
	// 10 path parameters.
	mux.HandleFunc("GET /{foo}/{bar}/{baz}/{qux}/{quux}/{corge}/{grault}/{garply}/{waldo}/{fred}", func(w http.ResponseWriter, r *http.Request) {
		_ = r.PathValue("foo")
		_ = r.PathValue("bar")
		_ = r.PathValue("baz")
		_ = r.PathValue("qux")
		_ = r.PathValue("quux")
		_ = r.PathValue("corge")
		_ = r.PathValue("grault")
		_ = r.PathValue("garply")
		_ = r.PathValue("waldo")
		_ = r.PathValue("fred")
	})
	return mux
}

func createChiMux() http.Handler {
	r := chi.NewRouter()
	// No path parameters.
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
	})
	// Only one path parameter.
	r.Get("/{foo}", func(w http.ResponseWriter, r *http.Request) {
		_ = chi.URLParam(r, "foo")
	})
	// 5 path parameters.
	r.Get("/{foo}/{bar}/{baz}/{qux}/{quux}", func(w http.ResponseWriter, r *http.Request) {
		_ = chi.URLParam(r, "foo")
		_ = chi.URLParam(r, "bar")
		_ = chi.URLParam(r, "baz")
		_ = chi.URLParam(r, "qux")
		_ = chi.URLParam(r, "quux")
	})
	// 10 path parameters.
	r.Get("/{foo}/{bar}/{baz}/{qux}/{quux}/{corge}/{grault}/{garply}/{waldo}/{fred}", func(w http.ResponseWriter, r *http.Request) {
		_ = chi.URLParam(r, "foo")
		_ = chi.URLParam(r, "bar")
		_ = chi.URLParam(r, "baz")
		_ = chi.URLParam(r, "qux")
		_ = chi.URLParam(r, "quux")
		_ = chi.URLParam(r, "corge")
		_ = chi.URLParam(r, "grault")
		_ = chi.URLParam(r, "garply")
		_ = chi.URLParam(r, "waldo")
		_ = chi.URLParam(r, "fred")
	})
	return r
}

func createGorillaMux() http.Handler {
	r := mux.NewRouter()
	// No path parameters.
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	}).Methods("GET")
	// Only one path parameter.
	r.HandleFunc("/{foo}", func(w http.ResponseWriter, r *http.Request) {
		_ = mux.Vars(r)["foo"]
	}).Methods("GET")
	// 5 path parameters.
	r.HandleFunc("/{foo}/{bar}/{baz}/{qux}/{quux}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		_ = vars["foo"]
		_ = vars["bar"]
		_ = vars["baz"]
		_ = vars["qux"]
		_ = vars["quux"]
	}).Methods("GET")
	// 10 path parameters.
	r.HandleFunc("/{foo}/{bar}/{baz}/{qux}/{quux}/{corge}/{grault}/{garply}/{waldo}/{fred}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		_ = vars["foo"]
		_ = vars["bar"]
		_ = vars["baz"]
		_ = vars["qux"]
		_ = vars["quux"]
		_ = vars["corge"]
		_ = vars["grault"]
		_ = vars["garply"]
		_ = vars["waldo"]
		_ = vars["fred"]
	}).Methods("GET")
	return r
}
