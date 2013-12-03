package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "<path>")
		return
	}
	path := os.Args[1]
	fmt.Println("Serving", path)
	panic(http.ListenAndServe(":9191", http.FileServer(http.Dir(path))))
}
