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
				"negative_prompt": map[string]string{
					"type":        "string",
					"description": "Items or concepts to exclude from the image.",
				},
				"width": map[string]interface{}{
					"type":        "integer",
					"description": "The width of the generated image.",
					"default":     1024,
				},
				"height": map[string]interface{}{
					"type":        "integer",
					"description": "The height of the generated image.",
					"default":     1024,
				},
				"aspect_ratio": map[string]string{
					"type":        "string",
					"description": "The aspect ratio of the image (e.g., '16:9', '1:1').",
					"enum":        []string{"1:1", "16:9", "4:3", "3:2"},
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
