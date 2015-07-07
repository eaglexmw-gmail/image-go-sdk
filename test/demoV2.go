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
	pic_test()
}

func pic_test(){
	var appid uint = 10000001
	sid := "AKIDNZwDVhbRtdGkMZQfWgl2Gnn1dhXs95C0"
	skey := "ZDdyyRLCLv1TkeYOl5OCMLbyH4sJ40wp"
	bucket := "testa"
	
	cloud := qcloud.PicCloud{appid, sid, skey, bucket}
	fmt.Println("=========================================")
	info, err := cloud.Upload("123456", "./test.jpg", "")
	if err != nil {
		fmt.Printf("pic upload failed, err = %s\n", err.Error())
	} else {
		fmt.Println("pic upload success")
		info.Print()
	}

	fmt.Println("******************")
	picInfo, err := cloud.Stat("123456", info.Fileid)
	if err != nil {
		fmt.Printf("pic stat failed, err = %s\n", err.Error())
	} else {
		fmt.Println("pic stat success")
		picInfo.Print()
	}

	fmt.Println("=========================================")
	info2, err := cloud.Copy("123456", info.Fileid)
	if err != nil {
		fmt.Printf("pic copy failed, err = %s\n", err.Error())
	} else {
		fmt.Println("pic copy success")
		info2.Print()
	}

	fmt.Println("******************")
	picInfo, err = cloud.Stat("123456", info2.Fileid)
	if err != nil {
		fmt.Printf("copy pic stat failed, err = %s\n", err.Error())
	} else {
		fmt.Println("copy pic stat success")
		picInfo.Print()
	}

	fmt.Println("=========================================")
	err = cloud.Download("123456", info2.Fileid, "./test2.jpg")
	if err != nil {
		fmt.Printf("pic download failed, err = %s\n", err.Error())
	} else {
		fmt.Println("pic download success")
	}

	fmt.Println("=========================================")
	err = cloud.Delete("123456", info.Fileid)
	if err != nil {
		fmt.Printf("pic delete failed, err = %s\n", err.Error())
	} else {
		fmt.Println("pic delete success")
	}

	err = cloud.Delete("123456", info2.Fileid)
	if err != nil {
		fmt.Printf("copy pic delete failed, err = %s\n", err.Error())
	} else {
		fmt.Println("copy pic delete success")
	}
	
	fmt.Println("=========================================")
	sign, _ := cloud.Sign("123456", 3600*24*7)
	fmt.Printf("gen sign with expire time, sign = %s\n", sign)
	sign, _ = cloud.SignOnce("123456", info.Fileid)
	fmt.Printf("gen sign with fileid, sign = %s\n", sign)

	fmt.Println("=========================================")
	var analyze qcloud.PicAnalyze
	analyze.Fuzzy = 1;
	analyze.Food = 1;
	//is fuzzy? is food?
	info, err = cloud.UploadBase("123456", "./fuzzy.jpg", "", analyze)
	if err != nil {
		fmt.Printf("pic analyze failed, err = %s\n", err.Error())
	} else {
		fmt.Println("pic analyze success, pic=./fuzzy.jpg")
		fmt.Printf("is fuzzy : %d\r\n", info.Analyze.Fuzzy)
		fmt.Printf("is food : %d\r\n", info.Analyze.Food)
	}
}

