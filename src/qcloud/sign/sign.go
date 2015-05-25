/**********************************************************************************************
 #
 # Author : solomonooo
 # Mail : hoshinight@gmail.com
 # Create time : 2015-05-25 11:05
 # Last modified : 2015-05-25 11:05
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

type pic_url_field struct {
	domain string
	appid  uint
	userid uint
	fileid string
}

func parse_pic_url(url string) (fields pic_url_field, e error) {
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

func app_sign(appid uint, secret_id string, secret_key string, expire uint, userid uint, url string) (string, error) {
	if "" == secret_id || "" == secret_key {
		return "", errors.New("invalid params, secret id or key is empty")
	}

	var fileid string
	if "" != url {
		fields, err := parse_pic_url(url)
		if nil != err {
			return "", err
		}

		fileid = fields.fileid
	}

	now := time.Now().Unix()
	r := rand.New(rand.NewSource(time.Now().Unix()))
	rdm := r.Int31()
	expire_time := expire
	if 0 != expire_time {
		expire_time += uint(now)
	}

	plain_str := fmt.Sprintf("a=%d&k=%s&e=%d&t=%d&r=%d&u=%d&f=%s",
		appid,
		secret_id,
		expire_time,
		now,
		rdm,
		userid,
		fileid)

	crypto_str := []byte(plain_str)
	h := hmac.New(sha1.New, []byte(secret_key))
	h.Write(crypto_str)
	hmac_str := h.Sum(nil)
	crypto_str = append(hmac_str, crypto_str...)
	sign := base64.StdEncoding.EncodeToString(crypto_str)
	return sign, nil
}

func Qc_app_sign(appid uint, secret_id string, secret_key string, expire uint, userid uint) (string, error) {
	return app_sign(appid, secret_id, secret_key, expire, userid, "")
}

func Qc_app_sign_once(appid uint, secret_id string, secret_key string, userid uint, url string) (string, error) {
	return app_sign(appid, secret_id, secret_key, 0, userid, url)
}
