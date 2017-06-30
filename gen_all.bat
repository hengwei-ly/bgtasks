call gengen.bat mvc  -view_tag=box -controller=App -projectPath=itsm -customPath=/self -root specs
@if errorlevel 1 goto failed
call gengen.bat db -override=true -root specs -output=app/models
@if errorlevel 1 goto failed
call gengen.bat test_base -override=true -projectPath=bgtasks -root specs
@if errorlevel 1 goto failed

FOR %%i IN (app\controllers\*.go) DO goimports -w %%i
FOR %%i IN (app\models\*.go) DO goimports -w %%i
FOR %%i IN (tests\*.go) DO goimports -w %%i
:failed
@exit /b -1