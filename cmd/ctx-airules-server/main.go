package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

type Rule struct {
	Pattern     string   `json:"pattern"`
	Description string   `json:"description"`
	Severity    string   `json:"severity"`
	Examples    []string `json:"examples,omitempty"`
}

type RuleSet struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Rules       []Rule `json:"rules"`
}

var (
	addr       = flag.String("addr", ":8081", "HTTP service address")
	outputDir  = flag.String("output-dir", "./outputs", "Directory to store generated rule files")
)

func main() {
	flag.Parse()
	
	if err := os.MkdirAll(*outputDir, 0755); err \!= nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}
	
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/generate", handleGenerate)
	http.HandleFunc("/health", handleHealth)
	
	port := os.Getenv("PORT")
	if port \!= "" {
		*addr = ":" + port
	}
	
	fmt.Printf("Starting server on %s\n", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "AI Rules Generator Service")
}

func handleGenerate(w http.ResponseWriter, r *http.Request) {
	if r.Method \!= http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	// Sample ruleset for demonstration
	ruleset := &RuleSet{
		Name:        "Generated Rules",
		Description: "Rules generated from source code",
		Rules: []Rule{
			{
				Pattern:     "type\\s+[A-Z][a-zA-Z0-9]*\\s+struct",
				Description: "Struct names should be in PascalCase and exported",
				Severity:    "error",
			},
		},
	}
	
	outputData, err := json.MarshalIndent(ruleset, "", "  ")
	if err \!= nil {
		http.Error(w, fmt.Sprintf("Failed to marshal rules: %v", err), http.StatusInternalServerError)
		return
	}
	
	filename := filepath.Join(*outputDir, "rules.json")
	if err := ioutil.WriteFile(filename, outputData, 0644); err \!= nil {
		http.Error(w, fmt.Sprintf("Failed to write rules: %v", err), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.Write(outputData)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

