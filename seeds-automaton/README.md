## Exercise A7 — Seeds Automaton

**Domain:** 2D cellular automaton with explosive growth.

**Why:** Extremely simple rules but **exponential state change**. Tests grid performance.

---

### Rules

**Grid:** 2D finite grid (no wrap-around).

**Cell States:** Binary (alive or dead).

**Transition Rules:**

1. **Dead → Alive** if **exactly 2 neighbors** are alive
2. **Alive → Dead** (always)

**Neighborhood:** 8-cell Moore neighborhood.

---

### Key Insight

Unlike Game of Life, a live cell **always dies**. This creates explosive, ephemeral patterns.

---

### Implementation Constraints

* Must handle rapid state changes efficiently
* Track population count per generation
* Terminate if grid becomes entirely dead
* Benchmark: 1000 generations on 500×500 grid in < 1 second

---

### Forces You to Learn

* Performance optimization
* Early termination conditions
* Allocation profiling
* Slice vs map for dense/sparse patterns

---
