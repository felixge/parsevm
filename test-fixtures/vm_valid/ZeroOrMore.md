# ZeroOrMore

## Programs

### Program abc1

```
0: fork 3
1: string "abc"
2: jump 0
```

### Program abc2

```
0: fork 5
1: string "a"
2: string "b"
3: string "c"
4: jump 0
```

## Inputs

## Input 1: ""

| PROGRAM | VALID | N |  ERR  | OPS | FORKS | CONCURRENCY |
|---------|-------|---|-------|-----|-------|-------------|
| abc1    | true  | 0 | <nil> |   3 |     1 |           0 |
| abc2    | true  | 0 | <nil> |   3 |     1 |           0 |

## Input 2: "abc"

| PROGRAM | VALID | N |  ERR  | OPS | FORKS | CONCURRENCY |
|---------|-------|---|-------|-----|-------|-------------|
| abc1    | true  | 3 | <nil> |  12 |     2 |           2 |
| abc2    | true  | 3 | <nil> |  12 |     2 |           2 |

## Input 3: "abcabc"

| PROGRAM | VALID | N |  ERR  | OPS | FORKS | CONCURRENCY |
|---------|-------|---|-------|-----|-------|-------------|
| abc1    | true  | 6 | <nil> |  21 |     3 |           2 |
| abc2    | true  | 6 | <nil> |  21 |     3 |           2 |

## Input 4: "abcabcabc"

| PROGRAM | VALID | N |  ERR  | OPS | FORKS | CONCURRENCY |
|---------|-------|---|-------|-----|-------|-------------|
| abc1    | true  | 9 | <nil> |  30 |     4 |           2 |
| abc2    | true  | 9 | <nil> |  30 |     4 |           2 |

## Input 5: "ab"

| PROGRAM | VALID | N |  ERR  | OPS | FORKS | CONCURRENCY |
|---------|-------|---|-------|-----|-------|-------------|
| abc1    | false | 2 | <nil> |   8 |     1 |           2 |
| abc2    | false | 2 | <nil> |   8 |     1 |           2 |

## Input 6: "def"

| PROGRAM | VALID | N |     ERR     | OPS | FORKS | CONCURRENCY |
|---------|-------|---|-------------|-----|-------|-------------|
| abc1    | false | 0 | short write |   4 |     1 |           2 |
| abc2    | false | 0 | short write |   4 |     1 |           2 |

