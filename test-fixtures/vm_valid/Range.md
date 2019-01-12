# Range

## Programs

### Program b-d1

```
0: range "b" "d"
```

### Program b-d2

```
0: fork 3
1: string "b"
2: jump 4
3: range "c" "d"
```

### Program b-d3

```
0: fork 6
1: fork 4
2: string "b"
3: jump 5
4: string "c"
5: jump 7
6: string "d"
```

## Inputs

## Input 1: "a"

| PROGRAM | VALID | N |     ERR     | OPS | FORKS |
|---------|-------|---|-------------|-----|-------|
| b-d1    | false | 0 | short write |   2 |     0 |
| b-d2    | false | 0 | short write |   5 |     1 |
| b-d3    | false | 0 | short write |   8 |     2 |

## Input 2: "b"

| PROGRAM | VALID | N |  ERR  | OPS | FORKS |
|---------|-------|---|-------|-----|-------|
| b-d1    | true  | 1 | <nil> |   3 |     0 |
| b-d2    | true  | 1 | <nil> |   7 |     1 |
| b-d3    | true  | 1 | <nil> |  11 |     2 |

## Input 3: "c"

| PROGRAM | VALID | N |  ERR  | OPS | FORKS |
|---------|-------|---|-------|-----|-------|
| b-d1    | true  | 1 | <nil> |   3 |     0 |
| b-d2    | true  | 1 | <nil> |   6 |     1 |
| b-d3    | true  | 1 | <nil> |  10 |     2 |

## Input 4: "d"

| PROGRAM | VALID | N |  ERR  | OPS | FORKS |
|---------|-------|---|-------|-----|-------|
| b-d1    | true  | 1 | <nil> |   3 |     0 |
| b-d2    | true  | 1 | <nil> |   6 |     1 |
| b-d3    | true  | 1 | <nil> |   9 |     2 |

## Input 5: "e"

| PROGRAM | VALID | N |     ERR     | OPS | FORKS |
|---------|-------|---|-------------|-----|-------|
| b-d1    | false | 0 | short write |   2 |     0 |
| b-d2    | false | 0 | short write |   5 |     1 |
| b-d3    | false | 0 | short write |   8 |     2 |

## Input 6: "ab"

| PROGRAM | VALID | N |     ERR     | OPS | FORKS |
|---------|-------|---|-------------|-----|-------|
| b-d1    | false | 0 | short write |   2 |     0 |
| b-d2    | false | 0 | short write |   5 |     1 |
| b-d3    | false | 0 | short write |   8 |     2 |

## Input 7: "bb"

| PROGRAM | VALID | N |     ERR     | OPS | FORKS |
|---------|-------|---|-------------|-----|-------|
| b-d1    | false | 1 | short write |   3 |     0 |
| b-d2    | false | 1 | short write |   7 |     1 |
| b-d3    | false | 1 | short write |  11 |     2 |

## Input 8: "cc"

| PROGRAM | VALID | N |     ERR     | OPS | FORKS |
|---------|-------|---|-------------|-----|-------|
| b-d1    | false | 1 | short write |   3 |     0 |
| b-d2    | false | 1 | short write |   6 |     1 |
| b-d3    | false | 1 | short write |  10 |     2 |

## Input 9: "dd"

| PROGRAM | VALID | N |     ERR     | OPS | FORKS |
|---------|-------|---|-------------|-----|-------|
| b-d1    | false | 1 | short write |   3 |     0 |
| b-d2    | false | 1 | short write |   6 |     1 |
| b-d3    | false | 1 | short write |   9 |     2 |

## Input 10: "ec"

| PROGRAM | VALID | N |     ERR     | OPS | FORKS |
|---------|-------|---|-------------|-----|-------|
| b-d1    | false | 0 | short write |   2 |     0 |
| b-d2    | false | 0 | short write |   5 |     1 |
| b-d3    | false | 0 | short write |   8 |     2 |

