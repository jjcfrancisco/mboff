package main

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/paulmach/orb/encoding/mvt"
	"github.com/paulmach/orb/geojson"
	"github.com/stretchr/testify/assert"
)

func TestCreateConn(t *testing.T) {
	_, err := createConn("tests/example.mbtiles")
	assert.NoError(t, err)
}

func TestFind(t *testing.T) {
	conn, err := createConn("tests/example.mbtiles")	
	assert.NoError(t, err)
	defer conn.db.Close()
	// results, err := conn.find(&kv{key: "ROUTE_NO", value: "137"}, nil)
	// assert.NoError(t, err)
	// assert.Equal(t, len(results), 1)	
} 

func TestUpdate(t *testing.T) {

	conn, err := createConn("tests/example.mbtiles")
	assert.NoError(t, err)
	defer conn.db.Close()

	// Check
	rows, err := conn.db.Query("SELECT zoom_level, tile_data, tile_id FROM images;")
	if err != nil {
		panic(err)
	}
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
		for i := range layers {
			layer := layers[i]
			for _, f := range layer.Features {
				newFeature := geojson.NewFeature(f.Geometry)
				newFeature.Properties = f.Properties
				id := int(newFeature.Properties["id"].(float64))
				newFeature.Properties["id"] = strconv.Itoa(id)
				if newFeature.Properties["id"] == "506484968" {
					fmt.Println("Present!", zoomLevel, tileId)
				}
			}
		}
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	assert.NoError(t, err)

}
