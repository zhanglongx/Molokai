package runner

import (
	"encoding/json"
	"errors"
)

type Runner struct {
	Symbol string

	// Runner name
	Name  string
	Param map[string]interface{}
}

type runnerOp interface {
	// if a runner is valid
	Valid() bool

	// trigger a runner to check
	Check() (Result, error)
}

type Result struct {
	IsShouldRemind bool
	Message        string
}

var (
	errRunnerNotFound = errors.New("runner not found")
	errRunnerNotValid = errors.New("runner not valid")
)

// Hacks: turn runner cfg -> runner struct dynamically.
//		  see mapToStruct()
var runnerMap = map[string]runnerOp{
	"MinMax": &minMax{},
}

func (r *Runner) Run() (Result, error) {
	if runnerMap[r.Name] == nil {
		return Result{}, errRunnerNotFound
	}

	r.Param["Symbol"] = r.Symbol

	runner := runnerMap[r.Name]
	if err := mapToStruct(r.Param, runner); err != nil {
		return Result{}, err
	}

	if !runner.Valid() {
		return Result{}, errRunnerNotValid
	}

	return runner.Check()
}

// turn a map into struct, use json.Marshal and json.Unmarshal
func mapToStruct(param map[string]interface{}, s interface{}) error {
	tmp, err := json.Marshal(param)
	if err != nil {
		return err
	}

	err = json.Unmarshal(tmp, s)
	if err != nil {
		return err
	}

	return nil
}
