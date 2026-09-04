# GitBrief

**A useful PR description before the model call.**

GitBrief turns a unified Git diff into a deterministic Markdown change map: files, categories, line counts, stable fingerprint, and review notes for large, untested, dependency, auth, or workflow changes.

```bash
go install github.com/al1re3a/gitbrief@latest
gitbrief -base main > PR_BODY.md
```

## Why

AI-written PR summaries can be helpful, but the factual layer should not depend on a model. GitBrief produces an instant local baseline that a human or agent can extend without losing the actual change shape.

## Usage

```bash
# Current branch versus main
gitbrief -base main

# Saved or piped diff
gitbrief -file change.diff
git diff --staged | gitbrief

# Automation
gitbrief -base origin/main -format json
```

## Output includes

- Per-file additions and deletions
- Source, tests, docs, config, dependency, and asset categories
- Missing-test and large-change notes
- Auth, permission, workflow, and dependency review flags
- Stable diff fingerprint for reruns

GitBrief never claims to understand behavior; it makes review scope visible.

## Development

```bash
go test ./...
go run . -file examples/change.diff
```

## License

MIT
