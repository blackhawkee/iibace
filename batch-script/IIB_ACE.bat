@echo off

set IIB_BACKUP_PATH=%1
set IIB_NODE=%2
set ACE_NODE=%3
set ACE_INSTALLED_PATH=%4

rem set ACE_INSTALLED_PATH="C:\Program Files\IBM\ACE\12.0.2.0\server\bin"

call %ACE_INSTALLED_PATH%\mqsiprofile.cmd

call %ACE_INSTALLED_PATH%\mqsilist.exe

call %ACE_INSTALLED_PATH%\mqsiextractcomponents.exe --backup-file %IIB_BACKUP_PATH% --source-integration-node %IIB_NODE% --target-integration-node %ACE_NODE%

echo ACE Integration Node %ACE_NODE% created successfully
echo :
echo :
echo Starting the ACE Integration Node %ACE_NODE%

call %ACE_INSTALLED_PATH%\mqsistart.exe %ACE_NODE%

REM %IIB_INSTALLED_PATH% && (
  REM echo "Profile already set"
  
  REM mqsilist
  REM mqsibackupbroker %IIB_NODE% -d %BACKUP_BROKER_PATH% -a %ARCHIVE_NAME%
  
  REM %ACE_INSTALLED_PATH% && (
	REM mqsiextractcomponents --backup-file %BACKUP_BROKER_PATH%\%ARCHIVE_NAME% --source-integration-node %IIB_NODE% --target-integration-node %ACE_NODE%
  REM ) || (
	REM mqsiextractcomponents --backup-file %BACKUP_BROKER_PATH%\%ARCHIVE_NAME% --source-integration-node %IIB_NODE% --target-integration-node %ACE_NODE%
  REM )

