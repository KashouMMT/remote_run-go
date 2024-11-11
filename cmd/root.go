/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/zukigit/remote_run-go/src/common"
	"github.com/zukigit/remote_run-go/src/dao"
	"github.com/zukigit/remote_run-go/src/lib"
	"github.com/zukigit/remote_run-go/src/tickets"

	"github.com/spf13/cobra"
)

var tkts, run_tickets []dao.Ticket
var run_testcases []dao.TestCase

func set_ticket_values(t []dao.Ticket) {
	for _, ticket := range t {
		ticket.Set_values()
	}
}

func check_duplicated_ticket() {
	seen := make(map[uint]bool)

	for _, tkt := range tkts {
		if seen[tkt.Get_no()] {
			fmt.Printf("Error: ticket[%d] is duplicated\n", tkt.Get_no())
			os.Exit(1)
		}
		seen[tkt.Get_no()] = true
	}
}

func add_run_tickets(ticket_number uint) {
	if ticket_number == 0 {
		run_tickets = tkts
	} else {
		for _, t := range tkts {
			if ticket_number == t.Get_no() {
				run_tickets = append(run_tickets, t)
				break
			}
		}
	}
}

func add_run_testcases(testcase_number uint) {
	for _, ticket := range run_tickets {
		ticket.Add_testcases()

		for _, testcase := range ticket.Get_testcases() {
			if testcase_number == 0 || testcase_number == testcase.Get_no() {
				testcase.Set_ticket_no(ticket.Get_no())
				run_testcases = append(run_testcases, testcase)

				if testcase_number != 0 {
					return
				}
			}
		}
	}
}

func run_tc() {
	for _, testcase := range run_testcases {
		dao.Run_testcase(testcase)
	}

	if len(run_testcases) > 0 {
		fmt.Println(lib.Formatted_log(common.INFO, "Logged File: %s", common.Log_filepath))
	} else {
		fmt.Println("There is no testcase to run.")
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "remote_run.exe user@host --with-mysql (or) --with-postgresql",
	Short: "Automated testing",
	Long:  "Automated testing",
	Args: func(cmd *cobra.Command, args []string) error {
		if err := cobra.MinimumNArgs(1)(cmd, args); err != nil {
			return err
		}

		if err := common.Set_db_type(); err != nil {
			return err
		}

		if common.Specific_testcase_no > 0 && common.Specific_ticket_no == 0 {
			return fmt.Errorf("specify the ticket number too by using --ticket")
		}

		return common.Set_usr_hst(args)
	},

	Run: func(cmd *cobra.Command, args []string) {
		common.Set_passwd()
		common.Set_client()
		defer common.Client.Close()

		common.Log_filepath = lib.Get_log_filepath()
		common.Set_sugar(common.Log_filepath)
		defer common.Sugar.Sync()

		// Initialize DB Connection
		common.Set_db_hostname()
		common.Set_default_db_port()
		lib.ConnectDB("zabbix", "zabbix", "zabbix")
		defer common.DB.Close()

		lib.Enable_common_jobnets()

		add_tickets(&tkts)
		set_ticket_values(tkts)
		check_duplicated_ticket()

		add_run_tickets(common.Specific_ticket_no)
		add_run_testcases(common.Specific_testcase_no)

		run_tc() // run test cases
	},
}

func Execute() {
	rootCmd.CompletionOptions.DisableDefaultCmd = true
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntVarP(&common.Login_info.Port, "port", "p", 22, "Port")
	rootCmd.Flags().BoolVar(&common.Is_mysql, "with-mysql", false, "Use MySQL database")
	rootCmd.Flags().BoolVar(&common.Is_psql, "with-postgresql", false, "Use PostgreSQL database")
	rootCmd.Flags().UintVar(&common.Specific_ticket_no, "ticket", 0, "Ticket number to run specific ticket")
	rootCmd.Flags().UintVar(&common.Specific_testcase_no, "testcase", 0, "Testcase number to run specific testcase")
	rootCmd.Flags().StringVar(&common.DB_hostname, "db-hostname", "", "Database specific hostname to connect.")
	rootCmd.Flags().UintVar(&common.DB_port, "db-port", 0, "Database specific port to connect.")
	rootCmd.Flags().UintVar(&common.Timeout, "timeout", 300, "Common timeout in seconds. ")
}

// Add your tickets here
func add_tickets(t *[]dao.Ticket) {
	// *t = append(*t, new(tickets.Ticket_000))
	*t = append(*t, new(tickets.Ticket_1318))
	*t = append(*t, new(tickets.Ticket_811))
	*t = append(*t, new(tickets.Ticket_800))
	*t = append(*t, new(tickets.Ticket_1225))
	*t = append(*t, new(tickets.Ticket_844))
	*t = append(*t, new(tickets.Ticket_794))
	*t = append(*t, new(tickets.Ticket_1021))
	*t = append(*t, new(tickets.Ticket_1091))
	*t = append(*t, new(tickets.Ticket_1089))
	*t = append(*t, new(tickets.Ticket_1264))
	*t = append(*t, new(tickets.Ticket_1292))
}
