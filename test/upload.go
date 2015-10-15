// copyright : tencent
// author : solomonooo
// github : github.com/tencentyun/go-sdk

// this is a demo for qcloud go sdk
package main

import (
	"fmt"
	"github.com/tencentyun/go-sdk"
)

func main() {
	var appid uint = 10000001
	sid := "AKIDNZwDVhbRtdGkMZQfWgl2Gnn1dhXs95C0"
	skey := "ZDdyyRLCLv1TkeYOl5OCMLbyH4sJ40wp"
	bucket := "testb"

	cloud := qcloud.PicCloud{appid, sid, skey, bucket}
	fmt.Println("=========================================")
	info, err := cloud.UploadFile("./test.jpg")
	if err != nil {
		fmt.Printf("pic upload failed, err = %s\n", err.Error())
	} else {
		fmt.Println("pic upload success")
		info.Print()
	}
}
