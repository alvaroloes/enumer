package main

// Arguments to format are:
//	[1]: type name
const valuer = `func (i %[1]s) Value() (driver.Value, error) {
	return i.String(), nil
}
`

const scanner = `func (i %[1]s) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		fmt.Errorf("value is not a byte slice")
	}

	str := string(bytes[:])

	val, err := %[1]sString(str)
	if err != nil {
		return err
	}
	
	i = val
	return nil
}
`

func (g *Generator) addValuerAndScanner(typeName string) {
	g.Printf("\n")
	g.Printf(valuer, typeName)
	g.Printf("\n\n")
	g.Printf(scanner, typeName)
}
