// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains simple golden tests for various examples.
// Besides validating the results when the implementation changes,
// it provides a way to look at the generated code without having
// to execute the print statements in one's head.

package main

import (
	"strings"
	"testing"
)

// Golden represents a test case.
type Golden struct {
	name   string
	input  string // input; the package clause is provided when running the test.
	output string // exected output.
}

var golden = []Golden{
	{"day", day_in, day_out},
	{"offset", offset_in, offset_out},
	{"gap", gap_in, gap_out},
	{"num", num_in, num_out},
	{"unum", unum_in, unum_out},
	{"prime", prime_in, prime_out},
}

var goldenJSON = []Golden{
	{"prime", prime_json_in, prime_json_out},
}

var goldenYAML = []Golden{
	{"prime", prime_yaml_in, prime_yaml_out},
}

var goldenSQL = []Golden{
	{"prime", prime_sql_in, prime_sql_out},
}

var goldenJSONAndSQL = []Golden{
	{"prime", prime_json_and_sql_in, prime_json_and_sql_out},
}

// Each example starts with "type XXX [u]int", with a single space separating them.

// Simple test: enumeration of type int starting at 0.
const day_in = `type Day int
const (
	Monday Day = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)
`

const day_out = `
const _Day_name = "MondayTuesdayWednesdayThursdayFridaySaturdaySunday"

var _Day_index = [...]uint8{0, 6, 13, 22, 30, 36, 44, 50}

func (i Day) String() string {
	if i < 0 || i >= Day(len(_Day_index)-1) {
		return fmt.Sprintf("Day(%d)", i)
	}
	return _Day_name[_Day_index[i]:_Day_index[i+1]]
}

var _DayNameToValue_map = map[string]Day{
	_Day_name[0:6]:   0,
	_Day_name[6:13]:  1,
	_Day_name[13:22]: 2,
	_Day_name[22:30]: 3,
	_Day_name[30:36]: 4,
	_Day_name[36:44]: 5,
	_Day_name[44:50]: 6,
}

func DayString(s string) (Day, error) {
	if val, ok := _DayNameToValue_map[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Day values", s)
}
`

// Enumeration with an offset.
// Also includes a duplicate.
const offset_in = `type Number int
const (
	_ Number = iota
	One
	Two
	Three
	AnotherOne = One  // Duplicate; note that AnotherOne doesn't appear below.
)
`

const offset_out = `
const _Number_name = "OneTwoThree"

var _Number_index = [...]uint8{0, 3, 6, 11}

func (i Number) String() string {
	i -= 1
	if i < 0 || i >= Number(len(_Number_index)-1) {
		return fmt.Sprintf("Number(%d)", i+1)
	}
	return _Number_name[_Number_index[i]:_Number_index[i+1]]
}

var _NumberNameToValue_map = map[string]Number{
	_Number_name[0:3]:  1,
	_Number_name[3:6]:  2,
	_Number_name[6:11]: 3,
}

func NumberString(s string) (Number, error) {
	if val, ok := _NumberNameToValue_map[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Number values", s)
}
`

// Gaps and an offset.
const gap_in = `type Gap int
const (
	Two Gap = 2
	Three Gap = 3
	Five Gap = 5
	Six Gap = 6
	Seven Gap = 7
	Eight Gap = 8
	Nine Gap = 9
	Eleven Gap = 11
)
`

