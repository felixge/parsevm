# OneOrMore

## Programs

### Program abc1

```
0: string "abc"
1: fork 0
```

### Program abc2

```
0: string "a"
1: string "b"
2: string "c"
3: fork 0
```

## Inputs

## Input 1: ""

| PROGRAM | VALID | N |  ERR  | OPS | FORKS | CONCURRENCY |
|---------|-------|---|-------|-----|-------|-------------|
| abc1    | false | 0 | <nil> |   2 |     0 |           0 |
| abc2    | false | 0 | <nil> |   2 |     0 |           0 |

## Input 2: "abc"

| PROGRAM | VALID | N |  ERR  | OPS | FORKS | CONCURRENCY |
|---------|-------|---|-------|-----|-------|-------------|
| abc1    | true  | 3 | <nil> |  10 |     1 |           1 |
| abc2    | true  | 3 | <nil> |  10 |     1 |           1 |

## Input 3: "abcabc"

| PROGRAM | VALID | N |  ERR  | OPS | FORKS | CONCURRENCY |
|---------|-------|---|-------|-----|-------|-------------|
| abc1    | true  | 6 | <nil> |  18 |     2 |           2 |
| abc2    | true  | 6 | <nil> |  18 |     2 |           2 |

## Input 4: "abcabcabc"

| PROGRAM | VALID | N |  ERR  | OPS | FORKS | CONCURRENCY |
|---------|-------|---|-------|-----|-------|-------------|
| abc1    | true  | 9 | <nil> |  26 |     3 |           2 |
| abc2    | true  | 9 | <nil> |  26 |     3 |           2 |

## Input 5: "ab"

| PROGRAM | VALID | N |  ERR  | OPS | FORKS | CONCURRENCY |
|---------|-------|---|-------|-----|-------|-------------|
| abc1    | false | 2 | <nil> |   6 |     0 |           1 |
| abc2    | false | 2 | <nil> |   6 |     0 |           1 |

## Input 6: "def"

| PROGRAM | VALID | N |     ERR     | OPS | FORKS | CONCURRENCY |
|---------|-------|---|-------------|-----|-------|-------------|
| abc1    | false | 0 | short write |   2 |     0 |           1 |
| abc2    | false | 0 | short write |   2 |     0 |           1 |

