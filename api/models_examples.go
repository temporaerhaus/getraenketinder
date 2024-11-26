package api_test

import "github.com/temporearhaus/getraenketinder/api"

var MinimalGetraenk = api.Getraenk{
	Name: "Blubberwasser Deluxe Plus Pro",
}

var MinimalGetraenkWithUUID = api.Getraenk{
	Name: "Blubberwasser Deluxe Plus Pro",
	Uuid: UUID.Parse(),
}
