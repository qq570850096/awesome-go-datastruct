# Textbook Style Guide

This guide defines the documentation standard for module-level textbook chapters. The goal is to make every major module teach a concept, not only list files.

## Chapter Contract

Every core module `README.md` should use this structure:

```markdown
# Topic Name

One short paragraph that explains why the topic matters.

## Learning Goals
## Prerequisites
## Mental Model
## Diagram
## Terminology
## Repository Code Map
## Core Invariants
## Operation Walkthrough
## Complexity
## Common Mistakes And Edge Cases
## Worked Example
## Practice Exercises
## Test And Benchmark Commands
## Next Topics
```

The headings are intentionally consistent. A learner should be able to move from linked lists to heaps to graphs without relearning the document shape.

## Exercise Levels

Use three exercise levels in each chapter:

| Level | Purpose | Typical task |
| --- | --- | --- |
| Warm-up | Confirm the mental model. | Draw state before and after an operation. |
| Drill | Build testing and implementation confidence. | Add edge-case coverage or rewrite one operation. |
| Challenge | Connect code to analysis. | Explain complexity, add benchmarks, or refactor tests. |

Each exercise should be small enough to complete inside the repository. Avoid tasks that require external articles, online judges, or new services.

## Hint And Answer Format

Embed hints and answers directly in the chapter with collapsible blocks:

```markdown
<details>
<summary>Hint</summary>

Give a small nudge without solving the task.

</details>

<details>
<summary>Reference answer</summary>

Show the expected reasoning, test shape, or implementation outline.

</details>
```

Reference answers should teach the reasoning. They do not need to paste a complete replacement file unless the exercise is specifically about writing code.

## Diagram Rules

- Prefer ASCII diagrams for pointer, array, tree, and queue state.
- Use Mermaid only when relationships or dependency graphs become clearer than text.
- Keep diagrams small enough to fit in a GitHub README without horizontal scrolling.
- Label indexes, roots, heads, colors, or visited state explicitly.

## Complexity Notation

Use a table with `Operation`, `Time`, `Space`, and `Why` columns.

- Use `n` for items in one collection.
- Use `h` for tree height.
- Use `V` and `E` for graph vertices and edges.
- Distinguish worst-case, average-case, and amortized cost when it matters.
- Add a short derivation in the `Why` column instead of listing Big-O alone.

## Linking Expectations

Each chapter should link to:

- Its prerequisites when they are in this repository.
- Its direct next topics.
- The exact test command for the module.
- Related catalog or roadmap pages only when they help the learner choose what to do next.

When adding a new chapter, update:

- [README.md](../README.md)
- [DIRECTORY.md](../DIRECTORY.md)
- [docs/catalog.md](catalog.md)
- [docs/exercise-matrix.md](exercise-matrix.md) when exercises change
- [docs/knowledge-graph.md](knowledge-graph.md) and [docs/data/knowledge-graph.json](data/knowledge-graph.json) when the learning path changes

## Language Standard

All Markdown files, examples, and code comments should be written in English. Keep prose direct and instructional. The repository can still teach beginners, but it should do so with precise engineering vocabulary.
