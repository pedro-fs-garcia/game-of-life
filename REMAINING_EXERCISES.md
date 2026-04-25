# Exercises — Extended Collection

This document contains **exercises** beyond the core Phase 1 exercises.
Each automaton has **explicit, deterministic rules** and forces deliberate practice of Go concepts.

---

## Exercise A4 — Falling Sand (Particle Automaton)

**Domain:** Gravity-based particle simulation using cellular automaton rules.

**Why:** Introduces **asymmetric rules** (gravity direction) and **multiple particle types** with interactions.

---

### Rules

**Grid:** 2D finite grid. Gravity points **downward** (increasing row index).

**Particle Types:**

| Type | Name | Behavior |
|------|------|----------|
| 0 | Empty | Nothing |
| 1 | Sand | Falls down, piles up |
| 2 | Water | Falls down, spreads horizontally |
| 3 | Wall | Immobile, blocks other particles |

**Transition Rules (applied **row by row from bottom to top**):**

**Sand:**
1. If cell below is **Empty** → move down
2. Else if cell **below-left** is Empty → move down-left
3. Else if cell **below-right** is Empty → move down-right
4. Else → stay in place

**Water:**
1. If cell below is **Empty** → move down
2. Else if cell **below-left** is Empty → move down-left
3. Else if cell **below-right** is Empty → move down-right
4. Else if cell **left** is Empty → move left
5. Else if cell **right** is Empty → move right
6. Else → stay in place

**Wall:** Never moves.

**Important:** When multiple particles compete for the same cell, the **topmost** particle wins (processed last).

---

### Implementation Constraints

* Process grid **bottom-to-top** to prevent double-moves
* Use `iota` for particle types
* Track particle movement with a **"moved this frame"** flag or separate pass
* Randomize left/right choice to avoid bias

---

### Test Scenarios

1. **Sand pile:** Drop sand from center, verify 45° slope
2. **Hour glass:** Sand through narrow gap
3. **Water level:** Water fills container, finds level
4. **Dam break:** Wall removed, water flows

---

### Forces You to Learn

* Non-uniform iteration order (bottom-to-top)
* Randomization in deterministic systems
* Multiple particle types with distinct behaviors
* State tracking across update passes

---

## Exercise A5 — 1D Totalistic Automaton

**Domain:** Generalization of elementary cellular automata using **sum-based rules**.

**Why:** Eliminates the need for 8-bit rule encoding. Introduces **totalistic** computation.

---

### Rules

**Grid:** 1D array of cells with states `0`, `1`, or `2` (ternary).

**Neighborhood:** 3 cells (left, self, right).

**Totalistic Rule:**
Instead of considering the exact pattern (L, C, R), compute the **sum** S = L + C + R.

For ternary states, S ∈ {0, 1, 2, 3, 4, 5, 6}.

A **totalistic rule** is encoded as a 7-digit ternary number, where digit i specifies the next state when S = i.

**Example Rule 777 (base 10 = 1001210 base 3):**
| Sum | 0 | 1 | 2 | 3 | 4 | 5 | 6 |
|-----|---|---|---|---|---|---|---|
| Next| 1 | 0 | 0 | 1 | 2 | 0 | 1 |

---

### Implementation Constraints

* Rule must be passed as an integer (0–2186 for ternary 3-cell)
* Decode rule at construction, **not** during each cell update
* Support both **periodic** (wrap) and **fixed** (0-padded) boundaries
* Implement rule decoding as: `func decodeRule(rule int, numSums int) []int`

---

### Interesting Rules

| Rule | Behavior |
|------|----------|
| 777 | Complex patterns |
| 600 | Replicator |
| 1041 | Chaotic |

---

### Forces You to Learn

* Base conversion (decode rule number to array)
* Lookup tables for fast evaluation
* Boundary condition abstraction
* Integer arithmetic precision

---

## Exercise A6 — Turmite (Generalized Langton's Ant)

**Domain:** Multi-state, multi-color extension of Langton's Ant.

**Why:** Forces explicit **state machine design** with **transition tables**.

---

### Rules

**Grid:** 2D infinite grid (simulated with bounds or map).

**Turmite State:**
- Position (x, y)
- Direction (N, E, S, W)
- Internal state (0 to S-1)

**Cell State:** Color (0 to C-1).

**Transition Table:**
A turmite is defined by a table `T[state][color] → (newColor, turn, newState)`

