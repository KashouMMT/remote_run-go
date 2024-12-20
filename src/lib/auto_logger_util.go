package lib

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/spf13/cast"
	"github.com/zukigit/remote_run-go/src/common"
)

type FunctionResult struct {
	ReturnResult bool
	ReturnValues []interface{}
}

func Run_function(func_name string, param ...interface{}) FunctionResult {

	if fn, exists := funcMap[func_name]; exists {
		fnValue := reflect.ValueOf(fn)
		if len(param) != fnValue.Type().NumIn() {
			fmt.Println(Logi(common.LOG_LEVEL_INFO, fmt.Sprintf("Error: Function %s() expects %d arguments but got %d", func_name, fnValue.Type().NumIn(), len(param))))
			return FunctionResult{ReturnResult: false}
		}

		args := make([]reflect.Value, len(param))
		for i, param := range param {
			args[i] = reflect.ValueOf(param)
		}

		results := fnValue.Call(args)

		returnValues := make([]interface{}, len(results))
		for i, result := range results {
			returnValues[i] = result.Interface()
		}

		for _, value := range returnValues {
			if err, ok := value.(error); ok && err != nil {
				fmt.Println(Logi(common.LOG_LEVEL_ERR, fmt.Sprintf("Error: %s() returned an error: %s", func_name, err.Error())))
				return FunctionResult{ReturnResult: false}
			}
		}

		log_string := fmt.Sprintf("Info: %s() called successfully. Parameter:", func_name)

		for iteration, parameter := range param {
			log_string += fmt.Sprintf(" (%d)%s = %s,", iteration+1, reflect.TypeOf(parameter).String(), CastToString(parameter))
		}

		fmt.Println(Logi(common.LOG_LEVEL_INFO, log_string))

		return FunctionResult{ReturnResult: true, ReturnValues: returnValues}

	} else {
		fmt.Println(Logi(common.LOG_LEVEL_ERR, fmt.Sprintf("Error: %s() not found.", func_name)))

		return FunctionResult{ReturnResult: false}
	}
}

func CastToStringMapString(input interface{}) map[string]string {

	value, err := cast.ToStringMapStringE(input)
	if err != nil {
		fmt.Println(Logi(common.LOG_LEVEL_ERR, "Error: Typecasting to map[string]string failed."))
		return nil
	}

	return value
}

func CastToStringArray(input []interface{}) []string {
	var stringArray []string
	for _, value := range input {
		stringArray = append(stringArray, CastToString(value))
	}
	return stringArray
}

func CastToString(input interface{}) string {
	if strings.Contains(reflect.TypeOf(input).String(), "[]") {
		return fmt.Sprintf("%v", input)
	}
	value, err := cast.ToStringE(input)
	if err != nil {
		value = "<TypeCast Error>"
	}
	return value
}
