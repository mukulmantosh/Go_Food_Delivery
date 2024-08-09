package database

import (
	"fmt"
	"log"
	"strings"
)

func (d *DB) loadModel(model any, queryType string) any {
	switch queryType {
	case "SELECT":
		return d.db.NewSelect().Model(model)
	case "DELETE":
		return d.db.NewDelete().Model(model)
	case "UPDATE":
		return d.db.NewUpdate().Model(model)
	case "INSERT":
		return d.db.NewInsert().Model(model)
	default:
		return nil
	}
}

func (d *DB) whereCondition(filter Filter) string {
	var whereClauses []string
	for key, value := range filter {
		var formattedValue string
		switch v := value.(type) {
		case string:
			// Quote string values
			formattedValue = fmt.Sprintf("'%s'", v)
		case int64:
			formattedValue = fmt.Sprintf("%d", v)
		default:
			log.Fatal("DB::Query:: Un-handled type for where condition!")

		}
		whereClauses = append(whereClauses, fmt.Sprintf("%s = %s", key, formattedValue))
	}

	var result string
	if len(whereClauses) > 0 {
		result = strings.Join(whereClauses, " AND ")
	}

	return result
}
