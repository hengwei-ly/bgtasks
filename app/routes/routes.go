// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tApp struct {}
var App tApp


func (_ tApp) UploadFile(
		qqfile []byte,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "qqfile", qqfile)
	return revel.MainRouter.Reverse("App.UploadFile", args).URL
}

func (_ tApp) GetBackgroundTasks(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("App.GetBackgroundTasks", args).URL
}

func (_ tApp) GetBaseBackgroundTaskProgressesById(
		task_id int64,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "task_id", task_id)
	return revel.MainRouter.Reverse("App.GetBaseBackgroundTaskProgressesById", args).URL
}

func (_ tApp) GetLastestProgressesWithId(
		task_id int64,
		currentLength int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "task_id", task_id)
	revel.Unbind(args, "currentLength", currentLength)
	return revel.MainRouter.Reverse("App.GetLastestProgressesWithId", args).URL
}

func (_ tApp) InsertTestProgresses(
		num int,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "num", num)
	return revel.MainRouter.Reverse("App.InsertTestProgresses", args).URL
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).URL
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).URL
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.Index", args).URL
}

func (_ tTestRunner) Suite(
		suite string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	return revel.MainRouter.Reverse("TestRunner.Suite", args).URL
}

func (_ tTestRunner) Run(
		suite string,
		test string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "suite", suite)
	revel.Unbind(args, "test", test)
	return revel.MainRouter.Reverse("TestRunner.Run", args).URL
}

func (_ tTestRunner) List(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("TestRunner.List", args).URL
}


type tHome struct {}
var Home tHome


func (_ tHome) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Home.Index", args).URL
}


