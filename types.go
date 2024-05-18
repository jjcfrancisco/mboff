package main

import "database/sql"

type result struct {
	zoomLevel int
	tileId    string
	tileData  []byte
}

type dbConn struct {
	db *sql.DB
}

type kv struct {
	key   string
	value string
}

