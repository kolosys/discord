# load_env

This example demonstrates basic usage of the library.

## Source Code

```go
package internal

import (
	"bufio"
	"os"
	"strings"
)

func LoadEnv(path string) {
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		key, value, ok := strings.Cut(line, "=")
		if !ok {
			continue
		}
		os.Setenv(strings.TrimSpace(key), strings.TrimSpace(value))
	}
}

```

## Running the Example

To run this example:

```bash
cd internal
go run load_env.go
```

## Expected Output

```
Hello from Proton examples!
```