const gap_out = `
const (
	_Gap_name_0 = "TwoThree"
	_Gap_name_1 = "FiveSixSevenEightNine"
	_Gap_name_2 = "Eleven"
)

var (
	_Gap_index_0 = [...]uint8{0, 3, 8}
	_Gap_index_1 = [...]uint8{0, 4, 7, 12, 17, 21}
	_Gap_index_2 = [...]uint8{0, 6}
)

func (i Gap) String() string {
	switch {
	case 2 <= i && i <= 3:
		i -= 2
		return _Gap_name_0[_Gap_index_0[i]:_Gap_index_0[i+1]]
	case 5 <= i && i <= 9:
		i -= 5
		return _Gap_name_1[_Gap_index_1[i]:_Gap_index_1[i+1]]
	case i == 11:
		return _Gap_name_2
	default:
		return fmt.Sprintf("Gap(%d)", i)
	}
}

var _GapNameToValue_map = map[string]Gap{
	_Gap_name_0[0:3]:   2,
	_Gap_name_0[3:8]:   3,
	_Gap_name_1[0:4]:   5,
	_Gap_name_1[4:7]:   6,
	_Gap_name_1[7:12]:  7,
	_Gap_name_1[12:17]: 8,
	_Gap_name_1[17:21]: 9,
	_Gap_name_2[0:6]:   11,
}

func GapString(s string) (Gap, error) {
	if val, ok := _GapNameToValue_map[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Gap values", s)
}
`

// Signed integers spanning zero.
const num_in = `type Num int
const (
	m_2 Num = -2 + iota
	m_1
	m0
	m1
	m2
)
`

const num_out = `
const _Num_name = "m_2m_1m0m1m2"

var _Num_index = [...]uint8{0, 3, 6, 8, 10, 12}

func (i Num) String() string {
	i -= -2
	if i < 0 || i >= Num(len(_Num_index)-1) {
		return fmt.Sprintf("Num(%d)", i+-2)
	}
	return _Num_name[_Num_index[i]:_Num_index[i+1]]
}

var _NumNameToValue_map = map[string]Num{
	_Num_name[0:3]:   -2,
	_Num_name[3:6]:   -1,
	_Num_name[6:8]:   0,
	_Num_name[8:10]:  1,
	_Num_name[10:12]: 2,
}

func NumString(s string) (Num, error) {
	if val, ok := _NumNameToValue_map[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Num values", s)
}
`

// Unsigned integers spanning zero.
const unum_in = `type Unum uint
const (
	m_2 Unum = iota + 253
	m_1
)

const (
	m0 Unum = iota
	m1
	m2
)
`

const unum_out = `
const (
	_Unum_name_0 = "m0m1m2"
	_Unum_name_1 = "m_2m_1"
)

var (
	_Unum_index_0 = [...]uint8{0, 2, 4, 6}
	_Unum_index_1 = [...]uint8{0, 3, 6}
)

func (i Unum) String() string {
	switch {
	case 0 <= i && i <= 2:
		return _Unum_name_0[_Unum_index_0[i]:_Unum_index_0[i+1]]
	case 253 <= i && i <= 254:
		i -= 253
		return _Unum_name_1[_Unum_index_1[i]:_Unum_index_1[i+1]]
	default:
		return fmt.Sprintf("Unum(%d)", i)
	}
}

var _UnumNameToValue_map = map[string]Unum{
	_Unum_name_0[0:2]: 0,
	_Unum_name_0[2:4]: 1,
	_Unum_name_0[4:6]: 2,
	_Unum_name_1[0:3]: 253,
	_Unum_name_1[3:6]: 254,
}

func UnumString(s string) (Unum, error) {
	if val, ok := _UnumNameToValue_map[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Unum values", s)
}
`

// Enough gaps to trigger a map implementation of the method.
// Also includes a duplicate to test that it doesn't cause problems
const prime_in = `type Prime int
const (
	p2 Prime = 2
	p3 Prime = 3
	p5 Prime = 5
	p7 Prime = 7
	p77 Prime = 7 // Duplicate; note that p77 doesn't appear below.
	p11 Prime = 11
	p13 Prime = 13
	p17 Prime = 17
	p19 Prime = 19
	p23 Prime = 23
	p29 Prime = 29
	p37 Prime = 31
	p41 Prime = 41
	p43 Prime = 43
)
`

