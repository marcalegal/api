package aws

// Access Key ID: AKIAIHN2WUGNKB6ODDYQ
// Secret Access Key: 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

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// AWSACCESSKEY ...
const (
	AWSACCESSKEY = "AKIAIHN2WUGNKB6ODDYQ"
	AWSSECRETKEY = "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"
	TOKEN        = ""
)

// S3Handler ...
type S3Handler struct {
	Bucket string
	Svc    *s3.S3
}

// GetImage ...
func (s *S3Handler) GetImage(path string) (string, error) {
	r, _ := s.Svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(path),
	})
	url, err := r.Presign(24 * 60 * time.Minute)
	if err != nil {
		fmt.Println("error presigning request", err)
		return "", err
	}

	return url, nil
}

// UploadImage ...
func (s *S3Handler) UploadImage(file *os.File, filename string, userID int, brandName string) (string, error) {
	defer file.Close()

	fileInfo, _ := file.Stat()
	var size = fileInfo.Size()

	path := fmt.Sprintf("/%d/%s/%s", userID, brandName, fileInfo.Name())
	ct := fileType(fileInfo.Name())

	params := &s3.PutObjectInput{
		Bucket:        aws.String(s.Bucket),
		Key:           aws.String(path),
		Body:          file,
		ContentLength: aws.Int64(size),
		ContentType:   aws.String(ct),
	}

	if _, err := s.Svc.PutObject(params); err != nil {
		fmt.Printf("bad response: %s", err)
		return "", err
	}

	return s.GetImage(path)
}

// NewAWSS3Handler ...
func NewAWSS3Handler(bucket string) *S3Handler {
	creds := credentials.NewStaticCredentials(AWSACCESSKEY, AWSSECRETKEY, TOKEN)
	_, err := creds.Get()
	if err != nil {
		fmt.Printf("bad credentials: %s", err)
		return nil
	}

	cfg := aws.NewConfig().WithRegion("us-west-2").WithCredentials(creds)
	svc := s3.New(session.New(), cfg)
	return &S3Handler{
		Bucket: bucket,
		Svc:    svc,
	}
}

// this function is not safe to create content-type,
// but for now is useful to continue with the development of the new API.
// TODO: find a better way to find file Content-Type
func fileType(ext string) string {
	if strings.Contains(ext, "png") {
		return "image/png"
	} else if strings.Contains(ext, "jpg") {
		return "image/jpeg"
	} else if strings.Contains(ext, "pdf") {
		return "application/pdf"
	}
	return "Unknown"
}
