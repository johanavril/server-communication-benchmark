package main

import (
	"net/http"
)

func run() error {
	http.HandleFunc("/ping", ping)
	http.HandleFunc("/small", small)
	http.HandleFunc("/big", big)
	if err := http.ListenAndServe(":8081", nil); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}
