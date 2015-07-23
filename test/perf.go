package main

import (
	"fmt"
	"github.com/tencentyun/go-sdk"
)

func main(){
	var appid uint = 10000216
	sid := "AKID9ppM2LVRhFERTCxUvfR3WMoW50eJlQpK"
	skey := "REV5OXppEFfDgtTT76WdnOViFOIXPTJn"
	bucket := "huchi"

	for i := 0; i< 100; i++ {
		cloud := qcloud.PicCloud{appid, sid, skey, bucket}
		info, err := cloud.Upload("0", "test.jpg")
		if err != nil {
			fmt.Printf("failed, err=%s", err.Error())
		}else{
			info.Print()
		}
	}

}
