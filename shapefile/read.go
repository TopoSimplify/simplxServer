package shapefile

import "github.com/jonas-p/go-shp"

func Read(filename string) []*ShapeRecord {
	var reader, err = shp.Open(filename)
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	var records []*ShapeRecord
	var fields = reader.Fields()

	for reader.Next() {
		var index, shape = reader.Shape()
		var shapes = flattenShp(index, shape)
		var attr = readAttribute(index, fields, reader)
		var record = &ShapeRecord{
			index:      index,
			shapes:     shapes,
			attributes: attr,
		}
		records = append(records, record)
	}

	return records
}
