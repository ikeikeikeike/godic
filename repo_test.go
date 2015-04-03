package main

import (
	"testing"

	"bitbucket.org/ikeikeikeike/godic/models"
)

func TestInit(t *testing.T) {
	dict := models.NewDict()
	dict.Init("./repo")
}

func TestCreateFile(t *testing.T) {
	dict := models.NewDict()
	dict.Init("./repo")

	_, err := dict.Create("fileone", "# aaaa\n- a\n- b\n")
	if err != nil {
		t.Fatalf("Add file error: %s", err)
	}
}

func TestUpdateFile(t *testing.T) {
	dict := models.NewDict()
	dict.Init("./repo")

	_, err := dict.Update("fileone", "# bbbb\n- a\n- b\n")
	if err != nil {
		t.Fatalf("Update file error: %s", err)
	}
}

func TestModifiedStats(t *testing.T) {
	dict := models.NewDict()
	dict.Init("./repo")

	stats, err := dict.ModifiedStats()
	if err != nil {
		t.Fatalf("Fetch status error: %s", err)
	}
	for _, stat := range stats {
		_ = stat
	}
}

func TestUntrackedStats(t *testing.T) {
	dict := models.NewDict()
	dict.Init("./repo")

	stats, err := dict.UntrackedStats()
	if err != nil {
		t.Fatalf("Fetch status error: %s", err)
	}
	for _, stat := range stats {
		_ = stat
	}
}
