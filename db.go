package main

import (
	"database/sql"
	"strconv"

	_ "github.com/mattn/go-sqlite3"
	"github.com/paulmach/orb/encoding/mvt"
	"github.com/paulmach/orb/geojson"
)

func createConn(fp string) (*dbConn, error) {
	db, err := sql.Open("sqlite3", fp)
	if err != nil {
		return nil, err
	}
	return &dbConn{db: db}, nil
}

func (conn *dbConn) find(keyValue *kv, userZoomLevel *int) ([]result, error) {

	rows, err := conn.db.Query("SELECT zoom_level, tile_data, tile_id FROM images;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []result

	for rows.Next() {
		var zoomLevel int
		var tileData []byte
		var tileId string
		err = rows.Scan(&zoomLevel, &tileData, &tileId)
		if err != nil {
			panic(err)
		}
		l, err := mvt.UnmarshalGzipped(tileData)
		if err != nil {
			panic(err)
		}
		layers := l.ToFeatureCollections()
		newLayers := make(map[string]*geojson.FeatureCollection, len(layers))
		for i := range layers {
			layer := layers[i]
			newLayer := geojson.NewFeatureCollection()
			for _, f := range layer.Features {
				newFeature := geojson.NewFeature(f.Geometry)
				newFeature.Properties = f.Properties
				// Keep copy of original id
				idOg := newFeature.Properties["id"]
				newFeature.Properties["id"] = strconv.FormatInt(int64(newFeature.Properties["id"].(float64)), 10)
				if userZoomLevel != nil {
					// Zoom level provided
					if newFeature.Properties[keyValue.key] != keyValue.value && zoomLevel == *userZoomLevel {
						newLayer.Append(newFeature)
					}
				} else {
					// No zoom level provided
					if newFeature.Properties[keyValue.key] != keyValue.value {
						newLayer.Append(newFeature)
					}
				}
				// Id back to original
				newFeature.Properties["id"] = idOg
			}
			newLayers[i] = newLayer
		}
		newTileData, err := mvt.MarshalGzipped(mvt.NewLayers(newLayers))
		if err != nil {
			return nil, err
		}
		results = append(results, result{tileId: tileId, tileData: newTileData})
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return results, nil

}

func (conn *dbConn) update(results []result, userZoomLevel *int) error {

	if userZoomLevel != nil {
		// Zoom level provided
		stmt, err := conn.db.Prepare("UPDATE images SET tile_data = ? WHERE tile_id = ? AND zoom_level = ?;")
		if err != nil {
			return err
		}
		defer stmt.Close()

		for _, result := range results {
			_, err = stmt.Exec(result.tileData, result.tileId, userZoomLevel)
			if err != nil {
				return err
			}
		}
	} else {
		// No zoom level provided
		stmt, err := conn.db.Prepare("UPDATE images SET tile_data = ? WHERE tile_id = ?;")
		if err != nil {
			return err
		}
		defer stmt.Close()

		for _, result := range results {
			_, err = stmt.Exec(result.tileData, result.tileId)
			if err != nil {
				return err
			}
		}
	}

	return nil

}
