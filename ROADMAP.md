# Roadmap

This roadmap keeps the repository moving toward the quality of strong open-source learning projects.

## Milestone 1: Stable Learning Entry

- Keep `go test ./...` green in CI.
- Maintain a short root README and a deeper learning path.
- Keep generated files, IDE files, and binaries out of git.
- Add table-driven tests for core structures.

## Milestone 2: Topic Quality

- Give each major topic a README with invariant, complexity, and usage.
- Add tests for currently untested packages and edge cases.
- Move console-heavy examples toward testable examples where practical.
- Add benchmark examples for sorting, heap, union-find, and graph traversal.

## Milestone 3: Exercise Mode

- Add beginner exercises with TODO markers and hidden/reference tests.
- Provide guided tasks for each learning pass.
- Add a progress checklist for learners.

## Milestone 4: API Consistency

- Normalize package names gradually without breaking all imports at once.
- Prefer constructors for exported structures.
- Replace panic-heavy teaching code with error-returning variants where the lesson is engineering robustness.
- Document intentional panics when the topic is about invariants.

## Milestone 5: Advanced Learning

- Add generic versions of selected structures.
- Add complexity visualizations or diagrams for trees and graphs.
- Add Dijkstra, topological sort, MST, dynamic programming, and string matching tracks.
