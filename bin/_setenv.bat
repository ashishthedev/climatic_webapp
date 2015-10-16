
@echo off

REM Entry point where the whole dev environment is set
REM Place the following line into the cmd promt shortcut and point to this setenv.bat
REM D:\Windows\system32\cmd.exe /k E:\Dev\WorkSpace\setenv.bat

REM XDATDOCSDIR=proj_ve\app THINK LIKE THIS
set XDATDOCSDIR=B:\GDrive\Appdir\climatic_webapp_ve\app
set RELATIVEPATH=bin\alias.txt

doskey /MACROFILE="%XDATDOCSDIR%\%RELATIVEPATH%"
call %XDATDOCSDIR%\..\scripts\activate.bat
echo _______________________________________
echo    Virtual Environment Activated
echo _______________________________________
set PATH=%XDATDOCSDIR%\bin;%PATH%
set PYTHONPATH=%XDATDOCSDIR%\django_norel_libs;%PYTHONPATH%

prompt [CLIMATIC]$P$_$_$G

pushd %XDATDOCSDIR%
color 73