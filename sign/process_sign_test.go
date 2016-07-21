// copyright : tencent
// author : solomonooo
// github : github.com/tencentyun/go-sdk

package sign

import (
	"fmt"
	"testing"
)

const BUCKET = "test"

func TestProcessSign(t *testing.T) {
	var expire uint = 3600 * 24 * 12
	sign, err := ProcessSign(APPID, SECRET_ID, SECRET_KEY, BUCKET, expire)
	if err != nil {
		t.Errorf("gen process sign failed, err = %s\n", err.Error())
	} else {
		fmt.Printf("gen process sign success, sign = %s\n", sign)
	}
}

func TestProcessDecode(t *testing.T) {
	//test1
	sign := "IeoWuK3FKOKX2LquHSpTkoeO251hPTIwMDk0MSZiPXRlc3Qmaz1BS0lEaDUxd0lGSEoxM01iYzVBV2QzN3o2V21Rd0lkVGdoQnUmdD0xNDQ0ODA5NjcwJmU9MTQ0NTg0NjQ3MCZsPWh0dHA6Ly9iLmhpcGhvdG9zLmJhaWR1LmNvbS9pbWFnZS9waWMvaXRlbS84YWQ0YjMxYzg3MDFhMThiMWVmZDUwYTg5YTJmMDcwODI5MzhmZWM3LmpwZw=="
	bucket, err := ProcessDecode(sign, APPID, SECRET_ID, SECRET_KEY)
	if err != nil {
		t.Error("decode error, err=%s\n", err.Error())
	}
	if bucket != BUCKET {
		t.Error("decode bucket info error\n")
	}
	fmt.Printf("Decode process sign success, bucket = %s\n", bucket)
}
