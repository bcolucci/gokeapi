package helpers

import (
	"encoding/json"
	"fmt"
)

func PrettyJSON(v interface{}) string {
	pretty, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		panic(err)
	}
	return string(pretty)
}

func PrintPrettyJSON(v interface{}) {
	fmt.Println(PrettyJSON(v))
}
