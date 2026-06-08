# Sorting Algorithms

Sorting makes algorithmic tradeoffs visible: comparisons, swaps, recursion, memory, stability, and input shape all matter.

## Learning Goals

- Compare simple, divide-and-conquer, and non-comparison sorts.
- Test ordering and element preservation.
- Reason about best, average, and worst cases.

## Prerequisites

- Arrays, loops, recursion, and comparisons.

## Mental Model

Sorting transforms a sequence into nondecreasing order without losing or inventing elements. Different algorithms pay different costs to discover where each element belongs.

## Diagram

```text
input:  [5, 1, 4, 1]
output: [1, 1, 4, 5]
required: ordered + same multiset
```

## Terminology

| Term | Meaning |
| --- | --- |
| Stable | Equal elements keep original relative order. |
| In-place | Uses constant or small extra storage. |
| Partition | Split data around a pivot. |
| Merge | Combine sorted subsequences. |

## Repository Code Map

| File or directory | What to read for |
| --- | --- |
| Bubble.go, Insert.go, Select.go, Shell.go | Simple comparison sorts. |
| Merge.go, Quick.go | Divide-and-conquer sorts. |
| Bucket.go | Range-aware non-comparison sorting. |
| Sort_test.go | Shared correctness checks. |

## Core Invariants

- The result is sorted.
- The result contains the same multiset as input.
- Recursive sorts reduce problem size.

## Operation Walkthrough

Start with insertion sort to understand local movement. Then read merge sort for divide-and-conquer with extra memory. Finish with quicksort and focus on partition correctness.

## Complexity

| Operation | Time | Space | Why |
| --- | --- | --- | --- |
| Bubble/insertion/selection | O(n^2) | O(1) | Nested scans dominate. |
| Merge sort | O(n log n) | O(n) | log n levels with n merge work each. |
| Quick sort | Average O(n log n), worst O(n^2) | O(log n) | Partition quality controls recursion depth. |
| Bucket/counting style | O(n + k) | O(k) | Depends on value range k. |

## Common Mistakes And Edge Cases

- Testing only random input.
- Forgetting duplicates and already-sorted input.
- Claiming stability without checking equal elements.
- Ignoring worst-case quicksort partitions.

## Worked Example

Sorting `[5,1,4,1]` must produce `[1,1,4,5]`. A test should verify both sorted order and that two `1` values remain.

## Practice Exercises

1. **Draw the state** (Warm-up)

   Use the diagram notation to trace partition or merge on a small input.

<details>
<summary>Hint</summary>

Write the state before the operation, after each important assignment or loop step, and after the invariant is restored.

</details>

<details>
<summary>Reference answer</summary>

A complete answer shows the same data before and after the operation, names the changed pointer, index, color, mark, or collection, and ends with the invariant visibly true.

</details>

2. **Add edge-case coverage** (Drill)

   Add or inspect tests for empty array, one value, duplicates, sorted input, reverse input.

<details>
<summary>Hint</summary>

Prefer table-driven tests. Give each case a name that explains the behavior under test.

</details>

<details>
<summary>Reference answer</summary>

The reference shape is a table with empty input, one minimal valid input, a normal case, and a failure or missing-value case when the module supports one.

</details>

3. **Explain the complexity** (Challenge)

   Explain why sorting has the complexity shown in the table.

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
go test ./Sorts
go test ./Sorts -bench=.
```

## Next Topics

- [Heap](../Heap/)
- [Graph_algo](../Graph_algo/)
