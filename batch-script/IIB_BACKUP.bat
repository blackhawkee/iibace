@echo off
rem set IIB_NODE=IIBNODE
set IIB_NODE=%1

set BACKUP_BROKER_PATH=%2
set IIB_INSTALLED_PATH=%3

set CHAR_TIME=%TIME%
set CHAR_TIME=%CHAR_TIME::=_%
set CHAR_TIME=%CHAR_TIME:.=_%
set ARCHIVE_NAME=%IIB_NODE%_%CHAR_TIME%.zip

rem set BACKUP_BROKER_PATH="C:\Users\Premson Karkare\Desktop\ACE Demo"

rem set IIB_INSTALLED_PATH="C:\Program Files\IBM\IIB\10.0.0.26\server\bin"


call %IIB_INSTALLED_PATH%\mqsiprofile.cmd
@echo on
echo Verifying the Integration Node %IIB_NODE% Status
@echo off
call %IIB_INSTALLED_PATH%\mqsilist.exe

mqsibackupbroker %IIB_NODE% -d %BACKUP_BROKER_PATH% -a %ARCHIVE_NAME%

echo Backup Created Successfully