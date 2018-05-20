package shapefile

import "github.com/jonas-p/go-shp"

func isPolygon(shape shp.Shape) bool {
	_, ok := shape.(*shp.Polygon)
	return ok
}

func isPolyline(shape shp.Shape) bool {
	_, ok := shape.(*shp.PolyLine)
	return ok
}

func isPolylineZ(shape shp.Shape) bool {
	_, ok := shape.(*shp.PolyLineZ)
	return ok
}