Where:
- `newColor` ∈ {0, ..., C-1}
- `turn` ∈ {N (none), L (left), R (right), U (u-turn)}
- `newState` ∈ {0, ..., S-1}

**Execution (each step):**
1. Read current cell color
2. Look up `(state, color)` in table
3. Write new color to cell
4. Turn according to table
5. Set new internal state
6. Move forward one cell

---

### Implementation Constraints

* Transition table must be a 2D slice: `[][]Transition`
* Direction must use `iota` enum with `TurnLeft()`, `TurnRight()` methods
* Support **sparse grid** (`map[Point]Color`) for potentially infinite expansion
* Validate transition table at construction (all entries defined)

---

### Notable Turmites

**Langton's Ant (1 state, 2 colors):**
```
T[0][0] = (1, R, 0)
T[0][1] = (0, L, 0)
```

**Fibonacci Spiral (2 states, 2 colors):**
```
T[0][0] = (1, R, 1)
T[0][1] = (1, L, 0)
T[1][0] = (1, L, 1)
T[1][1] = (0, N, 0)
```

---

### Forces You to Learn

* Multi-dimensional lookup tables
* State machine as data structure
* Validation of configuration
* Dynamic grid expansion

---


## Exercise B1 — Flood Fill (Multiple Variants)

**Domain:** Grid traversal and region filling.

**Why:** Teaches recursion limits and explicit stack management.

---

### Variants

1. **Recursive DFS** — Simple but stack-limited
2. **Iterative DFS** — Explicit stack, no recursion limit
3. **BFS** — Queue-based, level-order fill

---

### Rules

**Grid:** 2D finite grid of colors/values.

**Operation:** Given a starting cell and target color, fill all connected cells of the same original color with the target color.

**Connectivity:** 4-connected (orthogonal neighbors only) or 8-connected (includes diagonals).

---

### Implementation Constraints

* Implement all three variants
* Use slice-based stack/queue (not channels)
* Handle edge cases: start cell already target color, empty grid
* Benchmark recursive vs iterative on large grids

---

### Forces You to Learn

* Slice-based stacks/queues
* Bounds checking
* Struct reuse (avoid allocations in hot loop)
* Avoiding recursion overflow
* Comparing algorithm characteristics

---

## Exercise B2 — Union–Find (Disjoint Set Union)

**Domain:** Disjoint set data structure for tracking connected components.

