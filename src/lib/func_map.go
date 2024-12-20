package lib

// Auto-generated by gen_func_map.go. DO NOT EDIT.

var funcMap = map[string]interface{}{
    "ConvertParamPostgresToMysql": ConvertParamPostgresToMysql,
    "ConnectDB": ConnectDB,
    "ExecuteQuery": ExecuteQuery,
    "GetData": GetData,
    "GetSingleRow": GetSingleRow,
    "DBexec": DBexec,
    "JobProcessDBCountCheck": JobProcessDBCountCheck,
    "StopDatabaseService": StopDatabaseService,
    "CheckAndStopDBService": CheckAndStopDBService,
    "StartDatabaseService": StartDatabaseService,
    "Get_file_trunc": Get_file_trunc,
    "Get_hosts_from_jsonfile": Get_hosts_from_jsonfile,
    "Set_hosts_to_jsonfile": Set_hosts_to_jsonfile,
    "Jobarg_exec": Jobarg_exec,
    "Jobarg_get_JA_JOBNETSTATUS": Jobarg_get_JA_JOBNETSTATUS,
    "Jobarg_get_JA_JOBSTATUS": Jobarg_get_JA_JOBSTATUS,
    "Jobarg_get_LASTEXITCD": Jobarg_get_LASTEXITCD,
    "Jobarg_get_LASTSTDOUT": Jobarg_get_LASTSTDOUT,
    "Jobarg_get_LASTSTDERR": Jobarg_get_LASTSTDERR,
    "Jobarg_get_jobnet_run_info": Jobarg_get_jobnet_run_info,
    "Jobarg_get_jobnet_info": Jobarg_get_jobnet_info,
    "Jobarg_cleanup": Jobarg_cleanup,
    "Jobarg_exec_E": Jobarg_exec_E,
    "Jobarg_enable_jobnet": Jobarg_enable_jobnet,
    "Jobarg_server_check_log": Jobarg_server_check_log,
    "Enable_common_jobnets": Enable_common_jobnets,
    "GenerateExcelFile": GenerateExcelFile,
    "Check_service_status_windows": Check_service_status_windows,
    "Check_jazagent_status_windows": Check_jazagent_status_windows,
    "Restart_service_windows": Restart_service_windows,
    "Stop_service_windows": Stop_service_windows,
    "Restart_jaz_agent_windows": Restart_jaz_agent_windows,
    "Stop_jaz_agent_windows": Stop_jaz_agent_windows,
    "Execute_cmd_window": Execute_cmd_window,
    "Execute_cmd_window_str": Execute_cmd_window_str,
    "Ja_set_agent_config_windows": Ja_set_agent_config_windows,
    "Cleanup_agent_windows": Cleanup_agent_windows,
    "Jobarg_cleanup_windows": Jobarg_cleanup_windows,
    "Run_enable_jobnet": Run_enable_jobnet,
    "Run_Jobnet": Run_Jobnet,
    "Run_Jobnet_Exec": Run_Jobnet_Exec,
    "Run_Jobnet_Envs_And_Exec": Run_Jobnet_Envs_And_Exec,
    "Run_Restart_Linux_Jaz_agent": Run_Restart_Linux_Jaz_agent,
    "Run_Restart_Window_Jaz_agent": Run_Restart_Window_Jaz_agent,
    "Run_Restart_Linux_Jaz_server": Run_Restart_Linux_Jaz_server,
    "Run_Job_process_count": Run_Job_process_count,
    "Run_Jobarg_cleanup_linux": Run_Jobarg_cleanup_linux,
    "Run_Jobarg_cleanup_window": Run_Jobarg_cleanup_window,
    "Run_Jobarg_get_jobnet_run_info": Run_Jobarg_get_jobnet_run_info,
    "Run_Check_Jobnet_Finish_With_No_Zombie_Process": Run_Check_Jobnet_Finish_With_No_Zombie_Process,
    "Run_Jobarg_Get_LastSTDOUT": Run_Jobarg_Get_LastSTDOUT,
    "Run_Clear_Linux_Agent_log": Run_Clear_Linux_Agent_log,
    "Run_Clear_Linux_Server_log": Run_Clear_Linux_Server_log,
    "Run_Set_Config_Linux": Run_Set_Config_Linux,
    "Run_Linux_Command": Run_Linux_Command,
    "Run_Linux_Command_Str": Run_Linux_Command_Str,
    "Run_Window_Command": Run_Window_Command,
    "Run_Window_Command_Str": Run_Window_Command_Str,
    "Run_Sql_Script": Run_Sql_Script,
    "Run_Sql_Script_Return_Rows": Run_Sql_Script_Return_Rows,
    "Str_To_Int": Str_To_Int,
    "Run_Job_Status_Check_For_Error": Run_Job_Status_Check_For_Error,
    "JobProcessCountCheck_with_process_counter": JobProcessCountCheck_with_process_counter,
    "Run_SFTP_File_Transfer": Run_SFTP_File_Transfer,
    "Run_Timeout": Run_Timeout,
    "Get_res_no": Get_res_no,
    "Get_str_str_map": Get_str_str_map,
    "Ask_usrinput_string": Ask_usrinput_string,
    "Ask_usrinput_int": Ask_usrinput_int,
    "Ask_usrinput_passwd_string": Ask_usrinput_passwd_string,
    "Ja_set_config_linux": Ja_set_config_linux,
    "Ja_set_config_linux_str_replace": Ja_set_config_linux_str_replace,
    "Ja_set_agent_config_linux": Ja_set_agent_config_linux,
    "Ja_set_server_config_linux": Ja_set_server_config_linux,
    "Restart_jaz_agent_linux": Restart_jaz_agent_linux,
    "Stop_jaz_agent_linux": Stop_jaz_agent_linux,
    "Restart_jaz_server": Restart_jaz_server,
    "Stop_jaz_server": Stop_jaz_server,
    "Sleep_linux": Sleep_linux,
    "Clear_linux_jaz_agent_log": Clear_linux_jaz_agent_log,
    "Disable_jaz_server": Disable_jaz_server,
    "Clear_linux_jaz_server_log": Clear_linux_jaz_server_log,
    "Start_jaz_server": Start_jaz_server,
    "JobProcessCountCheck": JobProcessCountCheck,
    "JobSleepProcessCountCheck": JobSleepProcessCountCheck,
    "CheckZombieProcess": CheckZombieProcess,
    "Cleanup_agent_linux": Cleanup_agent_linux,
    "Jobarg_cleanup_linux": Jobarg_cleanup_linux,
    "ClearLogFile": ClearLogFile,
    "CheckRemoteDirectoryExists": CheckRemoteDirectoryExists,
    "RemoveRemoteDirectory": RemoveRemoteDirectory,
    "RemoveAllFilesInDirectory": RemoveAllFilesInDirectory,
    "CheckRemoteIndexFileExists": CheckRemoteIndexFileExists,
    "FindandTrimServerJobFilePrefix": FindandTrimServerJobFilePrefix,
    "FindandTrimAgentJobFilePrefix": FindandTrimAgentJobFilePrefix,
    "UpdateDebugLevel": UpdateDebugLevel,
    "WaitForPatternInLogFile": WaitForPatternInLogFile,
    "Get_formatted_time": Get_formatted_time,
    "Formatted_log": Formatted_log,
    "Get_log_folderpath": Get_log_folderpath,
    "Get_filepath": Get_filepath,
    "Spinner_log": Spinner_log,
    "Logi": Logi,
    "Get_session": Get_session,
    "Ssh_exec": Ssh_exec,
    "Ssh_exec_to_str": Ssh_exec_to_str,
    "Generate_sshkeys": Generate_sshkeys,
    "GetSSHClientWithKey": GetSSHClientWithKey,
    "GetSSHClient": GetSSHClient,
    "CheckSSHforRebootingAfterDelay": CheckSSHforRebootingAfterDelay,
    "ConnectWithRetry": ConnectWithRetry,
    "ExecuteSSHCommand": ExecuteSSHCommand,
    "GetOutputStrFromSSHCommand": GetOutputStrFromSSHCommand,
    "Set_common_client": Set_common_client,
    "Set_host_pool": Set_host_pool,
}
