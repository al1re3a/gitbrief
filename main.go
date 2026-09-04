package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	base := flag.String("base", "HEAD~1", "base git ref")
	format := flag.String("format", "markdown", "markdown or json")
	file := flag.String("file", "", "read a saved unified diff")
	flag.Parse()
	var data []byte
	var err error
	if *file != "" {
		data, err = os.ReadFile(*file)
	} else if info, _ := os.Stdin.Stat(); info != nil && info.Mode()&os.ModeCharDevice == 0 {
		data, err = io.ReadAll(os.Stdin)
	} else {
		data, err = exec.Command("git", "diff", *base+"...HEAD").Output()
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "gitbrief:", err)
		os.Exit(1)
	}
	summary := BuildSummary(string(data))
	if *format == "json" {
		encoded, _ := json.MarshalIndent(summary, "", "  ")
		fmt.Println(string(encoded))
	} else if *format == "markdown" {
		fmt.Print(Markdown(summary))
	} else {
		fmt.Fprintln(os.Stderr, "gitbrief: format must be markdown or json")
		os.Exit(1)
	}
}
