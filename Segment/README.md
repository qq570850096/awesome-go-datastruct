# Segment Tree

A segment tree stores aggregate information over ranges so queries and updates avoid scanning the whole array.

## Learning Goals

- Build a tree over array intervals.
- Use a merger function to define the aggregate.
- Explain range query and point update complexity.

## Prerequisites

- Arrays, recursion, and tree intervals.

## Mental Model

Each node owns an interval. A parent answer is produced by merging the answers of its two child intervals.

## Diagram

```text
[0,3]
 /   \
[0,1] [2,3]
/  \  /  \
0  1  2  3
```

## Terminology

| Term | Meaning |
| --- | --- |
| Interval | A contiguous index range. |
| Merger | Function combining two child values. |
| Query | Read the aggregate for a range. |
| Point update | Change one index and repair ancestors. |

## Repository Code Map

| File or directory | What to read for |
| --- | --- |
| Tree.go | Build, query, and update operations. |
| SegmentTree_test.go | Range-query tests. |

## Core Invariants

- Every internal node stores the merge of its children.
- Leaf nodes equal source array values.
- Queries split only when the requested range crosses a midpoint.

## Operation Walkthrough

Build recursively until each leaf represents one index. Query returns a node directly when its interval exactly matches the requested range; otherwise it splits and merges partial answers.

## Complexity

| Operation | Time | Space | Why |
| --- | --- | --- | --- |
| Build | O(n) | O(n) | Initializes tree nodes. |
| Range query | O(log n) for typical balanced splits | O(log n) | Visits relevant interval nodes. |
| Point update | O(log n) | O(log n) | Repairs one root-to-leaf path. |

## Common Mistakes And Edge Cases

- Using the wrong midpoint boundary.
- Returning a non-overlapping child interval.
- Forgetting to update ancestors.

## Worked Example

For array `[1,3,5,7]`, a sum segment tree stores `[0,3]=16`, `[0,1]=4`, and `[2,3]=12`.

## Practice Exercises

1. **Draw the state** (Warm-up)

   Use the diagram notation to trace range query split on a small input.

<details>
<summary>Hint</summary>

Write the state before the operation, after each important assignment or loop step, and after the invariant is restored.

</details>

<details>
<summary>Reference answer</summary>

A complete answer shows the same data before and after the operation, names the changed pointer, index, color, mark, or collection, and ends with the invariant visibly true.

</details>

2. **Add edge-case coverage** (Drill)

   Add or inspect tests for single-element arrays, full-range queries, left-only and right-only ranges.

<details>
<summary>Hint</summary>

Prefer table-driven tests. Give each case a name that explains the behavior under test.

</details>

<details>
<summary>Reference answer</summary>

The reference shape is a table with empty input, one minimal valid input, a normal case, and a failure or missing-value case when the module supports one.

</details>

3. **Explain the complexity** (Challenge)

   Explain why Range query has the complexity shown in the table.

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
go test ./Segment
```

## Next Topics

- [BinarySearch](../BinarySearch/)
- [Sorts](../Sorts/)
