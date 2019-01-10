# string

## Programs

### Program helloWorld1

```
0: string "hello world"
```

### Program helloWorld2

```
0: string "hello "
1: string "world"
```

### Program helloWorld3

```
0: string "hello"
1: string " "
2: string "world"
```

## Inputs

## Input 1: "hello world"

|   PROGRAM   | VALID | N  |  ERR  | OPS |
|-------------|-------|----|-------|-----|
| helloWorld1 | true  | 11 | <nil> |  11 |
| helloWorld2 | true  | 11 | <nil> |  11 |
| helloWorld3 | true  | 11 | <nil> |  11 |
## Input 2: "world hello"

|   PROGRAM   | VALID | N |     ERR     | OPS |
|-------------|-------|---|-------------|-----|
| helloWorld1 | false | 0 | short write |   1 |
| helloWorld2 | false | 0 | short write |   1 |
| helloWorld3 | false | 0 | short write |   1 |
## Input 3: "helloworld"

|   PROGRAM   | VALID | N |     ERR     | OPS |
|-------------|-------|---|-------------|-----|
| helloWorld1 | false | 5 | short write |   6 |
| helloWorld2 | false | 5 | short write |   6 |
| helloWorld3 | false | 5 | short write |   6 |
## Input 4: "hhello world"

|   PROGRAM   | VALID | N |     ERR     | OPS |
|-------------|-------|---|-------------|-----|
| helloWorld1 | false | 1 | short write |   2 |
| helloWorld2 | false | 1 | short write |   2 |
| helloWorld3 | false | 1 | short write |   2 |
## Input 5: ""

|   PROGRAM   | VALID | N |  ERR  | OPS |
|-------------|-------|---|-------|-----|
| helloWorld1 | false | 0 | <nil> |   0 |
| helloWorld2 | false | 0 | <nil> |   0 |
| helloWorld3 | false | 0 | <nil> |   0 |
## Input 6: "hello"

|   PROGRAM   | VALID | N |  ERR  | OPS |
|-------------|-------|---|-------|-----|
| helloWorld1 | false | 5 | <nil> |   5 |
| helloWorld2 | false | 5 | <nil> |   5 |
| helloWorld3 | false | 5 | <nil> |   5 |
