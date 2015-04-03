package main

import (
	"testing"

	"github.com/k0kubun/pp"

	"bitbucket.org/ikeikeikeike/godic/models"
)

func TestInit(t *testing.T) {
	dict := models.NewDict()
	dict.Init(".")
}

func TestStatus(t *testing.T) {
	dict := models.NewDict()
	dict.Init(".")

	stats, err := dict.Stats()
	if err != nil {
		t.Fatalf("Fetch status error: %s", err)
	}
	for _, stat := range stats {
		pp.Println(stat)
	}
}
