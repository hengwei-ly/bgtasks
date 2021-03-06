package controllers

import (
	"bgtasks/app"
	"bgtasks/app/libs"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"bgtasks/app/models"

	"time"

	"github.com/revel/revel"
	"github.com/runner-mei/orm"
	"github.com/three-plus-three/modules/web_ext"
)

type App struct {
	*revel.Controller
	Lifecycle *libs.Lifecycle
}

func (c *App) CurrentUser() web_ext.User {
	return c.Lifecycle.CurrentUser(c.Controller)
}

func (c *App) initLifecycle() revel.Result {
	c.Lifecycle = app.Lifecycle
	c.ViewArgs["menuList"] = c.Lifecycle.MenuList
	return nil
}

type UploadResult struct {
	Success bool   `json:"success"`
	File    string `json:"file"`
	Error   string `json:"error"`
}

func (c App) beforeInvoke() revel.Result {
	c.ViewArgs["controller"] = c.Name
	return nil
}

func (c *App) IsAjax() bool {
	return c.Request.Header.Get("X-Requested-With") == "XMLHttpRequest"
}

func (c App) UploadFile(qqfile []byte) revel.Result {
	errMsg := ""
	if err := os.MkdirAll("tmp/files", os.ModePerm); err != nil {
		errMsg = err.Error()
		revel.ERROR.Print(err)
	} else {
		if len(c.Params.Files["qqfile"]) > 0 {
			filename := c.Params.Files["qqfile"][0].Filename
			if filename, err := c.ensureFileName("tmp/files", filename); err == nil {
				if writeError := ioutil.WriteFile("tmp/files/"+filename, qqfile, os.ModeExclusive); writeError != nil {
					errMsg = writeError.Error()
					revel.ERROR.Print(writeError)
				} else {
					return c.RenderJSON(UploadResult{true, filename, ""})
				}
			} else {
				errMsg = err.Error()
			}
		}
	}

	return c.RenderJSON(UploadResult{false, "", errMsg})
}

func (c App) ensureFileName(dir string, file string) (string, error) {
	parts := strings.Split(file, ".")

	if len(parts) != 2 {
		return file, errors.New("invliad file name")
	}
	filename, ext := parts[0], parts[1]

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return dir, err
	}

	suffixes := []int64{}
	for _, fi := range files {
		if !fi.IsDir() {
			reg := regexp.MustCompile(filename + `(\(\d+\))?\.` + ext)
			matches := reg.FindStringSubmatch(fi.Name())
			if len(matches) == 2 {
				var idx int64
				if matches[1] != "" {
					idx, _ = strconv.ParseInt(matches[1][1:len(matches[1])-1], 10, 64)
				}

				suffixes = append(suffixes, idx)
			}
		}
	}

	if len(suffixes) > 0 {
		sort.Slice(suffixes, func(i, j int) bool { return suffixes[i] < suffixes[j] })
		return fmt.Sprintf("%s(%d).%s", filename, suffixes[len(suffixes)-1]+1, ext), nil
	}
	return file, nil
}

// func (c *ApplicationController) checkUser() revel.Result {
// 	return c.Lifecycle.CheckUser(c.Controller)
// }

func (c App) GetBackgroundTasks() revel.Result {
	backgroundTasks := []models.BackgroundTask{}
	err := c.Lifecycle.DB.BackgroundTasks().Where(orm.Cond{"ended_at is": "null"}).Desc("started_at").All(&backgroundTasks)
	if err != nil {
		c.Flash.Error("查询当前活动任务失败")
		c.FlashParams()
	}
	return c.RenderJSON(backgroundTasks)
}

//获取当前已有的进度信息 是查询所有
func (c App) GetBaseBackgroundTaskProgressesById(task_id int64) revel.Result {
	backgroundTaskProgresses := []models.BackgroundTaskProgress{}
	err := c.Lifecycle.DB.BackgroundTaskProgresses().Where(orm.Cond{"background_task_id = ": task_id}).All(&backgroundTaskProgresses)
	if err != nil {
		c.Flash.Error("查询当前任务进度失败")
		c.FlashParams()
	}
	return c.RenderJSON(backgroundTaskProgresses)
}

//获取最近的进度信息
func (c App) GetLastestProgressesWithId(task_id int64, currentLength int) revel.Result {
	backgroundTaskProgresses := []models.BackgroundTaskProgress{}
	err := c.Lifecycle.DB.BackgroundTaskProgresses().Where(orm.Cond{"background_task_id = ": task_id}).OrderBy("id").Offset(currentLength).Limit(99).All(&backgroundTaskProgresses)
	if err != nil {
		c.Flash.Error("获取任务最新进度失败 ")
		c.FlashParams()
	}
	return c.RenderJSON(backgroundTaskProgresses)
}

//测试用
func (c App) InsertOneProgress() {
	backgroundTaskProgress := models.BackgroundTaskProgress{
		BackgroundTaskId: 1,
		CreatedAt:        time.Now(),
	}
	backgroundTaskProgress.Description = "这里是任务描述撒当机立断卡萨分开弄撒了妇女说看房的科技三分开带伞疯狂圣诞快乐非农大烧烤九分裤删掉就阿凡达框架啊水泥厂小美女安居房女魔都是你的交罚款第三方会计师单" + strconv.Itoa(time.Now().Nanosecond())

	_, err := c.Lifecycle.DB.BackgroundTaskProgresses().Insert(&backgroundTaskProgress)
	if err != nil {
		c.Flash.Error("插入测试数据失败")
		c.FlashParams()
	}
}

func (c App) InsertTestProgresses(num int) revel.Result {
	for i := 0; i < num; i++ {
		c.InsertOneProgress()
	}
	return c.RenderJSON("ok")
}

//func (c App) Get() revel.Result {
//
//}

//

func init() {
	revel.InterceptMethod((*App).initLifecycle, revel.BEFORE)
	revel.InterceptMethod(func(c interface{}) revel.Result {
		if check, ok := c.(interface {
			CheckUser() revel.Result
		}); ok {
			return check.CheckUser()
		}
		return nil
	}, revel.BEFORE)

	revel.InterceptMethod((App).beforeInvoke, revel.BEFORE)
}
