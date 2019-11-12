package gfr13

import (
	"time"

	"github.com/gin-gonic/gin"
)

const (
	StatusCode = 513
)

type Opts struct {
	UseUTC         bool
	UseJSON        bool
	AbortJSONKey   string
	AbortJSONValue string
}

var DefaultOpts = Opts{
	UseUTC:         true,
	UseJSON:        true,
	AbortJSONKey:   "error",
	AbortJSONValue: "server too spooked, today is Friday 13th",
}

func Handler(opts *Opts) gin.HandlerFunc {
	return func(c *gin.Context) {
		if opts == nil {
			opts = &DefaultOpts
		}

		now := time.Now()
		if opts.UseUTC {
			now = now.UTC()
		}

		if now.Day() == 13 && now.Weekday() == time.Friday {
			if opts.UseJSON {
				c.AbortWithStatusJSON(StatusCode, gin.H{
					opts.AbortJSONKey: opts.AbortJSONValue,
					"code":            StatusCode,
				})
			} else {
				c.AbortWithStatus(StatusCode)
			}
			return
		}

		c.Next()
	}
}
