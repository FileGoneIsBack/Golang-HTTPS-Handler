package functions

import (
	"encoding/json"
	"net/http"
)

func GetQuery(r *http.Request, key string) (bool, string) {
	if !r.URL.Query().Has(key) {
		return false, ""
	}
	return true, r.URL.Query().Get(key)
}

func GetQuerys(w http.ResponseWriter, r *http.Request, query map[string]bool) map[string]string {
	values := make(map[string]string)
	for name, required := range query {
		if !r.URL.Query().Has(name) && !required {
			continue
		} else if !r.URL.Query().Has(name) && required {
			json.NewEncoder(w).Encode(map[string]any{"error": true, "message": "missing required query field \"" + name + "\""})
			return nil
		}
		values[name] = r.URL.Query().Get(name)
	}
	return values
}
