# Go Exercises — Algorithms, State, and Systems Thinking

This repository is a **collection of independent Go exercises**, each focused on a **well-defined problem domain** and designed to **force deliberate practice** of core Go concepts.

These are not toy examples, demos or MVPs.
Each exercise emphasizes **explicit data modeling**, **algorithmic rigor**, and **idiomatic Go**, with increasing friction as complexity grows.

Each exercise lives in its **own directory**, has its **own `go.mod`**, and can be built and run in isolation.

---

## Design Principles

Every exercise in this repository follows the same constraints:

* ✅ **Known problem domain** (e.g. Game of Life, A*, Union–Find)
* ✅ **Explicit, rule-driven behavior**
* ✅ **Algorithmic pressure** (time, space, or correctness)
* ✅ **No UI**
* ✅ **No I/O beyond basic printing or assertions**
* ✅ **Focus on reasoning, not frameworks**

The goal is to build **Go muscle memory**, not demos.

---

## Repository Structure

```
.
├── bin/                # Compiled binaries for exercises
├── gameoflife/
├── langtons-ant/
├── cellular-automata/
├── flood-fill/
├── union-find/
├── event-simulation/
├── astar/
├── immutability/
├── byte-parser/
├── Taskfile.yml        # Cross-platform automation
└── Makefile            # Optional Unix-based automation
```

* Each directory is a **self-contained Go module**
* No shared `go.mod` at the root
* Tooling (`Taskfile`, `Makefile`) is **optional and external**

---

## Tooling

### Taskfile (recommended, cross-platform)

From the repository root:

```bash
task <exercise>:run
task <exercise>:build
task <exercise>:tidy
```

Example:

```bash
task gameoflife:run
```

### Makefile (Unix-like systems)

```bash
make run EXERCISE=gameoflife
```

> Tooling is **for developer convenience only**.
> Each exercise can be run directly using `go run` inside its directory.

---

## Exercise Roadmap

The exercises are intentionally ordered by **conceptual and Go-specific difficulty**.

---

## Phase 1 — Deterministic State & Explicit Data Modeling

### 1. **Game of Life**

**Focus:** grid modeling, rules, iteration, termination conditions
**Key concepts:** slices vs maps, state transitions, invariants
→ See `gameoflife/`

---

### 2. **Langton’s Ant**

**Why:** Same cellular-automaton family as Game of Life, but introduces **stateful agents**.

**Rules:**

* Infinite grid (simulated with bounds)
* Ant has position + direction
* White → turn right, flip to black
* Black → turn left, flip to white
* Move forward

**Forces you to learn:**

* `struct` modeling (`Ant`, `Direction`)
* enums via `iota`
* `map` vs `[][]bool`
* mutability & value semantics

**Constraint:**
Implement **two versions**:

1. Dense grid (`[][]bool`)
2. Sparse grid (`map[Point]struct{}`)

→ See `langtons-ant/`

---

### 3. **Elementary Cellular Automata (Rule 30 / 90 / 110)**

**Why:** Simple rules, brutal indexing discipline.

**Rules:**

* 1D array of cells
* Next state depends on `(left, self, right)`
* Rule encoded as a number (e.g. 30)

**Forces you to learn:**

* bitwise operations
* slice copying vs reuse
* avoiding aliasing
* mutation vs purity

**Stretch goal:**
Generalize to arbitrary rule numbers without `if` chains.

---

## Phase 2 — Graphs, Maps, and Explicit Algorithms

### 4. **Flood Fill (Multiple Variants)**

**Why:** Teaches recursion limits and explicit stack management.

**Variants:**

1. Recursive DFS
2. Iterative DFS (explicit stack)
3. BFS (queue)

**Forces you to learn:**

* slice-based stacks/queues
* bounds checking
* struct reuse
* avoiding recursion overflow

---

### 5. **Union–Find (Disjoint Set Union)**

**Why:** Pure logic, extremely common in real systems.

**Operations:**

* `Find(x)`
* `Union(x, y)`
* Path compression

**Forces you to learn:**

* slices vs maps for parent tracking
* method receivers
* invariants and mutation safety

**Constraint:**
Implement **one slice-based** and **one map-based** version.

---

## Phase 3 — Algorithms That Punish Loose Typing

### 6. **Event Simulation Engine (Discrete Time)**

**Why:** Looks simple, becomes very Go-specific very fast.

**Problem:**

* Events have timestamps
* Events can schedule future events
* Engine processes events in time order

**Forces you to learn:**

* priority queues (`container/heap`)
* interfaces
* explicit state transitions
* error propagation

**No UI.** Final state is asserted.

---

### 7. **A* Pathfinding on a Grid**

**Why:** Well-known, algorithmically rich, zero fluff.

**Rules:**

* Grid with obstacles
* Manhattan distance heuristic
* Return path or error

**Forces you to learn:**

* maps keyed by structs
* priority queues (again)
* ownership of state
* error handling vs sentinel values

**Constraint:**
No global variables. Everything passed explicitly.

---

## Phase 4 — Go-Specific Muscle Memory

### 8. **Immutable vs Mutable Data Experiment**

**Exercise, not a program.**

Implement the same algorithm twice:

* once mutating in place
* once returning new values

Measure:

* allocations
* clarity
* bug surface

This is where Go thinking diverges sharply from Python/TypeScript.

---

### 9. **Byte-Level Parser (No Regex)**

**Why:** Go excels here; higher-level languages often hide this.

**Task:**

* Parse a tiny DSL or config format
* No regex
* No `strings.Split` abuse

**Forces you to learn:**

* `byte` vs `rune`
* explicit state machines
* precise error returns

---

## Philosophy

* **Thinking in state**
* **Making data explicit**
* **Paying the cost of correctness upfront**
* **Learning where Go pushes back**

---
