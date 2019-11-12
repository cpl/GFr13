# GFr13

## What

This package is a middleware/HandlerFunc for [Gin](https://github.com/gin-gonic/gin) that will abort any and all requests on Friday 13th with the following response:

```json
{
  "code": 513,
  "error": "server too spooked, today is Friday 13th"
}
```

## Why

Today, people are all about "high availability", but what about the web server wants? Maybe you don't want to risk a catastrophe on Friday 13th.
Well, if you ever had this problem, now you have a soltion.

No, but really, **this package is a joke**, don't use in any production environment.

# How

The package is simple to use. First you need to get it:

```bash
go get cpl.li/go/gfr13
```

Now you can use it in your web project:

```go
package main

import (
	"log"
	
	"cpl.li/go/gfr13"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	
	// if nil is used for Opts, then gfr13.DefaultOpts are used
	server.Use(gfr13.Handler(nil))
	
	if err := server.Run(":8099"); err != nil {
		log.Fatal(err)
	}
}
```

