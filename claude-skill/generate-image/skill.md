---
name: generate-image
description: Generate AI images using Nano Banana Pro API from text prompts. Use when the user requests creating images, illustrations, visual assets, or graphics from text descriptions.
disable-model-invocation: false
argument-hint: prompt="description" [aspect_ratio="16:9"] [image_size="2K"] [output_dir="./images/"]
---

# Generate Image Skill

Generate AI-powered images from text prompts using the Nano Banana Pro API via the `img-gen` tool.

## When to Use This Skill

Invoke this skill when the user requests:
- Creating or generating images from text descriptions
- Making visual assets, illustrations, or graphics
- Generating images with specific aspect ratios or sizes
- Creating images for documentation, presentations, or projects

## Prerequisites Check

Before invoking this skill, the following must be satisfied:
1. The `img-gen` tool must be installed and in the system PATH
2. The `NANOBANANA_API_KEY` environment variable must be set

**Important**: The skill will fail gracefully with clear error messages if prerequisites are missing. Do not attempt to verify prerequisites yourself—let the skill handle this and communicate any errors to the user.

## How to Invoke

Use the Skill tool with the following parameters:

```
Skill tool:
  skill: "generate-image"
  args: 'prompt="your image description" aspect_ratio="16:9" image_size="2K" output_dir="./images/"'
```

Directly invoke the executable with the proper parameters. You can get the required parameters by running `img-gen --describe` in the terminal.
For Windows, the executable is `img-gen.exe`.
do not invoke it with bash or shell commands. Always use the `img-gen` command with proper arguments.

### Parameters

- **prompt** (required): Text description of the image to generate
  - Be specific and detailed for better results
  - Include style, mood, composition details when relevant

- **aspect_ratio** (optional, default: "16:9"):
  - Valid values: `1:1`, `16:9`, `4:3`, `3:2`
  - Choose based on use case: `1:1` for social media, `16:9` for presentations, etc.

- **image_size** (optional, default: "2K"):
  - Valid values: `1K`, `2K`, `4K`
  - Larger sizes take longer but provide higher quality

- **output_dir** (optional, default: "./images/"):
  - Must be a relative path within the current repository
  - Directory will be created automatically if it doesn't exist
  - For security, paths outside the repository are rejected

## Communication Flow

**Before invoking**:
1. Acknowledge what image you're generating
2. Mention the key parameters (aspect ratio, size, output location)
3. If the user hasn't specified an output directory, use the default `./images/`

**After invoking**:
1. Check the JSON response status
2. If successful, inform the user of:
   - The saved file path (make it a clickable markdown link)
   - The image specifications used
3. If failed, explain the error and provide the solution from the error message

## Example Usage

**Simple generation**:
```
User: "Generate an image of a serene mountain landscape at sunset"

You: "I'll generate a serene mountain landscape at sunset image for you."

Skill tool:
  skill: "generate-image"
  args: 'prompt="A serene mountain landscape at sunset with a lake reflecting the golden sky, misty mountains in the background, peaceful atmosphere"'

After success: "I've generated your image and saved it to [images/mountain-landscape.png](images/mountain-landscape.png). The image was created in 16:9 aspect ratio at 2K resolution."
```

**With custom parameters**:
```
User: "Create a square profile picture of a cute robot, high quality"

You: "I'll create a square, high-quality image of a cute robot for you."

Skill tool:
  skill: "generate-image"
  args: 'prompt="A cute friendly robot with rounded edges, soft colors, appealing design, suitable for profile picture" aspect_ratio="1:1" image_size="4K" output_dir="./profile-images/"'
```

## Error Handling

The skill returns JSON responses. Handle these cases:

**Missing tool**:
- Error: "img-gen tool is not installed or not in PATH"
- Response: "The image generation tool isn't installed. You need to install the `img-gen` tool and add it to your system PATH. [See installation instructions](https://github.com/Parthipan-Natkunam/generate_image)."

**Missing API key**:
- Error: "NANOBANANA_API_KEY environment variable is not set"
- Response: "The NANOBANANA_API_KEY environment variable isn't configured. You need to add it to your shell profile. [Instructions provided in the error message]."

**Invalid output directory**:
- Error: "output_dir must be within the current repository"
- Response: "For security, images can only be saved within the current repository. Please use a relative path like `./images/` instead."

**Invalid parameters**:
- Explain what was invalid and the correct format
- Re-invoke with corrected parameters if you can infer the user's intent

## Important Notes

1. **Security**: Never attempt to save images outside the repository—the skill enforces this restriction

2. **Prompt Quality**: More detailed, specific prompts yield better results. When the user provides a brief description, consider enriching it with relevant details while preserving their core intent.

3. **Path Formatting**: Always format the saved image path as a clickable markdown link: `[filename.png](relative/path/to/filename.png)`

4. **No Prerequisites Verification**: Don't check for tool installation or API keys yourself—let the skill handle it and surface any errors to the user

5. **Async Operation**: Image generation may take a few seconds. Set appropriate expectations with the user.

## Verification

After successful generation:
- Confirm the file exists at the specified path
- Provide the user with the path as a markdown link they can click to view the image
