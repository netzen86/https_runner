package main

import (
	"fmt"
	"os"
	"log"
	"net/http"
	"os/exec"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	arg := os.Args[1]
	param := req.URL.Query().Get("param")
	fmt.Printf("!!! param %s %s\n", arg, param)
	cmd := exec.Command(arg, param)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(w, "This is an example server %s.\n", param)
}

func main() {
	http.HandleFunc("/", HelloServer)
	err := http.ListenAndServeTLS(":8888", "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
