# Knowledge Graph

This graph connects Go language mechanics, data structures, algorithms, and engineering skills. Edge meanings:

- `prerequisite`: learn this first.
- `practice_with`: use this module to practice the earlier concept.
- `deepens`: move from concept knowledge toward engineering judgment.

Structured data is available in [docs/data/knowledge-graph.json](data/knowledge-graph.json).

The core learning-spine nodes link to textbook-style module chapters. Those chapters contain the local mental model, invariants, complexity derivation, exercises, hints, reference answers, and test commands for the topic.

```mermaid
flowchart LR
  subgraph Beginner["Beginner"]
    go_basics["Go basics\nBasicGo/basics"]
    go_testing["Testing\nBasicGo/testingdemo"]
    go_pointers["Pointers\nBasicGo/pointers"]
    go_structs["Structs and methods\nBasicGo/structs"]
    go_interfaces["Interfaces\nBasicGo/interface"]
    go_errors["Errors and defer\nBasicGo/errors"]
  end

  subgraph Intermediate["Intermediate"]
    go_generics["Generics\nBasicGo/generics"]
    go_reflect["Reflection\nBasicGo/reflect"]
    conc_goroutine["Goroutines\nBasicGo/GoRoutine"]
    conc_channel["Channels/select\nBasicGo/channelselect"]
    conc_context["Context\nBasicGo/context"]
    conc_shared["Shared variables\nBasicGo/sharedvars"]
  end

  subgraph Advanced["Advanced"]
    ds_linear["Linear structures\nLinked stack queue"]
    ds_cache["Caches\nDoubleLinked"]
    ds_heap["Heap\nHeap"]
    ds_trees["Search trees\nBinarySearch AVL Red-Black"]
    ds_range["Range structures\nSegment"]
    ds_trie["Trie\nTrie"]
    ds_skiplist["SkipList\nskiplists"]
    ds_union["Union-find\nUnion"]
    algo_sort["Sorting\nSorts"]
    graph_repr["Graph representation\nGraph_algo/Adj"]
    graph_traverse["Graph traversal\nBFS DFS"]
    graph_search["Graph search\nGraph_algo/search"]
  end

  subgraph Expert["Expert"]
    eng_patterns["Design patterns\nDesignPatterns"]
    eng_web["Web engineering\nwebdemo"]
    eng_os["Systems practice\nOSExam"]
    eng_tooling["Tooling\nvet test race"]
    eng_performance["Performance\nbenchmark pprof"]
    eng_api["API boundaries\nCONTRIBUTING"]
  end

  go_basics -- prerequisite --> go_testing
  go_basics -- prerequisite --> go_pointers
  go_basics -- prerequisite --> go_structs
  go_structs -- prerequisite --> go_interfaces
  go_interfaces -- prerequisite --> go_errors
  go_testing -- practice_with --> ds_linear
  go_pointers -- practice_with --> ds_linear
  ds_linear -- practice_with --> ds_cache
  ds_linear -- prerequisite --> ds_heap
  ds_heap -- prerequisite --> algo_sort
  go_structs -- prerequisite --> ds_trees
  ds_trees -- practice_with --> ds_range
  ds_trees -- practice_with --> ds_trie
  ds_trees -- practice_with --> ds_skiplist
  ds_union -- practice_with --> graph_search
  algo_sort -- deepens --> eng_performance
  go_basics -- prerequisite --> conc_goroutine
  conc_goroutine -- prerequisite --> conc_channel
  conc_channel -- prerequisite --> conc_context
  conc_context -- deepens --> eng_web
  conc_shared -- deepens --> eng_tooling
  go_generics -- deepens --> eng_api
  go_reflect -- deepens --> eng_api
  graph_repr -- prerequisite --> graph_traverse
  graph_traverse -- prerequisite --> graph_search
  graph_search -- deepens --> eng_performance
  go_interfaces -- deepens --> eng_patterns
  eng_patterns -- deepens --> eng_api
  eng_web -- practice_with --> eng_api
  eng_tooling -- deepens --> eng_performance
  eng_os -- practice_with --> eng_tooling

  click go_basics "../BasicGo/basics/" "Go basics"
  click go_testing "../BasicGo/testingdemo/" "Testing"
  click go_pointers "../BasicGo/pointers/" "Pointers"
  click go_structs "../BasicGo/structs/" "Structs"
  click go_interfaces "../BasicGo/interface/" "Interfaces"
  click go_errors "../BasicGo/errors/" "Errors"
  click conc_context "../BasicGo/context/" "Context"
  click ds_linear "../Linked/" "Linear structures"
  click ds_heap "../Heap/" "Heap"
  click ds_trees "../BinarySearch/" "Search trees"
  click graph_search "../Graph_algo/search/" "Graph search"
  click eng_patterns "../DesignPatterns/" "Design patterns"
  click eng_web "../webdemo/" "Web demos"
```

## Node Index

| Level | Nodes | Repository paths |
| --- | --- | --- |
| Beginner | Basics, testing, pointers, structs, interfaces, errors | [BasicGo](../BasicGo/) |
| Intermediate | Goroutines, channels, context, shared variables, generics, reflection | [BasicGo](../BasicGo/) |
| Advanced | Lists, queues, heap, trees, graphs, sorting, Trie, SkipList, union-find | [Linked](../Linked/), [DoubleLinked](../DoubleLinked/), [Heap](../Heap/), [Union](../Union/), [Graph_algo](../Graph_algo/), [Sorts](../Sorts/) |
| Expert | Patterns, Web demos, systems practice, tooling, performance, API boundaries | [DesignPatterns](../DesignPatterns/), [webdemo](../webdemo/), [OSExam](../OSExam/), [CONTRIBUTING.md](../CONTRIBUTING.md) |

## Maintenance Rule

When a new module is added, update this Mermaid graph and [knowledge-graph.json](data/knowledge-graph.json) in the same change.
