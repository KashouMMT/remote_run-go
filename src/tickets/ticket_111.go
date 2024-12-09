package tickets

import (
	"github.com/zukigit/remote_run-go/src/common"
	"github.com/zukigit/remote_run-go/src/lib"
)

type Ticket_111 struct {
	Ticket_no                                   uint
	Ticket_description                          string
	PASSED_count, FAILED_count, MUSTCHECK_count int
	Testcases                                   []common.TestCase
}

func (t *Ticket_111) New_testcase(testcase_id uint, testcase_description string) *common.TestCase {
	return common.New_testcase(testcase_id, testcase_description)
}

func (t *Ticket_111) Get_no() uint {
	return t.Ticket_no
}

func (t *Ticket_111) Set_PASSED_count(passed_count int) {
	t.PASSED_count = passed_count
}

func (t *Ticket_111) Set_FAILED_count(failed_count int) {
	t.FAILED_count = failed_count
}

func (t *Ticket_111) Set_MUSTCHECK_count(mustcheck_count int) {
	t.MUSTCHECK_count = mustcheck_count
}

func (t *Ticket_111) Get_dsctn() string {
	return t.Ticket_description
}

func (t *Ticket_111) Add_testcase(tc common.TestCase) {
	t.Testcases = append(t.Testcases, tc)
}

func (t *Ticket_111) Get_testcases() []common.TestCase {
	return t.Testcases
}

// Enter your ticket information here
func (t *Ticket_111) Set_values() {
	t.Ticket_no = 111 // Enter your ticket id
	t.Ticket_description = "Enter your ticket description here."
}

// Add your test case here
func (t *Ticket_111) Add_testcases() {
	// TESTCASE 001
	tc_1 := t.New_testcase(111, "Enter your test case description here.")
	tc_func := func() common.Testcase_status {

		var jobnet_run_info string

		if lib.Run_function("Jobarg_cleanup_linux").ReturnResult &&
			lib.Run_function("Jobarg_enable_jobnet", "SIMPLE_JOB", "SIMPLE_JOB").ReturnResult &&
			func() bool {
				returnValues := lib.Run_function("Jobarg_exec", "SIMPLE_JOB")
				result := returnValues.ReturnResult
				jobnet_run_info = lib.CastToString(returnValues.ReturnValues[0])
				return result
			}() &&
			lib.Run_function("Jobarg_get_jobnet_run_info", jobnet_run_info).ReturnResult &&
			lib.Run_function("Clear_linux_jaz_agent_log").ReturnResult {
			return PASSED
		}
		return FAILED
	}
	tc_1.Set_function(tc_func)
	t.Add_testcase(*tc_1)
}
