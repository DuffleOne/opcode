# opcode

Started playing with the opcode interpreter from advent of code, and got carried away. Wanted to see if I could modularise the design a little.

It does include my puzzle input, but you can change that easily in `cmd/opcode.go`

## Applications

- `01`: add
- `02`: mul
- `03`: input
- `04`: output
- `05`: jump if true
- `06`: jump if false
- `07`: less than
- `08`: equals
- `99`: halt

## Run

`go run cmd/opcode.go`

## Test

`go test ./...`
