package querytools

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func QueryInToSql(c *gin.Context, prefix, db_field, query_field, def string) string {
	s := c.Query(query_field)
	if s == "" {
		s = def
	}
	if s == "all" {
		return ""
	}
	vals := strings.Split(s, ",")

	// Prevent SQL injection by escaping single quotes
	for i, v := range vals {
		vals[i] = "'" + strings.ReplaceAll(v, "'", "''") + "'"
	}

	stmt := fmt.Sprintf("%s %s IN (%s)", prefix, db_field, strings.Join(vals, ", "))
	return stmt
}
