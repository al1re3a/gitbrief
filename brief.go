package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"path/filepath"
	"sort"
	"strings"
)

type FileChange struct {
	Path      string `json:"path"`
	Additions int    `json:"additions"`
	Deletions int    `json:"deletions"`
	Category  string `json:"category"`
}

type Summary struct {
	Title        string         `json:"title"`
	FilesChanged int            `json:"files_changed"`
	Additions    int            `json:"additions"`
	Deletions    int            `json:"deletions"`
	Categories   map[string]int `json:"categories"`
	Files        []FileChange   `json:"files"`
	Risks        []string       `json:"risks"`
	Fingerprint  string         `json:"fingerprint"`
}

func category(path string) string {
	lower := strings.ToLower(filepath.ToSlash(path))
	base := filepath.Base(lower)
	switch {
	case strings.Contains(lower, "/test/") || strings.Contains(lower, "/tests/") || strings.Contains(lower, "_test.") || strings.Contains(lower, ".test.") || strings.Contains(lower, ".spec."):
		return "tests"
	case strings.HasSuffix(lower, ".md") || strings.HasPrefix(lower, "docs/") || strings.Contains(lower, "/docs/"):
		return "docs"
	case strings.HasPrefix(lower, ".github/") || base == "dockerfile" || strings.HasSuffix(lower, ".yml") || strings.HasSuffix(lower, ".yaml") || strings.HasSuffix(lower, ".toml"):
		return "config"
	case base == "package.json" || base == "go.mod" || base == "cargo.toml" || base == "pyproject.toml":
		return "dependencies"
	case strings.HasSuffix(lower, ".png") || strings.HasSuffix(lower, ".jpg") || strings.HasSuffix(lower, ".svg"):
		return "assets"
	default:
		return "source"
	}
}

func ParseDiff(input string) []FileChange {
	scanner := bufio.NewScanner(strings.NewReader(input))
	files := []FileChange{}
	current := -1
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "diff --git a/") {
			parts := strings.SplitN(strings.TrimPrefix(line, "diff --git a/"), " b/", 2)
			if len(parts) == 2 {
				files = append(files, FileChange{Path: parts[1], Category: category(parts[1])})
				current = len(files) - 1
			}
			continue
		}
		if current < 0 {
			continue
		}
		if strings.HasPrefix(line, "+") && !strings.HasPrefix(line, "+++") {
			files[current].Additions++
		} else if strings.HasPrefix(line, "-") && !strings.HasPrefix(line, "---") {
			files[current].Deletions++
		}
	}
	return files
}

func BuildSummary(input string) Summary {
	files := ParseDiff(input)
	result := Summary{Files: files, FilesChanged: len(files), Categories: map[string]int{}}
	for _, file := range files {
		result.Additions += file.Additions
		result.Deletions += file.Deletions
		result.Categories[file.Category]++
	}
	if result.FilesChanged == 0 {
		result.Title = "No code changes"
	} else {
		primary := "source"
		max := -1
		for name, count := range result.Categories {
			if count > max || (count == max && name < primary) {
				primary, max = name, count
			}
		}
		result.Title = fmt.Sprintf("Update %s across %d file(s)", primary, result.FilesChanged)
	}
	changed := result.Additions + result.Deletions
	if changed > 500 {
		result.Risks = append(result.Risks, fmt.Sprintf("Large change set: %d lines", changed))
	}
	if result.Categories["source"] > 0 && result.Categories["tests"] == 0 {
		result.Risks = append(result.Risks, "Source changed without test changes")
	}
	if result.Categories["dependencies"] > 0 {
		result.Risks = append(result.Risks, "Dependency manifest changed")
	}
	for _, file := range files {
		lower := strings.ToLower(filepath.ToSlash(file.Path))
		if strings.Contains(lower, "auth") || strings.Contains(lower, "permission") || strings.HasPrefix(lower, ".github/workflows/") {
			result.Risks = append(result.Risks, "Security or automation-sensitive files changed")
			break
		}
	}
	sort.Strings(result.Risks)
	digest := sha256.Sum256([]byte(input))
	result.Fingerprint = hex.EncodeToString(digest[:6])
	return result
}

func Markdown(summary Summary) string {
	var builder strings.Builder
	fmt.Fprintf(&builder, "## %s\n\n**%d files** · **+%d / -%d** · `%s`\n\n", summary.Title, summary.FilesChanged, summary.Additions, summary.Deletions, summary.Fingerprint)
	builder.WriteString("### Change map\n\n| File | Category | Delta |\n|---|---|---:|\n")
	for _, file := range summary.Files {
		fmt.Fprintf(&builder, "| `%s` | %s | +%d / -%d |\n", strings.ReplaceAll(file.Path, "|", "\\|"), file.Category, file.Additions, file.Deletions)
	}
	if len(summary.Risks) > 0 {
		builder.WriteString("\n### Review notes\n\n")
		for _, risk := range summary.Risks {
			fmt.Fprintf(&builder, "- %s\n", risk)
		}
	}
	return builder.String()
}
