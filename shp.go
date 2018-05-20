package main

import (
	"fmt"
	"github.com/jonas-p/go-shp"
	"log"
)

func readShp(filename string) {
	var shape, err = shp.Open(filename)
	if err != nil {
		log.Println("error reading shapefile : ", filename)
		panic(err)
	}
	defer shape.Close()

	// fields from the attribute table (DBF)
	var fields = shape.Fields()

	// loop through all features in the shapefile
	for shape.Next() {
		var n, p = shape.Shape()
		var polygon, ok = p.(*shp.Polygon)
		if !ok {
			panic("Failed to type assert.")
		}
		for _, point := range polygon.Points {
			fmt.Println(point)
		}

		// print attributes
		for k, f := range fields {
			var val = shape.ReadAttribute(n, k)
			fmt.Printf("\t%v: %v\n", f, val)
		}
		fmt.Println()
	}
}
