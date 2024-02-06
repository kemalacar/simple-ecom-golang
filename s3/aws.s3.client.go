package s3

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kemalacar/go-ecom/models"
	"log"
	"strings"
	"time"
)

func UploadBase64(ctx *gin.Context, pm *models.Product) {
	// The session the S3 Uploader will use
	sess := ctx.MustGet("sess").(*session.Session)
	uploader := s3manager.NewUploader(sess)

	for i, image := range pm.Images {
		response := UploadToS3(uploader, image)
		pm.Images[i].Big = response
	}

	for _, sp := range pm.StoreProducts {
		for i, image := range sp.Images {
			sp.Images[i].Big = UploadToS3(uploader, models.ProductImage{Big: image.Big, Extension: image.Extension})
		}
	}

}

func UploadToS3(uploader *s3manager.Uploader, image models.ProductImage) (response string) {

	b64data := image.Big[strings.IndexByte(image.Big, ',')+1:]

	decode, err := base64.StdEncoding.DecodeString(b64data)
	random, _ := uuid.NewRandom()
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String("my-market"),
		Key:    aws.String(random.String() + "." + image.Extension),
		Body:   bytes.NewReader(decode),
		ACL:    aws.String("public-read"),
	})
	if err != nil {
	}

	return aws.StringValue(&result.Location)
}

func DeleteObjects(ctx *gin.Context, list []string) {
	var objects []*s3.ObjectIdentifier
	for _, obj := range list {
		objects = append(objects, &s3.ObjectIdentifier{Key: aws.String(obj)})
	}

	sess := ctx.MustGet("sess").(*session.Session)
	svc := s3.New(sess)

	_, err := svc.DeleteObjects(&s3.DeleteObjectsInput{
		Bucket: aws.String("my-market"),
		Delete: &s3.Delete{
			Objects: objects,
		},
	})
	if err != nil {
		fmt.Printf("Failed: %s\n", err)
	}

	if err != nil {
	}

}

func GetPreSignURL(ctx *gin.Context) string {
	sess := ctx.MustGet("sess").(*session.Session)
	svc := s3.New(sess)

	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String("my-market"),
		Key:    aws.String(ctx.Param("name")),
		ACL:    aws.String("public-read"),
	})

	urlStr, err := req.Presign(2 * time.Minute)

	if err != nil {
		log.Println("Failed to sign request", err)
	}

	log.Println("The URL is", urlStr)
	return urlStr

}

var MYINDEX uint64 = 0

func MyTest(dto *models.BrandDto) {

	time.Sleep(5 * time.Second)
	MYINDEX = MYINDEX + 1
	dto.Id = MYINDEX
	fmt.Println(MYINDEX + 1)

}
