package lib

import (
	"fmt"
	"reflect"

	"github.com/zukigit/remote_run-go/src/common"
)

type FunctionResult struct {
	ReturnResult bool
	ReturnValues []interface{}
}

var funcMap = map[string]interface{}{
	"Restart_jaz_agent_linux":    Restart_jaz_agent_linux,
	"Jobarg_enable_jobnet":       Jobarg_enable_jobnet,
	"Jobarg_exec":                Jobarg_exec,
	"Clear_linux_jaz_agent_log":  Clear_linux_jaz_agent_log,
	"Jobarg_cleanup_linux":       Jobarg_cleanup_linux,
	"Jobarg_exec_E":              Jobarg_exec_E,
	"Jobarg_get_jobnet_run_info": Jobarg_get_jobnet_run_info,
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
				fmt.Println(Logi(common.LOG_LEVEL_ERR, fmt.Sprintf("Error: %s() returned an error: %v", func_name, err)))
				return FunctionResult{ReturnResult: false}
			}
		}

		fmt.Println(Logi(common.LOG_LEVEL_INFO, fmt.Sprintf("Info: %s() call succeed.", func_name)))

		return FunctionResult{ReturnResult: true, ReturnValues: returnValues}

	} else {
		fmt.Println(Logi(common.LOG_LEVEL_ERR, fmt.Sprintf("Error: %s() not found.", func_name)))

		return FunctionResult{ReturnResult: false}
	}
}
