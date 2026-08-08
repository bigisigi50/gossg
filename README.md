# GoSSG (Go Static Site Generator)

GoSSG is a simplified Command Line Interface (CLI) application written in Go that acts as a Static Site Generator (SSG). It reads a directory of Markdown files with YAML frontmatter, parses them, and injects the content into HTML templates to generate a fully static, deployable website.

## Prerequisites

- [Go](https://golang.org/doc/install) (1.20 or later recommended)

## Downloading / Installation

Since this project is a standard Go module, you can clone the repository to your local machine:

```bash
git clone https://github.com/<your-username>/gossg.git
cd gossg
```

*(Note: Replace `<your-username>` with your actual GitHub username once you've published the repository.)*

### Build the CLI

To compile the application into a standalone executable:

```bash
go build -o gossg main.go
```

This will create a `gossg` binary in your current directory.

## Usage

The CLI supports the following flags to customize the directory paths:

- `--source`: Directory containing markdown files (default: `./content`)
- `--template`: Directory containing HTML templates (default: `./templates`)
- `--output`: Directory for the generated static site (default: `./public`)

### Running the Generator

If you just want to run it without building:

```bash
go run main.go
```

Or, if you built the binary:

```bash
./gossg
```

### Specifying Custom Directories

If your content and templates are stored elsewhere, you can pass custom paths using flags:

```bash
./gossg --source ./my-content --template ./my-themes --output ./build
```

## Directory Structure Overview

When fully implemented, the project relies on the following structure:

- **content/**: Place your `.md` files here. Subdirectories will be mirrored in the output directory.
- **templates/**: Place your `layout.html` and any other HTML base templates here.
- **public/**: The generated HTML files will be placed here.

## Next Steps

Currently, the CLI parses flags and validates that the source and template directories exist. The next phase of development will introduce:
1. Markdown parsing using `github.com/yuin/goldmark`.
2. YAML frontmatter extraction using `gopkg.in/yaml.v3`.
3. HTML template injection using Go's standard `html/template` package.
