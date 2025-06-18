package querytools

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func QueryInt(c *gin.Context, key string, def int) (int, error) {
	val := c.Query(key)
	num, err := strconv.Atoi(val)
	if err != nil {
		return def, err
	}
	return num, nil
}

func QueryList(c *gin.Context, key string, def []string) []string {
	val := c.Query(key)
	return strings.Split(val, ",")
}
