package main

import (
	"net/http"
	"fmt"
  "flag"
)

func main() {
  directory := flag.String("d", "./", "directory")
  port := flag.String("p", "3000", "port")
  flag.Parse()

	http.Handle("/", http.FileServer(http.Dir(*directory)))
  done := make(chan bool)

  bind := "localhost:" + *port
	go http.ListenAndServe(bind, nil)
	fmt.Println("Serving at http://" + bind)
  <-done
}
