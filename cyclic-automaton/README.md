## Exercise A3 — Cyclic Cellular Automaton

**Domain:** A 2D grid simulation where each cell holds a numeric "color" that
can only increase by one step at a time, cycling back to zero after reaching
the maximum.

---

### Glossary

| Term              | Definition                                                                                            |
| ----------------- | ----------------------------------------------------------------------------------------------------- |
| **Cell**          | A single unit in the grid, identified by its row and column position                                  |
| **State**         | The current "color" of a cell, represented as an integer                                              |
| **N**             | The total number of distinct states available (e.g. N=4 means states 0, 1, 2, 3)                      |
| **Cycle**         | States form a loop: 0 → 1 → 2 → ... → (N-1) → 0 → ... A cell at the last state wraps back to 0        |
| **Target state**  | The state directly ahead of a cell in the cycle: `(current_state + 1) mod N`                          |
| **Neighbor**      | Any cell within a square ring of a given radius around the cell, excluding the cell itself            |
| **Range**         | The radius of that square ring (Range 1 = 1 cell out in all directions, Range 2 = 2 cells out)        |
| **T (Threshold)** | The minimum number of neighbors that must already be in the target state to allow the cell to advance |
| **Tick**          | One full update step where every cell in the grid is evaluated and updated simultaneously             |

---

### Parameters

| Parameter | Type    | Description                        | Typical Range |
| --------- | ------- | ---------------------------------- | ------------- |
| N         | `uint8` | Number of distinct states (colors) | 3–16          |
| T         | `uint8` | Advancement threshold              | 1–3           |
| Range     | `uint8` | Neighborhood radius                | 1–2           |

---

### Grid

- A finite 2D grid of cells with no wrap-around at the edges.
- Each cell stores its current state as a `uint8` value in the range `[0, N-1]`.
- Boundary cells (edges and corners) have fewer neighbors than interior cells
  because neighbors outside the grid are simply absent — they are not counted.

---

### Neighborhood

The neighbors of a cell are all cells within a square centered on it, at most
`Range` steps away in any direction (horizontally, vertically, or diagonally),
excluding the cell itself.

**Range 1 — Moore neighborhood (up to 8 neighbors):**
```
■ ■ ■
■ · ■
■ ■ ■
```

**Range 2 — Extended Moore neighborhood (up to 24 neighbors):**
```
■ ■ ■ ■ ■
■ ■ ■ ■ ■
■ ■ · ■ ■
■ ■ ■ ■ ■
■ ■ ■ ■ ■
```

A corner cell with Range 1 has only 3 neighbors. A corner cell with Range 2
has only 8. The threshold T is always compared against the actual neighbor
count, not the maximum possible.

---

### Transition Rule (applied to every cell on every tick)

Given a cell with current state `S`:

1. Compute its **target state**: `target = (S + 1) mod N`
2. Count how many of its neighbors currently hold state `target`
3. **If** that count is greater than or equal to T → the cell advances to `target`
4. **Otherwise** → the cell stays at `S`

> All cells are evaluated using the state of the grid at the **start** of the
> tick. No cell's new state affects the evaluation of any other cell in the
> same tick.

**Example** with N=4, T=2, Range=1:

A cell is in state 2. Its target state is `(2+1) mod 4 = 3`. If at least 2
of its (up to 8) neighbors are currently in state 3, it advances to 3.
Otherwise it stays at 2. Neighbors in states 0, 1, or 2 are irrelevant.

---

### Implementation Requirements

- Implement as a `type CyclicAutomaton struct` with all parameters passed at
  construction (N, T, Range), not hardcoded.
- Cell state must be stored as `uint8`.
- The grid update must be **synchronous**: compute the full next-generation
  grid before writing any new state back.
- Neighbor counting must correctly handle grid boundaries (no out-of-bounds
  access, no wrap-around).

---

### Suggested Configurations

| Name      | N   | T   | Range | Expected Behavior                                                                                       |
| --------- | --- | --- | ----- | ------------------------------------------------------------------------------------------------------- |
| Spiral    | 8   | 1   | 1     | Low threshold causes waves to spread fast, collide, and curl into persistent rotating spirals           |
| Turbulent | 14  | 3   | 2     | Many states and high threshold create competing fronts that never stabilize — continuous churn          |
| Slow      | 4   | 2   | 1     | High threshold relative to few states produces large, stable color regions with slow boundary migration |
