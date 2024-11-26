package tickets

import (
	"fmt"
	"os/user"

	"github.com/zukigit/remote_run-go/src/common"
	"github.com/zukigit/remote_run-go/src/dao"
	"github.com/zukigit/remote_run-go/src/lib"
)

type Ticket_000 struct {
	Ticket_no                                   uint
	Ticket_description                          string
	PASSED_count, FAILED_count, MUSTCHECK_count int
	Testcases                                   []dao.TestCase
}

func (t *Ticket_000) New_testcase(testcase_id uint, testcase_description string) *dao.TestCase {
	return dao.New_testcase(testcase_id, testcase_description)
}

func (t *Ticket_000) Get_no() uint {
	return t.Ticket_no
}

func (t *Ticket_000) Set_PASSED_count(passed_count int) {
	t.PASSED_count = passed_count
}

func (t *Ticket_000) Set_FAILED_count(failed_count int) {
	t.FAILED_count = failed_count
}

func (t *Ticket_000) Set_MUSTCHECK_count(mustcheck_count int) {
	t.MUSTCHECK_count = mustcheck_count
}

func (t *Ticket_000) Get_dsctn() string {
	return t.Ticket_description
}

func (t *Ticket_000) Add_testcase(tc dao.TestCase) {
	t.Testcases = append(t.Testcases, tc)
}

func (t *Ticket_000) Get_testcases() []dao.TestCase {
	return t.Testcases
}

// Enter your ticket information here
func (t *Ticket_000) Set_values() {
	t.Ticket_no = 1 // Enter your ticket id
	t.Ticket_description = "Enter your ticket description here."
}

// Add your test case here
func (t *Ticket_000) Add_testcases() {
	// TESTCASE 001
	tc_1 := t.New_testcase(1, "Enter your test case description here.")
	tc_func := func() common.Testcase_status {
		tc_1.Add_doc(common.PRE_OPT, "tested preopt")
		tc_1.Add_doc(common.OPT, "tested opt")
		tc_1.Add_doc(common.EXPT_RES, "tested expt_result")
		tc_1.Err_log("tested error log")
		lib.Logi(common.LOG_LEVEL_INFO, "testing new log function")
		lib.Logi(common.LOG_LEVEL_ERR, "testing new log function")

		currentUser, err := user.Current()
		if err != nil {
			fmt.Printf("Error fetching current user: %v\n", err)
			return FAILED
		}

		// Print the username and user details
		fmt.Printf("Current Username: %s\n", currentUser.Username)
		fmt.Printf("User UID: %s\n", currentUser.Uid)
		fmt.Printf("Home Directory: %s\n", currentUser.HomeDir)

		return PASSED
	}
	tc_1.Set_function(tc_func)
	t.Add_testcase(*tc_1)

	// TESTCASE 002
	tc_2 := t.New_testcase(2, "Enter your test case description here.")
	tc_func = func() common.Testcase_status {
		tc_1.Info_log("tested info log")
		return MUST_CHECK
	}
	tc_2.Set_function(tc_func)
	t.Add_testcase(*tc_2)

	// TESTCASE 003
	tc_3 := t.New_testcase(3, "Enter your test case description here.")
	tc_func = func() common.Testcase_status {
		return MUST_CHECK
	}
	tc_3.Set_function(tc_func)
	t.Add_testcase(*tc_3)
}
