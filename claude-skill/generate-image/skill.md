# Generate Image Skill

Use this skill to generate AI images using the Nano Banana Pro API based on text prompts.

## When to Use

Invoke this skill when the user requests:
- Creating or generating images from text descriptions
- Making visual assets, illustrations, or graphics
- Generating images with specific aspect ratios or sizes
- Creating images for documentation, presentations, or projects

## Prerequisites

**IMPORTANT**: Before using this skill, verify these requirements are met:
1. The `image-gen` tool must be installed and in the system PATH (on windows, ensure `image-gen.exe` is accessible and use `image-gen.exe` in commands)
2. The `NANOBANANA_API_KEY` environment variable must be set

If either prerequisite is missing, the skill will fail with a clear error message instructing the user how to fix it.

## How to Invoke

Use the `/generate-image` command or invoke via the Skill tool:

```
/generate-image prompt="your image description here"
```

### Parameters

- `prompt` (required): Text description of the image to generate
- `aspect_ratio` (optional): One of `1:1`, `16:9`, `4:3`, or `3:2` (default: `16:9`)
- `image_size` (optional): One of `1K`, `2K`, or `4K` (default: `2K`)
- `output_dir` (optional): Directory path within the current repo to save the image (default: `./images/`)

### Examples

**Simple generation:**
```
/generate-image prompt="A serene mountain landscape at sunset with a lake reflecting the golden sky"
```

**With all parameters:**
```
/generate-image prompt="A futuristic city skyline at night" aspect_ratio="16:9" image_size="4K" output_dir="./assets/generated/"
```

**Square image for social media:**
```
/generate-image prompt="A colorful abstract pattern" aspect_ratio="1:1" image_size="2K" output_dir="./social-media/"
```

## Important Notes

1. **Security**: The output directory must be within the current repository. Attempts to save outside the repo will be rejected.

2. **Path Creation**: If the output directory doesn't exist, it will be created automatically.

3. **Project Independence**: This skill works across all projects/repositories where you have the image-gen tool installed.

4. **Error Handling**: If prerequisites are missing or parameters are invalid, you'll receive a JSON error response with clear instructions for the user.

5. **API Key**: The API key is always read from the environment variable and never exposed in command-line arguments or logs.

## Response Format

The skill returns JSON output with:
- `status`: "success" or "error"
- `image_path`: Path to the generated image (on success)
- `error`: Error message (on failure)
- Additional metadata about the generation

## User Communication

When invoking this skill:
1. Confirm what image you're generating
2. Specify the parameters being used (aspect ratio, size, output location)
3. After generation, inform the user of the saved file location
4. If errors occur, explain the error and provide the solution from the error message
