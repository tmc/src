package main

import (
	"testing"
	"context"
	
	"rsc.io/script"
	"rsc.io/script/scripttest"
)

func TestCodeToGPT(t *testing.T) {
	engine := script.NewEngine()
	engine.Cmds["code-to-gpt"] = codeToGPTCmd
	ctx := context.Background()
	
	scripttest.Test(t, ctx, engine, nil, "testdata/*.txt")
}

var codeToGPTCmd = script.Command(
	script.CmdUsage{
		Summary: "Run code-to-gpt.sh script",
		Args:    "[args...]",
	}, func(s *script.State, args ...string) (script.WaitFunc, error) {
		// Implementation omitted for brevity
		return nil, nil
	})

