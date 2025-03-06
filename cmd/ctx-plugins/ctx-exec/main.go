package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s COMMAND\n", os.Args[0])
		os.Exit(1)
	}
	
	cmd := strings.Join(os.Args[1:], " ")
	execCmd := exec.Command("sh", "-c", cmd)
	
	output, err := execCmd.CombinedOutput()
	
	fmt.Printf("<exec-output cmd=%q>%s\n</exec-output>\n", cmd, output)
	
	if err \!= nil {
		os.Exit(1)
	}
}

