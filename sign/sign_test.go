/**********************************************************************************************
 #
 # Github : github.com/tencentyun/go-sdk
 # File name : sign_test.go
 # Description : unit test for qcloud sign
 #
**********************************************************************************************/
package sign

import (
	"fmt"
	"testing"
)

const APPID = 200941
const SECRET_ID = "AKIDh51wIFHJ13Mbc5AWd37z6WmQwIdTghBu"
const SECRET_KEY = "SU4Qn0GoK0YRNS97p0l5rAsxwxcN6Il3"

func TestQc_app_sign(t *testing.T) {
	var userid uint = 123456
	var expire uint = 3600
	sign, err := Qc_app_sign(APPID, SECRET_ID, SECRET_KEY, expire, userid)
	if err != nil {
		t.Errorf("gen sign failed, err = %s\n", err.Error())
	} else {
		fmt.Printf("gen sign success, sign = %s\n", sign)
	}
}

func TestQc_app_sign_once(t *testing.T) {
	var userid uint = 123456
	url := "http://web.image.myqcloud.com/2011541224/123456/442d8ddf-59a5-4dd4-b5f1-e38499fb33b4/orignal"
	sign, err := Qc_app_sign_once(APPID, SECRET_ID, SECRET_KEY, userid, url)
	if err != nil {
		t.Errorf("gen sign failed, err = %s\n", err.Error())
	} else {
		fmt.Printf("gen sign success, sign = %s\n", sign)
	}
}
