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
	"math/rand"
	"time"
	"testing"
)

const APPID_V1 = 200943
const SID_V1 = "AKIDOXkiS878nYFvc4sggDRxTU56UsmN3LMy"
const SKEY_V1 = "gMoR2lGvMWzxFGrxJCRoZMhU48f0tsdm"

const APPID_V2 = 10000001
const SID_V2 = "AKIDNZwDVhbRtdGkMZQfWgl2Gnn1dhXs95C0"
const SKEY_V2 = "ZDdyyRLCLv1TkeYOl5OCMLbyH4sJ40wp"
const BUCKET = "testa"

func TestUpload(t *testing.T) {
	var userid string = "123456"
	fileName := "./test/test.jpg"
	cloud := PicCloud{APPID_V1, SID_V1, SKEY_V1, ""}
	info, err := cloud.Upload(userid, fileName, "")
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

func TestUploadWithFileId(t *testing.T) {
	var userid string = "123456"
	fileName := "./test/test.jpg"
	cloud := PicCloud{APPID_V1, SID_V1, SKEY_V1, ""}
	var analyze PicAnalyze
	r := rand.New(rand.NewSource(time.Now().Unix()))
	fileid := fmt.Sprintf("goodnight%d", r.Int63())
	info, err := cloud.UploadBase(userid, fileName, fileid, analyze)
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

func TestUploadV2(t *testing.T) {
	var userid string = "123456"
	fileName := "./test/test.jpg"
	cloud := PicCloud{APPID_V2, SID_V2, SKEY_V2, BUCKET}
	info, err := cloud.Upload(userid, fileName, "")
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

func TestUploadWithFileIdV2(t *testing.T) {
	var userid string = "123456"
	fileName := "./test/test.jpg"
	cloud := PicCloud{APPID_V2, SID_V2, SKEY_V2, BUCKET}
	var analyze PicAnalyze
	r := rand.New(rand.NewSource(time.Now().Unix()))
	fileid := fmt.Sprintf("goodnight%d", r.Int63())
	info, err := cloud.UploadBase(userid, fileName, fileid, analyze)
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
	cloud := PicCloud{APPID_V1, SID_V1, SKEY_V1, ""}
	sign, err := cloud.Sign(userid, expire)
	if nil != err {
		t.Errorf("create sign fail, err=%s\n", err.Error())
	} else {
		fmt.Printf("create sign success, sign=%s\n", sign)
	}
}

func TestSignOnce(t *testing.T) {
	var userid string = "123456"
	fileid := "0fcfeeeb-461c-4693-913b-f32003de09a4"
	cloud := PicCloud{APPID_V1, SID_V1, SKEY_V1, ""}
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
