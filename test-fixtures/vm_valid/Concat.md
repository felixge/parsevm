# Concat

## Programs

### Program abc1

```
0: string "abc"
```

### Program abc2a

```
0: string "ab"
1: string "c"
```

### Program abc2b

```
0: string "a"
1: string "bc"
```

### Program abc3

```
0: string "a"
1: string "b"
2: string "c"
```

## Inputs

## Input 1: "hello world"

| PROGRAM | VALID | N |     ERR     | OPS | FORKS | CONCURRENCY |
|---------|-------|---|-------------|-----|-------|-------------|
| abc1    | false | 0 | short write |   2 |     0 |           1 |
| abc2a   | false | 0 | short write |   2 |     0 |           1 |
| abc2b   | false | 0 | short write |   2 |     0 |           1 |
| abc3    | false | 0 | short write |   2 |     0 |           1 |

## Input 2: "world hello"

| PROGRAM | VALID | N |     ERR     | OPS | FORKS | CONCURRENCY |
|---------|-------|---|-------------|-----|-------|-------------|
| abc1    | false | 0 | short write |   2 |     0 |           1 |
| abc2a   | false | 0 | short write |   2 |     0 |           1 |
| abc2b   | false | 0 | short write |   2 |     0 |           1 |
| abc3    | false | 0 | short write |   2 |     0 |           1 |

## Input 3: "helloworld"

| PROGRAM | VALID | N |     ERR     | OPS | FORKS | CONCURRENCY |
|---------|-------|---|-------------|-----|-------|-------------|
| abc1    | false | 0 | short write |   2 |     0 |           1 |
| abc2a   | false | 0 | short write |   2 |     0 |           1 |
| abc2b   | false | 0 | short write |   2 |     0 |           1 |
| abc3    | false | 0 | short write |   2 |     0 |           1 |

## Input 4: "hhello world"

| PROGRAM | VALID | N |     ERR     | OPS | FORKS | CONCURRENCY |
|---------|-------|---|-------------|-----|-------|-------------|
| abc1    | false | 0 | short write |   2 |     0 |           1 |
| abc2a   | false | 0 | short write |   2 |     0 |           1 |
| abc2b   | false | 0 | short write |   2 |     0 |           1 |
| abc3    | false | 0 | short write |   2 |     0 |           1 |

## Input 5: ""

| PROGRAM | VALID | N |  ERR  | OPS | FORKS | CONCURRENCY |
|---------|-------|---|-------|-----|-------|-------------|
| abc1    | false | 0 | <nil> |   2 |     0 |           0 |
| abc2a   | false | 0 | <nil> |   2 |     0 |           0 |
| abc2b   | false | 0 | <nil> |   2 |     0 |           0 |
| abc3    | false | 0 | <nil> |   2 |     0 |           0 |

## Input 6: "hello"

| PROGRAM | VALID | N |     ERR     | OPS | FORKS | CONCURRENCY |
|---------|-------|---|-------------|-----|-------|-------------|
| abc1    | false | 0 | short write |   2 |     0 |           1 |
| abc2a   | false | 0 | short write |   2 |     0 |           1 |
| abc2b   | false | 0 | short write |   2 |     0 |           1 |
| abc3    | false | 0 | short write |   2 |     0 |           1 |

