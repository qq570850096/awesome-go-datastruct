# Template Method Pattern

Category: Behavioral

Code: [Template.go](../Template.go)

## Intent

Define the skeleton of an algorithm while letting specific steps vary.

## Roles

- Template type defines the algorithm order.
- Concrete types implement variable steps.
- Client uses the common algorithm entrypoint.

## Use When

- Several workflows share the same sequence.
- Only selected steps vary.

## Tradeoffs

- Too many hooks make the algorithm hard to follow.
- Composition can be clearer than inheritance-style templates in Go.

## Test Command

```bash
go test ./DesignPatterns/BehavioralType -run TestTemplate
```
