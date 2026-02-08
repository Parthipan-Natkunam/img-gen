#!/usr/bin/env bash
# Image Generation Skill - Wrapper for img-gen tool
# This script executes the img-gen tool securely with input validation

set -euo pipefail

# Get the current working directory (the repository where the skill is being used)
REPO_ROOT="$(pwd)"

# Check if img-gen is available in PATH
if ! command -v img-gen &> /dev/null; then
    cat >&2 <<'EOF'
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
ERROR: img-gen tool is not installed or not in PATH
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

This skill requires the img-gen tool to be pre-installed and
available in your system PATH.

To install:
  1. Build the tool:
     cd /path/to/image-gen
     go build -o img-gen ./cmd/img-gen/main.go

  2. Add it to your PATH:
     - Linux/macOS: sudo mv img-gen /usr/local/bin/
     - Windows: Move img-gen.exe to a directory in your PATH

  3. Verify: img-gen --version

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
EOF
    echo '{"status":"error","error":"img-gen tool is not installed or not in PATH. Please install it first."}' >&2
    exit 1
fi

# Security check: Verify API key is set
if [[ -z "${NANOBANANA_API_KEY:-}" ]]; then
    cat >&2 <<'EOF'
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
ERROR: NANOBANANA_API_KEY environment variable is not set
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

To fix this, add the API key to your shell profile:

For Bash:
  echo 'export NANOBANANA_API_KEY="your-api-key-here"' >> ~/.bashrc
  source ~/.bashrc

For Zsh:
  echo 'export NANOBANANA_API_KEY="your-api-key-here"' >> ~/.zshrc
  source ~/.zshrc

For Windows Git Bash:
  echo 'export NANOBANANA_API_KEY="your-api-key-here"' >> ~/.bash_profile
  source ~/.bash_profile

Replace "your-api-key-here" with your actual Nano Banana API key.

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
EOF
    echo '{"status":"error","error":"NANOBANANA_API_KEY environment variable is not set. Please configure it in your shell profile."}' >&2
    exit 1
fi

# Parse JSON input from stdin
INPUT=$(cat)

# Extract parameters using jq (fallback to python if jq not available)
if command -v jq &> /dev/null; then
    PROMPT=$(echo "$INPUT" | jq -r '.prompt // empty')
    ASPECT_RATIO=$(echo "$INPUT" | jq -r '.aspect_ratio // "16:9"')
    IMAGE_SIZE=$(echo "$INPUT" | jq -r '.image_size // "2K"')
    OUTPUT_DIR=$(echo "$INPUT" | jq -r '.output_dir // "./images/"')
else
    # Fallback to python for JSON parsing
    PROMPT=$(echo "$INPUT" | python -c "import sys, json; data=json.load(sys.stdin); print(data.get('prompt', ''))")
    ASPECT_RATIO=$(echo "$INPUT" | python -c "import sys, json; data=json.load(sys.stdin); print(data.get('aspect_ratio', '16:9'))")
    IMAGE_SIZE=$(echo "$INPUT" | python -c "import sys, json; data=json.load(sys.stdin); print(data.get('image_size', '2K'))")
    OUTPUT_DIR=$(echo "$INPUT" | python -c "import sys, json; data=json.load(sys.stdin); print(data.get('output_dir', './images/'))")
fi

# Validate required parameter
if [[ -z "$PROMPT" ]]; then
    echo '{"status":"error","error":"prompt parameter is required"}' >&2
    exit 1
fi

# Security: Validate inputs to prevent injection attacks
# Check for suspicious characters in non-prompt parameters
if [[ "$ASPECT_RATIO" =~ [^\:0-9] ]] || [[ ! "$ASPECT_RATIO" =~ ^[0-9]+:[0-9]+$ ]]; then
    echo '{"status":"error","error":"Invalid aspect_ratio format. Must be in format N:N (e.g., 16:9)"}' >&2
    exit 1
fi

if [[ ! "$IMAGE_SIZE" =~ ^[0-9]+K$ ]]; then
    echo '{"status":"error","error":"Invalid image_size format. Must be in format NK (e.g., 2K)"}' >&2
    exit 1
fi

# Sanitize output directory path (prevent directory traversal)
OUTPUT_DIR=$(realpath -m "$OUTPUT_DIR")
if [[ "$OUTPUT_DIR" != "$REPO_ROOT"* ]]; then
    echo '{"status":"error","error":"output_dir must be within the current repository for security reasons"}' >&2
    exit 1
fi

# Create output directory if it doesn't exist
mkdir -p "$OUTPUT_DIR"

# Execute the img-gen tool with sanitized inputs
# Use --json flag for structured output
cd "$REPO_ROOT"
img-gen \
    --prompt "$PROMPT" \
    --aspect-ratio "$ASPECT_RATIO" \
    --image-size "$IMAGE_SIZE" \
    --output-dir "$OUTPUT_DIR" \
    --json

exit $?
