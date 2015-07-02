// copyright : tencent
// author : solomonooo
// github : github.com/tencentyun/go-sdk

// this is a demo for qcloud go sdk
package main

import (
	"os"
	"fmt"
	"time"
	"bufio"
	"strconv"
	"math/rand"
	"net/http"
)

func main() {
	if len(os.Args) < 4 {
		fmt.Println("usage : downPerf [data file] [thread] [round per thread]")
		return
	}

	var timeTotal int64 = 0
	var timeCnt int64 = 0
	failed := 0

	tcnt, _ := strconv.Atoi(os.Args[2])
	round, _ := strconv.Atoi(os.Args[3])

	urlArray, _ := readUrl(os.Args[1])

	chs := make([]chan int64, tcnt)
	for i, _ := range(chs) {
		chs[i] = make(chan int64)
		go do(urlArray, round, chs[i])
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

func readUrl(file string) (urlArray []string, err error) {
	f, err := os.Open(file)
	if err == nil {
		return 
	}
	defer f.Close()

	urlArray = make([]string, 0)
	bfRd := bufio.NewReader(f)
	delim := '\n'
	for {
		line, err := bfRd.ReadString(delim)
		if err != nil {
			break
		}
		append(urlArray, line)
		fmt.Println("read line, ", line)
	}
	err = nil
	return
}

func do(urlArray []string, round int, ch chan int64){
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for i := 0; i < round; i++ {
		url := urlArray[r.Int31n(len(urlArray))]
		fmt.Println("new test ", url)
		t, _ := get_pic(url)
		ch <- t
	}
	ch <- -1
}

func get_pic(url string) (t int64, err error) {
	t = 0
	start := time.Now().UnixNano()
	res, err := http.Get(url)
	defer res.Body.Close()
	if err != nil {
		return 
	}
	var data string
	data = string(res.Body)
	end := time.Now().UnixNano()
	t = (end - start) / 1000000
	err = nil
	return
}

