# Langton's Ant

Go implementation of **Langton's Ant**, a well-known cellular automaton with a stateful agent moving on a grid.

## Rules

- The grid is conceptually infinite (simulated using wrap-around logic).
- The ant has a **position** and a **direction**.
- At each step, the ant observes the **current cell**.
- On a white cell:
  - turn right
  - flip the **current cell** to black
- On a black cell:
  - turn left
  - flip the **current cell** to white
- After turning and flipping, the ant moves forward one cell.
