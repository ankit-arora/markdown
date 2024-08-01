//go:build linux || darwin

// Package main is generating markdown.
package main

import (
	"os"

	md "github.com/ankit-arora/markdown"
)

//go:generate go run main.go

func main() {
	f, err := os.Create("generated.md")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := md.NewMarkdown(f).
		H1("Alert example").
		Note("This is note").LF().
		Tip("This is tip").LF().
		Important("This is important").LF().
		Warning("This is warning").LF().
		Caution("This is caution").LF().
		Build(); err != nil {
		panic(err)
	}
}
