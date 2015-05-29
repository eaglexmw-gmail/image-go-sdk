/**********************************************************************************************
 #
 # Github : github.com/tencentyun/go-sdk
 # File name : sign.go
 # Description : qcloud sign
 #
**********************************************************************************************/
package sign

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type picUrlField struct {
	domain string
	appid  uint
	userid uint
	fileid string
}

func parsePicUrl(url string) (fields picUrlField, e error) {
	params := strings.Split(strings.TrimPrefix(url, "http://"), "/")
	if len(params) < 4 {
		desc := fmt.Sprintf("url format error, url=%s", url)
		e = errors.New(desc)
		return
	}
	fields.domain = params[0]
	var value int
	value, _ = strconv.Atoi(params[1])
	fields.appid = uint(value)
	value, _ = strconv.Atoi(params[2])
	fields.userid = uint(value)
	fields.fileid = params[3]
	e = nil
	return
}

func SignBase(appid uint, secretId string, secretKey string, expire uint, userid uint, url string) (string, error) {
	if "" == secretId || "" == secretKey {
		return "", errors.New("invalid params, secret id or key is empty")
	}

	var fileid string
	if "" != url {
		fields, err := parsePicUrl(url)
		if nil != err {
			return "", err
		}

		fileid = fields.fileid
	}

	now := time.Now().Unix()
	r := rand.New(rand.NewSource(time.Now().Unix()))
	rdm := r.Int31()
	expireTime := expire
	if 0 != expireTime {
		expireTime += uint(now)
	}

	plainStr := fmt.Sprintf("a=%d&k=%s&e=%d&t=%d&r=%d&u=%d&f=%s",
		appid,
		secretId,
		expireTime,
		now,
		rdm,
		userid,
		fileid)

	cryptoStr := []byte(plainStr)
	h := hmac.New(sha1.New, []byte(secretKey))
	h.Write(cryptoStr)
	hmacStr := h.Sum(nil)
	cryptoStr = append(hmacStr, cryptoStr...)
	sign := base64.StdEncoding.EncodeToString(cryptoStr)
	return sign, nil
}

func AppSign(appid uint, secretId string, secretKey string, expire uint, userid uint) (string, error) {
	return SignBase(appid, secretId, secretKey, expire, userid, "")
}

func AppSignOnce(appid uint, secretId string, secretKey string, userid uint, url string) (string, error) {
	return SignBase(appid, secretId, secretKey, 0, userid, url)
}
