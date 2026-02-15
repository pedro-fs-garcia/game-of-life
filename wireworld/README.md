# Wireworld

Go implementation of **Wireworld**, created by Brian Silverman in 1987. A cellular automaton designed to simulate digital logic circuits.
Different from Game of Life and Langton's Ant, Wireworld is designed to be a structured computational model.

---

## What is Wireworld?

Wireworld is a **Turing-complete** cellular automaton that models the flow of electrons through wires. Unlike Game of Life's chaotic emergent behavior, Wireworld produces **predictable, controllable patterns** that can implement actual digital logic.

The key insight: electrons travel along conductor cells as **head-tail pairs**, propagating signals that can be combined using geometric arrangements to form logic gates, clocks, and even full computers.

**Use cases:**
- Simulate AND, OR, XOR, NOT gates
- Build clocks (oscillators)
- Construct signal diodes (one-way paths)
- Design adders, multiplexers, and memory cells

---

## Rules

### Grid

- **2D finite grid** of cells
- Each cell updates **simultaneously** based on its current state and neighbors
- **Moore neighborhood**: 8 adjacent cells (orthogonal + diagonal)

### Cell States

| Value | State | Symbol | Description |
|-------|-------|--------|-------------|
| 0 | Empty | `.` | Background, never changes |
| 1 | Conductor | `#` | Wire that can carry electron signals |
| 2 | Electron Head | `H` | Leading edge of an electron |
| 3 | Electron Tail | `t` | Trailing edge, decays to conductor |

### Transition Rules

All cells update simultaneously. For each cell, apply:

```
┌─────────────────┬────────────────────────────────────────────────┐
│ Current State   │ Next State                                     │
├─────────────────┼────────────────────────────────────────────────┤
│ Empty (0)       │ Empty (0)                                      │
│ Electron Head (2)│ Electron Tail (3)                             │
│ Electron Tail (3)│ Conductor (1)                                 │
│ Conductor (1)   │ Electron Head (2) if 1 or 2 neighbors are Head │
│ Conductor (1)   │ Conductor (1) otherwise                        │
└─────────────────┴────────────────────────────────────────────────┘
```

### State Transition Diagram

```                                     
  ┌──────┐                                  
  │Empty │ ◄──── (never changes)            
  └──────┘                                  
                                            
  ┌──────────┐        ┌──────────┐    ┌────────────┐
  │Conductor │───────►│Head      │───►│Tail        │
  └──────────┘        └──────────┘    └────────────┘
        (1 or 2 Head neighbors)
        (0 or 3+ Head neighbors: stay Conductor)
```

---

## Example Patterns

### Signal Propagation on a Wire

An electron (Head-Tail pair) travels along a conductor wire, moving one cell per generation. The Head leads, the Tail follows.

**Initial state (t=0):**
```
. . . . . . . . . . . . . . .
. . # # # # t H # # # # # . .
. . . . . . . . . . . . . . .
```

**After 1 tick (t=1):**
```
. . . . . . . . . . . . . . .
. . # # # # # t H # # # # . .
. . . . . . . . . . . . . . .
```

**After 2 ticks (t=2):**
```
. . . . . . . . . . . . . . .
. . # # # # # # t H # # # . .
. . . . . . . . . . . . . . .
```

**After 3 ticks (t=3):**
```
. . . . . . . . . . . . . . .
. . # # # # # # # t H # # . .
. . . . . . . . . . . . . . .
```

The signal moves rightward at one cell per tick. The conductor **ahead** of the Head has exactly 1 Head neighbor, so it becomes a Head. Meanwhile: Head → Tail → Conductor.

---

### Clock (8-Tick Oscillator)

A loop of conductors with an electron pair creates a repeating clock signal.

**Initial state (t=0):**
```
. . . . . . .
. . t H # . .
. . #   # . .
. . # # # . .
. . . . . . .
```

**After 1 tick (t=1):**
```
. . . . . . .
. . # t H . .
. . #   # . .
. . # # # . .
. . . . . . .
```

**After 2 ticks (t=2):**
```
. . . . . . .
. . # # t . .
. . #   H . .
. . # # # . .
. . . . . . .
```

**After 3 ticks (t=3):**
```
. . . . . . .
. . # # # . .
. . #   t . .
. . # # H . .
. . . . . . .
```

The electron travels clockwise around the loop and returns to its starting position after 8 ticks (the loop perimeter).

---

### Example 1

A geometric structure that allows signals to pass in one direction only.

**Structure:**
```
. . . . . . . . . . . . . . . .
. . . . . # . . . . . . . . . .
. . # # # # # # # # # # # # . .
. . . . . # . . . . . . . . . .
. . . . . . . . . . . . . . . .
```

