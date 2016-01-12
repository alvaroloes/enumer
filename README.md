# Enumer
Enumer generates Go code to get string names from enum values and viceversa.
It is a fork of [Rob Pike’s Stringer tool](https://godoc.org/golang.org/x/tools/cmd/stringer) 
but adding a *"string to enum value"* method to the generated code.

This is useful when you need to read enum values from the command line arguments, from a configuration file, 
from a REST API request... In short, from those places where using the real enum value (an integer) would 
be almost meaningless or hard to trace or use by a human

For example, if we have an enum type called `Pill`,
```go
type Pill int

const (
	Placebo Pill = iota
	Aspirin
	Ibuprofen
	Paracetamol
	Acetaminophen = Paracetamol
)
```
executing `enumer -type=Pill` will generate a new file with two methods:
```go
func (i Pill) String() string {
    //...
}

func PillString(s string) (Pill, error) {
    //...
}
```
From now on, we can:
```go
// Convert any Pill value to string
var aspirinString string = Aspirin.String()
// (or use it in any place where a Stringer is accepted)
fmt.Println("I need ", Paracetamol) // Will print "I need Paracetamol"

// Convert a string with the enum name to the corresponding enum value
pill, err := PillString("Ibuprofen")
if err != nil {
    fmt.Println("Unrecognized pill: ", err)
    return
}
// Now pill == Ibuprofen
```

The generated code is exactly the same as the Stringer tool plus the `<Type>String` method, so you can use
**Enumer** where you are already using **Stringer** without any code change.

## How to use
The usage of Enumer is the same as Stringer, no changes were introduced.
For more information please refer to the [Stringer docs](https://godoc.org/golang.org/x/tools/cmd/stringer) 

## Additional functions of this fork
This fork additionally implements the Scanner and Valuer interface to use a enum seamlessly in a database model.
