package criterias

import (
	"net/http"
	"strings"
)

// DefaultCriteria to find data
func DefaultCriteria(r *http.Request) map[string][]interface{} {
	params := r.URL.Query()

	results := make(map[string][]interface{})
	searches := make(map[string]string)
	searchFields := make(map[string]string)

	for key := range params {
		if key == "search" {
			for _, pair := range strings.Split(params.Get(key), ";") {
				z := strings.Split(pair, ":")
				searches[z[0]] = z[1]
			}
		}
		if key == "searchFields" {
			for _, pair := range strings.Split(params.Get(key), ";") {
				z := strings.Split(pair, ":")
				searchFields[z[0]] = z[1]
			}
		}
	}
	for field, fieldValue := range searches {
		searchFieldKey, searchFieldsOk := searchFields[field]
		if searchFieldsOk == true && searchFieldKey != "" {
			fieldValueTypes := [1]string{fieldValue}
			if searchFieldKey == "like" {
				fieldValue = "%" + fieldValue + "%"
			}
			if searchFieldKey == "between" {
				fieldValueIn := strings.Split(fieldValue, ",")
				values := make([]interface{}, len(fieldValueIn))
				for i, v := range fieldValueIn {
					values[i] = v
				}
				results[field+" "+searchFieldKey+" ? AND ?"] = values
			} else if searchFieldKey == "in" {
				fieldValueIn := strings.Split(fieldValue, ",")
				values := make([]interface{}, len(fieldValueTypes))
				values[0] = fieldValueIn
				results[field+" "+searchFieldKey+" (?)"] = values
			} else {
				values := make([]interface{}, len(fieldValueTypes))
				values[0] = fieldValue
				results[field+" "+searchFieldKey+" ?"] = values
			}
		} else {
			fieldValueTypes := [1]string{fieldValue}
			values := make([]interface{}, len(fieldValueTypes))
			values[0] = fieldValue
			results[field+" = ?"] = values
		}
	}

	return results
}
