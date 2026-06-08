# Contributing

This repository is a learning project first. Prefer changes that make a concept easier to read, run, test, compare, or extend.

## Required Local Checks

Run these before submitting a change:

```bash
go vet ./...
go test ./...
```

For concurrency changes, also consider:

```bash
go test -race ./BasicGo/sharedvars
```

## New Topic Requirements

Every new topic should update:

- The topic README or an equivalent documentation file.
- `DIRECTORY.md`.
- `docs/catalog.md`.
- `docs/data/knowledge-graph.json` if the topic adds or changes a learning node.
- Relevant roadmap, exercise, or checklist entries when the topic changes a learner path.

## Topic README Template

Each module README should follow the [Textbook Style Guide](docs/textbook-style-guide.md). At minimum, include:

1. Goal: what the topic teaches.
2. Prerequisites: what the learner should know first.
3. Core invariant: what must always remain true.
4. Complexity: time and space cost of important operations.
5. Practice tasks: small exercises the learner can complete.
6. Hints and reference answers using collapsible `<details>` blocks.
7. Test command: the exact `go test` command for the module.

Recommended layout:

```text
TopicName/
  README.md
  implementation.go
  implementation_test.go
```

## Code Style

- Keep implementations small and explicit.
- Prefer table-driven tests.
- Avoid committing IDE files, coverage reports, compiled binaries, or module cache files.
- Keep console output in demos; tests should assert behavior directly.
- When a function panics, either test the panic or document the invariant that permits it.

## Documentation Style

- Use English for Markdown files and code comments.
- Start every topic README with a short concept summary.
- Include a small usage example when it helps.
- Add complexity notes in the form `Time: O(...)`, `Space: O(...)`.
- Link related topics, for example `Heap` to priority queues or `Union` to graph connectivity.

## Review Checklist

Before publishing a batch of changes:

- Does every new topic have tests?
- Does each README state the invariant and complexity?
- Can a learner run only that package?
- Did `DIRECTORY.md`, `docs/catalog.md`, and the knowledge graph stay in sync?
- Did generated files, IDE settings, or local build artifacts stay out of git?
