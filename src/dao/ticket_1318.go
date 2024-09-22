package dao

import (
	"fmt"
)

type Ticket_1318 struct {
	no          uint
	description string
	auth        *Auth
	testcases   []TestCase
}

func (t *Ticket_1318) Set_values(auth *Auth) {
	t.no = 1318
	t.description = "Fixed for negative JOB_EXT_CD return value."
	t.auth = auth
}

func (t *Ticket_1318) New_testcase(testcase_id uint, testcase_description string) *TestCase {
	return New_testcase(testcase_id, testcase_description, t.auth)
}

func (t *Ticket_1318) Get_no() uint {
	return t.no
}

func (t *Ticket_1318) Get_dsctn() string {
	return t.description
}

func (t *Ticket_1318) Set_testcase(tc TestCase) {
	t.testcases = append(t.testcases, tc)
}

func (t *Ticket_1318) Get_testcases() []TestCase {
	return t.testcases
}

func (t *Ticket_1318) Run() {
	for _, tc := range t.testcases {
		fmt.Println(tc.Info_log("running..."))
		if !tc.Is_function_nil() {
			tc.Set_is_passed(tc.Run_function())
		} else {
			fmt.Println(tc.Err_log("has no function. SKIPPED!"))
			tc.Set_is_passed(false)
		}

		fmt.Println(tc.Info_log("finished!"))
	}
}

func (t *Ticket_1318) Add_testcases() {
	// Add your test case here
	// ticket 168
	tc_168 := t.New_testcase(168, "Normal Case")
	tc_func := func() bool {
		registry_number, error := tc_168.Jobarg_exec("JOBNET_1")
		if error != nil {
			tc_168.Err_log("Error: %s", error.Error())
			return false
		}

		jobnet_info, error := tc_168.Jobarg_get_jobnet_run_info(registry_number)
		if error != nil {
			tc_168.Err_log("Error: %s", error.Error())
			return false
		}

		return jobnet_info.Status == "END" && jobnet_info.Job_status == "NORMAL"
	}
	tc_168.Set_function(tc_func)
	t.Set_testcase(*tc_168)

	// ticket 169
	tc_169 := t.New_testcase(169, "Abnormal Case") // create test case
	t.Set_testcase(*tc_169)                        // Add testcase to ticket
}
