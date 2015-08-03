// copyright : tencent
// author : solomonooo
// github : github.com/tencentyun/go-sdk

// Package qcloud implements go sdk for qcloud service of pic & video 
package qcloud

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"time"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/bitly/go-simplejson"
	"github.com/tencentyun/go-sdk/sign"
)

const QCLOUD_VERSION = "2.0.1"
const QCLOUD_DOMAIN = "image.myqcloud.com"

type PicCloud struct {
	Appid      uint
	SecretId  string
	SecretKey string
	Bucket	  string
}

type PicAnalyze struct {
	Fuzzy	int
	Food	int
}

type UrlInfo struct {
	Url          string
	DownloadUrl string
	Fileid       string
	Analyze		PicAnalyze
}

type PicInfo struct {
	Url         string
	Fileid      string
	UploadTime uint
	Size        uint
	Md5         string
	Width       uint
	Height      uint
}

func String2Uint(s string) uint {
	tmpInt, _ := strconv.Atoi(s)
	return uint(tmpInt)
}

func (ui *UrlInfo) Print() {
	fmt.Printf("url = %s\n", ui.Url)
	fmt.Printf("fileid = %s\n", ui.Fileid)
	fmt.Printf("download url = %s\n", ui.DownloadUrl)
}

func (pi *PicInfo) Version() string {
	return QCLOUD_VERSION
}

func (pi *PicInfo) Print() {
	fmt.Printf("url = %s\n", pi.Url)
	fmt.Printf("fileid = %s\n", pi.Fileid)
	fmt.Printf("upload time = %d\n", pi.UploadTime)
	fmt.Printf("size = %d\n", pi.Size)
	fmt.Printf("md5 = %s\n", pi.Md5)
	fmt.Printf("width = %d\n", pi.Width)
	fmt.Printf("height = %d\n", pi.Height)
}

func (pc *PicCloud) getUrl(userid string, fileid string) string {
	var req_url string
	//check version
	if "" == pc.Bucket {
		//v1
		//url = fmt.Sprintf("http://eleme.image.myqcloud.com/photos/v1/%d/%s", pc.Appid, userid)
		req_url = fmt.Sprintf("http://web.%s/photos/v1/%d/%s", QCLOUD_DOMAIN, pc.Appid, userid)
	}else {
		//v2
		//req_url = fmt.Sprintf("http://eleme.image.myqcloud.com/photos/v2/%d/%s/%s", pc.Appid, pc.Bucket, userid)
		req_url = fmt.Sprintf("http://web.%s/photos/v2/%d/%s/%s", QCLOUD_DOMAIN, pc.Appid, pc.Bucket, userid)
	}

	if "" != fileid {
		req_url += "/"+url.QueryEscape(fileid)
	}

	return req_url
}

func (pc *PicCloud) getDownloadUrl(userid string, fileid string) string {
	//check version
	if "" == pc.Bucket {
		//v1
		return fmt.Sprintf("http://%d.%s/%d/%s/%s/original", pc.Appid, QCLOUD_DOMAIN, pc.Appid, userid, fileid)
	}else {
		//v2
		return fmt.Sprintf("http://%s-%d.%s/%s-%d/%s/%s/original", pc.Bucket, pc.Appid, QCLOUD_DOMAIN, pc.Bucket, pc.Appid, userid, fileid)
	}
}

func (pc *PicCloud) parseRsp(rsp []byte) (code int, message string, js *simplejson.Json, err error) {
	//fmt.Printf("http rsp : %s\r\n", string(rsp))
	js, err = simplejson.NewJson(rsp)
	if nil != err {
		return
	}
	code, err = js.Get("code").Int()
	if nil != err {
		return
	}
	message, err = js.Get("message").String()
	if nil != err {
		return
	}
	return
}

func (pc *PicCloud) Upload(filename string) (UrlInfo, error) {
	var analyze PicAnalyze
	return pc.UploadBase(filename, "", analyze)
}

func (pc *PicCloud) UploadWithFileid(filename string, fileid string) (UrlInfo, error) {
	var analyze PicAnalyze
	return pc.UploadBase(filename, fileid, analyze)
}

