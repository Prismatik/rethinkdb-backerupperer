package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/defaults"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/davidbanham/required_env"
	"github.com/robfig/cron"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {
	required_env.Ensure(map[string]string{
		"AWS_ACCESS_KEY_ID":     "",
		"AWS_SECRET_ACCESS_KEY": "",
		"RETHINK_LOC":           "",
		"S3_BUCKET":             "",
	})

	if os.Getenv("CRON_STRING") != "" {
		c := cron.New()
		c.AddFunc(os.Getenv("CRON_STRING"), doBackup)
		c.Start()
		select {}
	} else {
		doBackup()
	}
}

func doBackup() {
	filename := time.Now().Format(time.RFC3339) + ".tar.gz"
	cmd := exec.Command("rethinkdb", "dump", "-c", os.Getenv("RETHINK_LOC"), "-f", filename)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	defaults.DefaultConfig.Region = aws.String("us-east-1")
	svc := s3.New(nil)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	params := &s3.PutObjectInput{
		Bucket: aws.String(os.Getenv("S3_BUCKET")),
		Key:    &filename,
		Body:   file,
	}

	if os.Getenv("SSE_KEY") != "" {
		params.SSECustomerKey = aws.String(os.Getenv("SSE_KEY"))
		params.SSECustomerAlgorithm = aws.String("AES256")
	}

	_, err = svc.PutObject(params)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully uploaded backup " + filename)

	err = os.Remove(filename)
	if err != nil {
		log.Fatal(err)
	}
}
