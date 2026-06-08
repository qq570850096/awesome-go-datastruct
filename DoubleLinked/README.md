# Doubly Linked List And Caches

Doubly linked lists add backward links. That extra pointer makes deletion and movement efficient, which is why cache implementations often combine a map with a doubly linked list.

## Learning Goals

- Maintain `prev` and `next` together.
- Explain LRU, LFU, and FIFO eviction.
- Test cache hit, miss, and eviction order.

## Prerequisites

- Singly linked lists.
- Maps and pointer receivers.

## Mental Model

A cache needs two synchronized views: a map finds entries quickly, and a list records order. Correct code keeps both views consistent.

## Diagram

```text
map[key] -> node
head <-> [key,value] <-> [key,value] <-> tail
```

## Terminology

| Term | Meaning |
| --- | --- |
| LRU | Evicts the least recently used item. |
| LFU | Evicts the least frequently used item. |
| FIFO | Evicts the oldest inserted item. |
| Sentinel | A dummy head or tail that simplifies edge cases. |

## Repository Code Map

| File or directory | What to read for |
| --- | --- |
| List.go | Doubly linked mechanics. |
| LRU.go, LFU.go, FIFO.go | Cache policies. |
| *_test.go | Policy tests. |

## Core Invariants

- For every adjacent pair, forward and backward links agree.
- The map points only to nodes currently in the list.
- Eviction removes from both map and list.

## Operation Walkthrough

On an LRU hit, the map finds the node in O(1), then the list moves it to the most-recent position. On a miss at full capacity, the least-recent node is removed from the list and deleted from the map before insertion.

## Complexity

| Operation | Time | Space | Why |
| --- | --- | --- | --- |
| LRU Get/Put | O(1) | O(1) | Map lookup plus constant list movement. |
| FIFO Put | O(1) | O(1) | Oldest node is known from list order. |
| LFU operations | Average O(1) | O(n) | Frequency buckets avoid full scans in the intended model. |

## Common Mistakes And Edge Cases

- Updating `next` but not `prev`.
- Leaving evicted keys in the map.
- Forgetting to refresh recency on read.

## Worked Example

Capacity 2: `Put(1)`, `Put(2)`, `Get(1)`, `Put(3)` evicts key 2 in LRU because key 1 was refreshed.

## Practice Exercises

1. **Draw the state** (Warm-up)

   Use the diagram notation to trace LRU cache order on a small input.

<details>
<summary>Hint</summary>

Write the state before the operation, after each important assignment or loop step, and after the invariant is restored.

</details>

<details>
<summary>Reference answer</summary>

A complete answer shows the same data before and after the operation, names the changed pointer, index, color, mark, or collection, and ends with the invariant visibly true.

</details>

2. **Add edge-case coverage** (Drill)

   Add or inspect tests for capacity one, repeated hits, misses, and eviction ties.

<details>
<summary>Hint</summary>

Prefer table-driven tests. Give each case a name that explains the behavior under test.

</details>

<details>
<summary>Reference answer</summary>

The reference shape is a table with empty input, one minimal valid input, a normal case, and a failure or missing-value case when the module supports one.

</details>

3. **Explain the complexity** (Challenge)

   Explain why LRU Get/Put has the complexity shown in the table.

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
go test ./DoubleLinked
```

## Next Topics

- [Heap](../Heap/)
- [DesignPatterns](../DesignPatterns/) later for policy design.
