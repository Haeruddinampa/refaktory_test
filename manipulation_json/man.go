package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type ListData struct {
	inventory_id int    `json:"id"`
	Name         string `json:"name"`
	type1        string `json:"type"`
	tags         string `json:"tags"`
}

func main() {
	fmt.Println(ReadJson())
}

func ReadJson() []ListData {
	var list_data []ListData
	var respData []ListData
	jsonFile, err := os.Open("test.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &list_data)

	for i := 0; i < len(list_data); i++ {
		lst := ListData{
			inventory_id: list_data[i].inventory_id,
			Name:         list_data[i].Name,
			type1:        list_data[i].type1,
			tags:         list_data[i].tags,
		}

		respData = append(respData, lst)
	}
	return respData
}