const prime_out = `
const _Prime_name = "p2p3p5p7p11p13p17p19p23p29p37p41p43"

var _Prime_map = map[Prime]string{
	2:  _Prime_name[0:2],
	3:  _Prime_name[2:4],
	5:  _Prime_name[4:6],
	7:  _Prime_name[6:8],
	11: _Prime_name[8:11],
	13: _Prime_name[11:14],
	17: _Prime_name[14:17],
	19: _Prime_name[17:20],
	23: _Prime_name[20:23],
	29: _Prime_name[23:26],
	31: _Prime_name[26:29],
	41: _Prime_name[29:32],
	43: _Prime_name[32:35],
}

func (i Prime) String() string {
	if str, ok := _Prime_map[i]; ok {
		return str
	}
	return fmt.Sprintf("Prime(%d)", i)
}

var _PrimeNameToValue_map = map[string]Prime{
	_Prime_name[0:2]:   2,
	_Prime_name[2:4]:   3,
	_Prime_name[4:6]:   5,
	_Prime_name[6:8]:   7,
	_Prime_name[8:11]:  11,
	_Prime_name[11:14]: 13,
	_Prime_name[14:17]: 17,
	_Prime_name[17:20]: 19,
	_Prime_name[20:23]: 23,
	_Prime_name[23:26]: 29,
	_Prime_name[26:29]: 31,
	_Prime_name[29:32]: 41,
	_Prime_name[32:35]: 43,
}

func PrimeString(s string) (Prime, error) {
	if val, ok := _PrimeNameToValue_map[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Prime values", s)
}
`
const prime_json_in = `type Prime int
const (
	p2 Prime = 2
	p3 Prime = 3
	p5 Prime = 5
	p7 Prime = 7
	p77 Prime = 7 // Duplicate; note that p77 doesn't appear below.
	p11 Prime = 11
	p13 Prime = 13
	p17 Prime = 17
	p19 Prime = 19
	p23 Prime = 23
	p29 Prime = 29
	p37 Prime = 31
	p41 Prime = 41
	p43 Prime = 43
)
`

const prime_json_out = `
const _Prime_name = "p2p3p5p7p11p13p17p19p23p29p37p41p43"

var _Prime_map = map[Prime]string{
	2:  _Prime_name[0:2],
	3:  _Prime_name[2:4],
	5:  _Prime_name[4:6],
	7:  _Prime_name[6:8],
	11: _Prime_name[8:11],
	13: _Prime_name[11:14],
	17: _Prime_name[14:17],
	19: _Prime_name[17:20],
	23: _Prime_name[20:23],
	29: _Prime_name[23:26],
	31: _Prime_name[26:29],
	41: _Prime_name[29:32],
	43: _Prime_name[32:35],
}

func (i Prime) String() string {
	if str, ok := _Prime_map[i]; ok {
		return str
	}
	return fmt.Sprintf("Prime(%d)", i)
}

var _PrimeNameToValue_map = map[string]Prime{
	_Prime_name[0:2]:   2,
	_Prime_name[2:4]:   3,
	_Prime_name[4:6]:   5,
	_Prime_name[6:8]:   7,
	_Prime_name[8:11]:  11,
	_Prime_name[11:14]: 13,
	_Prime_name[14:17]: 17,
	_Prime_name[17:20]: 19,
	_Prime_name[20:23]: 23,
	_Prime_name[23:26]: 29,
	_Prime_name[26:29]: 31,
	_Prime_name[29:32]: 41,
	_Prime_name[32:35]: 43,
}

func PrimeString(s string) (Prime, error) {
	if val, ok := _PrimeNameToValue_map[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Prime values", s)
}

func (i Prime) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

func (i *Prime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Prime should be a string, got %s", data)
	}

	var err error
	*i, err = PrimeString(s)
	return err
}
`

const prime_yaml_in = `type Prime int
const (
	p2 Prime = 2
	p3 Prime = 3
	p5 Prime = 5
	p7 Prime = 7
	p77 Prime = 7 // Duplicate; note that p77 doesn't appear below.
	p11 Prime = 11
	p13 Prime = 13
	p17 Prime = 17
	p19 Prime = 19
	p23 Prime = 23
	p29 Prime = 29
	p37 Prime = 31
	p41 Prime = 41
	p43 Prime = 43
)
`

