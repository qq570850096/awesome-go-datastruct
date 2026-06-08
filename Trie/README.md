# Trie

A Trie stores strings by sharing prefixes. It is a tree optimized around characters rather than numeric ordering.

## Learning Goals

- Represent words as paths.
- Distinguish full words from prefixes.
- Analyze lookup by key length.

## Prerequisites

- Strings, maps, and tree traversal.

## Mental Model

Each edge consumes one character. A node can represent a prefix even when no full word ends there.

## Diagram

```text
root
 ├─ c ─ a ─ t*
 │       └─ r*
 └─ d ─ o ─ g*
* marks a complete word
```

## Terminology

| Term | Meaning |
| --- | --- |
| Prefix | Characters consumed from root to a node. |
| Terminal marker | Flag marking a full stored word. |
| Alphabet | Possible next characters. |

## Repository Code Map

| File or directory | What to read for |
| --- | --- |
| Trie.go | Insert, search, and prefix behavior. |
| Trie_test.go | Word and corpus checks. |
| pride-and-prejudice.txt | Large text input for practice. |

## Core Invariants

- A path from root spells a prefix.
- A full word requires a terminal marker.
- Children maps contain valid next characters.

## Operation Walkthrough

Insert walks character by character, creating missing child nodes. Search walks the same path but succeeds only if the final node is terminal.

## Complexity

| Operation | Time | Space | Why |
| --- | --- | --- | --- |
| Insert | O(m) | O(m) | Processes each character. |
| Search | O(m) | O(1) | Follows one edge per character. |
| Prefix check | O(p) | O(1) | Stops after prefix length. |

## Common Mistakes And Edge Cases

- Treating every prefix as a full word.
- Ignoring byte versus rune choices.
- Forgetting empty-string behavior.

## Worked Example

After inserting `cat`, searching `ca` should fail as a word but succeed as a prefix.

## Practice Exercises

1. **Draw the state** (Warm-up)

   Use the diagram notation to trace insert and search on a small input.

<details>
<summary>Hint</summary>

Write the state before the operation, after each important assignment or loop step, and after the invariant is restored.

</details>

<details>
<summary>Reference answer</summary>

A complete answer shows the same data before and after the operation, names the changed pointer, index, color, mark, or collection, and ends with the invariant visibly true.

</details>

2. **Add edge-case coverage** (Drill)

   Add or inspect tests for empty string, prefix-only input, missing word, and repeated insertion.

<details>
<summary>Hint</summary>

Prefer table-driven tests. Give each case a name that explains the behavior under test.

</details>

<details>
<summary>Reference answer</summary>

The reference shape is a table with empty input, one minimal valid input, a normal case, and a failure or missing-value case when the module supports one.

</details>

3. **Explain the complexity** (Challenge)

   Explain why Search has the complexity shown in the table.

<details>
<summary>Hint</summary>

Count the number of nodes, array cells, characters, or edges that can be visited. Then count extra storage.

</details>

<details>
<summary>Reference answer</summary>

A good answer separates input size from auxiliary state. It mentions whether the operation follows one path, scans all elements, visits all edges, or allocates a helper structure.

</details>

## Test And Benchmark Commands

```bash
go test ./Trie
```

## Next Topics

- [skiplists](../skiplists/)
- [Graph_algo](../Graph_algo/)
