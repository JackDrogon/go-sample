# PROJECT KNOWLEDGE BASE

**Generated:** 2026-03-24 01:11:56 CST
**Commit:** 2795fd9
**Branch:** master

## OVERVIEW
Small Go CLI starter. Process entrypoint lives in `cmd/go-sample`, CLI behavior stays in `internal/app`, and build-time version metadata lives in `internal/version`.

Root-level automation matters as much as the code here: `justfile` is the canonical task runner, `.goreleaser.yml` owns release builds, `.golangci.yml` defines an explicit lint allowlist, and CI/spellcheck/coverage policy live under `.github/`, `typos.toml`, and `codecov.yml`.

## STRUCTURE
```text
./
├── cmd/go-sample/         # process entrypoint; stderr and exit behavior
├── internal/app/          # flag parsing, CLI flow, user-visible output
├── internal/version/      # build-time version injection target
├── .github/workflows/     # CI jobs; installs tools and runs canonical just-based checks
├── justfile               # canonical local build/test/lint/release tasks
├── .goreleaser.yml        # release matrix and ldflags version injection
├── .golangci.yml          # explicit linter allowlist
├── typos.toml             # spellcheck exceptions and ignore rules
└── codecov.yml            # coverage upload policy
```

## WHERE TO LOOK
| Task | Location | Notes |
|------|----------|-------|
| Change process startup or exit handling | `cmd/go-sample/main.go` | Thin wrapper around `app.Run`; prints errors to stderr and exits non-zero |
| Change CLI flags or output | `internal/app/app.go` | Owns `--name`, `--version`, arg validation, and stdout writes |
| Update CLI behavior tests | `internal/app/app_test.go` | Colocated same-package tests with `t.Run(...)` subtests |
| Change version reporting | `internal/version/version.go` | `Info()` reads the build-injected `Version` variable |
| Change version injection | `justfile`, `.goreleaser.yml` | Both inject `internal/version.Version` via `-ldflags -X ...` |
| Change local developer workflow | `justfile` | Source of truth for build/test/lint/check/pre-commit/ci |
| Change CI or quality gates | `.github/workflows/ci.yml`, `.golangci.yml`, `codecov.yml`, `typos.toml` | CI mirrors canonical `just` recipes as separate steps |
| Review maintainer expectations | `README.md`, `CONTRIBUTING.md` | Human-facing workflow and contribution guidance |

## CODE MAP
| Symbol | Type | Location | Refs | Role |
|--------|------|----------|------|------|
| `main` | Function | `cmd/go-sample/main.go` | 1 | Top-level process entry; delegates to `app.Run` and owns stderr/exit behavior |
| `Run` | Function | `internal/app/app.go` | 6 | CLI control flow; parses flags, validates args, prints greeting/version |
| `Info` | Function | `internal/version/version.go` | 3 | Version lookup used by the `--version` path |
| `Version` | Variable | `internal/version/version.go` | config-bound | Build-time injection target for local builds and GoReleaser |

## CONVENTIONS
- Use `just` recipes, not a parallel `make` workflow.
- Keep `cmd/go-sample` thin; CLI behavior belongs in `internal/app`.
- Keep version injection wired to `internal/version.Version`; local build and release tooling depend on that exact symbol path.
- Tests are colocated same-package white-box tests; standard execution path is `go test ./...` via `just` or CI.
- Keep CI aligned with `justfile`, but preserve step-level visibility in GitHub Actions by running the canonical recipes as separate steps.
- `golangci-lint` runs an explicit allowlist only: `errcheck`, `govet`, `ineffassign`, `staticcheck`, `unused`.
- Formatting is file-type specific: Go uses tabs, docs/config use 2 spaces, `justfile` uses 4 spaces.

## ANTI-PATTERNS (THIS PROJECT)
- Do not leave stale path references in automation; the CLI entrypoint is `./cmd/go-sample` and automation should keep using that path.
- Do not commit generated artifacts from normal workflows: `bin/`, `coverage.out`, `coverage.html`, `dist/`.
- Do not add more tests that mutate `internal/version.Version` unless state is saved and restored inside the test.
- Do not assume coverage upload is a blocking quality gate; current Codecov/CI policy is intentionally fail-open.
- Do not introduce a second task-runner convention when `justfile` already defines the canonical build/test/lint/release flow.
- Do not use `just check` or `just pre-commit` as the CI gate; they include `fmt` and can mutate the working tree.

## UNIQUE STYLES
- Binary name, flagset name, module suffix, and entrypoint folder all use `go-sample`.
- `Run(args, stdout)` is the seam for CLI behavior; `main.go` should stay focused on process concerns.
- `check` and `pre-commit` mean the same verification chain: `fmt -> lint -> test -> spellcheck`.
- `ci` is the non-mutating verification chain: `build -> cover -> lint -> spellcheck`; GitHub Actions mirrors it as separate steps for easier debugging.
- Release and local build paths both rely on ldflags version injection rather than runtime version discovery.

## COMMANDS
```bash
just build
just run -- --help
just run -- --version
just test
just cover
just ci
just lint
just spellcheck
just check
just pre-commit
just release-snapshot
```

## NOTES
- Actual code entrypoint is `./cmd/go-sample`, and CI runs the corresponding canonical `just` recipes as separate steps.
- Build output goes to `bin/go-sample`.
- Coverage output is `coverage.out`; Codecov uploads that file and ignores `*_test.go` plus `testdata` paths.
- No child `AGENTS.md` files yet: `cmd/go-sample`, `internal/app`, and `internal/version` are each too small to justify separate docs without duplicating this root file.
