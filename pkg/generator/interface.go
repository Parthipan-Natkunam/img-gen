package generator

import (
	"context"
)

type ImageGenerator interface {
	// Generate creates an image based on the prompt and options.
	// Returns the image data (bytes), the content type (e.g. "image/png"), and any error encountered.
	Generate(ctx context.Context, prompt string, opts ...Option) ([]byte, string, error)

	// Name returns the unique identifier for the provider.
	Name() string
}

type GenerateOptions struct {
	ImageSize string
	AspectRatio  string
}

// Option is a functional option for configuring GenerateOptions.
type Option func(*GenerateOptions)

func WithAspectRatio(ratio string) Option {
	return func(o *GenerateOptions) {
		o.AspectRatio = ratio
	}
}

func WithImageSize(size string) Option {
	return func(o *GenerateOptions) {
		o.ImageSize = size
	}
}
