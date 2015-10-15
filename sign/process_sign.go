// copyright : tencent
// author : solomonooo
// github : github.com/tencentyun/go-sdk

// Package sign implements sign for qcloud sdk
package sign

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func ProcessSign(appid uint, secretId string, secretKey string, bucket string, expire uint, url string) (string, error) {
	if "" == secretId || "" == secretKey {
		return "", errors.New("invalid params, secret id or key is empty")
	}

	now := time.Now().Unix()
	expireTime := expire + uint(now)

	plainStr := fmt.Sprintf("a=%d&b=%s&k=%s&t=%d&e=%d&l=%s",
		appid,
		bucket,
		secretId,
		now,
		expireTime,
		url)

	//fmt.Println("sign=", plainStr)
	cryptoStr := []byte(plainStr)
	h := hmac.New(sha1.New, []byte(secretKey))
	h.Write(cryptoStr)
	hmacStr := h.Sum(nil)
	cryptoStr = append(hmacStr, cryptoStr...)
	sign := base64.StdEncoding.EncodeToString(cryptoStr)
	return sign, nil
}

// decode a sign
func ProcessDecode(sign string, appid uint, secretId string, secretKey string) (url string, bucket string, e error) {
	if "" == sign {
		e = errors.New("invalid sign string")
		return
	}

	cryptoStr, e := base64.StdEncoding.DecodeString(sign)
	if nil != e {
		return
	} else if len(cryptoStr) <= HMAC_LENGTH {
		e = errors.New("sign is too short")
		return
	}

	hmacStr := cryptoStr[0:HMAC_LENGTH]
	cryptoStr = cryptoStr[HMAC_LENGTH:]

	//check hmac str
	h := hmac.New(sha1.New, []byte(secretKey))
	h.Write(cryptoStr)
	hmacStr2 := h.Sum(nil)
	if len(hmacStr) != len(hmacStr2) {
		desc := fmt.Sprintf("hmac check failed, hmac1=%s, hmac2=%s", hmacStr, hmacStr2)
		e = errors.New(desc)
		return
	}

	for i := range hmacStr {
		if hmacStr[i] != hmacStr2[i] {
			desc := fmt.Sprintf("hmac check failed, hmac1=%s, hmac2=%s", hmacStr, hmacStr2)
			e = errors.New(desc)
			return
		}
	}

	//check cryto string
	fields := strings.Split(string(cryptoStr), "&")
	cnt := 0
	//check appid
	if fields[cnt] != ("a=" + strconv.Itoa(int(appid))) {
		desc := fmt.Sprintf("invalid appid, appid=%d, sign=%s", appid, fields[0])
		e = errors.New(desc)
		return
	}
	cnt++
	//check skey
	bucket = strings.TrimLeft(fields[cnt], "b=")
	cnt++

	//check sid
	if fields[cnt] != ("k=" + secretId) {
		desc := fmt.Sprintf("invalid secret_id, sid=%s, sign=%s", secretId, fields[1])
		e = errors.New(desc)
		return
	}
	cnt++

	//check time
	//[3] is create time
	//[4] is expire time
	tmp, e := strconv.Atoi(strings.TrimLeft(fields[cnt], "t="))
	if nil != e {
		return
	}
	cnt++
	now := uint(tmp)
	tmp, e = strconv.Atoi(strings.TrimLeft(fields[cnt], "e="))
	if nil != e {
		return
	}
	cnt++
	expire := uint(tmp)
	if expire <= now {
		e = errors.New("sign exceed the time limit")
		return
	}

	//check url
	url = strings.TrimLeft(fields[cnt], "l=")
	e = nil
	return
}
