package runner

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/zhanglongx/Molokai/tsWrapper"
)

type Runner struct {
	// Runner TsCode. Same as core.TsCode
	// Normally, it will be assigned instead Unmarshal-ed
	TsCode tsWrapper.TsCode `yaml:"-"`

	// Runner name. It will be used to look up a runner
	Name string `yaml:"name"`

	// Runner's param
	Param map[string]interface{} `yaml:"param"`
}

type runnerOp interface {
	// if a runner is valid
	Valid() bool

	// launch a runner to check
	Check() (Result, error)
}

type Result struct {
	TsCode tsWrapper.TsCode

	// Used by runner to indict if Reminders should be launched
	IsShouldRemind bool

	// Reminder's Message
	Message string
}

var (
	errRunnerParamNil = errors.New("runner param is nil")
	errRunnerNotFound = errors.New("runner not found")
	errRunnerNotValid = errors.New("runner not valid")
)

func (r *Runner) Run() (Result, error) {
	if r.Param == nil {
		return Result{}, errRunnerParamNil
	}

	// XXX: TsCode is not in Molokai yaml, BUT is necessary for each runner
	r.Param["TsCode"] = r.TsCode

	runner, err := lookupRunner(r.Name)
	if err != nil {
		return Result{}, err
	}

	if err := mapToStruct(r.Param, runner); err != nil {
		return Result{}, err
	}

	if !runner.Valid() {
		return Result{}, errRunnerNotValid
	}

	m, err := runner.Check()
	if err != nil {
		return Result{}, err
	}

	return m, nil
}

// turn a map into struct, use json.Marshal and json.Unmarshal.
// s should be a pointer to the struct
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

// Hacks: turn runner name -> runner struct dynamically.
//		  see mapToStruct()
func lookupRunner(name string) (runnerOp, error) {
	switch name {
	case "MinMax":
		return &minMax{}, nil
	case "MA":
		return &ma{}, nil
	default:
		return nil, errRunnerNotFound
	}
}

func (r *Result) String() string {
	s, err := tsWrapper.SymbolName(string(r.TsCode), "name")
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s: %s", s, r.Message)
}
