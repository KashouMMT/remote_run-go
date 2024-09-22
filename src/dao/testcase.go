package dao

import (
	"fmt"
	"time"
	"zukigit/remote_run-go/src/lib"
)

const INFO = 1
const ERR = 2

type Testcase_status string

const (
	PASSED     Testcase_status = "PASSED"
	FAILED     Testcase_status = "FAILED"
	MUST_CHECK Testcase_status = "MUST_CHECK"
)

type TestCase struct {
	id          uint
	description string
	logs        *[]string
	status      *Testcase_status
	function    func() Testcase_status
	auth        *Auth
}

func New_testcase(testcase_id uint, testcase_description string, auth *Auth) *TestCase {
	status := FAILED
	logs := []string{}
	return &TestCase{
		id:          testcase_id,
		description: testcase_description,
		auth:        auth,
		status:      &status,
		logs:        &logs,
	}
}

func (t *TestCase) Set_log(log string) {
	*t.logs = append(*t.logs, log)
}

func (t *TestCase) Get_id() uint {
	return t.id
}

func (t *TestCase) Get_auth() Auth {
	return *t.auth
}

func (t *TestCase) Get_dsctn() string {
	return t.description
}

func (t *TestCase) Get_logs() []string {
	return *t.logs
}

func (t *TestCase) Set_status(status Testcase_status) {
	*t.status = status
}

func (t *TestCase) Get_status() Testcase_status {
	return *t.status
}

func (t *TestCase) Set_function(function func() Testcase_status) {
	t.function = function
}

func (t *TestCase) Run_function() Testcase_status {
	return t.function()
}

func (t *TestCase) Is_function_nil() bool {
	return t.function == nil
}

// From here is test case util functions

func (t *TestCase) Logi(level int, log string) string {
	log = fmt.Sprintf("Test_case no[%d], %s", t.Get_id(), log)
	log = lib.Formatted_log(level, log)
	t.Set_log(log)

	return log
}

func (t *TestCase) Err_log(unfmt string, arg ...any) string {
	log := fmt.Sprintf(unfmt, arg...)
	return t.Logi(ERR, log)
}

func (t *TestCase) Info_log(unfmt string, arg ...any) string {
	log := fmt.Sprintf(unfmt, arg...)
	return t.Logi(INFO, log)
}

func (t *TestCase) Ssh_exec(command string) ([]byte, error) {
	return lib.Ssh_exec(command, t.auth.Ssh_client)
}

func (t *TestCase) Ssh_exec_to_str(command string) (string, error) {
	return lib.Ssh_exec_to_str(command, t.auth.Ssh_client)
}

func (t *TestCase) Jobarg_exec(jobnet_id string) (string, error) {
	cmd := fmt.Sprintf("jobarg_exec -z %s -U Admin -P zabbix -j %s &> /tmp/moon_jobarg_exec_result", t.Get_auth().Hostname, jobnet_id)

	_, err := lib.Ssh_exec_to_str(cmd, t.auth.Ssh_client)

	cmd = "cat /tmp/moon_jobarg_exec_result"
	result, err1 := lib.Ssh_exec_to_str(cmd, t.auth.Ssh_client)
	if err1 != nil {
		return result, err1
	}

	if err != nil {
		return result, err
	}

	return lib.Get_res_no(result)
}

func (t *TestCase) Jobarg_get_JA_JOBNETSTATUS(registry_number string) (string, error) {
	cmd := fmt.Sprintf("eval $(jobarg_get -z %s -U Admin -P zabbix -r %s -e) && echo -n $JA_JOBNETSTATUS", t.Get_auth().Hostname, registry_number)
	return lib.Ssh_exec_to_str(cmd, t.auth.Ssh_client)
}

func (t *TestCase) Jobarg_get_JA_JOBSTATUS(registry_number string) (string, error) {
	cmd := fmt.Sprintf("eval $(jobarg_get -z %s -U Admin -P zabbix -r %s -e) && echo -n $JA_JOBSTATUS", t.Get_auth().Hostname, registry_number)
	return lib.Ssh_exec_to_str(cmd, t.auth.Ssh_client)
}

func (t *TestCase) Jobarg_get_LASTEXITCD(registry_number string) (string, error) {
	cmd := fmt.Sprintf("eval $(jobarg_get -z %s -U Admin -P zabbix -r %s -e) && echo -n $JA_LASTEXITCD", t.Get_auth().Hostname, registry_number)
	return lib.Ssh_exec_to_str(cmd, t.auth.Ssh_client)
}

func (t *TestCase) Jobarg_get_LASTSTDOUT(registry_number string) (string, error) {
	cmd := fmt.Sprintf("eval $(jobarg_get -z %s -U Admin -P zabbix -r %s -e) && echo -n $JA_LASTSTDOUT", t.Get_auth().Hostname, registry_number)
	return lib.Ssh_exec_to_str(cmd, t.auth.Ssh_client)
}

func (t *TestCase) Jobarg_get_LASTSTDERR(registry_number string) (string, error) {
	cmd := fmt.Sprintf("eval $(jobarg_get -z %s -U Admin -P zabbix -r %s -e) && echo -n $JA_LASTSTDERR", t.Get_auth().Hostname, registry_number)
	return lib.Ssh_exec_to_str(cmd, t.auth.Ssh_client)
}

// Jobarg_get_jobnet_run_info waits util the jobnet is done or get error and returns Jobnet run info.
func (t *TestCase) Jobarg_get_jobnet_run_info(registry_number string) (*Jobnet_run_info, error) {
	var status, job_status, exit_cd, std_out, std_error string
	var err error
	var index int

	for {
		status, err = t.Jobarg_get_JA_JOBNETSTATUS(registry_number)
		if err != nil {
			lib.Formatted_log(INFO, "Error:%s", err.Error())
		}

		job_status, err = t.Jobarg_get_JA_JOBSTATUS(registry_number)
		if err != nil {
			lib.Formatted_log(INFO, "Error:%s", err.Error())
		}

		if status == "END" || (status == "RUN" && job_status == "ERROR") {
			break
		}
		lib.Spinner_log(index, lib.Formatted_log(INFO, "Getting jobnet[%s] run info but jobnet is not finished yet", registry_number))
		time.Sleep(1 * time.Second)
		index++
	}

	exit_cd, err = t.Jobarg_get_LASTEXITCD(registry_number)
	if err != nil {
		lib.Formatted_log(INFO, "Error:%s", err.Error())
	}

	std_out, err = t.Jobarg_get_LASTSTDOUT(registry_number)
	if err != nil {
		lib.Formatted_log(INFO, "Error:%s", err.Error())
	}

	std_error, err = t.Jobarg_get_LASTSTDERR(registry_number)
	if err != nil {
		lib.Formatted_log(INFO, "Error:%s", err.Error())
	}

	fmt.Println()
	return New_Jobnet(status, job_status, exit_cd, std_out, std_error), nil
}
