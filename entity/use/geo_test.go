package use

import (
	"context"
	"fmt"
	"testing"

	"github.com/yohobala/taurus_go/encoding/geo"
	"github.com/yohobala/taurus_go/testutil/unit"
	"github.com/yohobala/taurus_go/tlog"
)

func TestGeo(t *testing.T) {

	// 全部geo已有字段类型创建测试。
	// All existing field type creation tests.
	t.Run("All existing field type creation tests.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()

		// Point
		p, err := geo.NewPoint(1, 2)
		unit.Must(t, err)

		// LineString
		line, err := geo.NewLineString([][]float64{{0, 0}, {4, 4}})
		unit.Must(t, err)

		// Polygon
		p1, err := geo.NewPoint(1, 1)
		unit.Must(t, err)
		p2, err := geo.NewPoint(5, 5)
		unit.Must(t, err)
		p3, err := geo.NewPoint(1, 9)
		unit.Must(t, err)
		polygon, err := geo.NewPolygonByPoint([]geo.Point{*p1, *p2, *p3, *p1})
		unit.Must(t, err)

		// MultiPoint
		p4, err := geo.NewPoint(2, 2)
		unit.Must(t, err)
		p5, err := geo.NewPoint(6, 6)
		unit.Must(t, err)
		multiPoint, err := geo.NewMultiPoint([]geo.Point{*p4, *p5})
		unit.Must(t, err)

		// MultiLineString
		line1, err := geo.NewLineString([][]float64{{0, 0}, {4, 4}})
		unit.Must(t, err)
		line2, err := geo.NewLineString([][]float64{{1, 1}, {5, 5}})
		unit.Must(t, err)
		multiLineString, err := geo.NewMultiLineString([]geo.LineString{*line1, *line2})
		unit.Must(t, err)

		// MultiPolygon
		polygon1, err := geo.NewPolygonByPoint([]geo.Point{*p1, *p2, *p3, *p1})
		unit.Must(t, err)
		polygon2, err := geo.NewPolygonByPoint([]geo.Point{*p4, *p5, *p3, *p4})
		unit.Must(t, err)
		multiPolygon, err := geo.NewMultiPolygon([]geo.Polygon{*polygon1, *polygon2})
		unit.Must(t, err)

		// CircularString
		circularString, err := geo.NewCircularString([][]float64{{1, 2}, {3, 4}, {5, 6}})
		unit.Must(t, err)

		g, err := db.Geos.Create(
			*p, *line, *polygon,
			db.Geos.WithMultiPoint(*multiPoint),
			db.Geos.WithMultiLineString(*multiLineString),
			db.Geos.WithMultiPolygon(*multiPolygon),
			db.Geos.WithCircularString(*circularString),
		)
		unit.Must(t, err)
		err = db.Save(ctx)
		unit.Must(t, err)
		tlog.Print(g)
	})

	// 查询第一条数据。
	// Query the first entry.
	t.Run("Query the first entry.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()

		g, err := db.Geos.First(ctx)
		unit.Must(t, err)
		fmt.Println(g)
	})

	// 查询所有数据。
	// Query all data.
	t.Run("Query all data.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()

		gs, err := db.Geos.ToList(ctx)
		unit.Must(t, err)
		fmt.Println(gs)
	})

	// 更新数据。
	// Update data.
	t.Run("Update data.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()

		g, err := db.Geos.First(ctx)
		unit.Must(t, err)

		np, err := geo.NewPoint(3, 4)
		unit.Must(t, err)
		g.Point.Set(*np)

		nl, err := geo.NewLineString([][]float64{{0, 0}, {4, 4}, {8, 8}})
		unit.Must(t, err)
		g.LineString.Set(*nl)

		nps, err := geo.NewPolygon([][]float64{{2, 2}, {6, 6}, {10, 2}, {2, 2}})
		unit.Must(t, err)
		g.Polygon.Set(*nps)

		err = db.Save(ctx)
		unit.Must(t, err)
	})

	// ST_Contains，点 - 多边形。
	// ST_Contains, Point - Polygon.
	t.Run("ST_Contains, Point - Polygon.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		g, err := db.Geos.Where(
			db.Geos.Point.Geo(
				db.Geos.Point.ST_Contains(
					db.Geos.Polygon.GeoColumn(),
					db.Geos.Point.ST_SetSRID(
						db.Geos.Point.ST_MakePoint(6, 5),
						new(geo.SDefault),
					),
				)),
		).Single(ctx)
		unit.Must(t, err)
		fmt.Println(g)
	})

	// ST_Contains，多边形 - 多边形
	// ST_Contains, Polygon - Polygon.
	t.Run("ST_Contains, Polygon - Polygon.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		l, err := geo.NewLineString([][]float64{{3, 3}, {5, 5}, {9, 3}, {3, 3}})
		unit.Must(t, err)
		g, err := db.Geos.Where(
			db.Geos.Polygon.Geo(
				db.Geos.Polygon.ST_Contains(
					db.Geos.Polygon.GeoColumn(),
					db.Geos.Polygon.ST_SetSRID(
						db.Geos.Polygon.ST_MakePolygon(
							db.Geos.Polygon.ST_GeomFromText(l.String()),
						),
						new(geo.SDefault),
					),
				),
			),
		).Single(ctx)
		unit.Must(t, err)
		fmt.Println(g)
	})

	// ST_Crosses，线 - 线。
	// ST_Crosses, Line - Line.
	t.Run("ST_Crosses, Line - Line.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		g, err := db.Geos.Where(
			db.Geos.LineString.Geo(
				db.Geos.LineString.ST_Crosses(
					db.Geos.LineString.GeoColumn(),
					db.Geos.LineString.ST_SetSRID(
						db.Geos.LineString.ST_MakeLine(
							db.Geos.LineString.ST_MakePoint(0, 3),
							db.Geos.LineString.ST_MakePoint(5, 3),
						),
						new(geo.SDefault),
					),
				)),
		).Single(ctx)
		unit.Must(t, err)
		fmt.Println(g)
	})

	// ST_Crosses，线 - 多边形。
	// ST_Crosses, Line - Polygon.
	t.Run("ST_Crosses, Line - Surface.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		g, err := db.Geos.Where(
			db.Geos.Polygon.Geo(
				db.Geos.Polygon.ST_Crosses(
					db.Geos.Polygon.GeoColumn(),
					db.Geos.Polygon.ST_MakeLine(
						db.Geos.LineString.ST_MakePoint(5, 5),
						db.Geos.LineString.ST_MakePoint(10, 5),
					),
				),
			),
		).Single(ctx)
		unit.Must(t, err)
		fmt.Println(g)
	})

	// ST_Disjoint，点 - 点
	// ST_Disjoint, Point - Point.
	t.Run("ST_Disjoint, Point - Point.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		g, err := db.Geos.Where(
			db.Geos.Point.Geo(
				db.Geos.Point.ST_Disjoint(
					db.Geos.Point.GeoColumn(),
					db.Geos.Point.ST_SetSRID(
						db.Geos.Point.ST_MakePoint(4, 3),
						new(geo.S4326),
					),
				)),
		).Single(ctx)
		unit.Must(t, err)
		fmt.Println(g)
	})

	// ST_Disjoint，点 - 线
	// ST_Disjoint, Point - Line.
	t.Run("ST_Disjoint, Point - Line.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		g, err := db.Geos.Where(
			db.Geos.Point.Geo(
				db.Geos.Point.ST_Disjoint(
					db.Geos.LineString.GeoColumn(),
					db.Geos.Point.ST_SetSRID(
						db.Geos.Point.ST_MakePoint(3, 4),
						new(geo.SDefault),
					),
				)),
		).Single(ctx)
		unit.Must(t, err)
		fmt.Println(g)
	})

	// ST_Disjoint，点 - 多边形
	// ST_Disjoint, Point - Polygon.
	t.Run("ST_Disjoint, Point - Polygon.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		g, err := db.Geos.Where(
			db.Geos.Point.Geo(
				db.Geos.Point.ST_Disjoint(
					db.Geos.Polygon.GeoColumn(),
					db.Geos.Point.ST_SetSRID(
						db.Geos.Point.ST_MakePoint(10, 10),
						new(geo.SDefault),
					),
				)),
		).Single(ctx)
		unit.Must(t, err)
		fmt.Println(g)
	})

	// ST_Disjoint，线 - 线
	// ST_Disjoint, Line - Line.
	t.Run("ST_Disjoint, Line - Line.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		g, err := db.Geos.Where(
			db.Geos.LineString.Geo(
				db.Geos.LineString.ST_Disjoint(
					db.Geos.LineString.GeoColumn(),
					db.Geos.LineString.ST_SetSRID(
						db.Geos.LineString.ST_MakeLine(
							db.Geos.LineString.ST_MakePoint(10, 10),
							db.Geos.LineString.ST_MakePoint(20, 20),
						),
						new(geo.SDefault),
					),
				)),
		).Single(ctx)
		unit.Must(t, err)
		fmt.Println(g)
	})

	// ST_Disjoint，线 - 多边形
	// ST_Disjoint, Line - Polygon.
	t.Run("ST_Disjoint, Line - Polygon.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		g, err := db.Geos.Where(
			db.Geos.LineString.Geo(
				db.Geos.LineString.ST_Disjoint(
					db.Geos.Polygon.GeoColumn(),
					db.Geos.LineString.ST_SetSRID(
						db.Geos.LineString.ST_MakeLine(
							db.Geos.LineString.ST_MakePoint(15, 15),
							db.Geos.LineString.ST_MakePoint(20, 20),
						),
						new(geo.SDefault),
					),
				)),
		).Single(ctx)
		unit.Must(t, err)
		fmt.Println(g)
	})

	// ST_Disjoint，多边形 - 多边形
	// ST_Disjoint, Polygon - Polygon.
	t.Run("ST_Disjoint, Polygon - Polygon.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		l, err := geo.NewLineString([][]float64{{10, 10}, {12, 12}, {14, 12}, {10, 10}})
		unit.Must(t, err)
		g, err := db.Geos.Where(
			db.Geos.Polygon.Geo(
				db.Geos.Polygon.ST_Disjoint(
					db.Geos.Polygon.GeoColumn(),
					db.Geos.Polygon.ST_SetSRID(
						db.Geos.Polygon.ST_MakePolygon(
							db.Geos.Polygon.ST_GeomFromText(l.String()),
						),
						new(geo.SDefault),
					),
				),
			),
		).Single(ctx)
		unit.Must(t, err)
		fmt.Println(g)
	})

	// ST_Equals，点 - 点
	// ST_Equals, Point - Point.
	t.Run("ST_Equals, Point - Point.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		g, err := db.Geos.Where(
			db.Geos.Point.Geo(
				db.Geos.Point.ST_Equals(
					db.Geos.Point.GeoColumn(),
					db.Geos.Point.ST_SetSRID(
						db.Geos.Point.ST_MakePoint(3, 4),
						new(geo.S4326),
					),
				)),
		).Single(ctx)
		unit.Must(t, err)
		fmt.Println(g)
	})

	// ST_Equals，线 - 线
	// ST_Equals, Line - Line.
	t.Run("ST_Equals, Line - Line.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		g, err := db.Geos.Where(
			db.Geos.LineString.Geo(
				db.Geos.LineString.ST_Equals(
					db.Geos.LineString.GeoColumn(),
					db.Geos.LineString.ST_SetSRID(
						db.Geos.LineString.ST_MakeLine(
							db.Geos.LineString.ST_MakePoint(0, 0),
							db.Geos.LineString.ST_MakePoint(8, 8),
						),
						new(geo.SDefault),
					),
				)),
		).Single(ctx)
		unit.Must(t, err)
		fmt.Println(g)
	})

	// ST_Equals，多边形 - 多边形
	// ST_Equals, Polygon - Polygon.
	t.Run("ST_Equals, Polygon - Polygon.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		l, err := geo.NewLineString([][]float64{{2, 2}, {6, 6}, {10, 2}, {2, 2}})
		unit.Must(t, err)
		g, err := db.Geos.Where(
			db.Geos.Polygon.Geo(
				db.Geos.Polygon.ST_Equals(
					db.Geos.Polygon.GeoColumn(),
					db.Geos.Polygon.ST_SetSRID(
						db.Geos.Polygon.ST_MakePolygon(
							db.Geos.Polygon.ST_GeomFromText(l.String()),
						),
						new(geo.SDefault),
					),
				),
			),
		).Single(ctx)
		unit.Must(t, err)
		fmt.Println(g)
	})

	// ST_Intersects，点 - 点
	// ST_Intersects, Point - Point.
	t.Run("ST_Intersects, Point - Point.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		g, err := db.Geos.Where(
			db.Geos.Point.Geo(
				db.Geos.Point.ST_Intersects(
					db.Geos.Point.GeoColumn(),
					db.Geos.Point.ST_SetSRID(
						db.Geos.Point.ST_MakePoint(3, 4),
						new(geo.S4326),
					),
				)),
		).Single(ctx)
		unit.Must(t, err)
		fmt.Println(g)
	})

	// ST_Intersects，点 - 线
	// ST_Intersects, Point - Line.
	t.Run("ST_Intersects, Point - Line.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		g, err := db.Geos.Where(
			db.Geos.Point.Geo(
				db.Geos.Point.ST_Intersects(
					db.Geos.LineString.GeoColumn(),
					db.Geos.Point.ST_SetSRID(
						db.Geos.Point.ST_MakePoint(4, 4),
						new(geo.SDefault),
					),
				)),
		).Single(ctx)
		unit.Must(t, err)
		fmt.Println(g)
	})

	// ST_Intersects，点 - 多边形
	// ST_Intersects, Point - Polygon.
	t.Run("ST_Intersects, Point - Polygon.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		g, err := db.Geos.Where(
			db.Geos.Point.Geo(
				db.Geos.Point.ST_Intersects(
					db.Geos.Polygon.GeoColumn(),
					db.Geos.Point.ST_SetSRID(
						db.Geos.Point.ST_MakePoint(4, 3),
						new(geo.SDefault),
					),
				)),
		).Single(ctx)
		unit.Must(t, err)
		fmt.Println(g)
	})

	// ST_Intersects，线 - 线
	// ST_Intersects, Line - Line.
	t.Run("ST_Intersects, Line - Line.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		g, err := db.Geos.Where(
			db.Geos.LineString.Geo(
				db.Geos.LineString.ST_Intersects(
					db.Geos.LineString.GeoColumn(),
					db.Geos.LineString.ST_SetSRID(
						db.Geos.LineString.ST_MakeLine(
							db.Geos.LineString.ST_MakePoint(0, 4),
							db.Geos.LineString.ST_MakePoint(4, 0),
						),
						new(geo.SDefault),
					),
				)),
		).Single(ctx)
		unit.Must(t, err)
		fmt.Println(g)
	})

	// ST_Intersects，线 - 多边形
	// ST_Intersects, Line - Polygon.
	t.Run("ST_Intersects, Line - Polygon.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		g, err := db.Geos.Where(
			db.Geos.LineString.Geo(
				db.Geos.LineString.ST_Intersects(
					db.Geos.Polygon.GeoColumn(),
					db.Geos.LineString.ST_SetSRID(
						db.Geos.LineString.ST_MakeLine(
							db.Geos.LineString.ST_MakePoint(0, 0),
							db.Geos.LineString.ST_MakePoint(8, 8),
						),
						new(geo.SDefault),
					),
				)),
		).Single(ctx)
		unit.Must(t, err)
		fmt.Println(g)
	})

	// ST_Intersects，多边形 - 多边形
	// ST_Intersects, Polygon - Polygon.
	t.Run("ST_Intersects, Polygon - Polygon.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		l, err := geo.NewLineString([][]float64{{2, 2}, {6, 6}, {10, 2}, {2, 2}})
		unit.Must(t, err)
		g, err := db.Geos.Where(
			db.Geos.Polygon.Geo(
				db.Geos.Polygon.ST_Intersects(
					db.Geos.Polygon.GeoColumn(),
					db.Geos.Polygon.ST_SetSRID(
						db.Geos.Polygon.ST_MakePolygon(
							db.Geos.Polygon.ST_GeomFromText(l.String()),
						),
						new(geo.SDefault),
					),
				),
			),
		).Single(ctx)
		unit.Must(t, err)
		fmt.Println(g)
	})

	// ST_Overlaps，线 - 线
	// ST_Overlaps, Line - Line.
	t.Run("ST_Overlaps, Line - Line.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		g, err := db.Geos.Where(
			db.Geos.LineString.Geo(
				db.Geos.LineString.ST_Overlaps(
					db.Geos.LineString.GeoColumn(),
					db.Geos.LineString.ST_SetSRID(
						db.Geos.LineString.ST_MakeLine(
							db.Geos.LineString.ST_MakePoint(5, 5),
							db.Geos.LineString.ST_MakePoint(10, 10),
						),
						new(geo.SDefault),
					),
				)),
		).Single(ctx)
		unit.Must(t, err)
		fmt.Println(g)
	})

	// ST_Overlaps，多边形 - 多边形
	// ST_Overlaps, Polygon - Polygon.
	t.Run("ST_Overlaps, Polygon - Polygon.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		l, err := geo.NewLineString([][]float64{{1, 1}, {5, 5}, {7, 1}, {1, 1}})
		unit.Must(t, err)
		g, err := db.Geos.Where(
			db.Geos.Polygon.Geo(
				db.Geos.Polygon.ST_Overlaps(
					db.Geos.Polygon.GeoColumn(),
					db.Geos.Polygon.ST_SetSRID(
						db.Geos.Polygon.ST_MakePolygon(
							db.Geos.Polygon.ST_GeomFromText(l.String()),
						),
						new(geo.SDefault),
					),
				),
			),
		).Single(ctx)
		unit.Must(t, err)
		fmt.Println(g)
	})

	// ST_Touches，点 - 线
	// ST_Touches, Point - Line.
	t.Run("ST_Touches, Point - Line.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		g, err := db.Geos.Where(
			db.Geos.Point.Geo(
				db.Geos.Point.ST_Touches(
					db.Geos.LineString.GeoColumn(),
					db.Geos.Point.ST_SetSRID(
						db.Geos.Point.ST_MakePoint(0, 0),
						new(geo.SDefault),
					),
				)),
		).Single(ctx)
		unit.Must(t, err)
		fmt.Println(g)
	})

	// ST_Touches，点 - 多边形
	// ST_Touches, Point - Polygon.
	t.Run("ST_Touches, Point - Polygon.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		g, err := db.Geos.Where(
			db.Geos.Point.Geo(
				db.Geos.Point.ST_Touches(
					db.Geos.Polygon.GeoColumn(),
					db.Geos.Point.ST_SetSRID(
						db.Geos.Point.ST_MakePoint(2, 2),
						new(geo.SDefault),
					),
				)),
		).Single(ctx)
		unit.Must(t, err)
		fmt.Println(g)
	})

	// ST_Touches，线 - 线
	// ST_Touches, Line - Line.
	t.Run("ST_Touches, Line - Line.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		g, err := db.Geos.Where(
			db.Geos.LineString.Geo(
				db.Geos.LineString.ST_Touches(
					db.Geos.LineString.GeoColumn(),
					db.Geos.LineString.ST_SetSRID(
						db.Geos.LineString.ST_MakeLine(
							db.Geos.LineString.ST_MakePoint(-2, -2),
							db.Geos.LineString.ST_MakePoint(0, 0),
						),
						new(geo.SDefault),
					),
				)),
		).Single(ctx)
		unit.Must(t, err)
		fmt.Println(g)
	})

	// ST_Touches，线 - 多边形
	// ST_Touches, Line - Polygon.
	t.Run("ST_Touches, Line - Polygon.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		g, err := db.Geos.Where(
			db.Geos.LineString.Geo(
				db.Geos.LineString.ST_Touches(
					db.Geos.Polygon.GeoColumn(),
					db.Geos.LineString.ST_SetSRID(
						db.Geos.LineString.ST_MakeLine(
							db.Geos.LineString.ST_MakePoint(2, 2),
							db.Geos.LineString.ST_MakePoint(7, 10),
						),
						new(geo.SDefault),
					),
				)),
		).Single(ctx)
		unit.Must(t, err)
		fmt.Println(g)
	})

	// ST_Touches，多边形 - 多边形
	// ST_Touches, Polygon - Polygon.
	t.Run("ST_Touches, Polygon - Polygon.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		l, err := geo.NewLineString([][]float64{{2, 2}, {6, 6}, {2, 10}, {2, 2}})
		unit.Must(t, err)
		g, err := db.Geos.Where(
			db.Geos.Polygon.Geo(
				db.Geos.Polygon.ST_Touches(
					db.Geos.Polygon.GeoColumn(),
					db.Geos.Polygon.ST_SetSRID(
						db.Geos.Polygon.ST_MakePolygon(
							db.Geos.Polygon.ST_GeomFromText(l.String()),
						),
						new(geo.SDefault),
					),
				),
			),
		).Single(ctx)
		unit.Must(t, err)
		fmt.Println(g)
	})

	// ST_Within，点 - 多边形
	// ST_Within, Point - Polygon.
	t.Run("ST_Within, Point - Polygon.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		g, err := db.Geos.Where(
			db.Geos.Point.Geo(
				db.Geos.Point.ST_Within(
					db.Geos.Polygon.GeoColumn(),
					db.Geos.Point.ST_SetSRID(
						db.Geos.Point.ST_MakePoint(3, 3),
						new(geo.SDefault),
					),
				)),
		).Single(ctx)
		unit.Must(t, err)
		fmt.Println(g)
	})

	// ST_Within，线 - 多边形
	// ST_Within, Line - Polygon.
	t.Run("ST_Within, Line - Polygon.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		g, err := db.Geos.Where(
			db.Geos.LineString.Geo(
				db.Geos.LineString.ST_Within(
					db.Geos.Polygon.GeoColumn(),
					db.Geos.LineString.ST_SetSRID(
						db.Geos.LineString.ST_MakeLine(
							db.Geos.LineString.ST_MakePoint(1, 1),
							db.Geos.LineString.ST_MakePoint(5, 5),
						),
						new(geo.SDefault),
					),
				)),
		).Single(ctx)
		unit.Must(t, err)
		fmt.Println(g)
	})

	// ST_Within，多边形 - 多边形
	// ST_Within, Polygon - Polygon.
	t.Run("ST_Within, Polygon - Polygon.", func(t *testing.T) {
		db := initDb()
		defer db.Close()
		ctx := context.Background()
		l, err := geo.NewLineString([][]float64{{2, 2}, {6, 6}, {10, 2}, {2, 2}})
		unit.Must(t, err)
		g, err := db.Geos.Where(
			db.Geos.Polygon.Geo(
				db.Geos.Polygon.ST_Within(
					db.Geos.Polygon.GeoColumn(),
					db.Geos.Polygon.ST_SetSRID(
						db.Geos.Polygon.ST_MakePolygon(
							db.Geos.Polygon.ST_GeomFromText(l.String()),
						),
						new(geo.SDefault),
					),
				),
			),
		).Single(ctx)
		unit.Must(t, err)
		fmt.Println(g)
	})
}
