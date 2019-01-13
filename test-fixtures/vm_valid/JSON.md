# JSON

## Programs

### Program json

```
  0: call "element"
  1: halt
  2: func "value"
  3: fork 21
  4: fork 19
  5: fork 17
  6: fork 15
  7: fork 13
  8: fork 11
  9: call "object"
 10: jump 12
 11: call "array"
 12: jump 14
 13: call "string"
 14: jump 16
 15: call "number"
 16: jump 18
 17: string "true"
 18: jump 20
 19: string "false"
 20: jump 22
 21: string "null"
 22: return
 23: func "object"
 24: string "{"
 25: fork 28
 26: call "ws"
 27: jump 29
 28: call "members"
 29: string "}"
 30: return
 31: func "members"
 32: call "member"
 33: fork 36
 34: string ","
 35: call "members"
 36: return
 37: func "member"
 38: call "ws"
 39: call "string"
 40: call "ws"
 41: string ":"
 42: call "element"
 43: return
 44: func "array"
 45: string "["
 46: fork 49
 47: call "ws"
 48: jump 50
 49: call "elements"
 50: string "]"
 51: return
 52: func "elements"
 53: call "element"
 54: fork 57
 55: string ","
 56: call "elements"
 57: return
 58: func "element"
 59: call "ws"
 60: call "value"
 61: call "ws"
 62: return
 63: func "string"
 64: string "\""
 65: call "characters"
 66: string "\""
 67: return
 68: func "characters"
 69: fork 87
 70: fork 85
 71: fork 83
 72: fork 81
 73: fork 79
 74: fork 77
 75: range "a" "z"
 76: jump 78
 77: range "A" "Z"
 78: jump 80
 79: range "0" "9"
 80: jump 82
 81: string " "
 82: jump 84
 83: string "."
 84: jump 86
 85: string ","
 86: jump 69
 87: return
 88: func "number"
 89: range "0" "9"
 90: fork 89
 91: return
 92: func "ws"
 93: fork 105
 94: fork 103
 95: fork 101
 96: fork 99
 97: string ""
 98: jump 100
 99: string "\t"
100: jump 102
101: string "\n"
102: jump 104
103: string "\r"
104: jump 106
105: string " "
106: return
```

## Inputs

## Input 1: "{}"

| PROGRAM | VALID | N |  ERR  | OPS | FORKS |
|---------|-------|---|-------|-----|-------|
| json    | true  | 2 | <nil> | 122 |    23 |

## Input 2: "[]"

| PROGRAM | VALID | N |  ERR  | OPS | FORKS |
|---------|-------|---|-------|-----|-------|
| json    | true  | 2 | <nil> | 143 |    29 |

## Input 3: "\"abc\""

| PROGRAM | VALID | N |  ERR  | OPS | FORKS |
|---------|-------|---|-------|-----|-------|
| json    | true  | 5 | <nil> | 176 |    38 |

## Input 4: "123"

| PROGRAM | VALID | N |  ERR  | OPS | FORKS |
|---------|-------|---|-------|-----|-------|
| json    | true  | 3 | <nil> | 135 |    25 |

## Input 5: "true"

| PROGRAM | VALID | N |  ERR  | OPS | FORKS |
|---------|-------|---|-------|-----|-------|
| json    | true  | 4 | <nil> |  76 |    14 |

## Input 6: "false"

| PROGRAM | VALID | N |  ERR  | OPS | FORKS |
|---------|-------|---|-------|-----|-------|
| json    | true  | 5 | <nil> |  77 |    14 |

## Input 7: "null"

| PROGRAM | VALID | N |  ERR  | OPS | FORKS |
|---------|-------|---|-------|-----|-------|
| json    | true  | 4 | <nil> |  74 |    14 |

## Input 8: "{\"foo\": \"bar\", \"hello\": {\"world\": true, \"yes\": 123}}"

| PROGRAM | VALID | N  |  ERR  | OPS  | FORKS |
|---------|-------|----|-------|------|-------|
| json    | true  | 52 | <nil> | 1433 |   297 |

## Input 9: "[{\"foo\": \"bar\", \"hello\": {\"world\": true, \"yes\": [123, false, null]}}]"

| PROGRAM | VALID | N  |  ERR  | OPS  | FORKS |
|---------|-------|----|-------|------|-------|
| json    | true  | 69 | <nil> | 1849 |   379 |

