package main

import (
	"fmt"
	"github.com/jonas-p/go-shp"
)

func readShp(filename string) {
	shape, err := shp.Open(filename)
	if err != nil {
		panic(err)
	}
	defer shape.Close()

	// fields from the attribute table (DBF)
	fields := shape.Fields()

	// loop through all features in the shapefile
	for shape.Next() {
		n, p := shape.Shape()
		polygon, ook := p.(*shp.Polygon)
		if !ook {
			panic("Failed to type assert.")
		}
		for _, point := range polygon.Points {
			fmt.Println(point)
		}

		// print attributes
		for k, f := range fields {
			val := shape.ReadAttribute(n, k)
			fmt.Printf("\t%v: %v\n", f, val)
		}
		fmt.Println()
	}
}
