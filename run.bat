call gengen_build.bat
@if errorlevel 1 goto failed
call gen_all.bat 
@if errorlevel 1 goto failed
revel run bgtasks
@goto :eof

:failed
@exit /b -1