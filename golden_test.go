// Copyright 2014 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains simple golden tests for various examples.
// Besides validating the results when the implementation changes,
// it provides a way to look at the generated code without having
// to execute the print statements in one's head.

package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// Golden represents a test case.
type Golden struct {
	name   string
	input  string // input; the package clause is provided when running the test.
	output string // expected output.
}

var golden = []Golden{
	{"day", dayIn, dayOut},
	{"offset", offsetIn, offsetOut},
	{"gap", gapIn, gapOut},
	{"num", numIn, numOut},
	{"unum", unumIn, unumOut},
	{"prime", primeIn, primeOut},
}

var goldenJSON = []Golden{
	{"prime with JSON", primeJsonIn, primeJsonOut},
}
var goldenText = []Golden{
	{"prime with Text", primeTextIn, primeTextOut},
}

var goldenYAML = []Golden{
	{"prime with YAML", primeYamlIn, primeYamlOut},
}

var goldenSQL = []Golden{
	{"prime with SQL", primeSqlIn, primeSqlOut},
}

var goldenGQL = []Golden{
	{"prime with GQL", primeGqlIn, primeGqlOut},
}

var goldenJSONAndSQL = []Golden{
	{"prime with JSONAndSQL", primeJsonAndSqlIn, primeJsonAndSqlOut},
}

var goldenPrefix = []Golden{
	{"prefix", prefixIn, dayOut},
}

var goldenWithLineComments = []Golden{
	{"primer with line Comments", primeWithLineCommentIn, primeWithLineCommentOut},
}

// Each example starts with "type XXX [u]int", with a single space separating them.

// Simple test: enumeration of type int starting at 0.
const dayIn = `type Day int
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

const dayOut = `
const _DayName = "MondayTuesdayWednesdayThursdayFridaySaturdaySunday"

var _DayIndex = [...]uint8{0, 6, 13, 22, 30, 36, 44, 50}

func (i Day) String() string {
	if i < 0 || i >= Day(len(_DayIndex)-1) {
		return fmt.Sprintf("Day(%d)", i)
	}
	return _DayName[_DayIndex[i]:_DayIndex[i+1]]
}

var _DayValues = []Day{0, 1, 2, 3, 4, 5, 6}

var _DayNameToValueMap = map[string]Day{
	_DayName[0:6]:   0,
	_DayName[6:13]:  1,
	_DayName[13:22]: 2,
	_DayName[22:30]: 3,
	_DayName[30:36]: 4,
	_DayName[36:44]: 5,
	_DayName[44:50]: 6,
}

// DayString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func DayString(s string) (Day, error) {
	if val, ok := _DayNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Day values", s)
}

// DayValues returns all values of the enum
func DayValues() []Day {
	return _DayValues
}

// IsADay returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Day) IsADay() bool {
	for _, v := range _DayValues {
		if i == v {
			return true
		}
	}
	return false
}
`

// Enumeration with an offset.
// Also includes a duplicate.
const offsetIn = `type Number int
const (
	_ Number = iota
	One
	Two
	Three
	AnotherOne = One  // Duplicate; note that AnotherOne doesn't appear below.
)
`

const offsetOut = `
const _NumberName = "OneTwoThree"

var _NumberIndex = [...]uint8{0, 3, 6, 11}

func (i Number) String() string {
	i -= 1
	if i < 0 || i >= Number(len(_NumberIndex)-1) {
		return fmt.Sprintf("Number(%d)", i+1)
	}
	return _NumberName[_NumberIndex[i]:_NumberIndex[i+1]]
}

var _NumberValues = []Number{1, 2, 3}

var _NumberNameToValueMap = map[string]Number{
	_NumberName[0:3]:  1,
	_NumberName[3:6]:  2,
	_NumberName[6:11]: 3,
}

// NumberString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func NumberString(s string) (Number, error) {
	if val, ok := _NumberNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Number values", s)
}

// NumberValues returns all values of the enum
func NumberValues() []Number {
	return _NumberValues
}

// IsANumber returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Number) IsANumber() bool {
	for _, v := range _NumberValues {
		if i == v {
			return true
		}
	}
	return false
}
`

// Gaps and an offset.
const gapIn = `type Gap int
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

