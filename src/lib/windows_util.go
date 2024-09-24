package lib

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/sys/windows/svc"
	"golang.org/x/sys/windows/svc/mgr"
)

func Check_service_status_windows(service string) (bool, error) {
	m, err := mgr.Connect()
	if err != nil {
		return false, err
	}
	defer m.Disconnect()

	s, err := m.OpenService(service)
	if err != nil {
		return false, err
	}
	defer s.Close()

	status, err := s.Query()
	if err != nil {
		return false, err
	}

	if status.State == svc.Running {
		return true, nil
	} else {
		return false, nil
	}
}

func Check_jazagent_status_windows(service string) (bool, error) {
	return Check_service_status_windows("Job Arranger Agent")
}

func Restart_service_windows(service string) error {

	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()

	s, err := m.OpenService(service)
	if err != nil {
		return err
	}
	defer s.Close()

	status, err := s.Query()
	if err != nil {
		return err
	}

	if status.State != svc.Running {
		err = s.Start()
	}

	return err
}

func Stop_service_windows(service string) error {

	m, err := mgr.Connect()
	if err != nil {
		return err
	}
	defer m.Disconnect()

	s, err := m.OpenService(service)
	if err != nil {
		return err
	}
	defer s.Close()

	status, err := s.Control(svc.Stop)
	if err != nil {
		return err
	}

	index := 0
	for status.State != svc.Stopped {
		status, err = s.Query()
		if err != nil {
			return err
		}
		Spinner_log(index, Formatted_log(INFO, "Jobarg agent service is stopping."))

		time.Sleep(2 * time.Second)
	}

	return err
}

func Restart_jaz_agent_windows() error {
	return Restart_service_windows("Job Arranger Agent")
}

func Stop_jaz_agent_windows() error {
	return Stop_service_windows("Job Arranger Agent")
}

// To use this function, you must have jobarg_agentd default filepath.
// Keys must include the following format.
//
// 1) #Javalue=
//
// 2) Javalue=
func Jaz_set_agent_config_windows(key string, value string) error {
	var lines []string
	var updated bool
	file_location := filepath.Join("C:\\", "Program Files", "Job Arranger", "Job Arranger Agent", "conf", "jobarg_agentd.conf")

	file, err := os.OpenFile(file_location, os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if strings.HasPrefix(line, key+"=") {
			fmt.Println("found key1", key)
			lines = append(lines, key+"="+value)
			updated = true
		} else {
			lines = append(lines, line)
		}
	}

	if !updated {
		for index, l := range lines {
			if strings.HasPrefix(l, "# "+key+"=") {
				fmt.Println("found key2", key)
				lines[index] = fmt.Sprintf("%s=%s", key, value)
				updated = true
			}
		}
	}

	// Check for scanner errors
	if err := scanner.Err(); err != nil {
		return err
	}

	if updated {
		// Seek to the beginning of the file to overwrite it
		if _, err := file.Seek(0, 0); err != nil {
			return err
		}

		// Truncate the file to remove any leftover content
		if err := file.Truncate(0); err != nil {
			return err
		}

		// Write the updated lines back to the file
		for _, line := range lines {
			if _, err := file.WriteString(line + "\n"); err != nil {
				return err
			}
		}
	}

	return nil
}
