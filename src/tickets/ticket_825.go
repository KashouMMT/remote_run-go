package tickets

import (
	"context"
	"fmt"
	"time"

	"github.com/zukigit/remote_run-go/src/common"
	"github.com/zukigit/remote_run-go/src/dao"
	"github.com/zukigit/remote_run-go/src/lib"
)

type Ticket_825 struct {
	no          uint
	description string
	testcases   []dao.TestCase
}

func (t *Ticket_825) New_testcase(testcase_id uint, testcase_description string) *dao.TestCase {
	return dao.New_testcase(testcase_id, testcase_description)
}

func (t *Ticket_825) Get_no() uint {
	return t.no
}

func (t *Ticket_825) Get_dsctn() string {
	return t.description
}

func (t *Ticket_825) Add_testcase(tc dao.TestCase) {
	t.testcases = append(t.testcases, tc)
}

func (t *Ticket_825) Get_testcases() []dao.TestCase {
	return t.testcases
}

// Enter your ticket information here
func (t *Ticket_825) Set_values() {
	t.no = 825 // Enter your ticket id
	t.description = "Synchronize hostlock state after reboot."
}

// Add your test case here
func (t *Ticket_825) Add_testcases() {
	// TESTCASE 001
	tc_1 := t.New_testcase(1, "Synchronize hostlock state after reboot.")
	tc_func := func() common.Testcase_status {

		if err := lib.Start_jaz_server(); err != nil {
			tc_1.Err_log("Failed to start jobarg-server, Error: %s", err.Error())
			return FAILED
		}
		fmt.Println(tc_1.Info_log("JAZ Server has been started."))

		if err := lib.Disable_jaz_server(); err != nil {
			tc_1.Err_log("Failed to disable jobarg-server, Error: %s", err.Error())
			return FAILED
		}
		fmt.Println(tc_1.Info_log("JAZ Server has been disabled."))

		return UpdateHostLockAfterRetryCountwithHostNameRebootAfter("Icon_1", tc_1)
	}

	tc_1.Set_function(tc_func)
	t.Add_testcase(*tc_1)

	// TESTCASE 002
	tc_2 := t.New_testcase(2, "Synchronize hostlock state force reboot.")
	tc_func = func() common.Testcase_status {

		if err := lib.Start_jaz_server(); err != nil {
			tc_2.Err_log("Failed to start jobarg-server, Error: %s", err.Error())
			return FAILED
		}
		fmt.Println(tc_2.Info_log("JAZ Server has been started."))

		if err := lib.Disable_jaz_server(); err != nil {
			tc_2.Err_log("Failed to disable jobarg-server, Error: %s", err.Error())
			return FAILED
		}
		fmt.Println(tc_2.Info_log("JAZ Server has been disabled."))

		return UpdateHostLockAfterRetryCountwithHostNameForceReboot("Icon_1", tc_2)
	}
	tc_2.Set_function(tc_func)
	t.Add_testcase(*tc_2)

	// TESTCASE 003
	tc_3 := t.New_testcase(3, "Update hostlock even retry count is over(environment variable)")
	tc_func = func() common.Testcase_status {

		if err := lib.Start_jaz_server(); err != nil {
			tc_3.Err_log("Failed to start jobarg-server, Error: %s", err.Error())
			return FAILED
		}
		fmt.Println(tc_3.Info_log("JAZ Server has been started."))

		if err := lib.Disable_jaz_server(); err != nil {
			tc_3.Err_log("Failed to disable jobarg-server, Error: %s", err.Error())
			return FAILED
		}
		fmt.Println(tc_3.Info_log("JAZ Server has been disabled."))

		return UpdateHostLockAfterRetryCountwithEnvironmentVariableRebootAfter("Icon_1", tc_3)
	}
	tc_3.Set_function(tc_func)
	t.Add_testcase(*tc_3)

	// TESTCASE 004
	tc_4 := t.New_testcase(4, "Synchronize hostlock state force reboot.")
	tc_func = func() common.Testcase_status {

		if err := lib.Start_jaz_server(); err != nil {
			tc_4.Err_log("Failed to start jobarg-server, Error: %s", err.Error())
			return FAILED
		}
		fmt.Println(tc_4.Info_log("JAZ Server has been started."))

		if err := lib.Disable_jaz_server(); err != nil {
			tc_4.Err_log("Failed to disable jobarg-server, Error: %s", err.Error())
			return FAILED
		}
		fmt.Println(tc_4.Info_log("JAZ Server has been disabled."))

		return UpdateHostLockAfterRetryCountwithEnvironmentVariableForceReboot("Icon_1", tc_4)
	}
	tc_4.Set_function(tc_func)
	t.Add_testcase(*tc_4)

}