const gapOut = `
const (
	_GapName_0 = "TwoThree"
	_GapName_1 = "FiveSixSevenEightNine"
	_GapName_2 = "Eleven"
)

var (
	_GapIndex_0 = [...]uint8{0, 3, 8}
	_GapIndex_1 = [...]uint8{0, 4, 7, 12, 17, 21}
	_GapIndex_2 = [...]uint8{0, 6}
)

func (i Gap) String() string {
	switch {
	case 2 <= i && i <= 3:
		i -= 2
		return _GapName_0[_GapIndex_0[i]:_GapIndex_0[i+1]]
	case 5 <= i && i <= 9:
		i -= 5
		return _GapName_1[_GapIndex_1[i]:_GapIndex_1[i+1]]
	case i == 11:
		return _GapName_2
	default:
		return fmt.Sprintf("Gap(%d)", i)
	}
}

var _GapValues = []Gap{2, 3, 5, 6, 7, 8, 9, 11}

var _GapNameToValueMap = map[string]Gap{
	_GapName_0[0:3]:   2,
	_GapName_0[3:8]:   3,
	_GapName_1[0:4]:   5,
	_GapName_1[4:7]:   6,
	_GapName_1[7:12]:  7,
	_GapName_1[12:17]: 8,
	_GapName_1[17:21]: 9,
	_GapName_2[0:6]:   11,
}

// GapString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func GapString(s string) (Gap, error) {
	if val, ok := _GapNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Gap values", s)
}

// GapValues returns all values of the enum
func GapValues() []Gap {
	return _GapValues
}

// IsAGap returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Gap) IsAGap() bool {
	for _, v := range _GapValues {
		if i == v {
			return true
		}
	}
	return false
}
`

// Signed integers spanning zero.
const numIn = `type Num int
const (
	m_2 Num = -2 + iota
	m_1
	m0
	m1
	m2
)
`

const numOut = `
const _NumName = "m_2m_1m0m1m2"

var _NumIndex = [...]uint8{0, 3, 6, 8, 10, 12}

func (i Num) String() string {
	i -= -2
	if i < 0 || i >= Num(len(_NumIndex)-1) {
		return fmt.Sprintf("Num(%d)", i+-2)
	}
	return _NumName[_NumIndex[i]:_NumIndex[i+1]]
}

var _NumValues = []Num{-2, -1, 0, 1, 2}

var _NumNameToValueMap = map[string]Num{
	_NumName[0:3]:   -2,
	_NumName[3:6]:   -1,
	_NumName[6:8]:   0,
	_NumName[8:10]:  1,
	_NumName[10:12]: 2,
}

// NumString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func NumString(s string) (Num, error) {
	if val, ok := _NumNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Num values", s)
}

// NumValues returns all values of the enum
func NumValues() []Num {
	return _NumValues
}

// IsANum returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Num) IsANum() bool {
	for _, v := range _NumValues {
		if i == v {
			return true
		}
	}
	return false
}
`

// Unsigned integers spanning zero.
const unumIn = `type Unum uint
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

const unumOut = `
const (
	_UnumName_0 = "m0m1m2"
	_UnumName_1 = "m_2m_1"
)

var (
	_UnumIndex_0 = [...]uint8{0, 2, 4, 6}
	_UnumIndex_1 = [...]uint8{0, 3, 6}
)

func (i Unum) String() string {
	switch {
	case 0 <= i && i <= 2:
		return _UnumName_0[_UnumIndex_0[i]:_UnumIndex_0[i+1]]
	case 253 <= i && i <= 254:
		i -= 253
		return _UnumName_1[_UnumIndex_1[i]:_UnumIndex_1[i+1]]
	default:
		return fmt.Sprintf("Unum(%d)", i)
	}
}

var _UnumValues = []Unum{0, 1, 2, 253, 254}

var _UnumNameToValueMap = map[string]Unum{
	_UnumName_0[0:2]: 0,
	_UnumName_0[2:4]: 1,
	_UnumName_0[4:6]: 2,
	_UnumName_1[0:3]: 253,
	_UnumName_1[3:6]: 254,
}

// UnumString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func UnumString(s string) (Unum, error) {
	if val, ok := _UnumNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Unum values", s)
}

// UnumValues returns all values of the enum
func UnumValues() []Unum {
	return _UnumValues
}

// IsAUnum returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Unum) IsAUnum() bool {
	for _, v := range _UnumValues {
		if i == v {
			return true
		}
	}
	return false
}
`

// Enough gaps to trigger a map implementation of the method.
// Also includes a duplicate to test that it doesn't cause problems
const primeIn = `type Prime int
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

