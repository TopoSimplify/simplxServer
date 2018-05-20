package shapefile

import "github.com/jonas-p/go-shp"

type Shp struct {
	Id          int
	Coordinates []shp.Point
	PartNum     int
	Type        shp.ShapeType
	ZArray      []float64
}

type Attr struct {
	Id   int
	Type shp.ShapeType
}

type ShapeRecord struct {
	index      int
	shapes     []*Shp
	attributes []string
}

func NewShp(id int, coords []shp.Point, partnum int, shptype shp.ShapeType) *Shp {
	return &Shp{
		Id:          id,
		Coordinates: coords,
		PartNum:     partnum,
		Type:        shptype,
	}
}
