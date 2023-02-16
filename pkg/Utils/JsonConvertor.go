package Utils

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"Trademarkia/pkg/Models"
	"io/ioutil"
)

func Converter(file []byte) {

	var TtabProceedings Models.TtabProceedings
	if err := xml.Unmarshal(file, &TtabProceedings); err != nil { //Unmarshaling the XML to the Struct
		panic(err)
	}
	jsonData, err := json.MarshalIndent(TtabProceedings.ProceedingInfo, "", "  ") //Converting the data to Json
	if err != nil {
		fmt.Println("Error marshalling to JSON:", err)
		return
	}

	
	err = ioutil.WriteFile("ProceedingInfo.json", jsonData, 0644) //Writing the Json Data into the file
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}
