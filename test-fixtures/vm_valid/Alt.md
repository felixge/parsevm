# Alt

## Programs

### Program abcdefghj1

```
0: fork 6
1: fork 4
2: string "abc"
3: jump 5
4: string "def"
5: jump 7
6: string "ghi"
```

### Program abcdefghi2

```
 0: fork 8
 1: fork 5
 2: string "a"
 3: string "bc"
 4: jump 7
 5: string "de"
 6: string "f"
 7: jump 11
 8: string "g"
 9: string "h"
10: string "i"
```

## Inputs

## Input 1: ""

|  PROGRAM   | VALID | N |  ERR  | OPS | FORKS | CONCURRENCY |
|------------|-------|---|-------|-----|-------|-------------|
| abcdefghj1 | false | 0 | <nil> |   8 |     2 |           0 |
| abcdefghi2 | false | 0 | <nil> |   8 |     2 |           0 |

## Input 2: "abc"

|  PROGRAM   | VALID | N |  ERR  | OPS | FORKS | CONCURRENCY |
|------------|-------|---|-------|-----|-------|-------------|
| abcdefghj1 | true  | 3 | <nil> |  15 |     2 |           3 |
| abcdefghi2 | true  | 3 | <nil> |  15 |     2 |           3 |

## Input 3: "def"

|  PROGRAM   | VALID | N |  ERR  | OPS | FORKS | CONCURRENCY |
|------------|-------|---|-------|-----|-------|-------------|
| abcdefghj1 | true  | 3 | <nil> |  14 |     2 |           3 |
| abcdefghi2 | true  | 3 | <nil> |  14 |     2 |           3 |

## Input 4: "ghi"

|  PROGRAM   | VALID | N |  ERR  | OPS | FORKS | CONCURRENCY |
|------------|-------|---|-------|-----|-------|-------------|
| abcdefghj1 | true  | 3 | <nil> |  13 |     2 |           3 |
| abcdefghi2 | true  | 3 | <nil> |  13 |     2 |           3 |

## Input 5: "abcabc"

|  PROGRAM   | VALID | N |     ERR     | OPS | FORKS | CONCURRENCY |
|------------|-------|---|-------------|-----|-------|-------------|
| abcdefghj1 | false | 3 | short write |  15 |     2 |           3 |
| abcdefghi2 | false | 3 | short write |  15 |     2 |           3 |

## Input 6: "adg"

|  PROGRAM   | VALID | N |     ERR     | OPS | FORKS | CONCURRENCY |
|------------|-------|---|-------------|-----|-------|-------------|
| abcdefghj1 | false | 1 | short write |  10 |     2 |           3 |
| abcdefghi2 | false | 1 | short write |  10 |     2 |           3 |

