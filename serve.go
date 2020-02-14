package main

import (
	"net/http"
	"fmt"
  "flag"
)

func main() {
  optDirectory := flag.String("d", "./", "directory")
  flag.Parse()
  directory := *optDirectory;

	http.Handle("/", http.FileServer(http.Dir(directory)))
  done := make(chan bool)
	go http.ListenAndServe(":3000", nil)
	fmt.Println("http://localhost:3000")
  <-done
}
