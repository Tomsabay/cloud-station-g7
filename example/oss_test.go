package example_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/joho/godotenv"
)

var (
	client *oss.Client
)

var (
	Endpoint        string
	AccessKeyId     string
	AccessKeySecret string
	BucketName      string
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

	bucket, err := client.Bucket(BucketName)
	if err != nil {
		fmt.Println("Error:", err)
		t.Log(err)
	}

	err = bucket.PutObjectFromFile("day1/1.txt", "./1.txt")
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func init() {
	if err := godotenv.Load("../etc/unit_test.env"); err != nil {
		panic(err)
	}

	Endpoint = os.Getenv("ENDPOINT")
	AccessKeyId = os.Getenv("ACCESS_KEY_ID")
	AccessKeySecret = os.Getenv("ACCESS_KEY_SECRET")
	BucketName = os.Getenv("BUCKET_NAME")

	// fmt.Printf("Current working directory: %s\n", getCurrentDir())
	// fmt.Printf("Environment variables:\n")
	// fmt.Printf("ENDPOINT: %s\n", Endpoint)
	// fmt.Printf("ACCESS_KEY_ID: %s\n", AccessKeyId)
	// fmt.Printf("ACCESS_KEY_SECRET: %s\n", AccessKeySecret)
	// fmt.Printf("BUCKET_NAME: %s\n", BucketName)

	if Endpoint == "" || AccessKeyId == "" || AccessKeySecret == "" {
		panic("Environment variables ENDPOINT, ACCESS_KEY_ID, and ACCESS_KEY_SECRET must be set")
	}
	c, err := oss.New(Endpoint, AccessKeyId, AccessKeySecret)
	fmt.Println(Endpoint, AccessKeyId, AccessKeySecret)
	if err != nil {
		panic(err)
	}
	client = c
}
