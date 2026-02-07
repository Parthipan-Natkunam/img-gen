# Generate Image Skill

A custom Claude Code skill that wraps the image-gen CLI tool for AI-powered image generation using the Nano Banana Pro API.

## Features

- **Project Independent**: Works in any repository - the skill can be used across all your projects
- **Secure API Key Handling**: Reads API key from environment variables only (never from command line or logs)
- **Input Validation**: Validates all inputs to prevent injection attacks
- **Path Security**: Restricts output directory to current repository to prevent directory traversal
- **Dynamic Schema**: Skill definition is generated from the image-gen tool's `--describe` command
- **User-Specified Output**: Asks for the output directory within the current repo for each image generation

## Setup

### Prerequisites

Before using this skill, you must have the `image-gen` tool installed and available in your system PATH.

**To install image-gen:**

1. Build the tool from source:
   ```bash
   cd /path/to/image-gen
   go build -o image-gen ./cmd/img-gen/main.go
   ```

2. Add it to your PATH:

   **For Linux/macOS:**
   ```bash
   sudo mv image-gen /usr/local/bin/
   # or add to your PATH in ~/.bashrc or ~/.zshrc:
   export PATH="$PATH:/path/to/directory/with/image-gen"
   ```

   **For Windows:**
   - Move `image-gen.exe` to a directory that's already in your PATH (like `C:\Windows\System32`)
   - Or add the directory containing `image-gen.exe` to your PATH environment variable

3. Verify installation:
   ```bash
   image-gen --version
   ```

### 1. Set API Key

The skill requires the `NANOBANANA_API_KEY` environment variable. Add it to your shell profile:

**For Bash:**
```bash
echo 'export NANOBANANA_API_KEY="your-api-key-here"' >> ~/.bashrc
source ~/.bashrc
```

**For Zsh:**
```bash
echo 'export NANOBANANA_API_KEY="your-api-key-here"' >> ~/.zshrc
source ~/.zshrc
```

**For Windows Git Bash:**
```bash
echo 'export NANOBANANA_API_KEY="your-api-key-here"' >> ~/.bash_profile
source ~/.bash_profile
```

### 2. Run Setup

The setup script verifies the tool is installed and generates the skill.json:

```bash
bash .claude/skills/generate-image/setup.sh
```

This command:
- Verifies that `image-gen` is in your PATH
- Generates skill.json from the tool's `--describe` output
- Adds the `output_dir` parameter to the schema

### 3. Verify Installation

Check that Claude Code recognizes the skill:

```bash
# The skill should appear in the available skills list
# Use it with: /generate-image
```

## Usage

### In Claude Code

Use the skill with the `/generate-image` command:

```
/generate-image prompt="A serene mountain landscape at sunset"
```

With optional parameters:

```
/generate-image prompt="A futuristic city" aspect_ratio="16:9" image_size="4K" output_dir="./images/"
```

### Parameters

| Parameter | Type | Required | Default | Description |
|-----------|------|----------|---------|-------------|
| `prompt` | string | Yes | - | Text description of the image to generate |
| `aspect_ratio` | string | No | `16:9` | Aspect ratio: `1:1`, `16:9`, `4:3`, or `3:2` |
| `image_size` | string | No | `2K` | Image size: `1K`, `2K`, or `4K` |
| `output_dir` | string | No | `./images/` | Directory to save the image (must be within current repository) |

## Security Features

### 1. Environment Variable for API Key
- API key is read from `NANOBANANA_API_KEY` environment variable only
- Never passed as command-line argument (which could appear in process lists)
- Never logged or printed to console

### 2. Input Validation
- **Aspect Ratio**: Must match format `N:N` (e.g., `16:9`)
- **Image Size**: Must match format `NK` (e.g., `2K`)
- **Output Directory**: Validated against directory traversal attacks
- **Path Restriction**: Output directory must be within the project root

### 3. Automatic Binary Build
- Checks if source files are newer than binary
- Rebuilds automatically when needed
- Prevents use of outdated or tampered binaries

### 4. Error Handling
- Clear error messages for missing API key or tool
- JSON-formatted errors for parsing by Claude Code
- Exit codes indicate success/failure

## Maintenance

### Regenerating skill.json

If the image-gen tool's schema changes, regenerate the skill.json:

```bash
bash .claude/skills/generate-image/setup.sh
```

This ensures the skill definition stays in sync with the tool's capabilities.

### Updating the Tool

Since the skill uses the `image-gen` tool from your PATH, to update it:

1. Rebuild the tool:
   ```bash
   cd /path/to/image-gen
   go build -o image-gen ./cmd/img-gen/main.go
   ```

2. Replace the existing binary in your PATH with the new one

## Troubleshooting

### "image-gen tool is not installed or not in PATH"

Solution: Install the image-gen tool and add it to your PATH (see Prerequisites section above).

### "NANOBANANA_API_KEY environment variable is not set"

Solution: Set the API key in your shell profile (see Setup section above).

### "output_dir must be within the current repository"

The skill restricts output directories to the current repository for security. Use a relative path within the repo:

```bash
# Good (within current repo)
output_dir="./images/"
output_dir="./assets/generated/"
output_dir="./public/images/"

# Bad (outside current repo)
output_dir="/tmp/images"
output_dir="../../../other-project"
output_dir="/home/user/Desktop"
```

## Architecture

```
.claude/skills/generate-image/
├── skill.json       # Generated skill definition (do not edit manually)
├── setup.sh         # Setup script (verifies tool, generates skill.json)
├── run.sh           # Runtime script (validates inputs, executes tool)
└── README.md        # This file
```

### Workflow

1. **Setup Phase** (`setup.sh`):
   - Verifies `image-gen` tool is in PATH
   - Runs `image-gen --describe` to get tool definition
   - Augments schema with `output_dir` parameter
   - Writes skill.json

2. **Runtime Phase** (`run.sh`):
   - Checks if `image-gen` tool is in PATH
   - Validates API key is set
   - Parses JSON input from stdin
   - Validates and sanitizes inputs
   - Ensures output directory is within current repository
   - Creates output directory if needed
   - Executes `image-gen` with `--json` flag
   - Returns JSON output to Claude Code

## License

This skill wrapper follows the same license as the parent project.
