# Flood Fill Algorithm (Multiple Variants) 

Go implementations of the flood fill algorithm using recursive DFS, iterative DFS, and BFS.

**Domain:** 2D grid traversal and connected-region replacement.

## Problem Definition

Given:

* A finite 2D grid `[][]T` (rectangular).
* A starting coordinate `(row, col)`.
* A `targetColor`.

Replace every cell in the same connected component as `(row, col)` that has the **original color** with `targetColor`.

## Formal Rules

### 1. Grid Constraints

* The grid must be rectangular: all rows have equal length.
* The grid may be empty (`len(grid) == 0`).
* Coordinates are zero-based.
* A coordinate is valid if:

  ```
  0 <= row < len(grid)
  0 <= col < len(grid[0])
  ```

### 2. Start Conditions

* If the grid is empty → no operation.
* If `(row, col)` is out of bounds → no operation.
* Let `originalColor = grid[row][col]`.
* If `originalColor == targetColor` → no operation (must exit immediately).

### 3. Connectivity

Connectivity must be explicitly selectable:

* **4-connected**: up, down, left, right
* **8-connected**: includes diagonals

The chosen connectivity must be applied consistently during traversal.

### 4. Fill Rule

A cell `(r, c)` is eligible for replacement if and only if:

* It is within bounds.
* Its value equals `originalColor`.
* It has not already been processed.

Eligible cells must be changed to `targetColor`.

## Expected Behavior

* The algorithm must only modify the connected component that contains the start cell.
* Cells with different colors must never be modified.
* The traversal must terminate for any finite grid.
* Time complexity must be `O(N)`, where `N` is the number of cells in the grid.
* Space complexity:

  * Recursive DFS: up to `O(N)` implicit stack.
  * Iterative DFS: up to `O(N)` explicit stack.
  * BFS: up to `O(N)` queue.

## Variants to Implement

### 1. Recursive DFS

* Use direct recursion.
* Must rely on the call stack.
* No global variables.
* Must short-circuit immediately when encountering:

  * Out-of-bounds cells
  * Cells not equal to `originalColor`

Limitation:

* May overflow stack on large grids or large connected components.

### 2. Iterative DFS

* Use a slice as an explicit stack:

  ```go
  stack := make([]Point, 0)
  ```
* Push start cell.
* Pop from the end (LIFO).
* Must not use recursion.
* Must not use channels.

Required behavior:

* Identical output to recursive DFS.

### 3. BFS

* Use a slice as a queue.
* Implement queue manually (head index or reslicing).
* FIFO order.
* Must not use channels.

Required behavior:

* Identical filled region to DFS variants.
* Order of traversal does not matter, only final grid state.

## Edge Cases (Must Handle)

1. Empty grid
2. Grid with one cell
3. Start cell already equals target color
4. Entire grid is one color
5. Very large connected region
6. Diagonal-only connection when using 4-connectivity (must not fill)
7. Diagonal connection when using 8-connectivity (must fill)

## Implementation Constraints

* Implement all three variants.
* Use only slice-based stack/queue (no channels, no container/list).
* No global mutable state.
* Do not allocate auxiliary 2D visited grid unless strictly necessary.

  * Prefer marking by changing color.
* Benchmark recursive vs iterative on large grids and report:

  * Execution time
  * Memory usage
  * Stack overflow behavior (if any)

## Validation Criteria

An implementation is correct if:

* All three variants produce identical final grids for the same input.
* No out-of-bounds panic occurs.
* No infinite loops occur.
* All required edge cases behave correctly.
* Performance scales linearly with grid size.

If any variant produces different output from the others for the same input and connectivity setting, the implementation is incorrect.
