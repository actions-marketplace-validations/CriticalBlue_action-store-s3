package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	action := Action{
		Action:    os.Getenv("ACTION"),
		Bucket:    os.Getenv("BUCKET"),
		S3Class:   os.Getenv("S3_CLASS"),
		Key:       fmt.Sprintf("%s.zip", os.Getenv("KEY")),
		Artifacts: strings.Split(strings.TrimSpace(os.Getenv("ARTIFACTS")), "\n"),
	}

	switch act := action.Action; act {
	case PutAction:
		if len(action.Artifacts[0]) <= 0 {
			log.Fatal("No artifacts patterns provided")
		}

		if err := Zip(action.Key, action.Artifacts); err != nil {
			log.Fatal(err)
		}

		if err := PutObject(action.Key, action.Bucket, action.S3Class); err != nil {
			log.Fatal(err)
		}