func UpdateHostLockAfterRetryCountwithHostNameRebootAfter(jobnetId string, testcase *dao.TestCase) common.Testcase_status {

	/******************
	Pre-Operation State
	******************/

	error := lib.Ja_set_agent_config_linux("AllowRoot", "1")
	if error != nil {
		fmt.Println(testcase.Err_log("Error Allow Root : %s", error))
		return FAILED
	}

	lib.Jobarg_cleanup_linux()
	logFilePath := "/var/log/jobarranger/jobarg_agentd.log"

	common.Client = lib.GetSSHClient(common.Login_info.Hostname, common.Login_info.Port, common.Login_info.Username, common.Login_info.Password)
	lib.ClearLogFile(common.Client, logFilePath)
	/**************
	Operation State
	***************/

	if err := lib.Jobarg_enable_jobnet("Icon_1", "jobicon_linux_hostname(sleep200)"); err != nil {
		fmt.Println(testcase.Err_log("Failed to enable jobnet, Error: %s", err))
		return FAILED
	}

	// Run jobnet
	run_jobnet_id, error := lib.Jobarg_exec(jobnetId)
	if error != nil {
		fmt.Println(testcase.Err_log("Error: %s, std_out: %s", error.Error(), run_jobnet_id))
		return FAILED
	}
	fmt.Println(testcase.Info_log("Job Icon %s has been successfully run with registry number: %s", jobnetId, run_jobnet_id))

	if err := lib.Jobarg_enable_jobnet("Icon_1", "reboot_icon_hostname(RebootAfter)"); err != nil {
		fmt.Println(testcase.Err_log("Failed to enable jobnet, Error: %s", err))
		return FAILED
	}

	// Run jobnet
	run_jobnet_id, error = lib.Jobarg_exec(jobnetId)
	if error != nil {
		fmt.Println(testcase.Err_log("Error: %s, std_out: %s", error.Error(), run_jobnet_id))
		return FAILED
	}
	fmt.Println(testcase.Info_log("Reboot Icon %s has been successfully run with registry number: %s", jobnetId, run_jobnet_id))

	fmt.Println(testcase.Info_log("Reboot after completing jobs. Waiting for Reboot..."))

	sleepDuration := 6 * time.Minute
	time.Sleep(sleepDuration)

	common.Client = lib.ConnectWithRetry(common.Login_info.Hostname, common.Login_info.Port, common.Login_info.Username, common.Login_info.Password, 60)

	result2, err := lib.GetOutputStrFromSSHCommand(common.Client, "hostname")

	if err != nil {
		fmt.Println(testcase.Err_log("error"))
		return FAILED
	}

	fmt.Print(testcase.Info_log("Successfully rebooted hostname : %s", result2))

	// defer client.Close()

	pattern := "[WARN] In ja_job_exec_close() agent close failed. retry count :[29]"
	timeout := 30 * time.Second // Timeout duration
	interval := 1 * time.Second // Polling interval

	_, err = lib.WaitForPatternInLogFile(common.Client, logFilePath, pattern, timeout, interval)
	if err != nil {
		fmt.Println(testcase.Err_log("Error:%s", err))
	} else {
		fmt.Println(testcase.Info_log("Agent try to connect server retry count is over."))
	}

	// Restart the jobarg server
	if err := lib.Restart_jaz_server(); err != nil {
		fmt.Println(testcase.Err_log("Faild to restart the JAZ server, Error: %s", err.Error()))
		return FAILED
	}
	fmt.Println(testcase.Info_log("JAZ server has been restarted."))

	// /***************
	// Expected Results
	// ****************/

	sleepDuration = 3 * time.Minute
	time.Sleep(sleepDuration)

	maxCount := 2
	_, errJobCountWithDone := runStatusJobnetProcess(nil, 1, &maxCount, testcase) // Correctly capturing count and error
	if errJobCountWithDone != nil {
		fmt.Println(testcase.Err_log("Job check status fail: %s", errJobCountWithDone))
		return FAILED
	}

	return PASSED

}

