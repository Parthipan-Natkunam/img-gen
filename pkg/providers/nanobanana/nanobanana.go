package nanobanana

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Parthipan-Natkunam/generate_image/pkg/generator"
)

const (
	defaultEndpoint = "https://generativelanguage.googleapis.com/v1beta/models/gemini-3-pro-image-preview:generateContent"
	providerName    = "nano-banana-pro"
)

type Provider struct {
	apiKey   string
	client   *http.Client
	endpoint string
}

func New(apiKey string, opts ...ProviderOption) *Provider {
	p := &Provider{
		apiKey:   apiKey,
		client:   &http.Client{Timeout: 60 * time.Second},
		endpoint: defaultEndpoint,
	}
	for _, opt := range opts {
		opt(p)
	}
	return p
}

type ProviderOption func(*Provider)

func WithEndpoint(url string) ProviderOption {
	return func(p *Provider) {
		p.endpoint = url
	}
}

func WithClient(client *http.Client) ProviderOption {
	return func(p *Provider) {
		p.client = client
	}
}

func (p *Provider) Name() string {
	return providerName
}

// Gemini Request Structure
type GenerateRequest struct {
	Contents         []Content         `json:"contents"`
	GenerationConfig *GenerationConfig `json:"generationConfig,omitempty"`
}

type Content struct {
	Role  string `json:"role"`
	Parts []Part `json:"parts"`
}

type Part struct {
	Text string `json:"text,omitempty"`
}

type ImageConfig struct {
	AspectRatio string `json:"aspectRatio,omitempty"`
	ImageSize   string `json:"imageSize,omitempty"`
}

type GenerationConfig struct {
	ResponseModalities []string    `json:"responseModalities"`
	ImageConfig        ImageConfig `json:"imageConfig"`
}

// Gemini Response Structure
type GenerateResponse struct {
	Candidates []Candidate `json:"candidates"`
}

type Candidate struct {
	Content CandidateContent `json:"content"`
}

type CandidateContent struct {
	Parts []ResponsePart `json:"parts"`
}

type ResponsePart struct {
	InlineData *InlineData `json:"inlineData,omitempty"`
}

type InlineData struct {
	MimeType string `json:"mimeType"`
	Data     string `json:"data"`
}

// Generate sends a request to the Nano Banana (Gemini) API.
func (p *Provider) Generate(ctx context.Context, prompt string, opts ...generator.Option) ([]byte, string, error) {
	genOpts := &generator.GenerateOptions{}
	for _, opt := range opts {
		opt(genOpts)
	}

	reqPayload := GenerateRequest{
		Contents: []Content{
			{
				Role: "user",
				Parts: []Part{
					{Text: prompt},
				},
			},
		},
		GenerationConfig: &GenerationConfig{
			ResponseModalities: []string{"TEXT", "IMAGE"},
			ImageConfig: ImageConfig{
				AspectRatio: genOpts.AspectRatio,
				ImageSize:   genOpts.ImageSize,
			},
		},
	}

	jsonBody, err := json.Marshal(reqPayload)
	if err != nil {
		return nil, "", fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", p.endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, "", fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("x-goog-api-key", p.apiKey)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "img-gen-cli/1.0")

	resp, err := p.client.Do(req)
	if err != nil {
		return nil, "", fmt.Errorf("api request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, "", fmt.Errorf("api returned error %d: %s", resp.StatusCode, string(body))
	}

	var genResp GenerateResponse
	if err := json.NewDecoder(resp.Body).Decode(&genResp); err != nil {
		return nil, "", fmt.Errorf("failed to decode response: %w", err)
	}

	if len(genResp.Candidates) == 0 || len(genResp.Candidates[0].Content.Parts) == 0 {
		return nil, "", fmt.Errorf("no candidates returned")
	}

	// Look for the image part
	for _, part := range genResp.Candidates[0].Content.Parts {
		if part.InlineData != nil {
			data, err := base64.StdEncoding.DecodeString(part.InlineData.Data)
			if err != nil {
				return nil, "", fmt.Errorf("failed to decode base64 image data: %w", err)
			}
			return data, part.InlineData.MimeType, nil
		}
	}

	return nil, "", fmt.Errorf("no image data found in response")
}
