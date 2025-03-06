package utils_test

import (
	"os"
	"testing"

	"github.com/zsp108/dbbmsql/pkg/utils"
)

func TestJson(t *testing.T) {
	jsonfilepath := "./jsontest.json"
	jsonStr := `{"name": "Bob", "age": 30, "city": "New York"}`

	var data Data
	// unmarshal to pointer
	err := utils.UnmarshalJSON([]byte(jsonStr), &data)
	if err != nil {
		t.Error(err)
	}
	t.Log(data.City)
	t.Log(data.Age)

	jsonStr1, err := utils.MarshalJSON(data)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(jsonStr1))

	// write to file
	err = utils.MarshalJSONFile(jsonfilepath, data)
	if err != nil {
		t.Error(err)
	}

	// read from file
	err = utils.UnmarshalJSONFile(jsonfilepath, &data)
	if err != nil {
		t.Error(err)
	}
	t.Log(data.City)
	t.Log(data.Age)

	t.Cleanup(func() {
		os.Remove(jsonfilepath)
	})

	// yaml test
	yamlfilepath := "./yamltest.yaml"
	yamlStr := `name: Bob
age: 30
city: New York`

	var data1 Data
	// unmarshal to pointer
	err = utils.UnmarshalYAML([]byte(yamlStr), &data1)
	if err != nil {
		t.Error(err)
	}
	t.Log(data1.City)
	t.Log(data1.Age)

	yamlStr1, err := utils.MarshalYAML(data1)
	if err != nil {
		t.Error(err)
	}
	t.Log(string(yamlStr1))

	// write to file
	err = utils.MarshalYAMLFile(yamlfilepath, data1)
	if err != nil {
		t.Error(err)
	}

	// read from file
	var data2 Data
	err = utils.UnmarshalYAMLFile(yamlfilepath, &data2)
	if err != nil {
		t.Error(err)
	}
	t.Log(data2.City)
	t.Log(data2.Age)
	t.Log("read from file:", yamlfilepath)
	t.Cleanup(func() {
		os.Remove(yamlfilepath)
	})
}

// Data struct for test
type Data struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}
