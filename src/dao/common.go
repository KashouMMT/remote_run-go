package dao

import "github.com/zukigit/remote_run-go/src/common"

const (
	PASSED     common.Testcase_status = "PASSED"
	FAILED     common.Testcase_status = "FAILED"
	MUST_CHECK common.Testcase_status = "MUST_CHECK"
)

var Tc_failed_cnt, Tc_passed_cnt, Tc_chk_cnt, Tc_unkown_cnt int

func Set_total_tc_results(t Ticket) {
	Tc_failed_cnt = 0
	Tc_chk_cnt = 0
	Tc_passed_cnt = 0
	Tc_unkown_cnt = 0

	for _, tc := range t.Get_testcases() {
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
}
