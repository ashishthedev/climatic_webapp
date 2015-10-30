
@echo off
REM This file will start the google app engine for you and will also start and smtp server on pot 8025.
start C:\Windows\system32\cmd.exe /k cd b:\\GDrive\Appdir\climatic_webapp_ve\app && goapp serve . --smtp_host=localhost --smtp_port=8025 --clear_datastore --skip_sdk_update_check true

REM start python SMTPDebugginsServer.py
