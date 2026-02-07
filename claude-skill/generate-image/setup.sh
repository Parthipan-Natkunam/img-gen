#!/usr/bin/env bash
# Setup script for generate_image skill
# This script verifies the image-gen tool is installed and generates skill.json

set -euo pipefail

# Get paths
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"

# Check if image-gen is available in PATH
if ! command -v image-gen &> /dev/null; then
    cat >&2 <<'EOF'
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
ERROR: image-gen tool is not installed or not in PATH
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

This skill requires the image-gen tool to be pre-installed.

To install:
  1. Build the tool from source:
     cd /path/to/image-gen
     go build -o image-gen ./cmd/img-gen/main.go

  2. Add it to your PATH:

     For Linux/macOS:
       sudo mv image-gen /usr/local/bin/
       # or add to your PATH in ~/.bashrc or ~/.zshrc:
       export PATH="$PATH:/path/to/directory/with/image-gen"

     For Windows:
       - Move image-gen.exe to a directory in your PATH
       - Or add the directory to your PATH environment variable

  3. Verify installation:
     image-gen --version

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
EOF
    exit 1
fi

echo "Found image-gen tool in PATH"
echo "Generating skill.json from tool definition..."
TOOL_DEF=$(image-gen --describe)

# Extract the base definition and add output_dir parameter
cat > "$SCRIPT_DIR/skill.json" <<EOF
{
  "name": "generate-image",
  "description": "Generate images using the Nano Banana Pro API based on text prompts. PREREQUISITE: Requires image-gen tool to be installed and available in PATH. Supports various aspect ratios and image sizes. Images are saved within the current repository.",
  "command": "bash",
  "args": [".claude/skills/generate_image/run.sh"],
  "input_schema": $(echo "$TOOL_DEF" | python -c "
import sys, json
data = json.load(sys.stdin)
schema = data['input_schema']
# Add output_dir parameter
schema['properties']['output_dir'] = {
    'type': 'string',
    'description': 'Directory within the current repository to save the generated image (e.g., ./images/, ./assets/generated/). Must be a relative path within the repo. Will be created if it does not exist.'
}
print(json.dumps(schema, indent=2))
")
}
EOF

echo "✓ Skill setup complete!"
echo ""
echo "To use this skill in Claude Code:"
echo "  - Ensure image-gen is installed and in your PATH"
echo "  - Ensure NANOBANANA_API_KEY is set in your environment"
echo "  - Use: /generate-image"
echo ""
echo "To regenerate skill.json after schema changes, run:"
echo "  bash .claude/skills/generate-image/setup.sh"
