# Game of Life

## Rules

1. A living cell with 2 or 3 neighbors stays alive.
2. A dead cell with exactly 3 neighbors becomes alive.
3. All other living cells die.
4. The simulation stops when no living cells remain.
5. The grid wraps around (toroidal).

## How to Run

```bash
# Run directly
make run

# Build and execute
make build
./bin/gameoflife
```
