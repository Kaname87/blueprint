package main

import (
	"flag"
	//"github.com/Kaname87/blueprint/chap1/src/trace"
	"log"
	"net/http"
	// "os"
	"path/filepath"
	"sync"
	"text/template"
)

// tmpl represents a single template
type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

// ServeHTTP handles the HTTP Request
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, r)
}
func main() {
	// root
	var addr = flag.String("addr", ":8080", "The addr of the application.")
	flag.Parse() //parse the flag
	r := newRoom()
	//r.tracer = trace.New(os.Stdout)
	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", r)
	// get the room going
	go r.run()
	// start web server
	log.Println("Starting web server on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
