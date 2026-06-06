# Project Standards

The repository should be useful to two readers:

- A beginner who wants a clear path and runnable examples.
- A maintainer who wants predictable structure and reliable tests.

## Quality Bar

Every topic should aim for:

- Clear invariant: what must always remain true.
- Runnable package: `go test ./Topic`.
- Edge-case tests: empty input, one element, duplicates, missing values, and boundary capacity where relevant.
- Complexity note: time and space cost of important operations.
- Local explanation: the README or test names should tell the learner what to notice.

## Code Expectations

- Keep teaching implementations direct and readable.
- Prefer table-driven tests.
- Avoid committing generated files, IDE settings, local binaries, module cache files, and coverage output.
- Keep console output in demos; tests should assert behavior directly.
- When a function panics, either test the panic or explain the invariant.

## Documentation Expectations

- Root README should be short and navigable.
- Deep notes should live under topic READMEs or `docs/`.
- New concepts should be linked from `docs/catalog.md`.
- New learning sequences should be linked from `docs/learning-path.md`.

## Release Checklist

Before publishing a batch of changes:

```bash
go vet ./...
go test ./...
git status --short
```

Then scan:

- Did any generated file slip in?
- Did a public constructor or method need documentation?
- Did a new package appear in the catalog?
- Did tests prove the behavior rather than only printing it?
