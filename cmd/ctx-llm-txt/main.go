package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	addr     = flag.String("addr", ":8080", "HTTP service address")
	domain   = flag.String("domain", "llm.txt.tmc.dev", "Domain to serve on")
	cacheDir = flag.String("cache-dir", "./cache", "Directory to cache files")
)

func main() {
	flag.Parse()
	
	if err := os.MkdirAll(*cacheDir, 0755); err \!= nil {
		log.Fatalf("Failed to create cache directory: %v", err)
	}
	
	http.HandleFunc("/", handleRoot)
	
	fmt.Printf("Starting server on %s\n", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "LLM Text Service")
}

