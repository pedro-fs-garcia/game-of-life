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

### 4. **Wireworld**

**Why:** Introduces **multiple cell states** (4) and **directional signal propagation**. Simulates digital logic circuits.

**Cell States:**

| State | Name | Description |
|-------|------|-------------|
| 0 | Empty | Blank space, never changes |
| 1 | Conductor | Wire that can carry electrons |
| 2 | Electron Head | Active electron, front of signal |
| 3 | Electron Tail | Decaying electron, back of signal |

**Transition Rules (simultaneous):**

1. Empty → Empty (always)
2. Electron Head → Electron Tail (always)
3. Electron Tail → Conductor (always)
4. Conductor → Electron Head if **exactly 1 or 2** Moore neighbors are Electron Heads
5. Conductor → Conductor otherwise

**Forces you to learn:**

* Multi-state enums with `iota`
* Double buffering for synchronous updates
* Neighbor counting with boundary handling
* Struct composition for grid + state

**Constraint:**
Implement at least one logic gate (OR, AND, or XOR) as a test pattern.

→ See `wireworld/`

---

### 5. **Brian's Brain**

**Why:** 3-state automaton with **chaotic, self-organizing behavior**. Simpler than Wireworld but forces precise state machine logic.

**Cell States:** Off (0), On (1), Dying (2)

**Transition Rules (simultaneous, toroidal grid):**

1. Off → On if **exactly 2** Moore neighbors are On
2. On → Dying (always)
3. Dying → Off (always)

**Forces you to learn:**

* Modular arithmetic for wrap-around indexing (toroidal topology)
* Value semantics vs pointer semantics for grid
* Method receivers on custom types
* Named types: `type State uint8`

**Constraint:**
Grid must support arbitrary dimensions passed at construction.

→ See `briansbrain/`

---

### 6. **Turmite (Generalized Langton's Ant)**

**Why:** Multi-state, multi-color extension of Langton's Ant. Forces explicit **state machine design** with **transition tables**.

**Turmite State:**

* Position (x, y)
* Direction (N, E, S, W)
* Internal state (0 to S-1)

**Cell State:** Color (0 to C-1)

**Transition Table:**
`T[state][color] → (newColor, turn, newState)`

Where `turn` ∈ {None, Left, Right, U-turn}

**Execution (each step):**

1. Read current cell color
2. Look up `(state, color)` in table
3. Write new color to cell
4. Turn according to table
5. Set new internal state
6. Move forward one cell

**Forces you to learn:**

* Multi-dimensional lookup tables
* State machine as data structure
* Validation of configuration at construction
* Sparse grid (`map[Point]Color`) for infinite expansion

**Constraint:**
Transition table must be a 2D slice: `[][]Transition`. Validate all entries at construction.

→ See `turmite/`

---

## Phase 2 — Graphs, Maps, and Explicit Algorithms

### 7. **Flood Fill (Multiple Variants)**

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

### 8. **Union–Find (Disjoint Set Union)**

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

### 9. **Event Simulation Engine (Discrete Time)**

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

### 10. **A* Pathfinding on a Grid**

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

### 11. **Immutable vs Mutable Data Experiment**

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

### 12. **Byte-Level Parser (No Regex)**

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