const prime_yaml_out = `
const _Prime_name = "p2p3p5p7p11p13p17p19p23p29p37p41p43"

var _Prime_map = map[Prime]string{
	2:  _Prime_name[0:2],
	3:  _Prime_name[2:4],
	5:  _Prime_name[4:6],
	7:  _Prime_name[6:8],
	11: _Prime_name[8:11],
	13: _Prime_name[11:14],
	17: _Prime_name[14:17],
	19: _Prime_name[17:20],
	23: _Prime_name[20:23],
	29: _Prime_name[23:26],
	31: _Prime_name[26:29],
	41: _Prime_name[29:32],
	43: _Prime_name[32:35],
}

func (i Prime) String() string {
	if str, ok := _Prime_map[i]; ok {
		return str
	}
	return fmt.Sprintf("Prime(%d)", i)
}

var _PrimeNameToValue_map = map[string]Prime{
	_Prime_name[0:2]:   2,
	_Prime_name[2:4]:   3,
	_Prime_name[4:6]:   5,
	_Prime_name[6:8]:   7,
	_Prime_name[8:11]:  11,
	_Prime_name[11:14]: 13,
	_Prime_name[14:17]: 17,
	_Prime_name[17:20]: 19,
	_Prime_name[20:23]: 23,
	_Prime_name[23:26]: 29,
	_Prime_name[26:29]: 31,
	_Prime_name[29:32]: 41,
	_Prime_name[32:35]: 43,
}

func PrimeString(s string) (Prime, error) {
	if val, ok := _PrimeNameToValue_map[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Prime values", s)
}

func (i Prime) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

func (i *Prime) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = PrimeString(s)
	return err
}
`

const prime_sql_in = `type Prime int
const (
	p2 Prime = 2
	p3 Prime = 3
	p5 Prime = 5
	p7 Prime = 7
	p77 Prime = 7 // Duplicate; note that p77 doesn't appear below.
	p11 Prime = 11
	p13 Prime = 13
	p17 Prime = 17
	p19 Prime = 19
	p23 Prime = 23
	p29 Prime = 29
	p37 Prime = 31
	p41 Prime = 41
	p43 Prime = 43
)
`

const prime_sql_out = `
const _Prime_name = "p2p3p5p7p11p13p17p19p23p29p37p41p43"

var _Prime_map = map[Prime]string{
	2:  _Prime_name[0:2],
	3:  _Prime_name[2:4],
	5:  _Prime_name[4:6],
	7:  _Prime_name[6:8],
	11: _Prime_name[8:11],
	13: _Prime_name[11:14],
	17: _Prime_name[14:17],
	19: _Prime_name[17:20],
	23: _Prime_name[20:23],
	29: _Prime_name[23:26],
	31: _Prime_name[26:29],
	41: _Prime_name[29:32],
	43: _Prime_name[32:35],
}

func (i Prime) String() string {
	if str, ok := _Prime_map[i]; ok {
		return str
	}
	return fmt.Sprintf("Prime(%d)", i)
}

var _PrimeNameToValue_map = map[string]Prime{
	_Prime_name[0:2]:   2,
	_Prime_name[2:4]:   3,
	_Prime_name[4:6]:   5,
	_Prime_name[6:8]:   7,
	_Prime_name[8:11]:  11,
	_Prime_name[11:14]: 13,
	_Prime_name[14:17]: 17,
	_Prime_name[17:20]: 19,
	_Prime_name[20:23]: 23,
	_Prime_name[23:26]: 29,
	_Prime_name[26:29]: 31,
	_Prime_name[29:32]: 41,
	_Prime_name[32:35]: 43,
}

func PrimeString(s string) (Prime, error) {
	if val, ok := _PrimeNameToValue_map[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Prime values", s)
}

func (i Prime) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *Prime) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	str, ok := value.(string)
	if !ok {
		bytes, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("value is not a byte slice")
		}

		str = string(bytes[:])
	}

	val, err := PrimeString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
`

