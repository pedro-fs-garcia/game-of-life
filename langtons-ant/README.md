# Langton's Ant

Go implementation of **Langton's Ant**, a well-known cellular automaton with a stateful agent moving on a grid.

---

## Rules

1. The grid is conceptually infinite (simulated using **wrap-around logic**).
2. The ant has a **position** and a **direction**.
3. At each step, the ant observes the **current cell**:
   - On a **white cell**: turn right, flip the cell to black.
   - On a **black cell**: turn left, flip the cell to white.
4. After turning and flipping, the ant **moves forward** one cell.

---

## Usage

When running, you can input initial coordinates for black cells in the format:

```
{(x,y),(x,y),...}
```

Or simply press **Enter** to start with an empty (all white) board.

The board is printed at each step, updating every 50 milliseconds.

---

## Build and Run

> From the **repository root**, using Taskfile (recommended, cross-platform):

```bash
# Run Langton's Ant
task langtons-ant:run

# Build binary to ./bin
task langtons-ant:build

# Clean and normalize dependencies
task langtons-ant:tidy
```

> Alternatively, from the repository root using Make (Unix-like systems):

```bash
# Run directly
make run EXERCISE=langtons-ant

# Build binary
make build EXERCISE=langtons-ant
./bin/langtons-ant

# Format code
make fmt EXERCISE=langtons-ant
```

---

## Example Output

```
        
  ■     
    ↑   
      ■ 
        
```

---

## Project Structure

```
langtons-ant/
├── cmd/langtonsant/
│   ├── main.go         # Entry point, handles input and runs the simulation
│   ├── ant.go          # Ant logic: position, direction, movement
│   ├── board.go        # Board logic and simulation steps (toroidal grid)
│   └── cell.go         # Cell state (white/black)
├── go.mod              # Go module definition (exercise is self-contained)
```

At the repository root:

```
.
├── bin/                # Compiled binaries for all exercises
├── langtons-ant/
├── ...
├── Taskfile.yml        # Cross-platform automation
└── Makefile            # Optional Unix-based automation
```

---

## Core Components

* **Ant**
  Represents the ant with its position and direction. Handles turning and movement logic.

* **Board**
  Manages the grid, applies simulation rules, and handles toroidal wrapping.

* **Cell**
  Represents individual cell state (white/black) and handles flipping.

---

## Notes

* This module is **completely independent** and can be run without Taskfile or Make.
* `Taskfile` and `Makefile` are **developer convenience tools only** and are not runtime dependencies.
* The project follows idiomatic Go structure (`cmd/`, isolated modules, explicit dependencies).
