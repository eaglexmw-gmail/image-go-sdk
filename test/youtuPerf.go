// copyright : tencent
// author : solomonooo
// github : github.com/tencentyun/go-sdk

// this is a demo for qcloud go sdk
package main

import (
	"os"
	"fmt"
	"time"
	"strconv"
	"math/rand"
	"github.com/tencentyun/go-sdk"
)

const appid uint = 200943
const sid = "AKIDOXkiS878nYFvc4sggDRxTU56UsmN3LMy"
const skey = "gMoR2lGvMWzxFGrxJCRoZMhU48f0tsdm"

var picArray = [3]string{
	"./test.jpg",
	"./food.jpg",
	"./fuzzy.jpg",
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage : performance [thread] [round per thread]")
		return
	}

	var timeTotal int64 = 0
	var timeCnt int64 = 0
	failed := 0

	tcnt, _ := strconv.Atoi(os.Args[1])
	round, _ := strconv.Atoi(os.Args[2])

	chs := make([]chan int64, tcnt)
	for i, _ := range(chs) {
		chs[i] = make(chan int64)
		go do(round, chs[i])
	}

	isLast := false
	for {
		for _, ch := range(chs) {
			t := <-ch
			if t == 0 {
				failed ++
			}else if t < 0 {
				isLast = true
				break
			}else{
				timeTotal += t
				timeCnt++
			}
		}
		fmt.Printf("total time=%dms cnt=%d failed=%d average=%fs\r\n", 
					timeTotal, timeCnt, failed, float32(timeTotal)/float32(timeCnt) / 1000)
	
		if isLast {
			break
		}
	}
}

func do(round int, ch chan int64){
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < round; i++ {
		pic := picArray[r.Int31n(3)]
		fmt.Println("new test ", pic)
		t, _ := pic_test(pic)
		ch <- t
	}
	ch <- -1
}

func pic_test(pic string) (t int64, err error) {
	cloud := qcloud.PicCloud{appid, sid, skey, ""}
	var analyze qcloud.PicAnalyze
	fmt.Println("=========================================")
	analyze.Fuzzy = 1;
	analyze.Food = 1;
	//is fuzzy? is food?
	start := time.Now().UnixNano()
	_, err = cloud.UploadBase("123456", pic, "", analyze)
	if err != nil {
		t = 0
		return
	}
	end := time.Now().UnixNano()
	t = (end - start) / 1000000
	err = nil
	return
}

