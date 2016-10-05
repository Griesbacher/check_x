package check_x

import (
	"errors"
	"sort"
	"strings"
)

//State represents an nagioskind returncode
type State struct {
	name string
	code int
}

//StateFromInt creates an known state if code is 0-3, else a new State will be returned
func StateFromInt(code int) State {
	switch code {
	case 0:
		return OK
	case 1:
		return Warning
	case 2:
		return Critical
	case 3:
		return Unknown
	default:
		return State{code: code}
	}
}

//StateFromString creates an known state if string is "ok|warning|critical|unknown", else a new State will be returned
func StateFromString(name string) State {
	lowerName := strings.ToLower(name)
	switch lowerName {
	case "ok":
		return OK
	case "warning":
		return Warning
	case "critical":
		return Critical
	case "unknown":
		return Unknown
	default:
		return State{name: name}
	}
}

//String prints the name of the state
func (s State) String() string {
	return s.name
}

var (
	//OK - returncode: 0
	OK = State{name: "OK", code: 0}
	//Warning - returncode: 1
	Warning = State{name: "Warning", code: 1}
	//Critical - returncode: 2
	Critical = State{name: "Critical", code: 2}
	//Unknown - returncode: 3
	Unknown = State{name: "Unknown", code: 3}
)

//States is a list of state
type States []State

//ErrEmptyStates will be thrown if no State was added to the States array
var ErrEmptyStates = errors.New("The given States do not contain a State")

//Len for Sort interface
func (s States) Len() int {
	return len(s)
}

//Less for Sort interface
func (s States) Less(i, j int) bool {
	return s[i].code < s[j].code
}

//Swap for Sort interface
func (s States) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s States) getSorted() error {
	if len(s) == 0 {
		return ErrEmptyStates
	}
	if !sort.IsSorted(s) {
		sort.Sort(s)
	}
	return nil
}

//GetBest returns the best State from the array
func (s States) GetBest() (*State, error) {
	if err := s.getSorted(); err != nil {
		return nil, err
	}
	return &s[0], nil
}

//GetWorst returns the worst State from the array
func (s States) GetWorst() (*State, error) {
	if err := s.getSorted(); err != nil {
		return nil, err
	}
	return &s[len(s)-1], nil
}
