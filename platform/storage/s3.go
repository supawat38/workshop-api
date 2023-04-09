package storage

import (
	"bytes"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	S3_REGION = os.Getenv("AWS_REGION")
	S3_BUCKET = os.Getenv("BUCKET_NAME")
)

func InitS3AWS() {
	S3_REGION = os.Getenv("AWS_REGION")
	S3_BUCKET = os.Getenv("BUCKET_NAME")
}

func ConnectToS3() (s *session.Session, err error) {

	// create an AWS session which can be
	// reused if we're uploading many files
	s, err = session.NewSession(&aws.Config{
		Region: aws.String(S3_REGION),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_ACCESS_KEY"),
			""),
	})

	if err != nil {
		log.Fatal(err)
	}

	return
}

func UploadFileToS3(s *session.Session, fileName string) error {

	// open the file for use
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// get the file size and read
	// the file content into a buffer
	fileInfo, _ := file.Stat()
	var size = fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	// config settings: this is where you choose the bucket,
	// filename, content-type and storage class of the file
	// you're uploading
	result, s3err := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(S3_BUCKET),
		Key:                  aws.String(fileName),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(size),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass:         aws.String("INTELLIGENT_TIERING"),
	})

	log.Println("Uploaded", result)

	return s3err
}
