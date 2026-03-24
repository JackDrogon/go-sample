# go-sample

A production-ready Go CLI starter with a small but maintainable project layout.

## Project Layout

```text
cmd/go-sample/     Program entrypoint
internal/app/         CLI behavior and application logic
internal/version/     Build-time version metadata
```

## Development

### Build and run locally

Use these when you want a fast local feedback loop while changing CLI behavior.

```bash
just build
just run -- --name team
just run -- --version
```

`just build` writes the binary to `bin/go-sample`. `just run -- ...` rebuilds first and then executes the compiled binary with your CLI arguments, which makes it a convenient default while iterating on flags or command output.

### Run tests and quality checks

Use the smaller commands while iterating, then the aggregate command before you push.

```bash
just test
just lint
just spellcheck
just cover
just ci
just check
just pre-commit
```

Start with `just test` when you only need a quick correctness check. Use `just lint` after refactors or API changes, and run `just spellcheck` when you want to catch repository-wide typos; it requires [`typos`](https://github.com/crate-ci/typos) to be installed locally. `just cover` generates `coverage.out` and prints function coverage, which is helpful when you want to inspect test reach. `just ci` is the CI-safe aggregate command and runs `build -> cover -> lint -> spellcheck` without mutating the working tree. GitHub Actions runs that same verification flow as separate `just` steps so build, coverage, lint, and spellcheck failures stay easy to identify in the workflow UI. Before pushing changes, `just check` is the main local verification command and runs `fmt -> lint -> test -> spellcheck`, while `just pre-commit` currently wraps the same flow as the safest final pass before committing.

### Release verification

Use this when you want to confirm the GoReleaser configuration can produce local artifacts without publishing anything.

```bash
just release-snapshot
```

This command is heavier than the normal edit/test loop, so it is best reserved for release preparation or for validating changes to the GoReleaser configuration.

## Module Path

This project uses `github.com/JackDrogon/go-sample` as its Go module path.

## Notes

- The binary is built to `bin/go-sample`.
- Version information is injected through `-ldflags` during builds.
- The default CLI supports `--name` and `--version` out of the box.
- A GitHub Actions workflow runs the same canonical `justfile` checks as separate steps on pushes and pull requests: `just build`, `just cover`, `just lint`, and `just spellcheck`.
- A repository-wide typos configuration is included for local and CI spellcheck runs.
- Dependabot and editor defaults are included so repository maintenance starts from day one.
- A GoReleaser config is included for local snapshot builds and future tagged releases.
- A lightweight Codecov config is included for coverage reporting from CI.
- Private repositories usually need a `CODECOV_TOKEN`; the default CI upload is fail-open so missing coverage uploads do not block builds.
- Contributor workflow and community behavior docs are included by default.
- Install [`typos`](https://github.com/crate-ci/typos) locally if you want `just spellcheck`, `just check`, or `just pre-commit` to pass.
- Use `just ci` when you want the same non-mutating verification flow that GitHub Actions mirrors in separate steps.
- Before accepting public contributions, update `CODE_OF_CONDUCT.md` with a real reporting contact or private escalation channel.
