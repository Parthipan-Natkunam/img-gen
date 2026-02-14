package schema

import "encoding/json"

// ToolDefinition represents a tool definition compatible with Claude/OpenAI.
type ToolDefinition struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	InputSchema InputSchema `json:"input_schema"`
}

type InputSchema struct {
	Type       string                 `json:"type"`
	Properties map[string]interface{} `json:"properties"`
	Required   []string               `json:"required"`
}

func GetToolDefinition() ToolDefinition {
	return ToolDefinition{
		Name:        "generate_image",
		Description: "Generate an image using the Nano Banana Pro API based on a text prompt.",
		InputSchema: InputSchema{
			Type: "object",
			Properties: map[string]interface{}{
				"prompt": map[string]string{
					"type":        "string",
					"description": "The text description of the image to generate.",
				},
				"aspect_ratio": map[string]interface{}{
					"type":        "string",
					"description": "The aspect ratio of the image (e.g., '16:9', '1:1').",
					"enum":        []string{"1:1", "16:9", "4:3", "3:2"},
				},
				"image_size": map[string]interface{}{
					"type":        "string",
					"description": "The size of the image (e.g., '1K', '2K', '4K').",
					"enum":        []string{"1K", "2K", "4K"},
				},
				"watermark_text": map[string]string{
					"type":        "string",
					"description": "Optional text to use as watermark on the generated image. Cannot be used with watermark_image.",
				},
				"watermark_image": map[string]string{
					"type":        "string",
					"description": "Optional path to an image file to use as watermark (supports PNG, JPEG, SVG). Cannot be used with watermark_text.",
				},
				"watermark_position": map[string]interface{}{
					"type":        "string",
					"description": "Position of the watermark on the image. Default: 'bottom-right'.",
					"enum":        []string{"top-left", "top-center", "top-right", "left-center", "center", "right-center", "bottom-left", "bottom-center", "bottom-right"},
				},
				"watermark_opacity": map[string]interface{}{
					"type":        "number",
					"description": "Opacity level of the watermark (0.0-1.0). Default: 0.7.",
					"minimum":     0.0,
					"maximum":     1.0,
				},
				"watermark_margin": map[string]interface{}{
					"type":        "integer",
					"description": "Margin from edge in pixels. Default: 20.",
					"minimum":     0,
				},
				"watermark_text_size": map[string]interface{}{
					"type":        "integer",
					"description": "Font size for text watermark in pixels. Default: 24.",
					"minimum":     1,
				},
				"watermark_text_color": map[string]string{
					"type":        "string",
					"description": "Hex color code for text watermark (e.g., '#FFFFFF'). Default: '#FFFFFF'.",
				},
				"watermark_scale": map[string]interface{}{
					"type":        "number",
					"description": "Scale factor for image watermark as a percentage of base image width (0.1-1.0). Default: 0.2.",
					"minimum":     0.1,
					"maximum":     1.0,
				},
			},
			Required: []string{"prompt"},
		},
	}
}

func GetJSON() (string, error) {
	def := GetToolDefinition()
	bytes, err := json.MarshalIndent(def, "", "  ")
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
