# Func_Recursive

## Programs

### Program pairs

```
 0: call "pair"
 1: halt
 2: func "pair"
 3: string "("
 4: fork 7
 5: call "pair"
 6: jump 8
 7: string ""
 8: string ")"
 9: return
```

## Inputs

## Input 1: "()"

| PROGRAM | VALID | N |  ERR  | OPS | FORKS |
|---------|-------|---|-------|-----|-------|
| pairs   | true  | 2 | <nil> |  13 |     1 |

## Input 2: "(())"

| PROGRAM | VALID | N |  ERR  | OPS | FORKS |
|---------|-------|---|-------|-----|-------|
| pairs   | true  | 4 | <nil> |  24 |     2 |

## Input 3: "((()))"

| PROGRAM | VALID | N |  ERR  | OPS | FORKS |
|---------|-------|---|-------|-----|-------|
| pairs   | true  | 6 | <nil> |  35 |     3 |

## Input 4: ""

| PROGRAM | VALID | N |  ERR  | OPS | FORKS |
|---------|-------|---|-------|-----|-------|
| pairs   | false | 0 | <nil> |   3 |     0 |

## Input 5: "("

| PROGRAM | VALID | N |  ERR  | OPS | FORKS |
|---------|-------|---|-------|-----|-------|
| pairs   | false | 1 | <nil> |  10 |     1 |

## Input 6: ")"

| PROGRAM | VALID | N |     ERR     | OPS | FORKS |
|---------|-------|---|-------------|-----|-------|
| pairs   | false | 0 | short write |   3 |     0 |

## Input 7: "(()"

| PROGRAM | VALID | N |  ERR  | OPS | FORKS |
|---------|-------|---|-------|-----|-------|
| pairs   | false | 3 | <nil> |  21 |     2 |

## Input 8: "())"

| PROGRAM | VALID | N |     ERR     | OPS | FORKS |
|---------|-------|---|-------------|-----|-------|
| pairs   | false | 2 | short write |  13 |     1 |

## Input 9: "()()"

| PROGRAM | VALID | N |     ERR     | OPS | FORKS |
|---------|-------|---|-------------|-----|-------|
| pairs   | false | 2 | short write |  13 |     1 |

