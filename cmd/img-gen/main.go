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
	"github.com/Parthipan-Natkunam/generate_image/pkg/watermark"
)

func main() {
	promptPtr := flag.String("prompt", "", "Text prompt for image generation")
	aspectRatioPtr := flag.String("aspect-ratio", "16:9", "Aspect ratio of the image")
	imageSizePtr := flag.String("image-size", "2K", "Size of the image")
	jsonPtr := flag.Bool("json", false, "Output result in JSON format")
	describePtr := flag.Bool("describe", false, "Output tool definition JSON")
	outputDirPtr := flag.String("output-dir", "./generated-images", "Directory to save generated images")

	// Watermark flags
	watermarkTextPtr := flag.String("watermark-text", "", "Text to use as watermark")
	watermarkImagePtr := flag.String("watermark-image", "", "Path to image file to use as watermark")
	watermarkPositionPtr := flag.String("watermark-position", "bottom-right", "Watermark position (top-left, top-center, top-right, left-center, center, right-center, bottom-left, bottom-center, bottom-right)")
	watermarkOpacityPtr := flag.Float64("watermark-opacity", 0.7, "Watermark opacity (0.0-1.0)")
	watermarkMarginPtr := flag.Int("watermark-margin", 20, "Watermark margin from edge in pixels")
	watermarkTextSizePtr := flag.Int("watermark-text-size", 24, "Font size for text watermark")
	watermarkTextColorPtr := flag.String("watermark-text-color", "#FFFFFF", "Text color in hex format (e.g., #FFFFFF)")
	watermarkScalePtr := flag.Float64("watermark-scale", 0.2, "Scale factor for image watermark (0.1-1.0)")

	flag.Parse()

	// Validate watermark flags (mutual exclusivity)
	if *watermarkTextPtr != "" && *watermarkImagePtr != "" {
		handleError("Cannot use both --watermark-text and --watermark-image",
			fmt.Errorf("flags are mutually exclusive"), *jsonPtr)
	}

	// Validate watermark image file exists and is valid (before image generation)
	if err := watermark.ValidateWatermarkImage(*watermarkImagePtr); err != nil {
		handleError("Invalid watermark image", err, *jsonPtr)
	}

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

	// Ensure output directory exists
	err = os.MkdirAll(*outputDirPtr, 0755)
	if err != nil {
		handleError("Failed to create output directory", err, *jsonPtr)
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

	// Apply watermark if requested
	finalImageData := imageData
	if *watermarkTextPtr != "" || *watermarkImagePtr != "" {
		wmConfig := watermark.Config{
			Text:      *watermarkTextPtr,
			Image:     *watermarkImagePtr,
			Position:  watermark.Position(*watermarkPositionPtr),
			Margin:    *watermarkMarginPtr,
			Opacity:   *watermarkOpacityPtr,
			TextSize:  *watermarkTextSizePtr,
			TextColor: *watermarkTextColorPtr,
			Scale:     *watermarkScalePtr,
		}

		watermarkedData, err := watermark.Apply(imageData, wmConfig)
		if err != nil {
			handleError("Failed to apply watermark", err, *jsonPtr)
		}
		finalImageData = watermarkedData

		if *jsonPtr == false {
			fmt.Println("Watermark applied successfully")
		}
	}

	ext := ".png"
	if contentType == "image/jpeg" {
		ext = ".jpg"
	}
	filename := fmt.Sprintf("img_%d%s", time.Now().Unix(), ext)
	outPath := filepath.Join(*outputDirPtr, filename)

	err = os.WriteFile(outPath, finalImageData, 0644)
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
