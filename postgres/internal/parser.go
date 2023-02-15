package handler

import (
	//"strings"
	"net/url"
	"log"
	"fmt"
	"context"
)


func GetFieldsAndValues(ctx context.Context, url *url.URL) ([]string, []string) {
	
	var fields, values []string

	qry := url.Query()
	for k, v := range qry {
	
		if k == "" || len(v) == 0 {
			log.Printf("Valye not found for key [%v]", k)
			continue
		}
	
		fields = append(fields, k)
		values = append(values, v[0])
	}

	if len(fields) == 0 || len(values) == 0 {		
		return nil, nil
	}

	return fields, values

}

func GetReadQryCond(ctx context.Context, url *url.URL) string {

	fields, values := GetFieldsAndValues(ctx, url)
	if len(fields) == 0 || len(values) == 0 || len(fields) != len(values) {
		return ""
	}

	var result string
	for i:=0; i<len(fields); i++ {
	
		if result != "" {
			result = result + " AND "
		}
	
		result = fmt.Sprintf(`%v%v='%v'`, result, fields[i], values[i])
	}

	return result
}
