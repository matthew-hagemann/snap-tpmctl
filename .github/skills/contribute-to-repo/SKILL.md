---
name: contribute-to-repo
description: >
  Full workflow for committing changes and opening a pull request against the
  upstream canonical/snap-tpmctl repository. Use this when asked to commit
  work, push, or open a PR on this repo.
---

## Repo layout

- **Upstream**: `git remote` named `upstream` → `github.com/canonical/snap-tpmctl`
- **Fork**: `git remote` named `origin` → `github.com/matthew-hagemann/snap-tpmctl`
- PRs must always target the **upstream** repo (`canonical/snap-tpmctl`), opened from the fork branch.

## Step-by-step workflow

### 1. Find recurring co-authors

Before committing, look up known collaborators in git history:

```bash
git log --format="%an <%ae>" | sort -u
```

Didier Roche-Tolomelli's email is `didier.roche@canonical.com`.
Always co-author commits with **both** Didier and Copilot.

### 2. Stage and commit

Stage all modified files and commit using **Conventional Commits** format:

```
<type>(<scope>): <short summary>

<body — what changed and why>

Co-authored-by: Didier Roche-Tolomelli <didier.roche@canonical.com>
Co-authored-by: Copilot <223556219+Copilot@users.noreply.github.com>
```

Common types: `feat`, `fix`, `refactor`, `chore`, `docs`, `test`.

```bash
git add -A
git commit -m "<message with trailers above>"
```

### 3. Push to the fork

```bash
git push origin <current-branch>
```

### 4. Open a PR against upstream

Use the GitHub MCP tool `create_pull_request` with:
- `owner`: `canonical`
- `repo`: `snap-tpmctl`
- `head`: `matthew-hagemann:<current-branch>`
- `base`: `main`
- Include a clear PR body summarising **what** changed and **why**.

### 5. Watch CI and fix failures

After opening the PR, check the CI status with `get_check_runs`.
If any check fails:

1. Fetch the job logs via the GitHub API:
   ```bash
   gh api /repos/canonical/snap-tpmctl/actions/jobs/<job-id>/logs
   ```
2. Read the linter output — the project uses **golangci-lint v2** with the config in `.golangci.yaml`.
3. Common issues in this repo:
   - `forcetypeassert`: suppress with `//nolint:forcetypeassert // <reason>` only when the assertion is genuinely safe (e.g. inside a type switch).
   - `gci` / `whitespace`: trailing spaces in comments or extra blank lines before `}`.
   - `godot`: comments must end with a period.
4. Fix, commit (same co-author trailers), and push to update the PR.
