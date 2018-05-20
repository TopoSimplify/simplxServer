package shapefile

import (
	"github.com/jonas-p/go-shp"
)

func readAttribute(index int, fields []shp.Field, reader *shp.Reader) []string {
	var columns []string
	for field := range fields {
		columns = append(columns, reader.ReadAttribute(index, field))
	}
	return columns
}
