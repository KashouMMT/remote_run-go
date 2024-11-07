package lib

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/zukigit/remote_run-go/src/common"
)

var spinner = []rune{'|', '/', '-', '\\'}

func Get_formatted_time() string {
	currentTime := time.Now()
	return currentTime.Format("20060102150405.000")
}

func Formatted_log(level int, unfmt string, arg ...any) string {
	log := fmt.Sprintf(unfmt, arg...)
	formattedTime := Get_formatted_time()

	switch level {
	case common.INFO:
		log = formattedTime + " [INFO] " + log
	case common.ERR:
		log = formattedTime + " [ERROR] " + log
	default:
		log = formattedTime + " [UNKNOWN] " + log

	}

	return log
}

func Get_log_filepath() string {
	file_name := fmt.Sprintf("%s_TK%d_TC%d.log", Get_formatted_time(), common.Specific_ticket_no, common.Specific_testcase_no)
	currentDir, err := os.Getwd()
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	}

	sub_dir := filepath.Join(currentDir, "logs")
	file_path := filepath.Join(sub_dir, file_name)

	if _, err := os.Stat(sub_dir); os.IsNotExist(err) {
		err = os.Mkdir(sub_dir, 0755) // Create the directory with read/write permissions
		if err != nil {
			fmt.Println("Error:", err.Error())
			os.Exit(1)
		}
	}

	return file_path
}

func Spinner_log(index int, log string) {
	fmt.Printf("\r%s %c", log, spinner[index%len(spinner)])
}
