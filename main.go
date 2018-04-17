package main

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func send(key string, b []byte) error {
	bucketName := os.Getenv("AWS_BUCKET")
	svc := s3.New(session.New(), aws.NewConfig())
	_, err := svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
		Body:   bytes.NewReader(b),
	})
	if err != nil {
		return err
	}
	return nil
}

func recv(key string) ([]byte, error) {
	bucketName := os.Getenv("AWS_BUCKET")
	svc := s3.New(session.New(), aws.NewConfig())
	output, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	result, err := ioutil.ReadAll(output.Body)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func main() {

	b := make([]byte, 1024)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println(err.Error())
	}

	key := "keyname"
	err = send(key, b)
	if err != nil {
		fmt.Println(err.Error())
	}
	b2, err := recv(key)
	if err != nil {
		fmt.Println(err.Error())
	}
	if bytes.Equal(b, b2) {
		fmt.Println("Worky")
	}

}