const prime_json_and_sql_in = `type Prime int
const (
	p2 Prime = 2
	p3 Prime = 3
	p5 Prime = 5
	p7 Prime = 7
	p77 Prime = 7 // Duplicate; note that p77 doesn't appear below.
	p11 Prime = 11
	p13 Prime = 13
	p17 Prime = 17
	p19 Prime = 19
	p23 Prime = 23
	p29 Prime = 29
	p37 Prime = 31
	p41 Prime = 41
	p43 Prime = 43
)
`

const prime_json_and_sql_out = `
const _Prime_name = "p2p3p5p7p11p13p17p19p23p29p37p41p43"

var _Prime_map = map[Prime]string{
	2:  _Prime_name[0:2],
	3:  _Prime_name[2:4],
	5:  _Prime_name[4:6],
	7:  _Prime_name[6:8],
	11: _Prime_name[8:11],
	13: _Prime_name[11:14],
	17: _Prime_name[14:17],
	19: _Prime_name[17:20],
	23: _Prime_name[20:23],
	29: _Prime_name[23:26],
	31: _Prime_name[26:29],
	41: _Prime_name[29:32],
	43: _Prime_name[32:35],
}

func (i Prime) String() string {
	if str, ok := _Prime_map[i]; ok {
		return str
	}
	return fmt.Sprintf("Prime(%d)", i)
}

var _PrimeNameToValue_map = map[string]Prime{
	_Prime_name[0:2]:   2,
	_Prime_name[2:4]:   3,
	_Prime_name[4:6]:   5,
	_Prime_name[6:8]:   7,
	_Prime_name[8:11]:  11,
	_Prime_name[11:14]: 13,
	_Prime_name[14:17]: 17,
	_Prime_name[17:20]: 19,
	_Prime_name[20:23]: 23,
	_Prime_name[23:26]: 29,
	_Prime_name[26:29]: 31,
	_Prime_name[29:32]: 41,
	_Prime_name[32:35]: 43,
}

func PrimeString(s string) (Prime, error) {
	if val, ok := _PrimeNameToValue_map[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Prime values", s)
}

func (i Prime) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

func (i *Prime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("Prime should be a string, got %s", data)
	}

	var err error
	*i, err = PrimeString(s)
	return err
}

func (i Prime) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *Prime) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	str, ok := value.(string)
	if !ok {
		bytes, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("value is not a byte slice")
		}

		str = string(bytes[:])
	}

	val, err := PrimeString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
`

func TestGolden(t *testing.T) {
	for _, test := range golden {
		runGoldenTest(t, test, false, false, false)
	}
	for _, test := range goldenJSON {
		runGoldenTest(t, test, true, false, false)
	}
	for _, test := range goldenYAML {
		runGoldenTest(t, test, false, true, false)
	}
	for _, test := range goldenSQL {
		runGoldenTest(t, test, false, false, true)
	}
	for _, test := range goldenJSONAndSQL {
		runGoldenTest(t, test, true, false, true)
	}
}

func runGoldenTest(t *testing.T, test Golden, generateJSON, generateYAML, generateSQL bool) {
	var g Generator
	input := "package test\n" + test.input
	file := test.name + ".go"
	g.parsePackage(".", []string{file}, input)
	// Extract the name and type of the constant from the first line.
	tokens := strings.SplitN(test.input, " ", 3)
	if len(tokens) != 3 {
		t.Fatalf("%s: need type declaration on first line", test.name)
	}
	g.generate(tokens[1], generateJSON, generateYAML, generateSQL, "noop", "", false)
	got := string(g.format())
	if got != test.output {
		t.Errorf("%s: got\n====\n%s====\nexpected\n====%s", test.name, got, test.output)
	}
}
