# tencentyun/image-go-sdk
----------------------------------- 
腾讯云 [万象优图（Cloud Image）](https://www.qcloud.com/product/ci.html) SDK for Go

===================================
简介
----------------------------------- 
go sdk for picture cloud service of tencentyun.

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

###PornDetectInfo
黄图识别的结果信息

		type PornDetectInfo struct {
			Result      int		//供参考的识别结果，0正常，1黄图，2疑似图片
			Confidence  float64 //识别为黄图的置信度，范围0-100；是normal_score, hot_score, porn_score的综合评分
			PornScore   float64 //图片为色情图片的评分
			NormalScore float64 //图片为正常图片的评分
			HotScore    float64	//图片为性感图片的评分
			ForbidStatus int    //封禁状态，0表示正常，1表示图片已被封禁（只有存储在万象优图的图片才会被封禁）
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

### 4. 调用对应的方法
在创建完对象后，根据实际需求，调用对应的操作方法就可以了。sdk提供的方法包括：签名计算、上传、复制、查询、下载和删除等。
#### 获得版本信息
		
	version := cloud.Version()
	
#### 上传数据
如果需要上传图片，根据不同的需求，可以选择不同的上传方法
			
	//pic_info是上传的返回结果
	//最简单的上传接口，提供图片路径即可
	pic_info, err := cloud.UploadFile(filepath)
	//支持自定义fileid的上传文件接口
	pic_info, err := cloud.UploadFileWithFileid(filepath, fileid)
	//使用字节数组[]byte的上传接口
	pic_info, err := cloud.Upload(picData)
	//使用字节数组[]byte且自定义fileid的上传接口
	pic_info, err := cloud.UploadWithFileid(picData, fileid)

#### 复制图片
		
	info, err := cloud.Copy(fileid)
	
#### 查询图片
		
	pic_info, err := cloud.Stat(fileid)

#### 删除图片
		
	err = cloud.Delete(userid, fileid)
	
#### 下载图片
下载图片直接利用图片的下载url即可，开发者可以自行处理。
如果开启了防盗链，还需要在下载url后面追加签名，如果要自行处理，请参考腾讯云的wiki页，熟悉鉴权签名的算法。

#### 黄图识别
	//单图片Url鉴黄
	url := "http://b.hiphotos.baidu.com/image/pic/item/8ad4b31c8701a18b1efd50a89a2f07082938fec7.jpg"
	detectInfo, err := cloud.PornDetect(url)

	//多图片Url鉴黄
	pornUrl := []string{
		"http://b.hiphotos.baidu.com/image/pic/item/8ad4b31c8701a18b1efd50a89a2f07082938fec7.jpg",
        "http://c.hiphotos.baidu.com/image/h%3D200/sign=7b991b465eee3d6d3dc680cb73176d41/96dda144ad3459829813ed730bf431adcaef84b1.jpg",
    }
	pornUrlRes, err := cloud.PornDetectUrl(pornUrl)

	//多图片文件鉴黄
	pornFile := []string{
        "D:/porn/test1.jpg",
        "D:/porn/test2.jpg",
        "../../../../../porn/测试.png",
    }
	pornFileRes, err := cloud.PornDetectFile(pornFile)

### demo示例
请阅读test/demo.go示例
对于v2版本的图片api，请参考demoV2.go
	
版本信息
----------------------------------- 

### v2.0.4
黄图识别新增多图片Url和多图片内容的支持。

### v2.0.3
修复黄图识别bug。

### v2.0.2
增加对黄图识别api的支持。

### v2.0.1
对fileid进行urlencode，支持slash
slash能力需要后台服务支持

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

