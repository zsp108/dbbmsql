package utils

import (
	"os"

	"gopkg.in/yaml.v3"
)

func UnmarshalYAML(data []byte, v interface{}) error {
	return yaml.Unmarshal(data, v)
}

func MarshalYAML(v interface{}) ([]byte, error) {
	return yaml.Marshal(v)
}

// UnmarshalJSONFile unmarshals a JSON file into a struct
func UnmarshalYAMLFile(filePath string, v interface{}) error {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, v)
}

// MarshalJSONFile marshals a struct into a JSON file
func MarshalYAMLFile(filePath string, v interface{}) error {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		f, err := os.Create(filePath)
		if err != nil {
			return err
		}
		defer f.Close()
	}
	data, err := yaml.Marshal(v)
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
