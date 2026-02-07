package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/Parthipan-Natkunam/generate_image/internal/config"
	"github.com/Parthipan-Natkunam/generate_image/pkg/generator"
	"github.com/Parthipan-Natkunam/generate_image/pkg/providers/nanobanana"
	"github.com/Parthipan-Natkunam/generate_image/pkg/schema"
)

func main() {
	promptPtr := flag.String("prompt", "", "Text prompt for image generation")
	aspectRatioPtr := flag.String("aspect-ratio", "16:9", "Aspect ratio of the image")
	imageSizePtr := flag.String("image-size", "2K", "Size of the image")
	jsonPtr := flag.Bool("json", false, "Output result in JSON format")
	describePtr := flag.Bool("describe", false, "Output tool definition JSON")
	outputDirPtr := flag.String("output-dir", ".", "Directory to save generated images")

	flag.Parse()

	if *describePtr {
		jsonSchema, err := schema.GetJSON()
		if err != nil {
			log.Fatalf("Error generating schema: %v", err)
		}
		fmt.Println(jsonSchema)
		return
	}

	if *promptPtr == "" {
		fmt.Println("Error: --prompt is required unless using --describe")
		flag.Usage()
		os.Exit(1)
	}

	cfg, err := config.LoadConfig()
	if err != nil {
		handleError("Failed to load configuration", err, *jsonPtr)
	}

	// Initialize Provider (Defaulting to Nano Banana for now)
	provider := nanobanana.New(cfg.NanoBananaAPIKey)

	ctx := context.Background()
	opts := []generator.Option{
		generator.WithAspectRatio(*aspectRatioPtr),
		generator.WithImageSize(*imageSizePtr),
	}

	if *jsonPtr == false {
		fmt.Printf("Generating image with prompt: %q...\n", *promptPtr)
	}

	imageData, contentType, err := provider.Generate(ctx, *promptPtr, opts...)
	if err != nil {
		handleError("Generation failed", err, *jsonPtr)
	}

	ext := ".png"
	if contentType == "image/jpeg" {
		ext = ".jpg"
	}
	filename := fmt.Sprintf("img_%d%s", time.Now().Unix(), ext)
	outPath := filepath.Join(*outputDirPtr, filename)

	err = os.WriteFile(outPath, imageData, 0644)
	if err != nil {
		handleError("Failed to save image", err, *jsonPtr)
	}

	if *jsonPtr {
		output := map[string]string{
			"status": "success",
			"path":   outPath,
			"prompt": *promptPtr,
		}
		jsonOut, _ := json.Marshal(output)
		fmt.Println(string(jsonOut))
	} else {
		fmt.Printf("Success! Image saved to: %s\n", outPath)
	}
}

func handleError(msg string, err error, jsonMode bool) {
	if jsonMode {
		out := map[string]string{
			"status": "error",
			"error":  fmt.Sprintf("%s: %v", msg, err),
		}
		jsonOut, _ := json.Marshal(out)
		fmt.Println(string(jsonOut))
	} else {
		log.Fatalf("%s: %v", msg, err)
	}
	os.Exit(1)
}
