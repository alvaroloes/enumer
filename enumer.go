package main
import "fmt"

// Arguments to format are:
//	[1]: type name
const stringValueToNameMap = `func %[1]sString(s string) (%[1]s, error) {
	if val, ok := _%[1]sNameToValue_map[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%%s does not belong to %[1]s values", s)
}
`

func (g *Generator) buildValueToNameMap(runs [][]Value, typeName string, runsThreshold int) {
	// At this moment, either "g.declareIndexAndNameVars()" or "g.declareNameVars()" has been called
	g.Printf("\nvar _%sNameToValue_map = map[string]%s{\n", typeName, typeName)
	thereAreRuns := len(runs) > 1 && len(runs) <= runsThreshold
	var n int
	var runID string
	for i, values := range runs {
		if thereAreRuns {
			runID = "_" + fmt.Sprintf("%d",i)
			n = 0
		} else {
			runID = ""
		}

		for _, value := range values {
			g.Printf("\t_%s_name%s[%d:%d]: %s,\n", typeName, runID, n, n+len(value.name), &value)
			n += len(value.name)
		}
	}
	g.Printf("}\n\n")
	g.Printf(stringValueToNameMap, typeName)
}
