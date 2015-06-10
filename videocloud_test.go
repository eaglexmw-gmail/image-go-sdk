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


func VideoTestUpload(t *testing.T) {
	var userid string = "123456"
	fileName := "./test/test.mp4"
	cloud := VideoCloud{APPID, SECRET_ID, SECRET_KEY}
	info, err := cloud.Upload(userid, fileName,"test_title","test_desc","test_magic_context")
	if err != nil {
		t.Errorf("video upload failed, userid=%s, video=%s, err=%s\n", userid, fileName, err.Error())
	} else {
		fmt.Printf("video upload success\n")
		info.Print()
	}

	if info.Url == "" || info.DownloadUrl == "" || info.Fileid == "" {
		t.Errorf("video info error\n")
	}
}

func VideoTestSign(t *testing.T) {
	var userid string = "123456"
	var expire uint = 3600
	cloud := VideoCloud{APPID, SECRET_ID, SECRET_KEY}
	sign, err := cloud.Sign(userid, expire)
	if nil != err {
		t.Errorf("create sign fail, err=%s\n", err.Error())
	} else {
		fmt.Printf("create sign success, sign=%s\n", sign)
	}
}

func VideoTestSignOnceWithUrl(t *testing.T) {
	var userid string = "123456"
	url := "http://200943.video.myqcloud.com/200943/123456/e7e4d587-e5fc-45c4-b5f8-ef0de5ce4f03/original"
	cloud := VideoCloud{APPID, SECRET_ID, SECRET_KEY}
	sign, err := cloud.SignOnceWithUrl(userid, url)
	if nil != err {
		t.Errorf("create sign fail, err=%s\n", err.Error())
	} else {
		fmt.Printf("create sign success, sign=%s\n", sign)
	}
}

func VideoTestSignOnce(t *testing.T) {
	var userid string = "123456"
	fileid := "0fcfeeeb-461c-4693-913b-f32003de09a4"
	cloud := VideoCloud{APPID, SECRET_ID, SECRET_KEY}
	sign, err := cloud.SignOnce(userid, fileid)
	if nil != err {
		t.Errorf("create sign fail, err=%s\n", err.Error())
	} else {
		fmt.Printf("create sign success, sign=%s\n", sign)
	}
}

func VideoTestCheckSign(t *testing.T) {
	var userid string = "123456"
	sign := "rxWs//7zkmp8q/YVYdWOPkKPGTthPTIwMDk0MyZrPUFLSURPWGtpUzg3OG5ZRnZjNHNnZ0RSeFRVNTZVc21OM0xNeSZlPTE0MzMyMjAzNjYmdD0xNDMzMjE2NzY2JnI9NTI2MzU3MjYxJnU9MTIzNDU2JmY9"
	cloud := VideoCloud{APPID, SECRET_ID, SECRET_KEY}
	err := cloud.CheckSign(userid, sign, "")
	if nil != err {
		t.Errorf("check sign failed, err=%s\n", err.Error())
	}
	err = cloud.CheckSign(userid, sign, "0fcfeeeb-461c-4693-913b-f32003de09a4")
	if nil != err {
		t.Errorf("check sign failed, err=%s\n", err.Error())
	}
	//
	sign = "9K3HFAMjj4pz1LE+Rh9CZrWI1UxhPTIwMDk0MyZrPUFLSURPWGtpUzg3OG5ZRnZjNHNnZ0RSeFRVNTZVc21OM0xNeSZlPTAmdD0xNDMzMjE2OTExJnI9MTgzNDk2NjQyJnU9MTIzNDU2JmY9MGZjZmVlZWItNDYxYy00NjkzLTkxM2ItZjMyMDAzZGUwOWE0"
	err = cloud.CheckSign(userid, sign, "0fcfeeeb-461c-4693-913b-f32003de09a4")
	if nil != err {
		t.Errorf("check sign failed, err=%s\n", err.Error())
	}
	err = cloud.CheckSign(userid, sign, "")
	if nil == err {
		t.Error("check sign failed, err should not nil\n")
	}
}
