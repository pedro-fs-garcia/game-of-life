# Game of Life

Go implementation of **Conway's Game of Life**, featuring a **toroidal (wrap-around) grid** and **interactive input** for defining the initial living cells.
---

## Rules

1. A living cell with **2 or 3 neighbors** stays alive.
2. A dead cell with **exactly 3 neighbors** becomes alive.
3. All other living cells die.
4. The simulation stops when **no living cells remain** or **no changes occur**.
5. The grid **wraps around** (toroidal topology).
---

## Usage

When running, you can input initial coordinates in the format:

```
{(x,y),(x,y),...}
```

Or simply press **Enter** to use a predefined cross pattern.

The board is printed at each generation, updating every second.

---

## Build and Run

> From the **repository root**, using Taskfile (recommended, cross-platform):

```bash
# Run the Game of Life
task gameoflife:run

# Build binary to ./bin
task gameoflife:build

# Clean and normalize dependencies
task gameoflife:tidy
```

> Alternatively, from the repository root using Make (Unix-like systems):

```bash
# Run directly
make run EXERCISE=gameoflife

# Build binary
make build EXERCISE=gameoflife
./bin/gameoflife

# Format code
make fmt EXERCISE=gameoflife
```

---

## Example Output

```
■ □ □ □ □
□ ■ □ □ □
□ □ ■ □ □
□ □ □ ■ □
□ □ □ □ ■
```

---

## Project Structure

```
gameoflife/
├── cmd/app/
│   ├── main.go         # Entry point, handles input and runs the simulation
│   └── game/
│       ├── board.go    # Board logic and simulation steps (toroidal grid)
│       ├── cell.go     # Cell state and neighbor rules
│       └── coord.go    # Coordinate utilities
├── go.mod              # Go module definition (exercise is self-contained)
```

At the repository root:

```
.
├── bin/                # Compiled binaries for all exercises
├── gameoflife/
├── ...
├── Taskfile.yml        # Cross-platform automation
└── Makefile            # Optional Unix-based automation
```

---

## Core Components

* **Board**
  Manages the grid, applies simulation rules, and handles toroidal wrapping.

* **Cell**
  Represents individual cell state (alive/dead) and computes transitions based on neighbors.

* **Coord**
  Utility type for working with grid coordinates, used during initialization and simulation.

---

## Notes

* This module is **completely independent** and can be run without Taskfile or Make.
* `Taskfile` and `Makefile` are **developer convenience tools only** and are not runtime dependencies.
* The project follows idiomatic Go structure (`cmd/`, isolated modules, explicit dependencies).
