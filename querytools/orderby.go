package querytools

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type OrderByField struct {
	Dir   string
	Field string
}

type OrderBy struct {
	Fields []OrderByField
}

func (ob OrderBy) SqlStmt() string {
	if len(ob.Fields) == 0 {
		return ""
	}
	stmt := "ORDER BY "
	for _, field := range ob.Fields {
		stmt += fmt.Sprintf("%s %s,", field.Field, field.Dir)
	}
	return stmt[:len(stmt)-1]
}

func (ob *OrderBy) Append(field OrderByField) {
	ob.Fields = append(ob.Fields, field)
}

func OrderByFromString(allowed_fields map[string]struct{}, s string) OrderBy {
	ob := OrderBy{}
	for raw := range strings.SplitSeq(s, ",") {
		dir := "ASC"
		if raw[0] == '-' {
			raw = raw[1:]
			dir = "DESC"
		}
		if _, exists := allowed_fields[raw]; exists {
			ob.Append(OrderByField{Field: raw, Dir: dir})
		}
	}
	return ob
}

// QueryOrderBy returns a direction and field name using default direction and default field to fall back on.
// Uses order_by query param or def if the query param does not exist.
// Each field must exist in allowed_fields and a dash (-) implies descending.
// Returns an OrderBy instance wich has the SqlStmt function.
func QueryOrderBy(c *gin.Context, allowed_fields map[string]struct{}, def string) OrderBy {
	if c.Query("order_by") == "" {
		return OrderByFromString(allowed_fields, def)
	}
	return OrderByFromString(allowed_fields, c.Query("order_by"))
}
