package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func init() {
	_, isexist := os.Stat("./tasks.sql")
	if isexist != nil {
		//文件不存在，创建文件
		_, err := os.Create("./tasks.sql")
		if err != nil {
			fmt.Println("文件创建错误")
		}
	}
}
func main() {
	file, _ := os.OpenFile("./tasks.sql", os.O_WRONLY, 0666)

	defer file.Close()
	//创建tasks
	p, _ := file.Seek(0, os.SEEK_END)
	str := "delete from tpt_background_tasks;"
	file.WriteAt([]byte(str+"\n"), p)
	for i := 1; i < 20; i++ {
		str = `INSERT INTO public.tpt_background_tasks (name, started_at, ended_at) VALUES (
		'任务` + strconv.Itoa(i) + `', to_timestamp('` + time.Now().Format("2006-01-02 15:04:05") + `','YYYY-MM-DD HH24:MI:SS'), null);`
		p, _ = file.Seek(0, os.SEEK_END)
		file.WriteAt([]byte(str+"\n"), p)
	}

	//创建process
	p, _ = file.Seek(0, os.SEEK_END)
	str = "delete from tpt_background_task_progresses;"
	file.WriteAt([]byte(str+"\n"), p)

	//任务1的进度
	for i := 1; i < 12; i++ {
		str = `INSERT INTO public.tpt_background_task_progresses (background_task_id, description, created_at)
		VALUES (1 ,'这里是任务描述。。。', to_timestamp('` + time.Now().Format("2006-01-02 15:04:05") + `','YYYY-MM-DD HH24:MI:SS'));`
		p, _ = file.Seek(0, os.SEEK_END)
		file.WriteAt([]byte(str+"\n"), p)
	}

}
