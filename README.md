# gignore-cli

A powerful CLI tool for managing ignore files (`.gitignore`, `.dockerignore`, etc.) with intelligent conflict detection, rule analysis, and automatic optimization.

Powered by [gignore](https://github.com/MoonMoon1919/gignore)

## Features

- **Create** ignore files from scratch
- **Add/Delete** rules by type (files, directories, extensions, glob patterns)
- **Analyze** ignore files for conflicts and redundancies
- **Auto-fix** conflicts and optimize rule ordering
- **Move** rules to specific positions
- Support for multiple ignore file formats

## Installation

```bash
go install github.com/MoonMoon1919/gignore-cli@latest
```

## Quick Start

```bash
# Create a new .gitignore file
gignore-cli create

# Add some rules
gignore-cli add file --filepath "config.json"
gignore-cli add extension --extension "log"
gignore-cli add directory --name "node_modules" --mode "directory"

# Analyze for conflicts and fix them
gignore-cli analyze --fix
```

## Usage

### Creating Ignore Files

```bash
# Create .gitignore in current directory
gignore-cli create

# Create ignore file at specific path
gignore-cli create --path .dockerignore
```

### Adding Rules

#### File Rules
```bash
# Ignore a specific file
gignore-cli add file --filepath "secrets.env"

# Allow a file (exclusion rule)
gignore-cli add file --filepath "important.log" --action exclude
```

#### Directory Rules
```bash
# Ignore directory contents
gignore-cli add directory --name "build" --mode "directory"

# Ignore directory recursively
gignore-cli add directory --name "cache" --mode "recursive"

# Ignore directories anywhere in the tree
gignore-cli add directory --name ".DS_Store" --mode "anywhere"
```

Available directory modes:
- `directory` - Match directory and its contents
- `recursive` - Match directory recursively
- `children` - Match only directory contents
- `anywhere` - Match directory name anywhere in tree
- `root` - Match only at repository root

#### Extension Rules
```bash
# Ignore all .log files
gignore-cli add extension --extension "log"

# Allow specific extension (exclusion)
gignore-cli add extension --extension "keep" --action exclude
```

#### Glob Pattern Rules
```bash
# Complex patterns
gignore-cli add glob --pattern "*.tmp.*"
gignore-cli add glob --pattern "test-*"
```

### Deleting Rules

Use the same syntax as adding, but with `delete`:

```bash
gignore-cli delete file --filepath "config.json"
gignore-cli delete extension --extension "log"
gignore-cli delete directory --name "build" --mode "directory"
gignore-cli delete glob --pattern "*.tmp.*"
```

### Analysis and Optimization

```bash
# Analyze file for conflicts
gignore-cli analyze

# Analyze and automatically fix issues
gignore-cli analyze --fix
```

The analyzer detects:
- Conflicting rules (ignore vs allow)
- Redundant patterns
- Suboptimal rule ordering
- Unreachable rules

### Manual Rule Management

```bash
# Move a rule before another rule
gignore-cli move --source-pattern "*.log" --destination-pattern "build/" --direction "before"

# Move a rule after another rule
gignore-cli move --source-pattern "cache/" --destination-pattern "*.tmp" --direction "after"
```

## Global Options

All commands support these options:

- `--path` - Path to ignore file (default: `.gitignore`)
- `--action` - Rule action: `include` (ignore) or `exclude` (allow) (default: `include`)

## Examples

### Setting up a Node.js project

```bash
gignore-cli create
gignore-cli add directory --name "node_modules" --mode "directory"
gignore-cli add directory --name "dist" --mode "directory"
gignore-cli add file --filepath ".env"
gignore-cli add extension --extension "log"
gignore-cli add glob --pattern "*.local"
```

### Setting up a Go project

```bash
gignore-cli create
gignore-cli add file --filepath ".env"
gignore-cli add extension --extension "exe"
gignore-cli add glob --pattern "*.test"
```

### Working with Docker

```bash
gignore-cli create --path .dockerignore
gignore-cli add directory --path .dockerignore --name ".git" --mode "recursive"
gignore-cli add file --path .dockerignore --filepath "README.md"
gignore-cli add glob --path .dockerignore --pattern "test*"
```

### Complex workflow with analysis

Start with an existing file with problematic rules:
```
debug.log
*.log  # makes debug.log rule redundant
!important.log
*.log # duplicate rule
```

Then analyze and fix the conflicts
```bash
# Analyze shows conflicts and redundancies
gignore-cli analyze

# Fix automatically
gignore-cli analyze --fix
```

## Rule Types

### File Rules
Target specific files by exact path.

### Directory Rules
Target directories with different matching modes:
- **directory**: `dirname/` - matches directory and contents
- **recursive**: `dirname/**` - matches recursively
- **children**: `dirname/*` - matches only direct children
- **anywhere**: `**/dirname/` - matches directory name anywhere
- **root**: `/dirname/` - matches only at repository root

### Extension Rules
Target all files with specific extensions.

### Glob Rules
Use glob patterns for complex matching rules.

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

If you encounter any issues or have questions, please [open an issue](https://github.com/MoonMoon1919/gignore-cli/issues) on GitHub.
