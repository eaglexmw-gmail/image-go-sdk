/**********************************************************************************************
 #
 # Author : solomonooo
 # Mail : hoshinight@gmail.com
 # Create time : 2015-05-25 11:05
 # Last modified : 2015-05-25 11:05
 # File name : piccloud_test.go
 # Description : unit test for picloud
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
	filename := "../test.jpg"
	cloud := PicCloud{APPID, SECRET_ID, SECRET_KEY}
	info, err := cloud.Upload(userid, filename)
	if err != nil {
		t.Errorf("pic upload failed, userid=%d, pic=%s, err=%s\n", userid, filename, err.Error())
	} else {
		fmt.Printf("pic upload success\n")
		info.Print()
	}

	if info.Url == "" || info.Download_url == "" || info.Fileid == "" {
		t.Errorf("pic info error\n")
	}
}
