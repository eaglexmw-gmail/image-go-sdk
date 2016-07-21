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
	"testing"
	"time"
)

const APPID_V1 = 200943
const SID_V1 = "AKIDOXkiS878nYFvc4sggDRxTU56UsmN3LMy"
const SKEY_V1 = "gMoR2lGvMWzxFGrxJCRoZMhU48f0tsdm"

const APPID_V2 = 10000001
const SID_V2 = "AKIDNZwDVhbRtdGkMZQfWgl2Gnn1dhXs95C0"
const SKEY_V2 = "ZDdyyRLCLv1TkeYOl5OCMLbyH4sJ40wp"
const BUCKET = "testb"

func TestUpload(t *testing.T) {
	fileName := "./test/pic/test.jpg"
	cloud := PicCloud{APPID_V1, SID_V1, SKEY_V1, ""}
	info, err := cloud.UploadFile(fileName)
	if err != nil {
		t.Errorf("pic upload failed, pic=%s, err=%s\n", fileName, err.Error())
	} else {
		fmt.Printf("pic upload success\n")
		info.Print()
	}

	if info.Url == "" || info.DownloadUrl == "" || info.Fileid == "" {
		t.Errorf("pic info error\n")
	}
}

func TestUploadWithFileId(t *testing.T) {
	fileName := "./test/pic/test.jpg"
	cloud := PicCloud{APPID_V1, SID_V1, SKEY_V1, ""}
	r := rand.New(rand.NewSource(time.Now().Unix()))
	fileid := fmt.Sprintf("goodnight%d", r.Int63())
	info, err := cloud.UploadFileWithFileid(fileName, fileid)
	if err != nil {
		t.Errorf("pic upload failed, pic=%s, err=%s\n", fileName, err.Error())
	} else {
		fmt.Printf("pic upload success\n")
		info.Print()
	}

	if info.Url == "" || info.DownloadUrl == "" || info.Fileid == "" {
		t.Errorf("pic info error\n")
	}
}

func TestUploadV2(t *testing.T) {
	fileName := "./test/pic/test.jpg"
	cloud := PicCloud{APPID_V2, SID_V2, SKEY_V2, BUCKET}
	info, err := cloud.UploadFile(fileName)
	if err != nil {
		t.Errorf("pic upload failed, pic=%s, err=%s\n", fileName, err.Error())
	} else {
		fmt.Printf("pic upload success\n")
		info.Print()
	}

	if info.Url == "" || info.DownloadUrl == "" || info.Fileid == "" {
		t.Errorf("pic info error\n")
	}
}

func TestUploadWithFileIdV2(t *testing.T) {
	fileName := "./test/pic/test.jpg"
	cloud := PicCloud{APPID_V2, SID_V2, SKEY_V2, BUCKET}
	r := rand.New(rand.NewSource(time.Now().Unix()))
	fileid := fmt.Sprintf("good/night%d.jpg", r.Int63())
	info, err := cloud.UploadFileWithFileid(fileName, fileid)
	if err != nil {
		t.Errorf("pic upload failed, pic=%s, err=%s\n", fileName, err.Error())
	} else {
		fmt.Printf("pic upload success\n")
		info.Print()
	}

	if info.Url == "" || info.DownloadUrl == "" || info.Fileid == "" {
		t.Errorf("pic info error\n")
	}
}

func TestPornDetect(t *testing.T) {
	cloud := PicCloud{APPID_V2, SID_V2, SKEY_V2, BUCKET}
	pornUrl = "http://b.hiphotos.baidu.com/image/pic/item/8ad4b31c8701a18b1efd50a89a2f07082938fec7.jpg"
	info, err := cloud.PornDetect(pornUrl)
	if nil != err {
		t.Errorf("porn detect failed, err=%s\n", err.Error())
	} else {
		fmt.Printf("porn detect success\n")
		info.Print()
	}
}

func TestPornDetectUrl(t *testing.T) {
	cloud := PicCloud{APPID_V2, SID_V2, SKEY_V2, BUCKET}
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
}

func TestPornDetectFile(t *testing.T) {
	cloud := PicCloud{APPID_V2, SID_V2, SKEY_V2, BUCKET}
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

func TestSign(t *testing.T) {
	var expire uint = 3600
	cloud := PicCloud{APPID_V2, SID_V2, SKEY_V2, BUCKET}
	sign, err := cloud.Sign(expire)
	if nil != err {
		t.Errorf("create sign fail, err=%s\n", err.Error())
	} else {
		fmt.Printf("create sign success, sign=%s\n", sign)
	}
}

func TestSignOnce(t *testing.T) {
	fileid := "0fcfeeeb-461c-4693-913b-f32003de09a4"
	cloud := PicCloud{APPID_V2, SID_V2, SKEY_V2, BUCKET}
	sign, err := cloud.SignOnce(fileid)
	if nil != err {
		t.Errorf("create sign fail, err=%s\n", err.Error())
	} else {
		fmt.Printf("create sign success, sign=%s\n", sign)
	}
}

func TestProcessSign(t *testing.T) {
	var expire uint = 3600
	cloud := PicCloud{APPID_V2, SID_V2, SKEY_V2, BUCKET}
	sign, err := cloud.ProcessSign(expire)
	if nil != err {
		t.Errorf("create process sign fail, err=%s\n", err.Error())
	} else {
		fmt.Printf("create process sign success, sign=%s\n", sign)
	}
}

/*
func TestCheckSign(t *testing.T) {
	sign := "rxWs//7zkmp8q/YVYdWOPkKPGTthPTIwMDk0MyZrPUFLSURPWGtpUzg3OG5ZRnZjNHNnZ0RSeFRVNTZVc21OM0xNeSZlPTE0MzMyMjAzNjYmdD0xNDMzMjE2NzY2JnI9NTI2MzU3MjYxJnU9MTIzNDU2JmY9"
	cloud := PicCloud{APPID, SECRET_ID, SECRET_KEY}
	err := cloud.CheckSign(sign, "")
	if nil != err {
		t.Errorf("check sign failed, err=%s\n", err.Error())
	}
	err = cloud.CheckSign(sign, "0fcfeeeb-461c-4693-913b-f32003de09a4")
	if nil != err {
		t.Errorf("check sign failed, err=%s\n", err.Error())
	}
	//
	sign = "9K3HFAMjj4pz1LE+Rh9CZrWI1UxhPTIwMDk0MyZrPUFLSURPWGtpUzg3OG5ZRnZjNHNnZ0RSeFRVNTZVc21OM0xNeSZlPTAmdD0xNDMzMjE2OTExJnI9MTgzNDk2NjQyJnU9MTIzNDU2JmY9MGZjZmVlZWItNDYxYy00NjkzLTkxM2ItZjMyMDAzZGUwOWE0"
	err = cloud.CheckSign(sign, "0fcfeeeb-461c-4693-913b-f32003de09a4")
	if nil != err {
		t.Errorf("check sign failed, err=%s\n", err.Error())
	}
	err = cloud.CheckSign(sign, "")
	if nil == err {
		t.Error("check sign failed, err should not nil\n")
	}
}
*/
