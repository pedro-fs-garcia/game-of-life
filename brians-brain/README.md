# Brian's Brain

Go implementation of **Brian's Brain**, featuring a **toroidal (wrap-around) grid** and **interactive input** for defining the initial living cells and grid size.
---

## Rules

**Grid:** 2D finite grid with wrap-around (toroidal topology).

**Cell States:**

| State | Name |
|-------|------|
| 0 | Off |
| 1 | On |
| 2 | Dying |

**Transition Rules (simultaneous):**

1. **Off → On** if **exactly 2 neighbors** are On
2. **On → Dying** (always)
3. **Dying → Off** (always)

**Neighborhood:** 8-cell Moore neighborhood.