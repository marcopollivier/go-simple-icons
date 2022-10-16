# go-simple-icons

<picture>
	<source media="(prefers-color-scheme: dark)" srcset="https://si-cdn.vercel.app/simpleicons/white">
	<source media="(prefers-color-scheme: light)" srcset="https://si-cdn.vercel.app/simpleicons">
	<img align="left" alt="Simple Icons" height="60" src="https://si-cdn.vercel.app/simpleicons">
</picture>

[Simple Icons](https://github.com/simple-icons/simple-icons) lib for Golang's projects

[![CircleCI](https://dl.circleci.com/status-badge/img/gh/marcopollivier/go-simple-icons/tree/main.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/marcopollivier/go-simple-icons/tree/main)

## Usage 

````shell
go get github.com/marcopollivier/go-simple-icons@v0.0.1
````

```go
package main

import (
	"fmt"
	"github.com/marcopollivier/go-simple-icons"
)

func main() {
	icon := gosimpleicons.GetIcon("Simple Icons")
	fmt.Println(gosimpleicons.IconToJson(icon))
}
```

## License

MIT
