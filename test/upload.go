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
	var appid uint = 200943
	sid := "AKIDOXkiS878nYFvc4sggDRxTU56UsmN3LMy"
	skey := "gMoR2lGvMWzxFGrxJCRoZMhU48f0tsdm"

	cloud := qcloud.PicCloud{appid, sid, skey, ""}
	fmt.Println("=========================================")
	var analyze qcloud.PicAnalyze
	info, err := cloud.Upload("123456", "./test.jpg", analyze)
	if err != nil {
		fmt.Printf("pic upload failed, err = %s\n", err.Error())
	} else {
		fmt.Println("pic upload success")
		info.Print()
	}
}