const primeOut = `
const _PrimeName = "p2p3p5p7p11p13p17p19p23p29p37p41p43"

var _PrimeMap = map[Prime]string{
	2:  _PrimeName[0:2],
	3:  _PrimeName[2:4],
	5:  _PrimeName[4:6],
	7:  _PrimeName[6:8],
	11: _PrimeName[8:11],
	13: _PrimeName[11:14],
	17: _PrimeName[14:17],
	19: _PrimeName[17:20],
	23: _PrimeName[20:23],
	29: _PrimeName[23:26],
	31: _PrimeName[26:29],
	41: _PrimeName[29:32],
	43: _PrimeName[32:35],
}

func (i Prime) String() string {
	if str, ok := _PrimeMap[i]; ok {
		return str
	}
	return fmt.Sprintf("Prime(%d)", i)
}

var _PrimeValues = []Prime{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43}

var _PrimeNameToValueMap = map[string]Prime{
	_PrimeName[0:2]:   2,
	_PrimeName[2:4]:   3,
	_PrimeName[4:6]:   5,
	_PrimeName[6:8]:   7,
	_PrimeName[8:11]:  11,
	_PrimeName[11:14]: 13,
	_PrimeName[14:17]: 17,
	_PrimeName[17:20]: 19,
	_PrimeName[20:23]: 23,
	_PrimeName[23:26]: 29,
	_PrimeName[26:29]: 31,
	_PrimeName[29:32]: 41,
	_PrimeName[32:35]: 43,
}

// PrimeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PrimeString(s string) (Prime, error) {
	if val, ok := _PrimeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Prime values", s)
}

// PrimeValues returns all values of the enum
func PrimeValues() []Prime {
	return _PrimeValues
}

// IsAPrime returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Prime) IsAPrime() bool {
	_, ok := _PrimeMap[i]
	return ok
}
`
const primeJsonIn = `type Prime int
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

const primeJsonOut = `
const _PrimeName = "p2p3p5p7p11p13p17p19p23p29p37p41p43"

var _PrimeMap = map[Prime]string{
	2:  _PrimeName[0:2],
	3:  _PrimeName[2:4],
	5:  _PrimeName[4:6],
	7:  _PrimeName[6:8],
	11: _PrimeName[8:11],
	13: _PrimeName[11:14],
	17: _PrimeName[14:17],
	19: _PrimeName[17:20],
	23: _PrimeName[20:23],
	29: _PrimeName[23:26],
	31: _PrimeName[26:29],
	41: _PrimeName[29:32],
	43: _PrimeName[32:35],
}

func (i Prime) String() string {
	if str, ok := _PrimeMap[i]; ok {
		return str
	}
	return fmt.Sprintf("Prime(%d)", i)
}

var _PrimeValues = []Prime{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43}

var _PrimeNameToValueMap = map[string]Prime{
	_PrimeName[0:2]:   2,
	_PrimeName[2:4]:   3,
	_PrimeName[4:6]:   5,
	_PrimeName[6:8]:   7,
	_PrimeName[8:11]:  11,
	_PrimeName[11:14]: 13,
	_PrimeName[14:17]: 17,
	_PrimeName[17:20]: 19,
	_PrimeName[20:23]: 23,
	_PrimeName[23:26]: 29,
	_PrimeName[26:29]: 31,
	_PrimeName[29:32]: 41,
	_PrimeName[32:35]: 43,
}

// PrimeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PrimeString(s string) (Prime, error) {
	if val, ok := _PrimeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Prime values", s)
}

// PrimeValues returns all values of the enum
func PrimeValues() []Prime {
	return _PrimeValues
}

// IsAPrime returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Prime) IsAPrime() bool {
	_, ok := _PrimeMap[i]
	return ok
}

// MarshalJSON implements the json.Marshaler interface for Prime
func (i Prime) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Prime
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

const primeTextIn = `type Prime int
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

const primeTextOut = `
const _PrimeName = "p2p3p5p7p11p13p17p19p23p29p37p41p43"

var _PrimeMap = map[Prime]string{
	2:  _PrimeName[0:2],
	3:  _PrimeName[2:4],
	5:  _PrimeName[4:6],
	7:  _PrimeName[6:8],
	11: _PrimeName[8:11],
	13: _PrimeName[11:14],
	17: _PrimeName[14:17],
	19: _PrimeName[17:20],
	23: _PrimeName[20:23],
	29: _PrimeName[23:26],
	31: _PrimeName[26:29],
	41: _PrimeName[29:32],
	43: _PrimeName[32:35],
}

func (i Prime) String() string {
	if str, ok := _PrimeMap[i]; ok {
		return str
	}
	return fmt.Sprintf("Prime(%d)", i)
}

