package main

import (
	"io/ioutil"
	"os"
	"runtime"
	"testing"

	"github.com/ikeikeikeike/godic/modules/git"
)

func TestDumpTree(t *testing.T) {
	repo := git.NewRepo()
	repo.Init(createTmp(t))

	defer os.RemoveAll(repo.Repo.Workdir())

	_, err := repo.SaveFile("fileone", "# aaaa\n- a\n- b\n", "first message")
	checkFatal(t, err)
	_, err = repo.SaveFile("filetwo", "# bbbb\n- a\n- b\n", "second message")
	checkFatal(t, err)
	_, err = repo.SaveFile("fileone", "# cccc\n- a\n- b\n", "first message2")
	checkFatal(t, err)
	_, err = repo.SaveFile("filetwo", "# dddd\n- a\n- b\n", "second message2")
	checkFatal(t, err)

	repo.DumpRepo()
}

func createTmp(t *testing.T) string {
	path, err := ioutil.TempDir("", "godic")
	checkFatal(t, err)
	return path
}

func TestInit(t *testing.T) {
	repo := git.NewRepo()

	repo.Init(createTmp(t))
	os.RemoveAll(repo.Repo.Workdir())
}

func TestCreateFile(t *testing.T) {
	repo := git.NewRepo()

	repo.Init(createTmp(t))
	defer os.RemoveAll(repo.Repo.Workdir())

	_, err := repo.SaveFile("fileone", "# aaaa\n- a\n- b\n", "first message")
	if err != nil {
		t.Fatalf("Add file error: %s", err)
	}

	stats, err := repo.ModifiedStats()
	if len(stats) > 0 {
		checkFatal(t, err)
	}
	stats, err = repo.UntrackedStats()
	if len(stats) > 0 {
		checkFatal(t, err)
	}
}

func TestUpdateFile(t *testing.T) {
	repo := git.NewRepo()
	repo.Init(createTmp(t))

	defer os.RemoveAll(repo.Repo.Workdir())

	_, err := repo.SaveFile("fileone", "# aaaa\n- a\n- b\n", "first message")
	checkFatal(t, err)
	_, err = repo.SaveFile("fileone", "# bbbb\n- a\n- b\n", "second message")
	checkFatal(t, err)

	stats, err := repo.ModifiedStats()
	if len(stats) > 0 {
		checkFatal(t, err)
	}
	stats, err = repo.UntrackedStats()
	if len(stats) > 0 {
		checkFatal(t, err)
	}
}

func TestGetFileBlob(t *testing.T) {
	repo := git.NewRepo()
	repo.Init(createTmp(t))

	defer os.RemoveAll(repo.Repo.Workdir())

	_, err := repo.SaveFile("fileone", "# aaaa\n- a\n- b\n", "first message")
	checkFatal(t, err)
	_, err = repo.SaveFile("fileone", "# bbbb\n- a\n- b\n", "second message")
	checkFatal(t, err)

	blob, err := repo.GetFileBlob("fileone")
	checkFatal(t, err)

	if string(blob.Contents()) != "# bbbb\n- a\n- b\n" {
		t.Fatalf("not matched string: %s", string(blob.Contents()))
	}
}

func TestGetCommitInfo(t *testing.T) {
	repo := git.NewRepo()
	repo.Init(createTmp(t))

	defer os.RemoveAll(repo.Repo.Workdir())

	_, err := repo.SaveFile("fileone", "# aaaa\n- a\n- b\n", "first message")
	checkFatal(t, err)
	_, err = repo.SaveFile("fileone", "# bbbb\n- a\n- b\n", "second message")
	checkFatal(t, err)

	commit, err := repo.GetCommit("fileone")
	checkFatal(t, err)

	if commit.Message() != "second message" {
		t.Fatalf("not matched string: %s", commit.Message())
	}
	if commit.Summary() != "second message" {
		t.Fatalf("not matched string: %s", commit.Summary())
	}
	if commit.Author().Name != "Author" {
		t.Fatalf("not matched string: %s", commit.Author().Name)
	}
	if commit.Committer().Name != "Committer" {
		t.Fatalf("not matched string: %s", commit.Committer().Name)
	}
}

func TestModifiedStats(t *testing.T) {
	repo := git.NewRepo()
	repo.Init(createTmp(t))

	defer os.RemoveAll(repo.Repo.Workdir())

	stats, err := repo.ModifiedStats()
	checkFatal(t, err)

	for _, stat := range stats {
		_ = stat
	}
}

func TestUntrackedStats(t *testing.T) {
	repo := git.NewRepo()
	repo.Init(createTmp(t))

	defer os.RemoveAll(repo.Repo.Workdir())

	stats, err := repo.UntrackedStats()
	checkFatal(t, err)

	for _, stat := range stats {
		_ = stat
	}
}

func TestGetTree(t *testing.T) {
	repo := git.NewRepo()
	repo.Init(createTmp(t))

	defer os.RemoveAll(repo.Repo.Workdir())

	_, err := repo.SaveFile("fileone", "# aaaa\n- a\n- b\n", "first message")
	checkFatal(t, err)
	_, err = repo.SaveFile("filetwo", "# bbbb\n- a\n- b\n", "second message")
	checkFatal(t, err)
	_, err = repo.SaveFile("fileone", "# cccc\n- a\n- b\n", "first message2")
	checkFatal(t, err)
	_, err = repo.SaveFile("filetwo", "# dddd\n- a\n- b\n", "second message2")
	checkFatal(t, err)

	names, err := repo.FolderFileNames(repo.Repo.Workdir())
	checkFatal(t, err)
	if len(names) != 2 {
		t.Fatalf("fail names length: %v", names)
	}
	for _, name := range names {
		if name == "fileone" {
			continue
		} else if name == "filetwo" {
			continue
		}
		t.Fatalf("Fail names at: %v", names)
	}
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
