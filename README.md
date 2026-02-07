# Image Generation Tool

A CLI tool for generating images using AI providers (currently supporting Nano Banana Pro).

## Prerequisites

- [Go](https://go.dev/dl/) installed (version 1.22+ recommended)
- A Nano Banana Pro API Key

## Setup

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/Parthipan-Natkunam/generate_image.git
    cd generate_image
    ```

2.  **Set the API Key:**
    Set the `NANOBANANA_API_KEY` environment variable.

    **Windows (Command Prompt):**
    ```cmd
    set NANOBANANA_API_KEY=your_api_key_here
    ```

    **Windows (PowerShell):**
    ```powershell
    $env:NANOBANANA_API_KEY="your_api_key_here"
    ```

    **Linux/macOS:**
    ```bash
    export NANOBANANA_API_KEY=your_api_key_here
    ```

## Usage

Run the tool using `go run`:

```bash
go run cmd/img-gen/main.go --prompt "A futuristic city with flying cars"
```

### Options

| Flag | Description | Default |
| :--- | :--- | :--- |
| `--prompt` | Text prompt for image generation (Required) | |
| `--negative-prompt` | Negative prompt to avoid certain elements | "" |
| `--width` | Image width | 1024 |
| `--height` | Image height | 1024 |
| `--model` | Model ID to use | "nano-banana-pro-v1" |
| `--output-dir` | Directory to save generated images | "." |
| `--describe` | Output tool definition JSON (for integration) | false |
| `--json` | Output result in JSON format | false |

### Examples

**Generate an image with specific dimensions:**
```bash
go run cmd/img-gen/main.go --prompt "A cat in space" --width 512 --height 512
```

**Save to a specific directory:**
```bash
go run cmd/img-gen/main.go --prompt "Sunset over mountains" --output-dir ./images
```
