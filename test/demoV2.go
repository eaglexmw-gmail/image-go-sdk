// copyright : tencent
// author : solomonooo
// github : github.com/tencentyun/go-sdk

// this is a demo for qcloud go sdk
package main

import (
	"fmt"
	"github.com/tencentyun/go-sdk"
	"io/ioutil"
	"os"
)

func main() {
	var appid uint = 111
	sid := "SECRET_ID"
	skey := "SECRET_KEY"
	bucket := "bucket"

	cloud := qcloud.PicCloud{appid, sid, skey, bucket}
	fmt.Println("=========================================")
	info, err := cloud.UploadFile("./pic/test.jpg")
	if err != nil {
		fmt.Printf("pic upload failed, err = %s\n", err.Error())
	} else {
		fmt.Println("pic upload success")
		info.Print()
	}

	fmt.Println("=========================================")
	picInfo, err := cloud.Stat(info.Fileid)
	if err != nil {
		fmt.Printf("pic stat failed, err = %s\n", err.Error())
	} else {
		fmt.Println("pic stat success")
		picInfo.Print()
	}

	fmt.Println("=========================================")
	info2, err := cloud.Copy(info.Fileid)
	if err != nil {
		fmt.Printf("pic copy failed, err = %s\n", err.Error())
	} else {
		fmt.Println("pic copy success")
		info2.Print()
	}

	fmt.Println("=========================================")
	err = cloud.Delete(info2.Fileid)
	if err != nil {
		fmt.Printf("pic delete failed, err = %s\n", err.Error())
	} else {
		fmt.Println("pic delete success")
	}

	fmt.Println("=========================================")
	url := "http://b.hiphotos.baidu.com/image/pic/item/8ad4b31c8701a18b1efd50a89a2f07082938fec7.jpg"
	sign, _ := cloud.Sign(3600 * 24 * 7)
	fmt.Printf("gen sign with expire time, sign = %s\n", sign)
	sign, _ = cloud.SignOnce(info.Fileid)
	fmt.Printf("gen sign with fileid, sign = %s\n", sign)
	sign, _ = cloud.ProcessSign(3600*24, url)
	fmt.Printf("gen process sign, sign = %s\n", sign)

	fmt.Println("=========================================")
	fi, err := os.Open("./pic/test.jpg")
	if nil != err {
		return
	}
	defer fi.Close()
	picData, err := ioutil.ReadAll(fi)
	if nil != err {
		return
	}
	var analyze qcloud.PicAnalyze
	analyze.Fuzzy = 1
	analyze.Food = 1
	//is fuzzy? is food?
	info, err = cloud.UploadBase(picData, "", analyze)
	if err != nil {
		fmt.Printf("pic analyze failed, err = %s\n", err.Error())
	} else {
		fmt.Println("pic analyze success, pic=./pic/fuzzy.jpg")
		fmt.Printf("is fuzzy : %d\r\n", info.Analyze.Fuzzy)
		fmt.Printf("is food : %d\r\n", info.Analyze.Food)
	}

	fmt.Println("=========================================")
	url := "http://b.hiphotos.baidu.com/image/pic/item/8ad4b31c8701a18b1efd50a89a2f07082938fec7.jpg"
	detectInfo, err := cloud.PornDetect(url)
	if err != nil {
		fmt.Printf("porn detect failed, err = %s\n", err.Error())
	} else {
		fmt.Printf("porn detect success\n")
		detectInfo.Print()
	}

	fmt.Println("=========================================")
	pornUrl := []string{
		"http://b.hiphotos.baidu.com/image/pic/item/8ad4b31c8701a18b1efd50a89a2f07082938fec7.jpg",
        "http://c.hiphotos.baidu.com/image/h%3D200/sign=7b991b465eee3d6d3dc680cb73176d41/96dda144ad3459829813ed730bf431adcaef84b1.jpg",
    }
	pornUrlRes, err := cloud.PornDetectUrl(pornUrl)
	if err != nil {
		fmt.Printf("porn detect failed, err = %s\n", err.Error())
	} else {
		fmt.Printf("porn detect success\n")
		fmt.Printf(pornUrlRes)
		fmt.Printf("\n")
	}

	fmt.Println("=========================================")
	pornFile := []string{
        "D:/porn/test1.jpg",
        "D:/porn/test2.jpg",
        "../../../../../porn/测试.png",
    }
	pornFileRes, err := cloud.PornDetectFile(pornFile)
	if err != nil {
		fmt.Printf("porn detect failed, err = %s\n", err.Error())
	} else {
		fmt.Printf("porn detect success\n")
		fmt.Printf(pornFileRes)
		fmt.Printf("\n")
	}
}