func UpdateHostLockAfterRetryCountwithHostNameForceReboot(jobnetId string, testcase *dao.TestCase) common.Testcase_status {

	/******************
	Pre-Operation State
	******************/

	lib.Jobarg_cleanup_linux()
	logFilePath := "/var/log/jobarranger/jobarg_agentd.log"

	common.Client = lib.GetSSHClient(common.Login_info.Hostname, common.Login_info.Port, common.Login_info.Username, common.Login_info.Password)
	lib.ClearLogFile(common.Client, logFilePath)
	/**************
	Operation State
	***************/

	if err := lib.Jobarg_enable_jobnet("Icon_1", "jobicon_linux_hostname(sleep200)"); err != nil {
		fmt.Println(testcase.Err_log("Failed to enable jobnet, Error: %s", err))
		return FAILED
	}

	// Run jobnet
	run_jobnet_id, error := lib.Jobarg_exec(jobnetId)
	if error != nil {
		fmt.Println(testcase.Err_log("Error: %s, std_out: %s", error.Error(), run_jobnet_id))
		return FAILED
	}
	fmt.Println(testcase.Info_log("Job Icon %s has been successfully run with registry number: %s", jobnetId, run_jobnet_id))

	if err := lib.Jobarg_enable_jobnet("Icon_1", "reboot_icon_hostname(ForceReboot)"); err != nil {
		fmt.Println(testcase.Err_log("Failed to enable jobnet, Error: %s", err))
		return FAILED
	}

	// Run jobnet
	run_jobnet_id, error = lib.Jobarg_exec(jobnetId)
	if error != nil {
		fmt.Println(testcase.Err_log("Error: %s, std_out: %s", error.Error(), run_jobnet_id))
		return FAILED
	}
	fmt.Println(testcase.Info_log("Reboot Icon %s has been successfully run with registry number: %s", jobnetId, run_jobnet_id))

	fmt.Println(testcase.Info_log("Force Reboot started. Waiting for Reboot..."))

	sleepDuration := 3 * time.Minute
	time.Sleep(sleepDuration)

	common.Client = lib.ConnectWithRetry(common.Login_info.Hostname, common.Login_info.Port, common.Login_info.Username, common.Login_info.Password, 60)

	result2, err := lib.GetOutputStrFromSSHCommand(common.Client, "hostname")

	if err != nil {
		fmt.Println(testcase.Err_log("error"))
		return FAILED
	}

	fmt.Print(testcase.Info_log("Successfully rebooted hostname : %s", result2))

	// defer client.Close()

	//------------------------

	pattern := "[WARN] In ja_job_exec_close() agent close failed. retry count :[29]"
	timeout := 2 * time.Second  // Timeout duration
	interval := 1 * time.Second // Polling interval

	_, err = lib.WaitForPatternInLogFile(common.Client, logFilePath, pattern, timeout, interval)
	if err != nil {
		fmt.Println(testcase.Err_log("Error:%s", err))
	} else {
		fmt.Println(testcase.Info_log("Agent try to connect server retry count is over."))
	}

	// Restart the jobarg server
	if err := lib.Restart_jaz_server(); err != nil {
		fmt.Println(testcase.Err_log("Faild to restart the JAZ server, Error: %s", err.Error()))
		return FAILED
	}
	fmt.Println(testcase.Info_log("JAZ server has been restarted."))

	// /***************
	// Expected Results
	// ****************/

	sleepDuration = 3 * time.Minute
	time.Sleep(sleepDuration)

	maxCount := 1
	_, errJobCountWithDone := runStatusJobnetProcess(nil, 1, &maxCount, testcase) // Correctly capturing count and error
	if errJobCountWithDone != nil {
		fmt.Println(testcase.Err_log("Job check status fail: %s", errJobCountWithDone))
		return FAILED
	}

	return PASSED

}

