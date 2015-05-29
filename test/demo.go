/**********************************************************************************************
 #
 # Github : github.com/tencentyun/go-sdk
 # File name : demo.go
 # Description : demo
 #
**********************************************************************************************/
package main

import (
	"fmt"
	"github.com/tencentyun/go-sdk"
)

func main() {
	var appid uint = 200943
	sid := "AKIDOXkiS878nYFvc4sggDRxTU56UsmN3LMy"
	skey := "gMoR2lGvMWzxFGrxJCRoZMhU48f0tsdm"

	cloud := qcloud.PicCloud{appid, sid, skey}
	fmt.Println("=========================================")
	info, err := cloud.Upload(123456, "./test.jpg")
	if err != nil {
		fmt.Printf("pic upload failed, err = %s\n", err.Error())
	} else {
		fmt.Println("pic upload success")
		info.Print()
	}

	fmt.Println("******************")
	picInfo, err := cloud.Stat(123456, info.Fileid)
	if err != nil {
		fmt.Printf("pic stat failed, err = %s\n", err.Error())
	} else {
		fmt.Println("pic stat success")
		picInfo.Print()
	}

	fmt.Println("=========================================")
	info2, err := cloud.Copy(123456, info.Fileid)
	if err != nil {
		fmt.Printf("pic copy failed, err = %s\n", err.Error())
	} else {
		fmt.Println("pic copy success")
		info2.Print()
	}

	fmt.Println("******************")
	picInfo, err = cloud.Stat(123456, info2.Fileid)
	if err != nil {
		fmt.Printf("copy pic stat failed, err = %s\n", err.Error())
	} else {
		fmt.Println("copy pic stat success")
		picInfo.Print()
	}

	fmt.Println("=========================================")
	err = cloud.Download(123456, info2.Fileid, "./test2.jpg")
	if err != nil {
		fmt.Printf("pic download failed, err = %s\n", err.Error())
	} else {
		fmt.Println("pic download success")
	}

	fmt.Println("=========================================")
	err = cloud.Delete(123456, info.Fileid)
	if err != nil {
		fmt.Printf("pic delete failed, err = %s\n", err.Error())
	} else {
		fmt.Println("pic delete success")
	}

	err = cloud.Delete(123456, info2.Fileid)
	if err != nil {
		fmt.Printf("copy pic delete failed, err = %s\n", err.Error())
	} else {
		fmt.Println("copy pic delete success")
	}

}