package main

import (
	"context"
	"fmt"
	"log"
	"os"

	vision "cloud.google.com/go/vision/apiv1"
)

func main() {

	ctx := context.Background()

	// Create Vision Client
	visionClient, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer visionClient.Close()

	// Get Picture
	pictureFileName := "testdata/squirrel-birdfeeder.png"
	file, err := os.Open(pictureFileName)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}
	defer file.Close()

	// Call API
	image, err := vision.NewImageFromReader(file)
	if err != nil {
		log.Fatalf("Failed to read image: %v", err)
	}

	// Check if squirrel
	labels, err := visionClient.DetectLabels(ctx, image, nil, 5)
	if err != nil {
		log.Fatalf("Failed to detect Labels: %v", err)
	}

	// Log event
	// TODO something useful
	fmt.Println(labels)
}