var _PrimeValues = []Prime{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43}

var _PrimeNameToValueMap = map[string]Prime{
	_PrimeName[0:2]:   2,
	_PrimeName[2:4]:   3,
	_PrimeName[4:6]:   5,
	_PrimeName[6:8]:   7,
	_PrimeName[8:11]:  11,
	_PrimeName[11:14]: 13,
	_PrimeName[14:17]: 17,
	_PrimeName[17:20]: 19,
	_PrimeName[20:23]: 23,
	_PrimeName[23:26]: 29,
	_PrimeName[26:29]: 31,
	_PrimeName[29:32]: 41,
	_PrimeName[32:35]: 43,
}

// PrimeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PrimeString(s string) (Prime, error) {
	if val, ok := _PrimeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Prime values", s)
}

// PrimeValues returns all values of the enum
func PrimeValues() []Prime {
	return _PrimeValues
}

// IsAPrime returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Prime) IsAPrime() bool {
	_, ok := _PrimeMap[i]
	return ok
}

// MarshalText implements the encoding.TextMarshaler interface for Prime
func (i Prime) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for Prime
func (i *Prime) UnmarshalText(text []byte) error {
	var err error
	*i, err = PrimeString(string(text))
	return err
}
`

const primeYamlIn = `type Prime int
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

const primeYamlOut = `
const _PrimeName = "p2p3p5p7p11p13p17p19p23p29p37p41p43"

var _PrimeMap = map[Prime]string{
	2:  _PrimeName[0:2],
	3:  _PrimeName[2:4],
	5:  _PrimeName[4:6],
	7:  _PrimeName[6:8],
	11: _PrimeName[8:11],
	13: _PrimeName[11:14],
	17: _PrimeName[14:17],
	19: _PrimeName[17:20],
	23: _PrimeName[20:23],
	29: _PrimeName[23:26],
	31: _PrimeName[26:29],
	41: _PrimeName[29:32],
	43: _PrimeName[32:35],
}

func (i Prime) String() string {
	if str, ok := _PrimeMap[i]; ok {
		return str
	}
	return fmt.Sprintf("Prime(%d)", i)
}

var _PrimeValues = []Prime{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43}

var _PrimeNameToValueMap = map[string]Prime{
	_PrimeName[0:2]:   2,
	_PrimeName[2:4]:   3,
	_PrimeName[4:6]:   5,
	_PrimeName[6:8]:   7,
	_PrimeName[8:11]:  11,
	_PrimeName[11:14]: 13,
	_PrimeName[14:17]: 17,
	_PrimeName[17:20]: 19,
	_PrimeName[20:23]: 23,
	_PrimeName[23:26]: 29,
	_PrimeName[26:29]: 31,
	_PrimeName[29:32]: 41,
	_PrimeName[32:35]: 43,
}

// PrimeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PrimeString(s string) (Prime, error) {
	if val, ok := _PrimeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Prime values", s)
}

// PrimeValues returns all values of the enum
func PrimeValues() []Prime {
	return _PrimeValues
}

// IsAPrime returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Prime) IsAPrime() bool {
	_, ok := _PrimeMap[i]
	return ok
}

// MarshalYAML implements a YAML Marshaler for Prime
func (i Prime) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for Prime
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

const primeSqlIn = `type Prime int
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

const primeSqlOut = `
const _PrimeName = "p2p3p5p7p11p13p17p19p23p29p37p41p43"

var _PrimeMap = map[Prime]string{
	2:  _PrimeName[0:2],
	3:  _PrimeName[2:4],
	5:  _PrimeName[4:6],
	7:  _PrimeName[6:8],
	11: _PrimeName[8:11],
	13: _PrimeName[11:14],
	17: _PrimeName[14:17],
	19: _PrimeName[17:20],
	23: _PrimeName[20:23],
	29: _PrimeName[23:26],
	31: _PrimeName[26:29],
	41: _PrimeName[29:32],
	43: _PrimeName[32:35],
}

func (i Prime) String() string {
	if str, ok := _PrimeMap[i]; ok {
		return str
	}
	return fmt.Sprintf("Prime(%d)", i)
}

var _PrimeValues = []Prime{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43}

var _PrimeNameToValueMap = map[string]Prime{
	_PrimeName[0:2]:   2,
	_PrimeName[2:4]:   3,
	_PrimeName[4:6]:   5,
	_PrimeName[6:8]:   7,
	_PrimeName[8:11]:  11,
	_PrimeName[11:14]: 13,
	_PrimeName[14:17]: 17,
	_PrimeName[17:20]: 19,
	_PrimeName[20:23]: 23,
	_PrimeName[23:26]: 29,
	_PrimeName[26:29]: 31,
	_PrimeName[29:32]: 41,
	_PrimeName[32:35]: 43,
}

// PrimeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PrimeString(s string) (Prime, error) {
	if val, ok := _PrimeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Prime values", s)
}

// PrimeValues returns all values of the enum
func PrimeValues() []Prime {
	return _PrimeValues
}

// IsAPrime returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Prime) IsAPrime() bool {
	_, ok := _PrimeMap[i]
	return ok
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
const primeGqlIn = `type Prime int
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

