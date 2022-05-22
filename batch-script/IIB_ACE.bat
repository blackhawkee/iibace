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
