# Repeat

## Programs

### Program abc1

```
0: string "abc"
1: fork 3
2: string "abc"
3: fork 5
4: string "abc"
```

### Program abc3

```
 0: string "a"
 1: string "b"
 2: string "c"
 3: fork 7
 4: string "a"
 5: string "b"
 6: string "c"
 7: fork 11
 8: string "a"
 9: string "b"
10: string "c"
```

## Inputs

## Input 1: ""

| PROGRAM | VALID | N |  ERR  | OPS | FORKS | CONCURRENCY |
|---------|-------|---|-------|-----|-------|-------------|
| abc1    | false | 0 | <nil> |   2 |     0 |           0 |
| abc3    | false | 0 | <nil> |   2 |     0 |           0 |

## Input 2: "ab"

| PROGRAM | VALID | N |  ERR  | OPS | FORKS | CONCURRENCY |
|---------|-------|---|-------|-----|-------|-------------|
| abc1    | false | 2 | <nil> |   6 |     0 |           1 |
| abc3    | false | 2 | <nil> |   6 |     0 |           1 |

## Input 3: "abc"

| PROGRAM | VALID | N |  ERR  | OPS | FORKS | CONCURRENCY |
|---------|-------|---|-------|-----|-------|-------------|
| abc1    | true  | 3 | <nil> |  11 |     2 |           1 |
| abc3    | true  | 3 | <nil> |  11 |     2 |           1 |

## Input 4: "abcabc"

| PROGRAM | VALID | N |  ERR  | OPS | FORKS | CONCURRENCY |
|---------|-------|---|-------|-----|-------|-------------|
| abc1    | true  | 6 | <nil> |  24 |     3 |           3 |
| abc3    | true  | 6 | <nil> |  24 |     3 |           3 |

## Input 5: "abcabcabc"

| PROGRAM | VALID | N |  ERR  | OPS | FORKS | CONCURRENCY |
|---------|-------|---|-------|-----|-------|-------------|
| abc1    | true  | 9 | <nil> |  30 |     3 |           3 |
| abc3    | true  | 9 | <nil> |  30 |     3 |           3 |

## Input 6: "abcabcabcabc"

| PROGRAM | VALID | N |     ERR     | OPS | FORKS | CONCURRENCY |
|---------|-------|---|-------------|-----|-------|-------------|
| abc1    | false | 9 | short write |  30 |     3 |           3 |
| abc3    | false | 9 | short write |  30 |     3 |           3 |

## Input 7: "adcd"

| PROGRAM | VALID | N |     ERR     | OPS | FORKS | CONCURRENCY |
|---------|-------|---|-------------|-----|-------|-------------|
| abc1    | false | 1 | short write |   4 |     0 |           1 |
| abc3    | false | 1 | short write |   4 |     0 |           1 |