const primeGqlOut = `
const _PrimeName = "p2p3p5p7p11p13p17p19p23p29p37p41p43"

var _PrimeMap = map[Prime]string{
	2:  _PrimeName[0:2],
	3:  _PrimeName[2:4],
	5:  _PrimeName[4:6],
	7:  _PrimeName[6:8],
	11: _PrimeName[8:11],
	13: _PrimeName[11:14],
	17: _PrimeName[14:17],
	19: _PrimeName[17:20],
	23: _PrimeName[20:23],
	29: _PrimeName[23:26],
	31: _PrimeName[26:29],
	41: _PrimeName[29:32],
	43: _PrimeName[32:35],
}

func (i Prime) String() string {
	if str, ok := _PrimeMap[i]; ok {
		return str
	}
	return fmt.Sprintf("Prime(%d)", i)
}

var _PrimeValues = []Prime{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43}

var _PrimeNameToValueMap = map[string]Prime{
	_PrimeName[0:2]:   2,
	_PrimeName[2:4]:   3,
	_PrimeName[4:6]:   5,
	_PrimeName[6:8]:   7,
	_PrimeName[8:11]:  11,
	_PrimeName[11:14]: 13,
	_PrimeName[14:17]: 17,
	_PrimeName[17:20]: 19,
	_PrimeName[20:23]: 23,
	_PrimeName[23:26]: 29,
	_PrimeName[26:29]: 31,
	_PrimeName[29:32]: 41,
	_PrimeName[32:35]: 43,
}

// PrimeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PrimeString(s string) (Prime, error) {
	if val, ok := _PrimeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Prime values", s)
}

// PrimeValues returns all values of the enum
func PrimeValues() []Prime {
	return _PrimeValues
}

// IsAPrime returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Prime) IsAPrime() bool {
	_, ok := _PrimeMap[i]
	return ok
}

func (i *Prime) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		bytes, ok := v.([]byte)
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

func (i Prime) MarshalGQL(w io.Writer) {
	_, _ = w.Write([]byte(strconv.Quote(f.String())))
}
`

const primeJsonAndSqlIn = `type Prime int
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

const primeJsonAndSqlOut = `
const _PrimeName = "p2p3p5p7p11p13p17p19p23p29p37p41p43"

var _PrimeMap = map[Prime]string{
	2:  _PrimeName[0:2],
	3:  _PrimeName[2:4],
	5:  _PrimeName[4:6],
	7:  _PrimeName[6:8],
	11: _PrimeName[8:11],
	13: _PrimeName[11:14],
	17: _PrimeName[14:17],
	19: _PrimeName[17:20],
	23: _PrimeName[20:23],
	29: _PrimeName[23:26],
	31: _PrimeName[26:29],
	41: _PrimeName[29:32],
	43: _PrimeName[32:35],
}

func (i Prime) String() string {
	if str, ok := _PrimeMap[i]; ok {
		return str
	}
	return fmt.Sprintf("Prime(%d)", i)
}

var _PrimeValues = []Prime{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43}

var _PrimeNameToValueMap = map[string]Prime{
	_PrimeName[0:2]:   2,
	_PrimeName[2:4]:   3,
	_PrimeName[4:6]:   5,
	_PrimeName[6:8]:   7,
	_PrimeName[8:11]:  11,
	_PrimeName[11:14]: 13,
	_PrimeName[14:17]: 17,
	_PrimeName[17:20]: 19,
	_PrimeName[20:23]: 23,
	_PrimeName[23:26]: 29,
	_PrimeName[26:29]: 31,
	_PrimeName[29:32]: 41,
	_PrimeName[32:35]: 43,
}

// PrimeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PrimeString(s string) (Prime, error) {
	if val, ok := _PrimeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Prime values", s)
}

// PrimeValues returns all values of the enum
func PrimeValues() []Prime {
	return _PrimeValues
}

