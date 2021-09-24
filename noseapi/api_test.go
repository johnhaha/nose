package noseapi_test

import (
	"testing"

	"github.com/johnhaha/nose/noseapi"
)

func TestAddParam(t *testing.T) {
	var x noseapi.ParamApi = "122"
	y := x.AddParam("xx").AddParam("xx")
	if y != "122/xx/xx" {
		t.Fatal(y)
	}
}
