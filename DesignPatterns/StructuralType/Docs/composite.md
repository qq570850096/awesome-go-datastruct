# Composite Pattern

Category: Structural

Code: [Component.go](../Component.go)

## Intent

Treat individual objects and object groups through the same interface.

## Roles

- Component defines shared behavior.
- Leaf has no children.
- Composite stores children and delegates work.

## Use When

- The domain is naturally tree-shaped.
- Callers should not care whether they have one item or a group.

## Tradeoffs

- The common interface may become too broad.
- Tree mutation rules need tests.

## Test Command

```bash
go test ./DesignPatterns/StructuralType -run TestComponent
```
