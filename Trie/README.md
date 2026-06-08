# Trie

## Goal

Learn prefix-tree storage for strings and prefix queries.

## Prerequisites

Maps, strings, and recursion or iteration.

## Core Invariant

Characters on a path represent a prefix, and terminal markers distinguish full words from prefixes.

## Complexity

Insert/search/prefix query: Time O(m), where m is key length. Space depends on total stored characters.

## Practice Tasks

- Add tests for prefix-only input.
- Add missing-word tests.
- Compare Trie lookup with map lookup.

## Test Command

```bash
go test ./Trie
```

## Related Topics

- [skiplists](../skiplists/)
- [Graph_algo](../Graph_algo/)
