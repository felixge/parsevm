# star

## Programs

### Program abcStar

```
0: fork 3
1: string "abc"
2: jmp -2
```

## Inputs

## Input 1: ""

| PROGRAM | VALID | N |  ERR  | OPS |
|---------|-------|---|-------|-----|
| abcStar | false | 0 | <nil> |   0 |
## Input 2: "abc"

| PROGRAM | VALID | N |  ERR  | OPS |
|---------|-------|---|-------|-----|
| abcStar | false | 3 | <nil> |   4 |
## Input 3: "abcabc"

| PROGRAM | VALID | N |  ERR  | OPS |
|---------|-------|---|-------|-----|
| abcStar | false | 6 | <nil> |   9 |
## Input 4: "abcabcabc"

| PROGRAM | VALID | N |  ERR  | OPS |
|---------|-------|---|-------|-----|
| abcStar | false | 9 | <nil> |  14 |
## Input 5: "ab"

| PROGRAM | VALID | N |  ERR  | OPS |
|---------|-------|---|-------|-----|
| abcStar | false | 2 | <nil> |   3 |
## Input 6: "def"

| PROGRAM | VALID | N |     ERR     | OPS |
|---------|-------|---|-------------|-----|
| abcStar | true  | 0 | short write |   2 |