**Why:** Pure logic, extremely common in real systems (network connectivity, Kruskal's algorithm, equivalence classes).

---

### Operations

| Operation | Description |
|-----------|-------------|
| `MakeSet(x)` | Create a new set containing only x |
| `Find(x)` | Return the representative of x's set |
| `Union(x, y)` | Merge the sets containing x and y |

---

### Optimizations

1. **Path Compression:** During `Find`, make each node point directly to root
2. **Union by Rank/Size:** Attach smaller tree under larger tree

---

### Implementation Constraints

* Implement **one slice-based version** (elements are integers 0 to N-1)
* Implement **one map-based version** (elements are arbitrary comparable types)
* Both must support path compression
* Track number of disjoint sets

---

### Forces You to Learn

* Slices vs maps for parent tracking
* Method receivers on custom types
* Invariants and mutation safety
* Generic-like patterns (before/after generics)

---

## Exercise B3 — Event Simulation Engine (Discrete Time)

**Domain:** Discrete event simulation engine.

**Why:** Looks simple, becomes very Go-specific very fast. Forces explicit handling of time, state, and event ordering.

---

### Problem

* Events have timestamps
* Events can schedule future events
* Engine processes events in time order
* Events can modify simulation state

---

### Core Types

```go
type Event interface {
    Time() float64
    Execute(sim *Simulation) []Event  // returns new events to schedule
}

type Simulation struct {
    clock  float64
    queue  EventQueue  // priority queue by time
    state  State       // domain-specific state
}
```

---

### Implementation Constraints

* Use `container/heap` for the event queue
* Events are interfaces, not concrete types
* Engine must handle events at the same timestamp deterministically
* No UI — final state is asserted in tests
* Support simulation termination conditions (max time, empty queue, custom predicate)

---

### Example Scenarios

1. **Bank queue:** Customers arrive, wait, get served, leave
2. **Network packets:** Messages sent, delayed, received, acknowledged
3. **Traffic light:** Lights change state on timer, cars arrive/depart

---

### Forces You to Learn

* Priority queues (`container/heap`)
* Interfaces for polymorphic events
* Explicit state transitions
* Error propagation
* Time handling precision

---

## Exercise B4 — A* Pathfinding on a Grid

**Domain:** Graph search algorithm for finding shortest paths.

**Why:** Well-known, algorithmically rich, zero fluff. Forces explicit state management and priority queue usage.

---

### Rules

* Grid with obstacles (passable/impassable cells)
* Start and goal positions
* Find shortest path or return error if none exists

---

### Algorithm

1. Maintain open set (priority queue) and closed set
2. For each cell, track g-score (cost from start) and f-score (g + heuristic)
3. Always expand the cell with lowest f-score
4. Use Manhattan distance as heuristic (admissible for 4-connected grid)

---

### Implementation Constraints

* Use `container/heap` for open set
* Use `map[Point]struct{}` for closed set
* Return the path as `[]Point` or error
* No global variables — everything passed explicitly
* Support diagonal movement as optional mode

---

### Forces You to Learn

* Maps keyed by structs
* Priority queues (again)
* Ownership of state
* Error handling vs sentinel values
* Reconstructing path from parent pointers

---

## Exercise B5 — Immutable vs Mutable Data Experiment

**Domain:** Comparative study of mutation strategies in Go.

**Why:** This is where Go thinking diverges sharply from Python/TypeScript. Understanding when to mutate vs copy is essential.

---

### Task

Implement the **same algorithm** twice:

1. **Mutable version:** Mutate data structures in place
2. **Immutable version:** Return new values, never modify inputs

**Suggested algorithms:**
* Tree insertion/deletion
* Graph transformation
* State machine transitions

---

### Measurements

| Metric | How to Measure |
|--------|----------------|
| Allocations | `go test -benchmem` |
| Performance | `go test -bench` |
| Clarity | Code review, line count |
| Bug surface | Edge cases, nil handling |

---

### Implementation Constraints

* Both versions must pass identical test suites
* Document trade-offs in comments
* Include benchmarks
* This is an **exercise, not a program** — no main function required

---

### Forces You to Learn

* Value semantics vs pointer semantics
* Allocation costs
* Defensive copying
* When immutability helps vs hurts in Go

---

## Exercise B6 — Byte-Level Parser (No Regex)

**Domain:** Low-level parsing of structured text.

**Why:** Go excels here; higher-level languages often hide this. Forces understanding of bytes, runes, and explicit state machines.

---

### Task

Parse a tiny DSL or config format:

* **No regex**
* **No `strings.Split` abuse**
* **No external parsing libraries**

**Suggested formats:**
* INI file (`[section]`, `key=value`)
* Simple CSV (handle quoted fields)
* Custom DSL (e.g., `SET x = 10; PRINT x;`)

---

### Implementation Constraints

* Work directly with `[]byte` or implement a scanner
* Handle UTF-8 correctly (if format allows non-ASCII)
* Return structured errors with line/column numbers
* Distinguish between `byte` (raw) and `rune` (Unicode code point)

---

### Core Types

```go
type Token struct {
    Type   TokenType
    Value  string
    Line   int
    Column int
}

type ParseError struct {
    Message string
    Line    int
    Column  int
}
```

---

### Forces You to Learn

* `byte` vs `rune`
* Explicit state machines
* Precise error returns
* Buffer management
* Position tracking

---

## Implementation Guidelines

For all exercises:

1. **No globals.** All state passed explicitly.
2. **Pure functions** where possible. Side effects isolated to grid mutation.
3. **Double buffering** for synchronous updates.
4. **Separate concerns:** Grid storage, transition logic, neighbor counting.
5. **Test with known patterns.** Verify against expected behavior.

---

## Difficulty Progression

| Exercise               | States | Dimensions | Complexity |
|------------------------|--------|------------|------------|
| 1D Totalistic (A5)     | 3      | 1D         | ★★☆        |
| Flood Fill (B1)        | N      | 2D         | ★★☆        |
| Union–Find (B2)        | —      | —          | ★★☆        |
| Immutable/Mutable (B5) | —      | —          | ★★☆        |
| Falling Sand (A4)      | 4      | 2D         | ★★★        |
| Turmite (A6)           | S×C    | 2D         | ★★★        |
| A* Pathfinding (B4)    | 2      | 2D         | ★★★        |
| Event Simulation (B3)  | —      | —          | ★★★        |
| Byte Parser (B6)       | —      | —          | ★★★        |

---