func UpdateHostLockAfterRetryCountwithEnvironmentVariableRebootAfter(jobnetId string, testcase *dao.TestCase) common.Testcase_status {

	/******************
	Pre-Operation State
	******************/
	lib.Jobarg_cleanup_linux()
	logFilePath := "/var/log/jobarranger/jobarg_agentd.log"

	common.Client = lib.GetSSHClient(common.Login_info.Hostname, common.Login_info.Port, common.Login_info.Username, common.Login_info.Password)
	lib.ClearLogFile(common.Client, logFilePath)
	/**************
	Operation State
	***************/

	if err := lib.Jobarg_enable_jobnet("Icon_1", "jobicon_linux_environment_variable(sleep200)"); err != nil {
		fmt.Println(testcase.Err_log("Failed to enable jobnet, Error: %s", err))
		return FAILED
	}

	// Run jobnet
	run_jobnet_id, error := lib.Jobarg_exec(jobnetId)
	if error != nil {
		fmt.Println(testcase.Err_log("Error: %s, std_out: %s", error.Error(), run_jobnet_id))
		return FAILED
	}
	fmt.Println(testcase.Info_log("%s has been successfully run with registry number: %s", jobnetId, run_jobnet_id))

	if err := lib.Jobarg_enable_jobnet("Icon_1", "reboot_icon_environment_variable(RebootAfter)"); err != nil {
		fmt.Println(testcase.Err_log("Failed to enable jobnet, Error: %s", err))
		return FAILED
	}

	// Run jobnet
	run_jobnet_id, error = lib.Jobarg_exec(jobnetId)
	if error != nil {
		fmt.Println(testcase.Err_log("Error: %s, std_out: %s", error.Error(), run_jobnet_id))
		return FAILED
	}
	fmt.Println(testcase.Info_log("%s has been successfully run with registry number: %s", jobnetId, run_jobnet_id))

	fmt.Println(testcase.Info_log("Reboot after completing jobs started. Waiting for Reboot..."))

	sleepDuration := 6 * time.Minute
	time.Sleep(sleepDuration)

	common.Client = lib.ConnectWithRetry(common.Login_info.Hostname, common.Login_info.Port, common.Login_info.Username, common.Login_info.Password, 60)

	result2, err := lib.GetOutputStrFromSSHCommand(common.Client, "hostname")

	if err != nil {
		fmt.Println(testcase.Err_log("error"))
		return FAILED
	}

	fmt.Print(testcase.Info_log("Successfully rebooted hostname : %s", result2))

	// defer client.Close()

	pattern := "[WARN] In ja_job_exec_close() agent close failed. retry count :[29]"
	timeout := 30 * time.Second // Timeout duration
	interval := 1 * time.Second // Polling interval

	_, err = lib.WaitForPatternInLogFile(common.Client, logFilePath, pattern, timeout, interval)
	if err != nil {
		fmt.Println(testcase.Err_log("Error:%s", err))
	} else {
		fmt.Println(testcase.Info_log("Agent try to connect server retry count is over."))
	}

	// Restart the jobarg server
	if err := lib.Restart_jaz_server(); err != nil {
		fmt.Println(testcase.Err_log("Faild to restart the JAZ server, Error: %s", err.Error()))
		return FAILED
	}
	fmt.Println(testcase.Info_log("JAZ server has been restarted."))

	// /***************
	// Expected Results
	// ****************/

	sleepDuration = 3 * time.Minute
	time.Sleep(sleepDuration)

	maxCount := 2
	_, errJobCountWithDone := runStatusJobnetProcess(nil, 1, &maxCount, testcase) // Correctly capturing count and error
	if errJobCountWithDone != nil {
		fmt.Println(testcase.Err_log("Job check status fail: %s", errJobCountWithDone))
		return FAILED
	}

	return PASSED

}

