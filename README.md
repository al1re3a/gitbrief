<!-- readme-refresh:start -->
<p align="center">
  <picture>
    <source media="(prefers-color-scheme: dark)" srcset="assets/readme-banner.png">
    <source media="(prefers-color-scheme: light)" srcset="assets/readme-banner.png">
    <img alt="GitBrief project banner" src="assets/readme-banner.png" width="100%">
  </picture>
</p>

<h1 align="center">📝 GitBrief</h1>

<p align="center"><strong>Turn any Git diff into a deterministic review brief and risk hints.</strong></p>

<p align="center">
  <a href="https://github.com/al1re3a/gitbrief/actions/workflows/ci.yml"><img alt="CI" src="https://github.com/al1re3a/gitbrief/actions/workflows/ci.yml/badge.svg"></a>
  <a href="https://go.dev/"><img alt="Go" src="https://img.shields.io/badge/Go-1.23%2B-00ADD8?logo=go&logoColor=white"></a>
  <a href="LICENSE"><img alt="License: MIT" src="https://img.shields.io/badge/License-MIT-fbbf24.svg"></a>
  <a href="https://github.com/al1re3a/gitbrief/stargazers"><img alt="GitHub stars" src="https://img.shields.io/github/stars/al1re3a/gitbrief?style=flat&color=8b5cf6"></a>
  <a href="https://github.com/al1re3a/gitbrief/issues"><img alt="Open issues" src="https://img.shields.io/github/issues/al1re3a/gitbrief?style=flat&color=06b6d4"></a>
</p>

<p align="center">
  <a href="https://github.com/al1re3a/gitbrief"><img alt="Source" src="https://img.shields.io/badge/Source-open-111827?style=for-the-badge&logo=github&logoColor=white"></a>
  <a href="#usage"><img alt="Quick Start" src="https://img.shields.io/badge/Quick_Start-open-0f766e?style=for-the-badge&logo=gnubash&logoColor=white"></a>
  <a href="CONTRIBUTING.md"><img alt="Contribute" src="https://img.shields.io/badge/Contribute-open-7c3aed?style=for-the-badge&logo=github&logoColor=white"></a>
  <a href="SECURITY.md"><img alt="Security" src="https://img.shields.io/badge/Security-open-b91c1c?style=for-the-badge&logo=securityscorecard&logoColor=white"></a>
</p>

<p align="center">
  <img src="https://skillicons.dev/icons?i=go,githubactions" alt="Go and GitHub Actions" height="42">
</p>

> [!NOTE]
> Risk hints help focus review attention and do not replace a security review or project-specific test plan.

## 📑 Contents

- [At a glance](#-at-a-glance)
- [Why](#why)
- [Usage](#usage)
- [Output includes](#output-includes)
- [Development](#development)

---

## 🔎 At a glance

| | |
|---|---|
| **Purpose** | Deterministic pull-request summaries and review-risk hints from any Git diff. Zero-dependency Go CLI. |
| **Input** | Git diff |
| **Output** | Markdown or JSON brief |
| **Runtime** | Go 1.23+ |
| **CI** | ✅ Linux · Windows |
| **Status** | ✅ Maintained |

<details>
<summary><strong>🧭 How it works</strong></summary>

```mermaid
flowchart LR
    A["Git diff"] --> B["Summarize changes"]
    B --> C["Markdown or JSON brief"]
```

</details>

<details>
<summary><strong>📁 Repository layout</strong></summary>

```text
gitbrief/
├── .github/
├── examples/
├── go.mod
├── main.go
└── README.md
```

</details>

<details>
<summary><strong>🤝 Contributors</strong></summary>

<br>
<a href="https://github.com/al1re3a/gitbrief/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=al1re3a/gitbrief" alt="Contributors">
</a>

</details>
<!-- readme-refresh:end -->

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
