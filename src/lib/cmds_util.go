package lib

func Ssh_exec(command string) ([]byte, error) {
	return test_case.Get_session().Output(command)
}

func Restart_jaz_server() ([]byte, error) {
	return Ssh_exec("systemctl restart jobarg-server")
}

func Stop_jaz_server() ([]byte, error) {
	return Ssh_exec("systemctl stop jobarg-server")
}

func Start_jaz_server() ([]byte, error) {
	return Ssh_exec("systemctl start jobarg-server")
}
