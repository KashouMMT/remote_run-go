package dao

import (
	"fmt"

	"github.com/zukigit/remote_run-go/src/common"
)

const (
	PASSED     common.Testcase_status = "PASSED"
	FAILED     common.Testcase_status = "FAILED"
	MUST_CHECK common.Testcase_status = "MUST_CHECK"
)

var Tc_failed_cnt, Tc_passed_cnt, Tc_chk_cnt, Tc_unkown_cnt int

func Set_total_tc_results_(tc TestCase) {
	switch tc.Get_status() {
	case PASSED:
		Tc_passed_cnt++
	case FAILED:
		Tc_failed_cnt++
	case MUST_CHECK:
		Tc_chk_cnt++
	default:
		Tc_unkown_cnt++
	}
}

func Set_total_tc_results(t Ticket) {
	Tc_failed_cnt = 0
	Tc_chk_cnt = 0
	Tc_passed_cnt = 0
	Tc_unkown_cnt = 0
	for _, tc := range t.Get_testcases() {
		if common.Specific_testcase_no == 0 {
			Set_total_tc_results_(tc)
		} else if common.Specific_testcase_no == tc.Get_id() {
			Set_total_tc_results_(tc)
		}
	}
}

func Run_testcase_(tc TestCase) {
	fmt.Println(tc.Info_log("running..."))
	if !tc.Is_function_nil() {
		tc.Set_status(tc.Run_function())
	} else {
		fmt.Println(tc.Err_log("has no function. SKIPPED!"))
		tc.Set_status(FAILED)
	}
	fmt.Println(tc.Info_log("finished!"))

	tc.Write_log()
}

func Run_testcase(t Ticket) {
	for _, tc := range t.Get_testcases() {
		if common.Specific_testcase_no == 0 {
			Run_testcase_(tc)
		} else if common.Specific_testcase_no == tc.Get_id() {
			Run_testcase_(tc)
		}
	}
}
