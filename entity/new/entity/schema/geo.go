package schema

import (
	"github.com/yohobala/taurus_go/encoding/geo"
	"github.com/yohobala/taurus_go/entity"
	"github.com/yohobala/taurus_go/entity/field"
)

type GeoEntity struct {
	entity.Entity
	ID              *field.Int64
	Point           *geo.PostGIS[geo.Point, geo.S4326, geo.GeomFromText]
	LineString      *geo.PostGIS[geo.LineString, geo.SDefault, geo.GeomFromText]
	Polygon         *geo.PostGIS[geo.Polygon, geo.SDefault, geo.GeomFromText]
	MultiPoint      *geo.PostGIS[geo.MultiPoint, geo.SDefault, geo.GeomFromText]
	MultiLineString *geo.PostGIS[geo.MultiLineString, geo.SDefault, geo.GeomFromText]
	MultiPolygon    *geo.PostGIS[geo.MultiPolygon, geo.SDefault, geo.GeomFromText]
	CircularString  *geo.PostGIS[geo.CircularString, geo.SDefault, geo.GeomFromText]
}

func (e *GeoEntity) Config() entity.EntityConfig {
	return entity.EntityConfig{
		AttrName: "geo_demo",
		Comment:  "Geo的类型测试",
	}
}

func (e *GeoEntity) Fields() []entity.FieldBuilder {
	return []entity.FieldBuilder{
		e.ID.
			Name("id").
			Primary(1).
			Comment("主键。").
			Sequence(entity.NewSequence("geo_id_seq")),
		e.Point.
			Name("point").
			Required().
			Comment("点"),
		e.LineString.
			Name("line_string").
			Required().
			Comment("线"),
		e.Polygon.
			Name("polygon").
			Required().
			Comment("多边形"),
		e.MultiPoint.
			Name("multi_point").
			Comment("多点"),
		e.MultiLineString.
			Name("multi_line_string").
			Comment("多线"),
		e.MultiPolygon.
			Name("multi_polygon").
			Comment("多多边形"),
		e.CircularString.
			Name("circular_string").
			Comment("圆弧"),
	}
}
