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

const APPID = 200943
const SECRET_ID = "AKIDOXkiS878nYFvc4sggDRxTU56UsmN3LMy"
const SECRET_KEY = "gMoR2lGvMWzxFGrxJCRoZMhU48f0tsdm"

func TestUpload(t *testing.T) {
	var userid string = "123456"
	fileName := "./test/test.jpg"
	cloud := PicCloud{APPID, SECRET_ID, SECRET_KEY}
	info, err := cloud.Upload(userid, fileName)
	if err != nil {
		t.Errorf("pic upload failed, userid=%s, pic=%s, err=%s\n", userid, fileName, err.Error())
	} else {
		fmt.Printf("pic upload success\n")
		info.Print()
	}

	if info.Url == "" || info.DownloadUrl == "" || info.Fileid == "" {
		t.Errorf("pic info error\n")
	}
}

func TestSign(t *testing.T) {
	var userid string = "123456"
	var expire uint = 3600
	cloud := PicCloud{APPID, SECRET_ID, SECRET_KEY}
	sign, err := cloud.Sign(userid, expire)
	if nil != err {
		t.Errorf("create sign fail, err=%s\n", err.Error())
	} else {
		fmt.Printf("create sign success, sign=%s\n", sign)
	}
}

func TestSignOnceWithUrl(t *testing.T) {
	var userid string = "123456"
	url := "http://200943.image.myqcloud.com/200943/123456/e7e4d587-e5fc-45c4-b5f8-ef0de5ce4f03/original"
	cloud := PicCloud{APPID, SECRET_ID, SECRET_KEY}
	sign, err := cloud.SignOnceWithUrl(userid, url)
	if nil != err {
		t.Errorf("create sign fail, err=%s\n", err.Error())
	} else {
		fmt.Printf("create sign success, sign=%s\n", sign)
	}
}

func TestSignOnce(t *testing.T) {
	var userid string = "123456"
	fileid := "0fcfeeeb-461c-4693-913b-f32003de09a4"
	cloud := PicCloud{APPID, SECRET_ID, SECRET_KEY}
	sign, err := cloud.SignOnce(userid, fileid)
	if nil != err {
		t.Errorf("create sign fail, err=%s\n", err.Error())
	} else {
		fmt.Printf("create sign success, sign=%s\n", sign)
	}
}

/*
func TestCheckSign(t *testing.T) {
	var userid string = "123456"
	sign := "rxWs//7zkmp8q/YVYdWOPkKPGTthPTIwMDk0MyZrPUFLSURPWGtpUzg3OG5ZRnZjNHNnZ0RSeFRVNTZVc21OM0xNeSZlPTE0MzMyMjAzNjYmdD0xNDMzMjE2NzY2JnI9NTI2MzU3MjYxJnU9MTIzNDU2JmY9"
	cloud := PicCloud{APPID, SECRET_ID, SECRET_KEY}
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
*/
