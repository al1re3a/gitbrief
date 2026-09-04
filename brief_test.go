package main

import (
	"strings"
	"testing"
)

func sample(path string, additions int, deletions int) string {
	var b strings.Builder
	b.WriteString("diff --git a/" + path + " b/" + path + "\n--- a/" + path + "\n+++ b/" + path + "\n@@ -1 +1 @@\n")
	for i := 0; i < deletions; i++ {
		b.WriteString("-old\n")
	}
	for i := 0; i < additions; i++ {
		b.WriteString("+new\n")
	}
	return b.String()
}

func TestParseDiff(t *testing.T) {
	files := ParseDiff(sample("src/app.go", 3, 2))
	if len(files) != 1 || files[0].Additions != 3 || files[0].Deletions != 2 {
		t.Fatalf("unexpected parse: %#v", files)
	}
}

func TestCategories(t *testing.T) {
	input := sample("src/app.go", 1, 0) + sample("tests/app_test.go", 1, 0) + sample("README.md", 1, 0)
	summary := BuildSummary(input)
	if summary.Categories["source"] != 1 || summary.Categories["tests"] != 1 || summary.Categories["docs"] != 1 {
		t.Fatalf("unexpected categories: %#v", summary.Categories)
	}
}

func TestMissingTestsRisk(t *testing.T) {
	summary := BuildSummary(sample("src/auth.go", 2, 1))
	joined := strings.Join(summary.Risks, " ")
	if !strings.Contains(joined, "without test") || !strings.Contains(joined, "Security") {
		t.Fatalf("missing risks: %#v", summary.Risks)
	}
}

func TestMarkdown(t *testing.T) {
	output := Markdown(BuildSummary(sample("main.go", 1, 1)))
	if !strings.Contains(output, "| `main.go` | source | +1 / -1 |") {
		t.Fatalf("unexpected markdown: %s", output)
	}
}
