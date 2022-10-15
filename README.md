# go-simple-icons

<img src="https://raw.githubusercontent.com/simple-icons/simple-icons/develop/icons/simpleicons.svg#gh-light-mode-only" alt="Simple Icons" width=70><img src="https://raw.githubusercontent.com/simple-icons/simple-icons/develop/assets/readme/simpleicons-white.svg#gh-dark-mode-only" alt="Simple Icons" width=70>
Simple Icons lib for Golang's projects 

[![CircleCI](https://dl.circleci.com/status-badge/img/gh/marcopollivier/go-simple-icons/tree/main.svg?style=svg)](https://dl.circleci.com/status-badge/redirect/gh/marcopollivier/go-simple-icons/tree/main)

## Usage 

### Get the git-simple-icons module

````shell
$ go get github.com/marcopollivier/go-simple-icons@v0.0.1
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

