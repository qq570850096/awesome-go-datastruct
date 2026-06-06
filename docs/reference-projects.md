# Reference Projects

This repository borrows ideas from mature learning repositories without copying their structure wholesale.

## Go by Example

Repository: <https://github.com/mmcgrana/gobyexample>

What to learn:

- Small runnable examples.
- Short explanations close to code.
- A reader can jump into one topic without reading the whole project.

Applied here:

- Root README points to small package-level test commands.
- Topic pages should stay focused and runnable.

## Rustlings

Repository: <https://github.com/rust-lang/rustlings>

What to learn:

- Ordered exercises.
- A visible learner journey.
- The project feels like a workshop, not only a code dump.

Applied here:

- `docs/learning-path.md` defines study passes.
- Roadmap includes a future exercise mode.

## JavaScript Algorithms

Repository: <https://github.com/trekhleb/javascript-algorithms>

What to learn:

- Topic catalog with data structures and algorithms.
- Each topic has local explanation and references.
- Complexity is part of the learning surface.

Applied here:

- `docs/catalog.md` summarizes topic, path, core idea, and complexity.
- `CONTRIBUTING.md` asks new topics to include invariant and complexity notes.

## The Algorithms - Go

Repository: <https://github.com/TheAlgorithms/Go>

What to learn:

- Go packages are tested as the core quality gate.
- Algorithms are discoverable as packages.
- Automated checks protect the repository from broken examples.

Applied here:

- The required local quality gate is `go vet ./...` plus `go test ./...`.
- New tests cover previously untested core packages.
