# Installation

This guide will help you install and set up discord in your Go project.

## Prerequisites

Before installing discord, ensure you have:

- **Go 1.21** or later installed
- A Go module initialized in your project (run `go mod init` if needed)
- Access to the GitHub repository (for private repositories)

## Installation Steps

### Step 1: Install the Package

Use `go get` to install discord:

```bash
go get github.com/kolosys/discord
```

This will download the package and add it to your `go.mod` file.

### Step 2: Import in Your Code

Import the package in your Go source files:

```go
import "github.com/kolosys/discord"
```

### Multiple Packages

discord includes several packages. Import the ones you need:

```go
// 
import "github.com/kolosys/discord/commands"
```

```go
// 
import "github.com/kolosys/discord"
```

```go
// 
import "github.com/kolosys/discord/components"
```

```go
// 
import "github.com/kolosys/discord/embed"
```

```go
// 
import "github.com/kolosys/discord/events"
```

```go
// 
import "github.com/kolosys/discord/examples/basic/bot"
```

```go
// 
import "github.com/kolosys/discord/examples/basic/commands"
```

```go
// 
import "github.com/kolosys/discord/examples/basic/events"
```

```go
// 
import "github.com/kolosys/discord/examples/basic/routes"
```

```go
// 
import "github.com/kolosys/discord/examples/internal"
```

```go
// 
import "github.com/kolosys/discord/gateway"
```

```go
// 
import "github.com/kolosys/discord/models"
```

```go
// 
import "github.com/kolosys/discord/rest/internal"
```

```go
// 
import "github.com/kolosys/discord/rest"
```

```go
// 
import "github.com/kolosys/discord/server"
```

```go
// 
import "github.com/kolosys/discord/state"
```

```go
// 
import "github.com/kolosys/discord/types"
```

```go
// 
import "github.com/kolosys/discord/voice"
```

### Step 3: Verify Installation

Create a simple test file to verify the installation:

```go
package main

import (
    "fmt"
    "github.com/kolosys/discord"
)

func main() {
    fmt.Println("discord installed successfully!")
}
```

Run the test:

```bash
go run main.go
```

## Updating the Package

To update to the latest version:

```bash
go get -u github.com/kolosys/discord
```

To update to a specific version:

```bash
go get github.com/kolosys/discord@v1.2.3
```

## Installing a Specific Version

To install a specific version of the package:

```bash
go get github.com/kolosys/discord@v1.0.0
```

Check available versions on the [GitHub releases page](https://github.com/kolosys/discord/releases).

## Development Setup

If you want to contribute or modify the library:

### Clone the Repository

```bash
git clone https://github.com/kolosys/discord.git
cd discord
```

### Install Dependencies

```bash
go mod download
```

### Run Tests

```bash
go test ./...
```

## Troubleshooting

### Module Not Found

If you encounter a "module not found" error:

1. Ensure your `GOPATH` is set correctly
2. Check that you have network access to GitHub
3. Try running `go clean -modcache` and reinstall

### Private Repository Access

For private repositories, configure Git to use SSH or a personal access token:

```bash
git config --global url."git@github.com:".insteadOf "https://github.com/"
```

Or set up GOPRIVATE:

```bash
export GOPRIVATE=github.com/kolosys/discord
```

## Next Steps

Now that you have discord installed, check out the [Quick Start Guide](quick-start.md) to learn how to use it.

## Additional Resources

- [Quick Start Guide](quick-start.md)
- [API Reference](../api-reference/)
- [Examples](../examples/README.md)

