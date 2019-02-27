package main

import "fmt"

// Arguments to format are:
//	[1]: type name
//	[2]: numeric value check code (or "")
const stringNameToValueMethod = `// %[1]sString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func %[1]sString(s string) (%[1]s, error) {
	if val, ok := _%[1]sNameToValueMap[s]; ok {
		return val, nil
	}%[2]s
	return 0, fmt.Errorf("%%s does not belong to %[1]s values", s)
}
`
const stringIgnoreCaseNameToValueMethod = `// %[1]sString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func %[1]sString(s string) (%[1]s, error) {
	if val, ok := _%[1]sNameToValueMap[s]; ok {
		return val, nil
	}
	for k, v := range _%[1]sNameToValueMap {
		if strings.EqualFold(s, k) {
			return v, nil
		}
	}%[2]s
	return 0, fmt.Errorf("%%s does not belong to %[1]s values", s)
}
`
const stringUpperNameToValueMethod = `// %[1]sString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func %[1]sString(s string) (%[1]s, error) {
	if val, ok := _%[1]sNameToValueMap[strings.ToUpper(s)]; ok {
		return val, nil
	}%[2]s
	return 0, fmt.Errorf("%%s does not belong to %[1]s values", s)
}
`
const stringLowerNameToValueMethod = `// %[1]sString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func %[1]sString(s string) (%[1]s, error) {
	if val, ok := _%[1]sNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}%[2]s
	return 0, fmt.Errorf("%%s does not belong to %[1]s values", s)
}
`

// Arguments to format are:
//      [1]: type name
const stringNumericCheck = `
	i, err := strconv.Atoi(s)
	if err == nil {
		for _, v := range _%[1]sNameToValueMap {
			if int(v) == i {
				return v, nil
			}
		}
	}`

type CaseMatch int

const (
	CaseNone = iota
	CaseLower
	CaseUpper
	CaseMixed
)

// Arguments to format are:
//	[1]: type name
const stringValuesMethod = `// %[1]sValues returns all values of the enum
func %[1]sValues() []%[1]s {
	return _%[1]sValues
}
`

// Arguments to format are:
//	[1]: type name
const stringBelongsMethodLoop = `// IsA%[1]s returns "true" if the value is listed in the enum definition. "false" otherwise
func (i %[1]s) IsA%[1]s() bool {
	for _, v := range _%[1]sValues {
		if i == v {
			return true
		}
	}
	return false
}
`

// Arguments to format are:
//	[1]: type name
const stringBelongsMethodSet = `// IsA%[1]s returns "true" if the value is listed in the enum definition. "false" otherwise
func (i %[1]s) IsA%[1]s() bool {
	_, ok := _%[1]sMap[i] 
	return ok
}
`

func (g *Generator) buildBasicExtras(runs [][]Value, typeName string, runsThreshold int, ignoreCase CaseMatch, numeric bool) {
	// At this moment, either "g.declareIndexAndNameVars()" or "g.declareNameVars()" has been called

	// Print the slice of values
	g.Printf("\nvar _%sValues = []%s{", typeName, typeName)
	for _, values := range runs {
		for _, value := range values {
			g.Printf("\t%s, ", value.str)
		}
	}
	g.Printf("}\n\n")

	// Print the map between name and value
	g.Printf("\nvar _%sNameToValueMap = map[string]%s{\n", typeName, typeName)
	thereAreRuns := len(runs) > 1 && len(runs) <= runsThreshold
	var n int
	var runID string
	for i, values := range runs {
		if thereAreRuns {
			runID = "_" + fmt.Sprintf("%d", i)
			n = 0
		} else {
			runID = ""
		}

		for _, value := range values {
			g.Printf("\t_%sName%s[%d:%d]: %s,\n", typeName, runID, n, n+len(value.name), &value)
			n += len(value.name)
		}
	}
	g.Printf("}\n\n")

	// Print the basic extra methods
	numCheck := ""
	if numeric {
		numCheck = fmt.Sprintf(stringNumericCheck, typeName)
	}
	switch ignoreCase {
	case CaseLower:
		g.Printf(stringLowerNameToValueMethod, typeName, numCheck)
	case CaseUpper:
		g.Printf(stringUpperNameToValueMethod, typeName, numCheck)
	case CaseMixed:
		g.Printf(stringIgnoreCaseNameToValueMethod, typeName, numCheck)
	default:
		g.Printf(stringNameToValueMethod, typeName, numCheck)
	}

	g.Printf(stringValuesMethod, typeName)
	if len(runs) < runsThreshold {
		g.Printf(stringBelongsMethodLoop, typeName)
	} else { // There is a map of values, the code is simpler then
		g.Printf(stringBelongsMethodSet, typeName)
	}
}

// Arguments to format are:
//	[1]: type name
const jsonMethods = `
// MarshalJSON implements the json.Marshaler interface for %[1]s
func (i %[1]s) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for %[1]s
func (i *%[1]s) UnmarshalJSON(data []byte) error {
	var s string
	var err error
	if err = json.Unmarshal(data, &s); err != nil {%[2]s	}

	*i, err = %[1]sString(s)
	return err
}
`

const jsonNumericCheck = `
		var val int
		if err = json.Unmarshal(data, &val); err != nil {
			return fmt.Errorf("%[1]s should be a string, got %%s", data)
		}
		*i = %[1]s(val)
		if !i.IsA%[1]s() {
			return fmt.Errorf("Invalid value for %[1]s (%%d)", val)
		}
		return nil
`

const jsonNoNumericCheck = `
		return fmt.Errorf("%[1]s should be a string, got %%s", data)
`

func (g *Generator) buildJSONMethods(runs [][]Value, typeName string, runsThreshold int, numeric bool) {
	var numCheck string
	if numeric {
		numCheck = fmt.Sprintf(jsonNumericCheck, typeName)
	} else {
		numCheck = fmt.Sprintf(jsonNoNumericCheck, typeName)
	}
	g.Printf(jsonMethods, typeName, numCheck)
}

// Arguments to format are:
//	[1]: type name
const textMethods = `
// MarshalText implements the encoding.TextMarshaler interface for %[1]s
func (i %[1]s) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for %[1]s
func (i *%[1]s) UnmarshalText(text []byte) error {
	var err error
	*i, err = %[1]sString(string(text))
	return err
}
`

func (g *Generator) buildTextMethods(runs [][]Value, typeName string, runsThreshold int) {
	g.Printf(textMethods, typeName)
}

// Arguments to format are:
//	[1]: type name
const yamlMethods = `
// MarshalYAML implements a YAML Marshaler for %[1]s
func (i %[1]s) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for %[1]s
func (i *%[1]s) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = %[1]sString(s)
	return err
}
`

func (g *Generator) buildYAMLMethods(runs [][]Value, typeName string, runsThreshold int) {
	g.Printf(yamlMethods, typeName)
}
