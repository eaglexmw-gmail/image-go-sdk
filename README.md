qcloud-go-sdk
===================================
简介
----------------------------------- 
go sdk for picture cloud service of tencentyun.

版本信息
----------------------------------- 
### v1.2.0
new features:
增加视频上传、查询、删除功能

### v1.1.0
new features:
增加本地签名生成和校验的函数

### v1.0.1
调整github上的包结构
调整代码规范

### v1.0.0
稳定版本，支持图片云的基本api。
包括图片的上传、下载、复制、查询和删除。

安装
----------------------------------- 
		
	go get github.com/tencentyun/go-sdk

数据结构
----------------------------------- 
###PicUrlInfo
上传和复制api返回的图片资源链接信息
		
		type PicUrlInfo struct {
			Url          string	//图片的资源url
			DownloadUrl  string	//图片的下载url
			Fileid       string	//图片资源的唯一标识
		}

###PicInfo
图片本身的属性信息，可以通过查询api获得
		
		type PicInfo struct {
			Url         string	//图片的资源url
			Fileid      string	//图片资源的唯一标识
			UploadTime  uint	//图片的上传时间
			Size        uint	//图片大小，单位字节
			Md5         string	//图片的md5
			Width       uint	//图片宽度
			Height      uint	//图片高度
		}

###VideoUrlInfo
上传和复制api返回的图片资源链接信息
		
		type VideoUrlInfo struct {
			Url          string	//视频的资源url
			DownloadUrl  string	//视频的下载url
			Fileid       string	//视频资源的唯一标识
			Fileid       string	//视频资源的封面url，只有使用视频转码的业务才会有封面 
		}

###VideoInfo
视频本身的属性信息，可以通过查询api获得
		
		type VideoInfo struct {
			Url         string	//视频的下载url 
			Fileid      string	//视频资源的唯一标识 
			UploadTime  uint	//视频上传时间，unix时间戳 
			Size        uint	//视频大小，单位byte
			Sha         string	//视频的sha1摘要 
			Status      uint	//视频状态码，0-初始化, 1-转码中, 2-转码结束,3-转码失败,4-未审核,5-审核通过,6-审核未通过,7-审核失败 
			StatusMsg	string	//视频状态字符串 
			PlayTime	uint	//视频视频播放时长,只有使用视频转码的业务才有 
			Title		string	//视频标题 
			Desc		string	//视频描述 
			CoverUrl	string	//视频封面url,只有使用视频转码的业务才会有封面 
		}
		
How to start
----------------------------------- 
### 1. 在[腾讯云](http://app.qcloud.com) 申请业务的授权
授权包括：
		
	APP_ID 
	SECRET_ID
	SECRET_KEY

### 2. 引入qlcoud包
		
	import "github.com/tencentyun/go-sdk"

### 3. 创建PicCloud对象
		
	p_cloud := qcloud.PicCloud{appid, sid, skey}
	v_cloud := qcloud.VideoCloud{appid, sid, skey}

### 4. 调用对应的方法
上传

	pic_info, err := p_cloud.upload(123456, "./test.jpg")
	video_info, err := v_cloud.upload(123456, "./test.mp4")	
查询

	pic_info, err := p_cloud.Stat(123456, fileid)
	video_info, err := v_cloud.Stat(123456, fileid)
复制
		
	info, err := p_cloud.Copy(123456, fileid)
删除

	err = p_cloud.Delete(123456, fileid)
	err = v_cloud.Delete(123456, fileid)	
下载

	err = p_cloud.Download(123456, info2.Fileid, "./test2.jpg")
	err = v_cloud.Download(123456, info2.Fileid, "./test2.mp4")
### demo示例
请阅读test/demo.go示例
	
