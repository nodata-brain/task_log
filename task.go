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

	//file書き込み
	err := ioutil.WriteFile("task.txt", []byte(*task), 0755)

	if err != nil {
		panic(err)
	}

}
