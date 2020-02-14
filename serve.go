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
	go http.ListenAndServe(":" + *port, nil)
	fmt.Println("http://localhost:" + *port)
  <-done
}
