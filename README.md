qcloud-go-sdk
===================================
简介
----------------------------------- 
go sdk for picture cloud service of tencentyun.

版本信息
----------------------------------- 
### v2.0.0
支持2.0版本的图片restful api。内部实现了高度封装，对开发者透明。

### v1.2.1
new features:
增加上传图片的模糊识别和美食识别功能接口

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
2.0版本的云服务在使用前，还需要先创建空间。在使用2.0 api时，需要使用空间名（Bucket）。

### 2. 引入qlcoud包
		
	import "github.com/tencentyun/go-sdk"

### 3. 创建对应操作类的对象
如果要使用图片，需要创建图片操作类对象
		
	//v1版本
	cloud := qcloud.PicCloud{appid, sid, skey, ""}
	//v2版本
	cloud := qcloud.PicCloud{appid, sid, skey, bucket}
如果要使用视频，需要创建视频操作类对象
		
	cloud := qcloud.VideoCloud{appid, sid, skey}

### 4. 调用对应的方法
在创建完对象后，根据实际需求，调用对应的操作方法就可以了。sdk提供的方法包括：签名计算、上传、复制、查询、下载和删除等。
#### 获得版本信息
		
	version := cloud.Version()
	
#### 上传数据
如果需要上传图片，根据不同的需求，可以选择不同的上传方法
			
	//pic_info是上传的返回结果
	//最简单的上传接口，提供userid和图片路径即可
	pic_info, err := cloud.upload(userid, filepath)
	//可以自定义fileid的上传接口
	pic_info, err := cloud.uploadWithFileid(userid, filepath, fileid)
如果需要上传视频
		
	video_info, err := cloud.upload(userid, filepath)

#### 复制图片
		
	info, err := cloud.Copy(userid, fileid)
	
#### 查询图片(视频)
		
	//图片查询
	pic_info, err := cloud.Stat(userid, fileid)
	//视频查询
	video_info, err := cloud.Stat(userid, fileid)

#### 删除图片(视频)
		
	err = cloud.Delete(userid, fileid)
	
#### 下载图片
下载图片直接利用图片的下载url即可，开发者可以自行处理，这里提供的是本地下载的方法。
如果开启了防盗链，还需要在下载url后面追加签名，如果要自行处理，请参考腾讯云的wiki页，熟悉鉴权签名的算法。
		
	//filename是要保存的文件路径	
	//不开启防盗链
	err = cloud.Download(userid, fileid, filepath)
	//开启防盗链
    	err = cloud.DownloadWithSign(userid, fileid, filepath)
	//直接提供url下载
	err = cloud.DownloadByUrl(url, filepath)

### demo示例
请阅读test/demo.go示例
对于v2版本的图片api，请参考demoV2.go
	
