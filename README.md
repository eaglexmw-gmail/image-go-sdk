qcloud-go-sdk
===================================
简介
----------------------------------- 
go sdk for picture cloud service of tencentyun.

版本信息
----------------------------------- 
### v1.0.0
稳定版本，支持图片云的基本api。
包括图片的上传、下载、复制、查询和删除。

依赖信息
----------------------------------- 
###simplejson
在使用sdk前，请确保已经安装simplejson
安装方法如下
		
	go get github.com/bitly/go-simplejson

How to start
----------------------------------- 
### 1. 在[腾讯云](http://app.qcloud.com) 申请业务的授权
授权包括：
		
	APP_ID 
	SECRET_ID
	SECRET_KEY

### 2. 创建PicCloud对象
		
	cloud := qcloud.PicCloud{appid, sid, skey}

### 3. 调用对应的方法
上传
		
	pic_info, err := cloud.Stat(123456, fileid)
复制
		
	info, err := cloud.Copy(123456, fileid)
查询
		
	pic_info, err = cloud.Stat(123456, fileid)
删除
		
	err = cloud.Delete(123456, fileid)
下载
		
	err = cloud.Download(123456, info2.Fileid, "./test2.jpg")

### demo示例
请阅读src/demo.go示例
	
