package shapefile

import "github.com/jonas-p/go-shp"

func flattenShp(id int, shape shp.Shape) []*Shp {
	var shapes []*Shp
	if isPolygon(shape) {
		shapes = polygonShps(id, shape.(*shp.Polygon))
	} else if isPolyline(shape) {
		shapes = polylineShps(id, shape.(*shp.PolyLine))
	} else if isPolylineZ(shape) {
		shapes = polylineZShps(id, shape.(*shp.PolyLineZ))
	} else {
		panic("not implemented")
	}
	return shapes
}

func polylineShps(id int, shape *shp.PolyLine) []*Shp {
	var shapes []*Shp
	var shpPoints = shape.Points
	for partNum, o := range partRanges(shape.Parts, shape.NumPoints) {
		var i, j = o[0], o[1]
		shapes = append(shapes, NewShp(id, shpPoints[i:j], partNum, shp.POLYLINE))
	}
	return shapes
}

func polylineZShps(id int, shape *shp.PolyLineZ) []*Shp {
	var shapes []*Shp
	var shpPoints = shape.Points
	for partNum, o := range partRanges(shape.Parts, shape.NumPoints) {
		var i, j = o[0], o[1]
		var s = NewShp(id, shpPoints[i:j], partNum, shp.POLYLINEZ)
		s.ZArray = shape.ZArray[i:j]
		shapes = append(shapes, s)
	}
	return shapes
}

func polygonShps(id int, shape *shp.Polygon) []*Shp {
	var shapes []*Shp
	var shpPoints = shape.Points
	for partNum, o := range partRanges(shape.Parts, shape.NumPoints) {
		var i, j = o[0], o[1]
		shapes = append(shapes, NewShp(id, shpPoints[i:j], partNum, shp.POLYGON))
	}
	return shapes
}
