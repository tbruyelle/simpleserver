package main

import (
	"fmt"
	"net/http"
	"os"
)

const port = "9191"

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "<path>")
		return
	}
	path := os.Args[1]
	fmt.Printf("Serving %s to http://localhost:%s\n", path, port)
	panic(http.ListenAndServe(":"+port, http.FileServer(http.Dir(path))))
}
