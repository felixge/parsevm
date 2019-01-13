# Func_Simple

## Programs

### Program abc1

```
0: call "abc"
1: halt
2: func "abc"
3: string "abc"
4: return
```

## Inputs

## Input 1: ""

| PROGRAM | VALID | N |  ERR  | OPS | FORKS |
|---------|-------|---|-------|-----|-------|
| abc1    | false | 0 | <nil> |   3 |     0 |

## Input 2: "a"

| PROGRAM | VALID | N |  ERR  | OPS | FORKS |
|---------|-------|---|-------|-----|-------|
| abc1    | false | 1 | <nil> |   5 |     0 |

## Input 3: "abc"

| PROGRAM | VALID | N |  ERR  | OPS | FORKS |
|---------|-------|---|-------|-----|-------|
| abc1    | true  | 3 | <nil> |  10 |     0 |

## Input 4: "abcd"

| PROGRAM | VALID | N |     ERR     | OPS | FORKS |
|---------|-------|---|-------------|-----|-------|
| abc1    | false | 3 | short write |  10 |     0 |

## Input 5: "b"

| PROGRAM | VALID | N |     ERR     | OPS | FORKS |
|---------|-------|---|-------------|-----|-------|
| abc1    | false | 0 | short write |   3 |     0 |