func (pc *PicCloud) UploadBase(filename string, fileid string, analyze PicAnalyze) (info UrlInfo, err error) {
	if "" == filename {
		err = errors.New("invliad filename")
		return
	}

	reqUrl := pc.getUrl("0", fileid)
	boundary := "-------------------------abcdefg1234567"
	expire := uint(3600)

	var queryString string 
	if analyze.Fuzzy != 0 {
		queryString += "fuzzy."
	}
	if analyze.Food != 0 {
		queryString += "food."
	}
	if queryString != "" {
		reqUrl += "?analyze="+strings.TrimRight(queryString, ".")
	}

	fmt.Println(reqUrl)

	sign, err := sign.AppSignV2(pc.Appid, pc.SecretId, pc.SecretKey, pc.Bucket, expire)
	if nil != err {
		return
	}

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	bodyWriter.SetBoundary(boundary)
	fileWriter, err := bodyWriter.CreateFormFile("FileContent", filename)
	if nil != err {
		return
	}
	fh, err := os.Open(filename)
	if nil != err {
		return
	}
	_, err = io.Copy(fileWriter, fh)
	if nil != err {
		return
	}
	bodyWriter.Close()

	//req, err := http.NewRequest("POST", reqUrl, bodyBuf)
	req, err := http.NewRequest("POST", "http://web."+QCLOUD_DOMAIN, bodyBuf)
	if nil != err {
		return
	}
	req.Header.Set("HOST", "web.image.myqcloud.com")
	req.Header.Set("user-agent", "qcloud-go-sdk")
	req.Header.Set("Authorization", "QCloud "+sign)
	req.Header.Set("Content-Type", "multipart/form-data; boundary="+boundary)
	req.URL.Opaque = strings.TrimPrefix(reqUrl, "http://web."+QCLOUD_DOMAIN)

	var client http.Client
	resp, err := client.Do(req)
	if nil != err {
		fmt.Printf("http error, err=%s", err.Error())
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return
	}

	code, message, js, err := pc.parseRsp(data)
	if nil != err {
		return
	}
	if code != 0 {
		desc := fmt.Sprintf("rsp error, code=%d, message=%s", code, message)
		err = errors.New(desc)
		return
	}

	info.Url, _ = js.Get("data").Get("url").String()
	info.DownloadUrl, _ = js.Get("data").Get("download_url").String()
	info.Fileid, _ = js.Get("data").Get("fileid").String()
	if nil != js.Get("data").Get("is_fuzzy") {
		info.Analyze.Fuzzy, _ = js.Get("data").Get("is_fuzzy").Int()
	}
	if nil != js.Get("data").Get("is_food") {
		info.Analyze.Food, _ = js.Get("data").Get("is_food").Int()
	}
	return
}

func (pc *PicCloud) Download(url string, filename string) error {
	if "" == url || "" == filename {
		return errors.New("invalid param")
	}

	req, err := http.NewRequest("GET", url, nil)
	if nil != err {
		return err
	}
	req.Header.Set("HOST", "web.image.myqcloud.com")
	req.Header.Set("user-agent", "qcloud-go-sdk")

	var client http.Client
	resp, err := client.Do(req)
	if nil != err {
		fmt.Printf("http error, err=%s", err.Error())
		return err
	}
	defer resp.Body.Close()

	fh, err := os.Create(filename)
	defer fh.Close()
	if nil != err {
		return err
	}
	_, err = io.Copy(fh, resp.Body)
	if nil != err {
		return err
	}

	return nil
}

func (pc *PicCloud) Stat(fileid string) (info PicInfo, err error) {
	reqUrl := pc.getUrl("0", fileid)
	//req, err := http.NewRequest("GET", reqUrl, nil)
	req, err := http.NewRequest("GET", "http://web."+QCLOUD_DOMAIN, nil)
	if nil != err {
		return
	}
	req.Header.Set("HOST", "web.image.myqcloud.com")
	req.Header.Set("user-agent", "qcloud-go-sdk")
	req.URL.Opaque = strings.TrimPrefix(reqUrl, "http://web."+QCLOUD_DOMAIN)

	var client http.Client
	resp, err := client.Do(req)
	if nil != err {
		fmt.Printf("http error, err=%s", err.Error())
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return
	}

	code, message, js, err := pc.parseRsp(data)
	if nil != err {
		return
	}
	if code != 0 {
		desc := fmt.Sprintf("rsp error, code=%d, message=%s", code, message)
		err = errors.New(desc)
		return
	}

	var tmp string
	info.Url, _ = js.Get("data").Get("file_url").String()
	info.Fileid, _ = js.Get("data").Get("file_fileid").String()
	tmp, _ = js.Get("data").Get("file_upload_time").String()
	info.UploadTime = String2Uint(tmp)
	tmp, _ = js.Get("data").Get("file_size").String()
	info.Size = String2Uint(tmp)
	info.Md5, _ = js.Get("data").Get("file_md5").String()
	tmp, _ = js.Get("data").Get("photo_width").String()
	info.Width = String2Uint(tmp)
	tmp, _ = js.Get("data").Get("photo_height").String()
	info.Height = String2Uint(tmp)
	return
}

