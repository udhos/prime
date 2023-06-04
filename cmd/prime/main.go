// Package main implements the program.
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"time"

	"github.com/jbarham/primegen"

	"github.com/udhos/boilerplate/boilerplate"
	"github.com/udhos/boilerplate/envconfig"
)

const version = "1.0.0"

type appConfig struct {
	port   string
	route  string
	health string
}

func main() {

	var showVersion bool
	flag.BoolVar(&showVersion, "version", showVersion, "show version")
	flag.Parse()

	me := filepath.Base(os.Args[0])

	{
		v := boilerplate.LongVersion(me + " version=" + version)
		if showVersion {
			fmt.Print(v)
			fmt.Println()
			return
		}
		log.Print(v)
	}

	env := envconfig.NewSimple(me)

	app := appConfig{
		port:   env.String("PORT", ":8080"),
		route:  env.String("ROUTE", "/prime/"),
		health: env.String("HEALTH", "/health"),
	}

	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    app.port,
		Handler: mux,
	}

	const root = "/"

	register(mux, app.port, root, func(w http.ResponseWriter, r *http.Request) { handlerRoot(&app, w, r) })
	register(mux, app.port, app.health, func(w http.ResponseWriter, r *http.Request) { handlerHealth(&app, w, r) })
	register(mux, app.port, app.route, func(w http.ResponseWriter, r *http.Request) { handlerRoute(&app, w, r) })

	go listenAndServe(server, app.port)

	<-chan struct{}(nil)
}

func register(mux *http.ServeMux, addr, path string, handler http.HandlerFunc) {
	mux.HandleFunc(path, handler)
	log.Printf("registered on port %s path %s", addr, path)
}

func listenAndServe(s *http.Server, addr string) {
	log.Printf("listening on port %s", addr)
	err := s.ListenAndServe()
	log.Printf("listening on port %s: %v", addr, err)
}

func handlerRoot(app *appConfig, w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s - 404 not found", r.RemoteAddr, r.Method, r.RequestURI)
	http.Error(w, "not found", http.StatusNotFound)
}

func handlerRoute(app *appConfig, w http.ResponseWriter, r *http.Request) {

	base := path.Base(r.RequestURI)

	n, errConv := strconv.ParseInt(base, 10, 64)
	if errConv != nil {
		log.Printf("%s %s %s - requested:%s error:%v",
			r.RemoteAddr, r.Method, r.RequestURI, base, errConv)
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	log.Printf("%s %s %s - requested:%s",
		r.RemoteAddr, r.Method, r.RequestURI, base)

	begin := time.Now()

	gen := primegen.New()

	var prime uint64

	for i := int64(1); i <= n; i++ {
		prime = gen.Next()
	}

	primeStr := strconv.FormatUint(prime, 10)

	elap := time.Since(begin)

	log.Printf("%s %s %s - requested:%s found:%s elapsed:%v",
		r.RemoteAddr, r.Method, r.RequestURI, base, primeStr, elap)

	http.Error(w, primeStr, http.StatusOK)
}

func handlerHealth(app *appConfig, w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s - 200 health ok", r.RemoteAddr, r.Method, r.RequestURI)
	http.Error(w, "health ok", http.StatusOK)
}
