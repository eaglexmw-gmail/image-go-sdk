// copyright : tencent
// author : solomonooo
// github : github.com/tencentyun/go-sdk

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
	"os"

	"github.com/bitly/go-simplejson"
	"github.com/tencentyun/go-sdk/sign"
)

const QCLOUD_VIDEO_DOMAIN = "web.video.myqcloud.com/videos/v1"
const QCLOUD_VIDEO_DOWNLOAD_DOMAIN = "video.myqcloud.com"

type VideoCloud struct {
	Appid      uint
	SecretId  string
	SecretKey string
}

type VideoUrlInfo struct {
	Url          string
	DownloadUrl  string
	Fileid       string
	CoverUrl     string
}

type VideoInfo struct {
	Url         string
	Fileid      string
	UploadTime  uint
	Size        uint
	Sha         string
	Status      uint
	StatusMsg	string
	PlayTime	uint
	Title		string
	Desc		string
	CoverUrl	string
}

func (ui *VideoUrlInfo) Print() {
	fmt.Printf("url = %s\n", ui.Url)
	fmt.Printf("fileid = %s\n", ui.Fileid)
	fmt.Printf("download url = %s\n", ui.DownloadUrl)
	fmt.Printf("cover url = %s\n", ui.CoverUrl)
}

func (vi *VideoInfo) Version() string {
	return QCLOUD_VERSION
}

func (vi *VideoInfo) Print() {
	fmt.Printf("url = %s\n", vi.Url)
	fmt.Printf("fileid = %s\n", vi.Fileid)
	fmt.Printf("upload time = %d\n", vi.UploadTime)
	fmt.Printf("size = %d\n", vi.Size)
	fmt.Printf("sha = %s\n", vi.Sha)
	fmt.Printf("status = %d\n", vi.Status)
	fmt.Printf("status msg = %s\n", vi.StatusMsg)
	fmt.Printf("play time = %d\n", vi.PlayTime)
	fmt.Printf("title = %s\n", vi.Title)
	fmt.Printf("desc = %s\n", vi.Desc)	
	fmt.Printf("cover url = %s\n", vi.CoverUrl)
}

func (vc *VideoCloud) parseRsp(rsp []byte) (code int, message string, js *simplejson.Json, err error) {
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

func (vc *VideoCloud) Upload(userid string, filename string, title string, desc string, magicContext string) (info VideoUrlInfo, err error) {
	if "" == filename {
		err = errors.New("invliad filename")
		return
	}
	reqUrl := fmt.Sprintf("http://%s/%d/%s", QCLOUD_VIDEO_DOMAIN, vc.Appid, userid)
	boundary := "-------------------------abcdefg1234567"
	expire := uint(3600)

	sign, err := sign.AppSign(vc.Appid, vc.SecretId, vc.SecretKey, expire, userid)
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
	//write desc
	bodyWriter.WriteField("desc", desc)
	bodyWriter.WriteField("magicContext", magicContext)
	bodyWriter.WriteField("title", title)


	bodyWriter.Close()

	req, err := http.NewRequest("POST", reqUrl, bodyBuf)
	if nil != err {
		return
	}
	req.Header.Set("HOST", "web.video.myqcloud.com")
	req.Header.Set("user-agent", "qcloud-go-sdk")
	req.Header.Set("Authorization", sign)
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

	code, message, js, err := vc.parseRsp(data)
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
	return
}

func (vc *VideoCloud) Download(userid string, fileid string, filename string) error {
	reqUrl := fmt.Sprintf("http://%d.%s/%d/%s/%s/original", vc.Appid, QCLOUD_VIDEO_DOWNLOAD_DOMAIN, vc.Appid, userid, fileid)
	return vc.DownloadByUrl(reqUrl, filename)
}

func (vc *VideoCloud) DownloadWithSign(userid string, fileid string, filename string) error {

	reqUrl := fmt.Sprintf("http://%d.%s/%d/%s/%s/original", vc.Appid, QCLOUD_VIDEO_DOWNLOAD_DOMAIN, vc.Appid, userid, fileid)
	sign, err := sign.AppSignOnce(vc.Appid, vc.SecretId, vc.SecretKey, userid, fileid)
	if nil != err {
		return err
	}

	return vc.DownloadByUrl(reqUrl+"?sign="+sign, filename)
}

func (vc *VideoCloud) DownloadByUrl(url string, filename string) error {
	if "" == url || "" == filename {
		return errors.New("invalid param")
	}

	req, err := http.NewRequest("GET", url, nil)
	if nil != err {
		return err
	}
	req.Header.Set("HOST", "web.video.myqcloud.com")
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

func (vc *VideoCloud) Stat(userid string, fileid string) (info VideoInfo, err error) {
	reqUrl := fmt.Sprintf("http://%s/%d/%s/%s", QCLOUD_VIDEO_DOMAIN, vc.Appid, userid, fileid)

	req, err := http.NewRequest("GET", reqUrl, nil)
	if nil != err {
		return
	}
	req.Header.Set("HOST", "web.video.myqcloud.com")
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

	code, message, js, err := vc.parseRsp(data)
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
	info.Sha, _ = js.Get("data").Get("file_sha").String()
	tmp, _ = js.Get("data").Get("video_play_time").String()
	info.PlayTime = String2Uint(tmp)
	info.StatusMsg, _ = js.Get("data").Get("video_status_msg").String()
	tmp, _ = js.Get("data").Get("video_status").String()
	info.Status = String2Uint(tmp)
	info.Title, _ = js.Get("data").Get("video_title").String()
	info.Desc, _ = js.Get("data").Get("video_desc").String()
	info.CoverUrl, _ = js.Get("data").Get("video_cover_url").String()
	return
}

func (vc *VideoCloud) Delete(userid string, fileid string) error {
	reqUrl := fmt.Sprintf("http://%s/%d/%s/%s/del", QCLOUD_VIDEO_DOMAIN, vc.Appid, userid, fileid)
	sign, err := sign.AppSignOnce(vc.Appid, vc.SecretId, vc.SecretKey, userid, fileid)
	if nil != err {
		return err
	}

	req, err := http.NewRequest("POST", reqUrl, nil)
	if nil != err {
		return err
	}
	req.Header.Set("HOST", "web.video.myqcloud.com")
	req.Header.Set("user-agent", "qcloud-go-sdk")
	req.Header.Set("Authorization", sign)

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

	code, message, _, err := vc.parseRsp(data)
	if nil != err {
		return err
	}
	if code != 0 {
		desc := fmt.Sprintf("rsp error, code=%d, message=%s", code, message)
		return errors.New(desc)
	}

	return nil
}

func (vc *VideoCloud) Sign(userid string, expire uint) (string, error) {
	return sign.AppSign(vc.Appid, vc.SecretId, vc.SecretKey, expire, userid)
}

func (vc *VideoCloud) SignOnce(userid string, fileid string) (string, error) {
	return sign.AppSignOnce(vc.Appid, vc.SecretId, vc.SecretKey, userid, fileid)
}

func (vc *VideoCloud) CheckSign(userid string, picSign string, fileid string) error {
	if "" == picSign {
		return errors.New("empty sign")
	}
	
	uid, expire, fid, err := sign.Decode(picSign, vc.Appid, vc.SecretId, vc.SecretKey)
	if nil != err {
		return err
	}else if uid != userid {
		desc := fmt.Sprintf("userid conflict, userid=%s, userid in sign=%d", userid, uid)
		return errors.New(desc)
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
