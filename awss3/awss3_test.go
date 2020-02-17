package awss3

import (
	"fmt"
	"net/http"
	"os"
	"testing"
)

var (
	testS3 *AwsS3
	err    error

	bucket   = "myBucket"
	basePath = "/test/"

	region          = "ap-northeast-1"
	credentialsFile = "./credentials"
	accessKeyID     = "xxxxxx"
	secretAccessKey = "xxxxxx"
)

func init() {
	testS3, err = NewAwsS3(bucket, basePath, region, credentialsFile)
	if err != nil {
		panic(err)
	}
}

func TestNewAwsS3(t *testing.T) {
	// 使用配置文件初始化
	as3, err := NewAwsS3(bucket, basePath, region, credentialsFile)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(*as3)

	// 使用参数初始化
	//as3, err = NewAwsS3(bucket, basePath, region, "", accessKeyID, secretAccessKey)
	//if err != nil {
	//	t.Error(err)
	//	return
	//}
	//pp.Println(as3)
}

func TestUseS3(t *testing.T) {
	err := InitS3(bucket, basePath, region, credentialsFile)
	if err != nil {
		t.Error(err)
		return
	}

	err = GetS3().CheckFileIsExist("uploadTest1.txt")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestUploadFromFile(t *testing.T) {
	//localFile := "./uploadTest1.txt"
	//localFile := "uploadTest2.jpg"
	localFile := "uploadTest3.csv"
	url, err := testS3.UploadFromFile(localFile)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("upload success, url =", url)
}

func TestUploadFromReader(t *testing.T) {
	localFile := "./uploadTest4.zip"
	f, err := os.Open(localFile)
	if err != nil {
		t.Error(err)
		return
	}
	defer f.Close()

	url, err := testS3.UploadFromReader(f, localFile)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println("upload success, url =", url)
}

func TestDownloadToFile(t *testing.T) {
	awsFile := "uploadTest1.txt"
	localFile := "./download/" + awsFile
	n, err := testS3.DownloadToFile(awsFile, localFile)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println("download file success, size =", n)
}

func TestCheckFileIsExist(t *testing.T) {
	awsFile := "uploadTest1.zip"
	err := testS3.CheckFileIsExist(awsFile)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(awsFile, "is exist")
}

func TestDeleteFile(t *testing.T) {
	awsFile := "uploadTest1.txt"
	err := testS3.DeleteFile(awsFile)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Printf("delete file %s success.\n", awsFile)
}

func TestPreSignedURL(t *testing.T) {
	errMsg := ""
	awsFiles := []string{"uploadTest1.txt", "uploadTest2.jpg", "uploadTest3.csv", "uploadTest4.zip"}
	for _, awsFile := range awsFiles {
		url, err := testS3.GetPreSignedURL(http.MethodGet, awsFile, 300)
		if err != nil {
			errMsg += err.Error() + "\n"
			continue
		}
		fmt.Println(url)
	}

	if errMsg != "" {
		t.Error(errMsg)
	}
}
