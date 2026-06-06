# Contributing

This repository is a learning project first. Prefer changes that make a concept easier to read, run, test, or compare.

## Topic Template

Each new topic should answer four questions:

1. What problem does this structure or algorithm solve?
2. What invariant makes it correct?
3. What are the time and space costs?
4. Which tests prove the important edge cases?

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
- Use `go vet ./...` and `go test ./...` as the required local checks.
- Avoid committing IDE files, coverage reports, compiled binaries, or module cache files.
- When adding a LeetCode-style exercise, keep the problem note close to the implementation.

## Documentation Style

- Start every topic README with a short concept summary.
- Include a tiny usage example when it helps.
- Add complexity notes in the form `Time: O(...)`, `Space: O(...)`.
- Link related topics, for example `Heap` to priority queues or `Union` to graph connectivity.
