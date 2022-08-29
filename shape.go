package main

import (
	"encoding/json"
	"github.com/buger/jsonparser"
	"github.com/everystreet/go-shapefile"
	"os"
)

func main() {
	file, _ := os.Open("C:\\Users\\zeus\\Desktop\\오늘\\용도별건물정보\\AL_11110_D198_20220118.zip")
	//file, err := os.Open("C:\\Users\\zeus\\Desktop\\오늘\\용도별건물정보\\AL_11140_D198_20220118.zip")
	//file, err := os.Open("C:\\ Users\\zeus\\Desktop\\오늘\\용도별건물정보\\AL_11170_D198_20220118.zip")
	//file, err := os.Open("C:\\test\\LARD_ADM_SECT_SGG_제주.zip")
	stat, _ := file.Stat()
	scanner, _ := shapefile.NewZipScanner(file, stat.Size(), stat.Name())

	_ = scanner.Scan()

	for {
		record := scanner.Record()

		feature := record.GeoJSONFeature()
		jsonData, _ := json.Marshal(feature.Geometry)
		get, _, _, _ := jsonparser.Get(jsonData, "coordinates")
		var t1 [][][]float64
		_ = json.Unmarshal(get, &t1)
		//polygonResult := ""
		//for _, multiPolygon := range t1[0] {
		//	polygonResult += fmt.Sprintf(", %v %v", multiPolygon[0], multiPolygon[1])
		//
		//
		//	// Each record contains a shape (from .shp file) and attributes (from .dbf file)
		//	fmt.Println(record)
		//}
		if record == nil {
			break
		}
	}
}
