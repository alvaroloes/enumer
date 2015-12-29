h1. Enumer

Enumer generates Go code to get string names from enum values and viceversa.
It is a fork of [Rob Pike’s Stringer tool](https://godoc.org/golang.org/x/tools/cmd/stringer) 
but adding a *"string to enum value"* method to the generated code.

For example, if we have an enum type called `Pill`, executing `enumer -type=Pill` will generate two methods:

```
func (i Pill) String() string {
    //...
}

func PillString(s string) (Pill, error) {
    //...
}
```

For more information on how to use, please go to the [Stringer docs](https://godoc.org/golang.org/x/tools/cmd/stringer) 