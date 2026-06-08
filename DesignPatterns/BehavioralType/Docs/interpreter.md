# Interpreter Pattern

Category: Behavioral

Code: [Interpreter.go](../Interpreter.go)

## Intent

Represent grammar rules as objects that can evaluate an expression.

## Roles

- Expression defines evaluation.
- Terminal expression represents a value.
- Nonterminal expression combines expressions.
- Context provides runtime values.

## Use When

- A small language or expression format is stable.
- Grammar rules are simple enough to model directly.

## Tradeoffs

- Complex grammars need parsers instead.
- Class/type count can grow quickly.

## Test Command

```bash
go test ./DesignPatterns/BehavioralType -run TestInterpreter
```