t=0:
```
. . . . . . . . . . . . . . . .
. . . . . # . . . . . . . . . .
. . t H # # # # # # # # # # . .
. . . . . # . . . . . . . . . .
. . . . . . . . . . . . . . . .
```

t=1:
```
. . . . . . . . . . . . . . . .
. . . . . # . . . . . . . . . .
. . # t H # # # # # # # # # . .
. . . . . # . . . . . . . . . .
. . . . . . . . . . . . . . . .
```

t=2:
```
. . . . . . . . . . . . . . . .
. . . . . H . . . . . . . . . .
. . # # t H # # # # # # # # . .
. . . . . H . . . . . . . . . .
. . . . . . . . . . . . . . . .
```

t=3:
```
. . . . . . . . . . . . . . . .
. . . . . t . . . . . . . . . .
. . # # # t # # # # # # # # . .
. . . . . t . . . . . . . . . .
. . . . . . . . . . . . . . . .
```

t=4 (signal dies):
```
. . . . . . . . . . . . . . . .
. . . . . # . . . . . . . . . .
. . # # # # # # # # # # # # . .
. . . . . # . . . . . . . . . .
. . . . . . . . . . . . . . . .
```

Since conductors only become Heads with exactly 1 or 2 Head neighbors, the signal stops.

---

### Example 2:

**Structure:**
```
. . . . . . . . . . . . . . .
. . . . # # # . . . . . . . .
. . . #       # . . . . . . .
. # # #   .   # # # # # # . .
. . . #       # . . . . . . .
. . . . # # # . . . . . . . .
. . . . . . . . . . . . . . .
```

**Input A only (top):**

t=0:
```
. . . . . . . . . . . . . . .
. . . . t H # . . . . . . . .
. . . #       # . . . . . . .
. . . #   .   # # # # # # . .
. . . #       # . . . . . . .
. . . . # # # . . . . . . . .
. . . . . . . . . . . . . . .
```

t=2 (output fires):
```
. . . . . . . . . . . . . . .
. . . . # # t . . . . . . . .
. . . #       H . . . . . . .
. . . #   .   # # # # # # . .
. . . #       # . . . . . . .
. . . . # # # . . . . . . . .
. . . . . . . . . . . . . . .
```

t=3
```
. . . . . . . . . . . . . . .
. . . . # # # . . . . . . . .
. . . #       t . . . . . . .
. . . #   .   H H # # # # . .
. . . #       # . . . . . . .
. . . . # # # . . . . . . . .
. . . . . . . . . . . . . . .
```

t=4
```
. . . . . . . . . . . . . . .
. . . . # # # . . . . . . . .
. . . #       # . . . . . . .
. . . #   .   t t H # # # . .
. . . #       H . . . . . . .
. . . . # # # . . . . . . . .
. . . . . . . . . . . . . . .
```

t=5
```
. . . . . . . . . . . . . . .
. . . . # # # . . . . . . . .
. . . #       # . . . . . . .
. . . #   .   # # t H # # . .
. . . #       t . . . . . . .
. . . . # # H . . . . . . . .
. . . . . . . . . . . . . . .
```
---

## Create Basic Logic Gates
By arranging conductors in specific patterns, we can create logic gates that perform basic Boolean operations.
1. **AND gate**: Output is true if both inputs are true
2. **OR gate**: Output is true if at least one input is true
3. **NOT gate**: Output is the inverse of the input
4. **XOR gate**: Output is true if exactly one input is true
5. **Diode**: Allows signal to pass in one direction only

## Test Scenarios

1. **Signal propagation**: Place `Ht` on a wire, verify it moves one cell per tick
2. **Clock period**: Verify oscillator returns to initial state after N ticks
3. **Diode blocking**: Send signal backward through diode, verify it stops
4. **OR gate truth table**: Test all input combinations (00, 01, 10, 11)
5. **Stability**: Empty grid remains empty forever
3. **Diode blocking**: Send signal backward through diode, verify it stops
4. **OR gate truth table**: Test all input combinations (00, 01, 10, 11)
5. **Stability**: Empty grid remains empty forever

---

## Stretch Goals

- [ ] Build a binary counter
- [ ] Load patterns from file (RLE or plaintext format)
- [ ] Detect stable states and oscillators

---

## References

- [Wireworld - Wikipedia](https://en.wikipedia.org/wiki/Wireworld)
- Brian Silverman's original 1987 paper
- [Wireworld Computer](http://www.quinapalus.com/wi-index.html) — A full computer built in Wireworld
