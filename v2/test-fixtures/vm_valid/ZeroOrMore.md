# ZeroOrMore

## Programs

### Program abc1

```
0: fork +3
1: string "abc"
2: jmp -2
```

### Program abc2

```
0: fork +5
1: string "a"
2: string "b"
3: string "c"
4: jmp -4
```

## Inputs

## Input 1: ""

| PROGRAM | VALID | N |  ERR  | OPS | FORKS |
|---------|-------|---|-------|-----|-------|
| abc1    | true  | 0 | <nil> |   1 |     1 |
| abc2    | true  | 0 | <nil> |   1 |     1 |
## Input 2: "abc"

| PROGRAM | VALID | N |  ERR  | OPS | FORKS |
|---------|-------|---|-------|-----|-------|
| abc1    | true  | 3 | <nil> |   6 |     2 |
| abc2    | true  | 3 | <nil> |   6 |     2 |
## Input 3: "abcabc"

| PROGRAM | VALID | N |  ERR  | OPS | FORKS |
|---------|-------|---|-------|-----|-------|
| abc1    | true  | 6 | <nil> |  11 |     3 |
| abc2    | true  | 6 | <nil> |  11 |     3 |
## Input 4: "abcabcabc"

| PROGRAM | VALID | N |  ERR  | OPS | FORKS |
|---------|-------|---|-------|-----|-------|
| abc1    | true  | 9 | <nil> |  16 |     4 |
| abc2    | true  | 9 | <nil> |  16 |     4 |
## Input 5: "ab"

| PROGRAM | VALID | N |  ERR  | OPS | FORKS |
|---------|-------|---|-------|-----|-------|
| abc1    | false | 2 | <nil> |   3 |     1 |
| abc2    | false | 2 | <nil> |   3 |     1 |
## Input 6: "def"

| PROGRAM | VALID | N |     ERR     | OPS | FORKS |
|---------|-------|---|-------------|-----|-------|
| abc1    | false | 0 | short write |   2 |     1 |
| abc2    | false | 0 | short write |   2 |     1 |
