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
	for i := 1; i < 20; i++ {
		str := `INSERT INTO public.background_task(id, name, started_at, ended_at) VALUES (` +
			strconv.Itoa(i) + `, '任务` + strconv.Itoa(i) + `', to_timestamp('` + time.Now().Format("2006-01-02 15:04:05") + `','YYYY-MM-DD HH24:MI:SS'), null);`
		p, _ := file.Seek(0, os.SEEK_END)
		file.WriteAt([]byte(str+"\n"), p)
	}

	//创建process

}
