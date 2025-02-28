package utils

import (
	"encoding/json"
	"os"
)

func UnmarshalJSON(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}

func MarshalJSON(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

// UnmarshalJSONFile unmarshals a JSON file into a struct
func UnmarshalJSONFile(filePath string, v interface{}) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

// MarshalJSONFile marshals a struct into a JSON file
func MarshalJSONFile(filePath string, v interface{}) error {
	data, err := json.MarshalIndent(v, " ", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)

	// jsonStr := `{"name": "John", "age": 30}`
	// var data Data
	// json.Unmarshal([]byte(jsonStr), &data)
	// fmt.Println(data.Name, data.Age)
	// jdata, _ := json.Marshal(data)

	// fmt.Println(string(jdata))

	// filepath := "/root/workspace/github/pgy/aaa.json"

	// utils.UnmarshalJSONFile(filepath, &data)
	// fmt.Println(data.Name, data.Age)
}
