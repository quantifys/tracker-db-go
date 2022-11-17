package models

import (
	"context"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type GeoJsonPolygon struct {
	Type        string        `json:"type"`
	Coordinates [][][]float32 `json:"coordinates"`
}

func (point GeoJsonPolygon) GormDataType() string {
	return "geometry(Polygon)"
}

func (point GeoJsonPolygon) GormValue(ctx context.Context, db *gorm.DB) clause.Expr {
	return clause.Expr{
		SQL:  "ST_PointFromText(?)",
		Vars: []interface{}{fmt.Sprintf("POINT(%f %f)", point.Coordinates[0], point.Coordinates[1])},
	}
}

func (point *GeoJsonPolygon) Scan(v interface{}) error {
	data := GeoJsonPolygon{}
	err := json.Unmarshal([]byte(v.(string)), &data)
	*point = data
	return err
}

func (point *GeoJsonPolygon) GetCoordinates() [][]float32 {
	coords := point.Coordinates[0]
	coords = append(coords, point.Coordinates[0][0])
	return coords
}

func (point *GeoJsonPolygon) GetLength() int {
	return len(point.GetCoordinates())
}
