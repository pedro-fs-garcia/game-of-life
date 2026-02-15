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
├── wireworld/
├── briansbrain/
├── turmite/
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
make run EXERCISE=cellular-automata
make run EXERCISE=langtons-ant
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

**Focus:** stateful agent movement, direction handling, wrap-around logic
**Key concepts:** agent state (position + direction), cell toggling, toroidal grid

→ See `langtons-ant/`

---

### 3. **Elementary Cellular Automata (Rule 30 / 90 / 110)**

**Focus:** 1D cellular automata, Wolfram rule encoding, pattern-to-index conversion
**Key concepts:** indexing discipline, bitwise operations, wrap-around boundaries
→ See `cellular-automata/`

---

### 4. **Wireworld**

**Focus:** Introduces **multiple cell states** (4) and **directional signal propagation**. Simulates digital logic circuits.
**Key concepts:** multiple cell states (4); directional signal propagation; Simulates digital logic circuits.
→ See `wireworld/`

---

### 5. **Brian's Brain**

**Focus:** 3-state automaton with chaotic, self-organizing behavior
**Key concepts:** modular arithmetic, toroidal topology, named types (`type State uint8`)
→ See `briansbrain/`

---

### 6. **Turmite (Generalized Langton's Ant)**

**Focus:** multi-state, multi-color agent with transition tables
**Key concepts:** multi-dimensional lookup tables, state machine design, sparse grid (`map[Point]Color`)
→ See `turmite/`

---

## Phase 2 — Graphs, Maps, and Explicit Algorithms

### 7. **Flood Fill (Multiple Variants)**

**Focus:** recursion limits, explicit stack management
**Key concepts:** slice-based stacks/queues, bounds checking, avoiding recursion overflow
→ See `flood-fill/`

---

### 8. **Union–Find (Disjoint Set Union)**

**Focus:** disjoint set data structure, path compression
**Key concepts:** slices vs maps for parent tracking, method receivers, invariants
→ See `union-find/`

---

## Phase 3 — Algorithms That Punish Loose Typing

### 9. **Event Simulation Engine (Discrete Time)**

**Focus:** discrete event simulation, time-ordered processing
**Key concepts:** priority queues (`container/heap`), interfaces, error propagation
→ See `event-simulation/`

---

### 10. **A* Pathfinding on a Grid**

**Focus:** pathfinding algorithm, heuristic search
**Key concepts:** maps keyed by structs, priority queues, error handling vs sentinel values
→ See `astar/`

---

## Phase 4 — Go-Specific Muscle Memory

### 11. **Immutable vs Mutable Data Experiment**

**Focus:** comparing mutation vs value semantics
**Key concepts:** allocations, clarity vs performance, Go idioms
→ See `immutability/`

---

### 12. **Byte-Level Parser (No Regex)**

**Focus:** low-level parsing without regex or string splitting
**Key concepts:** `byte` vs `rune`, explicit state machines, precise error returns
→ See `byte-parser/`

---

## Philosophy

* **Thinking in state**
* **Making data explicit**
* **Paying the cost of correctness upfront**
* **Learning where Go pushes back**

---
