package example_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	client *oss.Client
)

var (
	Endpoint        = os.Getenv("ENDPOINT")
	AccessKeyId     = os.Getenv("ACCESS_KEY_ID")
	AccessKeySecret = os.Getenv("ACCESS_KEY_SECRET")
	BucketName      = os.Getenv("BUCKET_NAME")
)

func TestBucketList(t *testing.T) {
	lsRes, err := client.ListBuckets()
	if err != nil {
		t.Log(err)
	}

	for _, bucket := range lsRes.Buckets {
		fmt.Println("Buckets:", bucket.Name)
	}
}

func TestUploadFile(t *testing.T) {

	bucket, err := client.Bucket("my-bucket")
	if err != nil {
		t.Log(err)
	}

	err = bucket.PutObjectFromFile("my-object", "LocalFile")
	if err != nil {
		// HandleError(err)
	}
}

func init() {
	c, err := oss.New("Endpoint", "AccessKeyId", "AccessKeySecret")
	fmt.Println(Endpoint, AccessKeyId, AccessKeySecret)
	if err != nil {
		panic(err)
	}
	client = c
}
