package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	//コマンドラインのオプションを取得する
	time := flag.Int("t", 0, "task_time")
	task := flag.String("l", "Other" ,"task_type")
	flag.Parse()

	//ファイルを開く
	//todo:ファイル名可変にする
	fp, err := os.Open("task.txt")

	if err != nil {
		panic(err)
	}

	//開いたファイルを閉じる
	defer fp.Close()

	//タスクと時間を出力
	fmt.Println(*time)
	fmt.Println(*task)

}