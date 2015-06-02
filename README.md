qcloud-go-sdk
===================================
简介
----------------------------------- 
go sdk for picture cloud service of tencentyun.

版本信息
----------------------------------- 
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
###UrlInfo
上传和复制api返回的图片资源链接信息
		
		type UrlInfo struct {
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
		
	cloud := qcloud.PicCloud{appid, sid, skey}

### 4. 调用对应的方法
上传
		
	pic_info, err := cloud.Stat(123456, fileid)
复制
		
	info, err := cloud.Copy(123456, fileid)
查询
		
	picInfo, err = cloud.Stat(123456, fileid)
删除
		
	err = cloud.Delete(123456, fileid)
下载
		
	err = cloud.Download(123456, info2.Fileid, "./test2.jpg")

### demo示例
请阅读test/demo.go示例
	
