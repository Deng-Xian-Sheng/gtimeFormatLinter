package main

import (
	"github.com/gogf/gf/os/gtime"
	"time"
)

func A() {
	//expectToGetWrong：

	//viaVariable
	var t gtime.Time
	t.Format("2006-01-02 15:04:05")

	//directCall
	gtime.New().Format("2006-01-02 15:04:05")

	//moreComplexDirectCalls
	gtime.NewFromStr("2006-01-02 15:04:05").Format("2006-01-02 15:04:05")

	//expectNoMistakes：
	//viaVariable
	t.Format("%Y-m-d")

	//directCall
	gtime.New().Format("%Y-m-d")

	//moreComplexDirectCalls
	gtime.NewFromStr("2006-01-02 15:04:05").Format("%Y-m-d")

	//expectNoMistakes：
	//viaVariable
	var sy time.Time
	sy.Format("2006-01-02 15:04:05")

	//directCall
	time.Now().Format("2006-01-02 15:04:05")

	//moreTests

	//expectNoMistakes：
	//viaVariable
	t.Time.Format("2006-01-02 15:04:05")

	//directCall
	gtime.New().Time.Format("2006-01-02 15:04:05")

	//moreComplexDirectCalls
	gtime.NewFromStr("2006-01-02 15:04:05").Time.Format("2006-01-02 15:04:05")
}
