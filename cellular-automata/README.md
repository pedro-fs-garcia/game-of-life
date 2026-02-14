# Cellular Automata

Go implementation of **Elementary Cellular Automata (ECA)** using Wolfram's rule classification.
Implemented rules: **30**, **90**, and **110**.

---

## What is an Elementary Cellular Automaton?

A one-dimensional, binary-state cellular automaton where:

- Cells are arranged in a **1D array** (each cell is `0` or `1`)
- At each generation, all cells update **simultaneously**
- Each cell's next state depends only on itself and its two immediate neighbors

---

## How It Works

### 1. Neighborhood

Each cell has a **radius-1 neighborhood**: left neighbor, self, right neighbor.

```
... [L] [C] [R] ...
      |   |   |
      v   v   v
    3-bit pattern
```

### 2. Pattern to Index

The 3 cells form a 3-bit binary number, which converts directly to a decimal index (0-7):

| Pattern | 111 | 110 | 101 | 100 | 011 | 010 | 001 | 000 |
|---------|-----|-----|-----|-----|-----|-----|-----|-----|
| Index   |  7  |  6  |  5  |  4  |  3  |  2  |  1  |  0  |

### 3. Rule Encoding (Wolfram Code)

Each rule is an integer **0-255**, interpreted as an 8-bit binary value.

- Bit position `N` determines the output for pattern index `N`
- If bit `N` is `1` -> next state is `1`
- If bit `N` is `0` -> next state is `0`

### 4. Boundary Conditions

Uses **wrap-around (toroidal)** boundaries:
- Left neighbor of first cell = last cell
- Right neighbor of last cell = first cell

---

## Example: Applying Rule 30

Rule 30 in binary: `00011110`

```
Bit value:   0   0   0   1   1   1   1   0
             |   |   |   |   |   |   |   |
Bit index:   7   6   5   4   3   2   1   0
```

**Step-by-step:**

1. Current cell and neighbors are `011`
2. Pattern `011` in binary = index **3**
3. Check bit 3 of Rule 30 -> `1`
4. Cell's next state = **1**

Another example:

1. Current cell and neighbors are `110`
2. Pattern `110` = index **6**
3. Check bit 6 of Rule 30 -> `0`
4. Cell's next state = **0**

---

## Implemented Rules

### Rule 30 (`00011110`)

| Pattern | 111 | 110 | 101 | 100 | 011 | 010 | 001 | 000 |
|---------|-----|-----|-----|-----|-----|-----|-----|-----|
| Result  |  0  |  0  |  0  |  1  |  1  |  1  |  1  |  0  |

Produces **chaotic, pseudo-random** patterns. Used historically as a random number generator in Mathematica.

### Rule 90 (`01011010`)

| Pattern | 111 | 110 | 101 | 100 | 011 | 010 | 001 | 000 |
|---------|-----|-----|-----|-----|-----|-----|-----|-----|
| Result  |  0  |  1  |  0  |  1  |  1  |  0  |  1  |  0  |

Equivalent to `left XOR right`. Produces the **Sierpinski triangle** fractal.

### Rule 110 (`01101110`)

| Pattern | 111 | 110 | 101 | 100 | 011 | 010 | 001 | 000 |
|---------|-----|-----|-----|-----|-----|-----|-----|-----|
| Result  |  0  |  1  |  1  |  0  |  1  |  1  |  1  |  0  |

Proven to be **Turing-complete** — capable of universal computation.
