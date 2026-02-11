# Game of Life

Go implementation of Conway's Game of Life, featuring a toroidal (wrap-around) grid and interactive input for initial cell coordinates.

## Rules

1. A living cell with 2 or 3 neighbors stays alive.
2. A dead cell with exactly 3 neighbors becomes alive.
3. All other living cells die.
4. The simulation stops when no living cells remain or no changes occur.
5. The grid wraps around (toroidal).

## Usage

When running, you can input initial coordinates in the format:

```
{(x,y),(x,y),...}
```

Or press Enter to use a standard diagonal/cross pattern.

The board is printed each round, updating every second.

## Build and Run

```bash
# Run directly
make run

# Build and execute
make build
./bin/gameoflife

# Format code
make fmt

# Run tests
make test
```

## Example Output

```
■ □ □ □ □
□ ■ □ □ □
□ □ ■ □ □
□ □ □ ■ □
□ □ □ □ ■
```

## Project Structure

```
gameoflife/
├── cmd/app/
│   ├── main.go         # Entry point, handles input and runs simulation
│   └── game/
│       ├── board.go    # Board logic and simulation
│       ├── cell.go     # Cell state and neighbor rules
│       └── coord.go    # Coordinate utilities
├── bin/                # Built binary output
├── Makefile            # Build, test, run commands
├── go.mod              # Go module definition
```

## Core Components

- **Board**: Manages the grid, cell states, and simulation rounds. Implements toroidal wrapping.
- **Cell**: Represents each cell, tracks alive/dead state, and determines next state based on neighbors.
- **Coord**: Utility for cell coordinates, used for initial setup and board construction.