// IsAPrime returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Prime) IsAPrime() bool {
	_, ok := _PrimeMap[i]
	return ok
}

// MarshalJSON implements the json.Marshaler interface for Prime
func (i Prime) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for Prime
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

const prefixIn = `type Day int
const (
	DayMonday Day = iota
	DayTuesday
	DayWednesday
	DayThursday
	DayFriday
	DaySaturday
	DaySunday
)
`

const primeWithLineCommentIn = `type Prime int
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

const primeWithLineCommentOut = `
const _PrimeName = "p2p3GoodPrimep7p11p13p17p19p23p29p37TwinPrime41Twin prime 43"

var _PrimeMap = map[Prime]string{
	2:  _PrimeName[0:2],
	3:  _PrimeName[2:4],
	5:  _PrimeName[4:13],
	7:  _PrimeName[13:15],
	11: _PrimeName[15:18],
	13: _PrimeName[18:21],
	17: _PrimeName[21:24],
	19: _PrimeName[24:27],
	23: _PrimeName[27:30],
	29: _PrimeName[30:33],
	31: _PrimeName[33:36],
	41: _PrimeName[36:47],
	43: _PrimeName[47:60],
}

func (i Prime) String() string {
	if str, ok := _PrimeMap[i]; ok {
		return str
	}
	return fmt.Sprintf("Prime(%d)", i)
}

var _PrimeValues = []Prime{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43}

var _PrimeNameToValueMap = map[string]Prime{
	_PrimeName[0:2]:   2,
	_PrimeName[2:4]:   3,
	_PrimeName[4:13]:  5,
	_PrimeName[13:15]: 7,
	_PrimeName[15:18]: 11,
	_PrimeName[18:21]: 13,
	_PrimeName[21:24]: 17,
	_PrimeName[24:27]: 19,
	_PrimeName[27:30]: 23,
	_PrimeName[30:33]: 29,
	_PrimeName[33:36]: 31,
	_PrimeName[36:47]: 41,
	_PrimeName[47:60]: 43,
}

// PrimeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PrimeString(s string) (Prime, error) {
	if val, ok := _PrimeNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Prime values", s)
}

// PrimeValues returns all values of the enum
func PrimeValues() []Prime {
	return _PrimeValues
}

// IsAPrime returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Prime) IsAPrime() bool {
	_, ok := _PrimeMap[i]
	return ok
}
`

func TestGolden(t *testing.T) {
	for _, test := range golden {
		runGoldenTest(t, test, false, false, false, false, false, "")
	}
	for _, test := range goldenJSON {
		runGoldenTest(t, test, true, false, false, false, false, "")
	}
	for _, test := range goldenText {
		runGoldenTest(t, test, false, false, false, false, true, "")
	}
	for _, test := range goldenYAML {
		runGoldenTest(t, test, false, true, false, false, false, "")
	}
	for _, test := range goldenSQL {
		runGoldenTest(t, test, false, false, true, false, false, "")
	}
	for _, test := range goldenGQL {
		runGoldenTest(t, test, false, false, false, true, false, "")
	}
	for _, test := range goldenJSONAndSQL {
		runGoldenTest(t, test, true, false, true, false, false, "")
	}
	for _, test := range goldenPrefix {
		runGoldenTest(t, test, false, false, false, false, false, "Day")
	}
}

func runGoldenTest(t *testing.T, test Golden, generateJSON, generateYAML, generateSQL, generateGQL, generateText bool, prefix string) {
	var g Generator
	input := "package test\n" + test.input
	file := test.name + ".go"

	dir, err := ioutil.TempDir("", "stringer")
	if err != nil {
		t.Error(err)
	}
	defer func() {
		err = os.RemoveAll(dir)
		if err != nil {
			t.Error(err)
		}
	}()

	absFile := filepath.Join(dir, file)
	err = ioutil.WriteFile(absFile, []byte(input), 0644)
	if err != nil {
		t.Error(err)
	}
	g.parsePackage([]string{absFile})
	// Extract the name and type of the constant from the first line.
	tokens := strings.SplitN(test.input, " ", 3)
	if len(tokens) != 3 {
		t.Fatalf("%s: need type declaration on first line", test.name)
	}
	g.generate(tokens[1], generateJSON, generateYAML, generateSQL, generateGQL, generateText, "noop", prefix, false)
	got := string(g.format())
	if got != test.output {
		t.Errorf("%s: got\n====\n%s====\nexpected\n====%s", test.name, got, test.output)
	}
}
