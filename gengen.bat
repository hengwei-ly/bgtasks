if "%gengen_path%" EQU "" (
  if exist D:\developing\go\meijing\tpt_vendor\src\github.com\three-plus-three\gengen (
    set gengen_path=D:\developing\go\meijing\tpt_vendor\src\github.com\three-plus-three\gengen
  )
  if exist D:\myWorkspace\tpt_vendor\src\github.com\three-plus-three\gengen (
    set gengen_path=D:\myWorkspace\tpt_vendor\src\github.com\three-plus-three\gengen
  )
)

%gengen_path%\gengen.exe %*