func UpdateHostLockAfterRetryCountwithEnvironmentVariableForceReboot(jobnetId string, testcase *dao.TestCase) common.Testcase_status {

	/******************
	Pre-Operation State
	******************/
	lib.Jobarg_cleanup_linux()
	logFilePath := "/var/log/jobarranger/jobarg_agentd.log"

	common.Client = lib.GetSSHClient(common.Login_info.Hostname, common.Login_info.Port, common.Login_info.Username, common.Login_info.Password)
	lib.ClearLogFile(common.Client, logFilePath)
	/**************
	Operation State
	***************/

	common.Client = lib.GetSSHClient(common.Login_info.Hostname, common.Login_info.Port, common.Login_info.Username, common.Login_info.Password)
	lib.ClearLogFile(common.Client, logFilePath)
	/**************
	Operation State
	***************/

	if err := lib.Jobarg_enable_jobnet("Icon_1", "jobicon_linux_environment_variable(sleep200)"); err != nil {
		fmt.Println(testcase.Err_log("Failed to enable jobnet, Error: %s", err))
		return FAILED
	}

	// Run jobnet
	run_jobnet_id, error := lib.Jobarg_exec(jobnetId)
	if error != nil {
		fmt.Println(testcase.Err_log("Error: %s, std_out: %s", error.Error(), run_jobnet_id))
		return FAILED
	}
	fmt.Println(testcase.Info_log("%s has been successfully run with registry number: %s", jobnetId, run_jobnet_id))

	if err := lib.Jobarg_enable_jobnet("Icon_1", "reboot_icon_environment_variable(ForceReboot)"); err != nil {
		fmt.Println(testcase.Err_log("Failed to enable jobnet, Error: %s", err))
		return FAILED
	}

	// Run jobnet
	run_jobnet_id, error = lib.Jobarg_exec(jobnetId)
	if error != nil {
		fmt.Println(testcase.Err_log("Error: %s, std_out: %s", error.Error(), run_jobnet_id))
		return FAILED
	}
	fmt.Println(testcase.Info_log("%s has been successfully run with registry number: %s", jobnetId, run_jobnet_id))

	fmt.Println(testcase.Info_log("Force Reboot started. Waiting for Reboot..."))

	sleepDuration := 3 * time.Minute
	time.Sleep(sleepDuration)

	common.Client = lib.ConnectWithRetry(common.Login_info.Hostname, common.Login_info.Port, common.Login_info.Username, common.Login_info.Password, 60)

	result2, err := lib.GetOutputStrFromSSHCommand(common.Client, "hostname")

	if err != nil {
		fmt.Println(testcase.Err_log("error"))
		return FAILED
	}

	fmt.Print(testcase.Info_log("Successfully rebooted hostname : %s", result2))

	pattern := "[WARN] In ja_job_exec_close() agent close failed. retry count :[29]"
	timeout := 2 * time.Second  // Timeout duration
	interval := 1 * time.Second // Polling interval

	_, err = lib.WaitForPatternInLogFile(common.Client, logFilePath, pattern, timeout, interval)
	if err != nil {
		fmt.Println(testcase.Err_log("Error:%s", err))
	} else {
		fmt.Println(testcase.Info_log("Agent try to connect server retry count is over."))
	}

	// Restart the jobarg server
	if err := lib.Restart_jaz_server(); err != nil {
		fmt.Println(testcase.Err_log("Faild to restart the JAZ server, Error: %s", err.Error()))
		return FAILED
	}
	fmt.Println(testcase.Info_log("JAZ server has been restarted."))

	// /***************
	// Expected Results
	// ****************/

	sleepDuration = 3 * time.Minute
	time.Sleep(sleepDuration)

	maxCount := 1
	_, errJobCountWithDone := runStatusJobnetProcess(nil, 1, &maxCount, testcase) // Correctly capturing count and error
	if errJobCountWithDone != nil {
		fmt.Println(testcase.Err_log("Job check status fail: %s", errJobCountWithDone))
		return FAILED
	}

	return PASSED

}

func GetJobnetStatusFromDB(query string) (int, error) {
	dbQuery := lib.DBQuery(query) // Ensure to use the correct DBQuery from lib

	rows, err := lib.GetData(dbQuery) // Ensure GetData accepts this type
	if err != nil {
		return 0, fmt.Errorf("error fetching count: %w", err)
	}
	defer rows.Close()

	var count int
	if rows.Next() {
		err = rows.Scan(&count)
		if err != nil {
			return 0, fmt.Errorf("error scanning count: %w", err)
		}
	} else {
		return 0, fmt.Errorf("no rows found")
	}

	return count, nil
}

// runProcess monitors the count and aborts if it exceeds a threshold
func runStatusJobnetProcess(query *string, processCheckTimeout int, maxCount *int, testcase *dao.TestCase) (int, error) {
	// Use default query if none provided
	defaultQuery := "SELECT COUNT(*) FROM ja_run_jobnet_table WHERE status = 3;"
	actualQuery := defaultQuery
	if query != nil {
		actualQuery = *query
	}

	// Use default maxCount of 0 if none provided
	actualMaxCount := 0
	if maxCount != nil {
		actualMaxCount = *maxCount
	}

	timeoutDuration := time.Duration(processCheckTimeout) * time.Minute

	ctx, cancel := context.WithTimeout(context.Background(), timeoutDuration)
	defer cancel()

	count := 0
	var err error

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return 0, fmt.Errorf("error: timeout reached, exiting loop")

		case <-ticker.C:
			count, err = GetJobnetStatusFromDB(actualQuery)
			if err != nil {
				fmt.Println(err) // Log and continue
				continue
			}

			if count == actualMaxCount {
				if actualMaxCount == 2 {
					fmt.Println(testcase.Info_log("Both jobnet execution is complete successfully"))
				} else if actualMaxCount == 1 {
					fmt.Println(testcase.Info_log("Reboot jobnet execution is complete successfully"))
				}
				return count, nil
			}
		}
	}
}
