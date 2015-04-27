package main

import (
	"io/ioutil"
	"os"
	"runtime"
	"testing"

	"github.com/ikeikeikeike/godic/models"
)

func createTmp(t *testing.T) string {
	path, err := ioutil.TempDir("", "godic")
	checkFatal(t, err)
	return path
}

func TestInit(t *testing.T) {
	dict := models.NewDict()

	dict.Init(createTmp(t))
	os.RemoveAll(dict.Repo.Workdir())
}

func TestCreateFile(t *testing.T) {
	dict := models.NewDict()

	dict.Init(createTmp(t))
	defer os.RemoveAll(dict.Repo.Workdir())

	_, err := dict.SaveFile("fileone", "# aaaa\n- a\n- b\n", "first message")
	if err != nil {
		t.Fatalf("Add file error: %s", err)
	}

	stats, err := dict.ModifiedStats()
	if len(stats) > 0 {
		checkFatal(t, err)
	}
	stats, err = dict.UntrackedStats()
	if len(stats) > 0 {
		checkFatal(t, err)
	}
}

func TestUpdateFile(t *testing.T) {
	dict := models.NewDict()
	dict.Init(createTmp(t))

	defer os.RemoveAll(dict.Repo.Workdir())

	_, err := dict.SaveFile("fileone", "# aaaa\n- a\n- b\n", "first message")
	checkFatal(t, err)
	_, err = dict.SaveFile("fileone", "# bbbb\n- a\n- b\n", "second message")
	checkFatal(t, err)

	stats, err := dict.ModifiedStats()
	if len(stats) > 0 {
		checkFatal(t, err)
	}
	stats, err = dict.UntrackedStats()
	if len(stats) > 0 {
		checkFatal(t, err)
	}
}

func TestModifiedStats(t *testing.T) {
	dict := models.NewDict()
	dict.Init(createTmp(t))

	defer os.RemoveAll(dict.Repo.Workdir())

	stats, err := dict.ModifiedStats()
	checkFatal(t, err)

	for _, stat := range stats {
		_ = stat
	}
}

func TestUntrackedStats(t *testing.T) {
	dict := models.NewDict()
	dict.Init(createTmp(t))

	defer os.RemoveAll(dict.Repo.Workdir())

	stats, err := dict.UntrackedStats()
	checkFatal(t, err)

	for _, stat := range stats {
		_ = stat
	}
}

func TestDumpTree(t *testing.T) {
	dict := models.NewDict()
	dict.Init(createTmp(t))

	defer os.RemoveAll(dict.Repo.Workdir())

	_, err := dict.SaveFile("fileone", "# aaaa\n- a\n- b\n", "first message")
	checkFatal(t, err)
	_, err = dict.SaveFile("filetwo", "# bbbb\n- a\n- b\n", "second message")
	checkFatal(t, err)
	_, err = dict.SaveFile("fileone", "# cccc\n- a\n- b\n", "first message2")
	checkFatal(t, err)
	_, err = dict.SaveFile("filetwo", "# dddd\n- a\n- b\n", "second message2")
	checkFatal(t, err)

	dict.DumpRepo()
}

func TestGetTree(t *testing.T) {
	dict := models.NewDict()
	dict.Init(createTmp(t))

	defer os.RemoveAll(dict.Repo.Workdir())

	_, err := dict.SaveFile("fileone", "# aaaa\n- a\n- b\n", "first message")
	checkFatal(t, err)
	_, err = dict.SaveFile("filetwo", "# bbbb\n- a\n- b\n", "second message")
	checkFatal(t, err)
	_, err = dict.SaveFile("fileone", "# cccc\n- a\n- b\n", "first message2")
	checkFatal(t, err)
	_, err = dict.SaveFile("filetwo", "# dddd\n- a\n- b\n", "second message2")
	checkFatal(t, err)
}

func checkFatal(t *testing.T, err error) {
	if err == nil {
		return
	}

	// The failure happens at wherever we were called, not here
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		t.Fatal()
	}

	t.Fatalf("Fail at %v:%v; %v", file, line, err)
}
