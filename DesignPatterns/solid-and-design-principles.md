# SOLID And Design Principles

Design patterns are easier to use well when the underlying design principles are clear.

| Principle | Meaning | Practical check |
| --- | --- | --- |
| Single Responsibility | A unit should have one reason to change. | Can this type be described with one clear job? |
| Open Closed | Extend behavior without modifying stable code. | Can new behavior be added through a new type or function? |
| Liskov Substitution | Subtypes should be usable through the parent contract. | Does an implementation surprise callers of the interface? |
| Interface Segregation | Prefer small focused interfaces. | Are callers forced to depend on methods they do not use? |
| Dependency Inversion | Depend on abstractions rather than concrete details. | Can tests or alternative implementations replace dependencies? |
| Least Knowledge | Talk only to close collaborators. | Does a method reach through too many objects? |
| Composition Reuse | Prefer composition before inheritance-style reuse. | Can behavior be built by combining smaller parts? |

## Repository Practice

- Use interfaces when they create a stable boundary.
- Avoid adding a pattern when a simple function is clearer.
- Keep examples small enough that the responsibility boundary is visible.
- Use tests to show the pattern's behavior, not just its structure.
