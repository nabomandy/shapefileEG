package main

import (
	"encoding/json"
	"fmt"
	"github.com/everystreet/go-shapefile"
	"os"
)

func main() {
	file, err := os.Open("C:\\test\\Z_NGII_N3A_G0010000.zip")
	//file, err := os.Open("C:\\test\\LARD_ADM_SECT_SGG_제주.zip")
	fmt.Println(err)
	stat, err := file.Stat()
	fmt.Println(err)
	scanner, err := shapefile.NewZipScanner(file, stat.Size(), stat.Name())
	fmt.Println(stat.Name())
	fmt.Println(err)

	//info, err := scanner.Info()
	//fmt.Println(info)

	err = scanner.Scan()

	for {
		record := scanner.Record()
		feature := record.GeoJSONFeature()
		jsonData, _ := json.Marshal(feature)
		fmt.Println(string(jsonData))
		if record == nil {
			break
		}

		// Each record contains a shape (from .shp file) and attributes (from .dbf file)
		fmt.Println(record)
	}

}
