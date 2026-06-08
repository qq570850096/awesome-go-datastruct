# Roadmap

This roadmap keeps the repository moving toward the quality of strong open-source learning projects.

## Milestone 1: Stable Learning Entry

- Keep the root README as the main English learning portal.
- Keep `README.zh.md` only as a legacy compatibility pointer.
- Maintain [DIRECTORY.md](DIRECTORY.md), [docs/catalog.md](docs/catalog.md), and the knowledge graph together.
- Keep generated files, IDE files, binaries, coverage output, and module cache files out of git.

## Milestone 2: Topic Quality

- Give each major topic a README with goal, prerequisites, invariant, complexity, practice tasks, and test command.
- Add tests for currently weak edge cases.
- Move console-heavy examples toward testable examples where practical.
- Keep Markdown files and code comments in English.

## Milestone 3: Exercise Mode

- Expand [docs/exercise-matrix.md](docs/exercise-matrix.md) with more tasks per module.
- Add beginner exercises with TODO markers and reference tests where appropriate.
- Add progress checklists for each stage.
- Connect exercises back to `DIRECTORY.md` and the knowledge graph.

## Milestone 4: Expert Engineering

- Add benchmark examples for sorting, heap, union-find, and graph traversal.
- Add documented `go test -race` examples.
- Add pprof walkthroughs after benchmark coverage exists.
- Improve package boundary documentation and API review examples.

## Milestone 5: Automation And Maintenance

- Generate `DIRECTORY.md` from structured metadata after the manual format stabilizes.
- Validate `docs/data/knowledge-graph.json` in local tooling.
- Add link-checking scripts if the repository adopts automation.
- Address Dependabot-reported vulnerabilities.

## Current Constraints

- Do not add GitHub Actions workflow files until the available OAuth token has `workflow` scope.
- Do not move or rename Go packages as part of documentation-only roadmap work.
- Prefer additive improvements and focused refactors over broad package API churn.
