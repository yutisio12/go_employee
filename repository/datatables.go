package repository

import (
	"net/url"
	"strconv"
)

type DataTablesQuery struct {
	Draw int
	Start int
	Length int
	Search string
	Sort string
}

func ParseDataTables(values url.Values, columns map[int]string) DataTablesQuery {
	getInt := func(key string, defaultValue int) int {
		v := values.Get(key)
		i, err := strconv.Atoi(v)
		if err != nil { return defaultValue }
		return i
	}
	q := DataTablesQuery{
		Draw: getInt("draw", 0),
		Start: getInt("start", 0),
		Length: getInt("length", 10),
		SearchValue: values.Get("search[value]"),
		OrderDir: values.Get("order[0][dir]"),
	}
	colIdx := getInt("order[0][column]", 0)
	if name, ok := columns[colIdx]; ok {
		q.OrderColumnName = name
	} else {
		q.OrderColumnName = "id"
	}
	if q.OrderDir != "desc" { q.OrderDir = "asc" }
	return q
}