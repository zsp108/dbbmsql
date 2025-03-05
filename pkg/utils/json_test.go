package utils_test

import (
	"os"
	"testing"

	"github.com/zsp108/dbbmsql/pkg/utils"
)

func TestJson(t *testing.T) {
	filepath := "./jsontest.json"
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
	err = utils.MarshalJSONFile(filepath, data)
	if err != nil {
		t.Error(err)
	}
	t.Log("write to file:", filepath)
	t.Cleanup(func() {
		os.Remove(filepath)
	})
}

// Data struct for test
type Data struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	City string `json:"city"`
}
