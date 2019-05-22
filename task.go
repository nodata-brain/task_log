package main

import (
	"flag"
	"io/ioutil"
)

func main() {

	//コマンドラインのオプションを取得する
	//time := flag.Int("t", 0, "task_time")
	task := flag.String("l", "Other", "task_type")
	flag.Parse()

	//file読み込み
	read_task, err := ioutil.ReadFile("task.txt")
	if err != nil {
		panic(err)
	}

	new_task := append(read_task, *task...)

	//file書き込み
	file_err := ioutil.WriteFile("task.txt", new_task, 0755)

	if file_err != nil {
		panic(err)
	}

}