func (pc *PicCloud) Copy(fileid string) (info UrlInfo, err error) {
	reqUrl := pc.getUrl("0", fileid) + "/copy"
	sign, err := sign.AppSignOnceV2(pc.Appid, pc.SecretId, pc.SecretKey, pc.Bucket, fileid)
	if nil != err {
		return
	}

	//req, err := http.NewRequest("POST", reqUrl, nil)
	req, err := http.NewRequest("POST", "http://web."+QCLOUD_DOMAIN, nil)
	if nil != err {
		return
	}
	req.Header.Set("HOST", "web.image.myqcloud.com")
	req.Header.Set("user-agent", "qcloud-go-sdk")
	req.Header.Set("Authorization", "QCloud "+sign)
	req.URL.Opaque = strings.TrimPrefix(reqUrl, "http://web."+QCLOUD_DOMAIN)

	var client http.Client
	resp, err := client.Do(req)
	if nil != err {
		fmt.Printf("http error, err=%s", err.Error())
		return
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return
	}

	code, message, js, err := pc.parseRsp(data)
	if nil != err {
		return
	}
	if code != 0 {
		desc := fmt.Sprintf("rsp error, code=%d, message=%s", code, message)
		err = errors.New(desc)
		return
	}

	info.Url, _ = js.Get("data").Get("url").String()
	info.DownloadUrl, _ = js.Get("data").Get("download_url").String()
	info.Fileid = info.Url[strings.LastIndex(info.Url, "/")+1:]
	return
}

func (pc *PicCloud) Delete(fileid string) error {
	reqUrl := pc.getUrl("0", fileid) + "/del"
	sign, err := sign.AppSignOnceV2(pc.Appid, pc.SecretId, pc.SecretKey, pc.Bucket, fileid)
	if nil != err {
		return err
	}

	//req, err := http.NewRequest("POST", reqUrl, nil)
	req, err := http.NewRequest("POST", "http://web."+QCLOUD_DOMAIN, nil)
	if nil != err {
		return err
	}
	req.Header.Set("HOST", "web.image.myqcloud.com")
	req.Header.Set("user-agent", "qcloud-go-sdk")
	req.Header.Set("Authorization", "QCloud "+sign)
	req.URL.Opaque = strings.TrimPrefix(reqUrl, "http://web."+QCLOUD_DOMAIN)

	var client http.Client
	resp, err := client.Do(req)
	if nil != err {
		fmt.Printf("http error, err=%s", err.Error())
		return err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		return err
	}

	code, message, _, err := pc.parseRsp(data)
	if nil != err {
		return err
	}
	if code != 0 {
		desc := fmt.Sprintf("rsp error, code=%d, message=%s", code, message)
		return errors.New(desc)
	}

	return nil
}

func (pc *PicCloud) Sign(expire uint) (string, error) {
	return sign.AppSignV2(pc.Appid, pc.SecretId, pc.SecretKey, pc.Bucket, expire)
}

func (pc *PicCloud) SignOnce(fileid string) (string, error) {
	return sign.AppSignOnceV2(pc.Appid, pc.SecretId, pc.SecretKey, pc.Bucket, fileid)
}

func (pc *PicCloud) CheckSign(picSign string, fileid string) error {
	if "" == picSign {
		return errors.New("empty sign")
	}
	
	expire, fid, _, err := sign.Decode(picSign, pc.Appid, pc.SecretId, pc.SecretKey)
	if nil != err {
		return err
	}
	//check time
	if expire != 0 {
		//
		now := uint(time.Now().Unix())
		if expire <= now {
			desc := fmt.Sprintf("sign expire, expire time=%d, now=%d", expire, now)
			return errors.New(desc)
		}
	}else{
		//check file id
		if fileid != fid {
			desc := fmt.Sprintf("sign fileid conflict, fileid=%s, fileid in sign=%s", fileid, fid)
			return errors.New(desc)
		}
	}

	return nil
}
