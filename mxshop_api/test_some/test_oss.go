package main

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"os"
)

func handleError(err error) {
	fmt.Println("Error:", err)
	os.Exit(-1)
}
func main() {
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	endpoint := "oss-cn-hangzhou-internal.aliyuncs.com"
	// 阿里云账号AccessKey拥有所有API的访问权限，风险很高。强烈建议您创建并使用RAM用户进行API访问或日常运维，请登录RAM控制台创建RAM用户。
	accessKeyId := "LTAI5tBKqi3bCEzvQ2ayR7e3"
	accessKeySecret := "B9M5dPq9aL4PJmh6HAckJ41y81U90d"
	// yourBucketName填写存储空间名称。
	bucketName := "zcc-shop"
	// yourObjectName填写Object完整路径，完整路径不包含Bucket名称。
	objectName := "gougou.jpg"
	// yourLocalFileName填写本地文件的完整路径。
	localFileName := "C:/Users/张宸铖/Desktop/gougou.jpg"
	// 创建OSSClient实例。
	client, err := oss.New(endpoint, accessKeyId, accessKeySecret)
	if err != nil {
		handleError(err)
	}
	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		handleError(err)
	}
	// 上传文件。
	err = bucket.PutObjectFromFile(objectName, localFileName)
	if err != nil {
		handleError(err)
	}
}
