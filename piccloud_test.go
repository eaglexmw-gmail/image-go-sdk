/**********************************************************************************************
 #
 # Github : github.com/tencentyun/go-sdk
 # File name : picloud.go
 # Description : unit test for tencent pic cloud sdk
 #
**********************************************************************************************/
package qcloud

import (
	"fmt"
	"testing"
)

const APPID = 200941
const SECRET_ID = "AKIDh51wIFHJ13Mbc5AWd37z6WmQwIdTghBu"
const SECRET_KEY = "SU4Qn0GoK0YRNS97p0l5rAsxwxcN6Il3"

func TestUpload(t *testing.T) {
	var userid uint = 123456
	fileName := "test/test.jpg"
	cloud := PicCloud{APPID, SECRET_ID, SECRET_KEY}
	info, err := cloud.Upload(userid, fileName)
	if err != nil {
		t.Errorf("pic upload failed, userid=%d, pic=%s, err=%s\n", userid, fileName, err.Error())
	} else {
		fmt.Printf("pic upload success\n")
		info.Print()
	}

	if info.Url == "" || info.DownloadUrl == "" || info.Fileid == "" {
		t.Errorf("pic info error\n")
	}
}
