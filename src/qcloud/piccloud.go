/**********************************************************************************************
 #
 # Author : solomonooo
 # Mail : hoshinight@gmail.com
 # Create time : 2015-05-25 11:05
 # Last modified : 2015-05-25 11:05
 # File name : piccloud.go
 # Description : qcloud sdk for go
 #
**********************************************************************************************/
package qcloud

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/bitly/go-simplejson"
	"qcloud/sign"
)

const QCLOUD_DOMAIN = "web.image.myqcloud.com/photos/v1"
const QCLOUD_DOWNLOAD_DOMAIN = "image.myqcloud.com"

type PicCloud struct {
	Appid      uint
	Secret_id  string
	Secret_key string
}

type UrlInfo struct {
	Url          string
	Download_url string
	Fileid       string
}

type PicInfo struct {
	Url         string
	Fileid      string
	Upload_time uint
	Size        uint
	Md5         string
	Width       uint
	Height      uint
}

func String2Uint(s string) uint {
	tmp_int, _ := strconv.Atoi(s)
	return uint(tmp_int)
}

func (ui *UrlInfo) Print() {
	fmt.Printf("url = %s\n", ui.Url)
	fmt.Printf("fileid = %s\n", ui.Fileid)
	fmt.Printf("download url = %s\n", ui.Download_url)
}

func (pi *PicInfo) Print() {
	fmt.Printf("url = %s\n", pi.Url)
	fmt.Printf("fileid = %s\n", pi.Fileid)
	fmt.Printf("upload time = %d\n", pi.Upload_time)
	fmt.Printf("size = %d\n", pi.Size)
	fmt.Printf("md5 = %s\n", pi.Md5)
	fmt.Printf("width = %d\n", pi.Width)
	fmt.Printf("height = %d\n", pi.Height)
}

func (pc *PicCloud) parseRsp(rsp []byte) (code int, message string, js *simplejson.Json, err error) {
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

func (pc *PicCloud) Upload(userid uint, filename string) (info UrlInfo, err error) {
	if "" == filename {
		err = errors.New("invliad filename")
		return
	}

	req_url := fmt.Sprintf("http://%s/%d/%d", QCLOUD_DOMAIN, pc.Appid, userid)
	boundary := "-------------------------abcdefg1234567"
	expire := uint(3600)

	sign, err := sign.Qc_app_sign(pc.Appid, pc.Secret_id, pc.Secret_key, expire, userid)
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

	req, err := http.NewRequest("POST", req_url, bodyBuf)
	if nil != err {
		return
	}
	req.Header.Set("HOST", "web.image.myqcloud.com")
	req.Header.Set("user-agent", "qcloud-go-sdk")
	req.Header.Set("Authorization", "QCloud "+sign)
	req.Header.Set("Content-Type", "multipart/form-data; boundary="+boundary)

	var client http.Client
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if nil != err {
		return
	}

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
	info.Download_url, _ = js.Get("data").Get("download_url").String()
	info.Fileid, _ = js.Get("data").Get("fileid").String()
	return
}

func (pc *PicCloud) Download(userid uint, fileid string, filename string) error {
	req_url := fmt.Sprintf("http://%d.%s/%d/%d/%s/original", pc.Appid, QCLOUD_DOWNLOAD_DOMAIN, pc.Appid, userid, fileid)
	return pc.DownloadByUrl(req_url, filename)
}

func (pc *PicCloud) DownloadWithSign(userid uint, fileid string, filename string) error {

	req_url := fmt.Sprintf("http://%d.%s/%d/%d/%s/original", pc.Appid, QCLOUD_DOWNLOAD_DOMAIN, pc.Appid, userid, fileid)
	sign, err := sign.Qc_app_sign_once(pc.Appid, pc.Secret_id, pc.Secret_key, userid, req_url)
	if nil != err {
		return err
	}

	return pc.DownloadByUrl(req_url+"?sign="+sign, filename)
}

func (pc *PicCloud) DownloadByUrl(url string, filename string) error {
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
	defer resp.Body.Close()
	if nil != err {
		return err
	}

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

func (pc *PicCloud) Stat(userid uint, fileid string) (info PicInfo, err error) {
	req_url := fmt.Sprintf("http://%s/%d/%d/%s", QCLOUD_DOMAIN, pc.Appid, userid, fileid)

	req, err := http.NewRequest("GET", req_url, nil)
	if nil != err {
		return
	}
	req.Header.Set("HOST", "web.image.myqcloud.com")
	req.Header.Set("user-agent", "qcloud-go-sdk")

	var client http.Client
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if nil != err {
		return
	}

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
	info.Upload_time = String2Uint(tmp)
	tmp, _ = js.Get("data").Get("file_size").String()
	info.Size = String2Uint(tmp)
	info.Md5, _ = js.Get("data").Get("file_md5").String()
	tmp, _ = js.Get("data").Get("photo_width").String()
	info.Width = String2Uint(tmp)
	tmp, _ = js.Get("data").Get("photo_height").String()
	info.Height = String2Uint(tmp)
	return
}

func (pc *PicCloud) Copy(userid uint, fileid string) (info UrlInfo, err error) {
	req_url := fmt.Sprintf("http://%s/%d/%d/%s/copy", QCLOUD_DOMAIN, pc.Appid, userid, fileid)
	download_url := fmt.Sprintf("http://%d.%s/%d/%d/%s/original", pc.Appid, QCLOUD_DOWNLOAD_DOMAIN, pc.Appid, userid, fileid)
	sign, err := sign.Qc_app_sign_once(pc.Appid, pc.Secret_id, pc.Secret_key, userid, download_url)
	if nil != err {
		return
	}

	req, err := http.NewRequest("POST", req_url, nil)
	if nil != err {
		return
	}
	req.Header.Set("HOST", "web.image.myqcloud.com")
	req.Header.Set("user-agent", "qcloud-go-sdk")
	req.Header.Set("Authorization", "QCloud "+sign)

	var client http.Client
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if nil != err {
		return
	}

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
	info.Download_url, _ = js.Get("data").Get("download_url").String()
	info.Fileid = info.Url[strings.LastIndex(info.Url, "/")+1:]
	return
}

func (pc *PicCloud) Delete(userid uint, fileid string) error {
	req_url := fmt.Sprintf("http://%s/%d/%d/%s/del", QCLOUD_DOMAIN, pc.Appid, userid, fileid)
	download_url := fmt.Sprintf("http://%d.%s/%d/%d/%s/original", pc.Appid, QCLOUD_DOWNLOAD_DOMAIN, pc.Appid, userid, fileid)
	sign, err := sign.Qc_app_sign_once(pc.Appid, pc.Secret_id, pc.Secret_key, userid, download_url)
	if nil != err {
		return err
	}

	req, err := http.NewRequest("POST", req_url, nil)
	if nil != err {
		return err
	}
	req.Header.Set("HOST", "web.image.myqcloud.com")
	req.Header.Set("user-agent", "qcloud-go-sdk")
	req.Header.Set("Authorization", "QCloud "+sign)

	var client http.Client
	resp, err := client.Do(req)
	defer resp.Body.Close()
	if nil != err {
		return err
	}

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
