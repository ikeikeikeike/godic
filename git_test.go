package main

import (
	"io/ioutil"
	"os"
	"runtime"
	"strings"
	"testing"

	"github.com/ikeikeikeike/godic/modules/git"
	git2go "github.com/libgit2/git2go"
)

func TestDumpTree(t *testing.T) {
	repo := git.NewRepo()
	repo.Init(createTmp(t))

	defer os.RemoveAll(repo.Repo.Workdir())

	saveSomeFiles(t, repo)

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

func TestGetFileHistory(t *testing.T) {
	repo := git.NewRepo()
	repo.Init(createTmp(t))

	defer os.RemoveAll(repo.Repo.Workdir())

	saveSomeFiles(t, repo)

	l, err := repo.GetFileHistory("fileone", 1)
	checkFatal(t, err)

	for e := l.Front(); e != nil; e = e.Next() {
		c := e.Value.(*git2go.Commit)
		if !strings.HasSuffix(c.Message(), "one") {
			t.Fatalf("not matched suffix string: %s", c.Message())
		}
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

func TestFolderFileNames(t *testing.T) {
	repo := git.NewRepo()
	repo.Init(createTmp(t))

	defer os.RemoveAll(repo.Repo.Workdir())

	saveSomeFiles(t, repo)

	names, err := repo.FolderFileNames()
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

func TestGetDiffCommit(t *testing.T) {
	repo := git.NewRepo()
	repo.Init(createTmp(t))

	defer os.RemoveAll(repo.Repo.Workdir())

	saveSomeFiles(t, repo)

	oid1, err := repo.SaveFile("fileone", "# aaa\n- a\n- a\n", "first one")
	checkFatal(t, err)
	oid2, err := repo.SaveFile("fileone", "# bbbb\n- a\n- b\n", "second one")
	checkFatal(t, err)
	oid3, err := repo.SaveFile("fileone", "# cccc\n- a\n- c\n", "third one")
	checkFatal(t, err)

	_, _ = oid2, oid3

	diff, err := repo.GetDiffCommit(oid1.String(), 0)
	checkFatal(t, err)

	if diff.TotalAddition != 3 {
		t.Fatalf("Fail at TotalAddition: %d", diff.TotalAddition)
	}
	if diff.TotalDeletion != 3 {
		t.Fatalf("Fail at TotalDeletion: %d", diff.TotalDeletion)
	}
	for _, f := range diff.Files {
		if f.Name != "fileone" {
			t.Fatalf("Fail at diff.Name: %s", f.Name)
		}
		if f.Name != "fileone" {
			t.Fatalf("Fail :%s", f.Name)
		}
		if f.Index != 1 {
			t.Fatalf("Fail :%s", f.Index)
		}
		if f.Addition != 3 {
			t.Fatalf("Fail :%s", f.Addition)
		}
		if f.Deletion != 3 {
			t.Fatalf("Fail :%s", f.Deletion)
		}
		if f.Type != 2 {
			t.Fatalf("Fail :%s", f.Type)
		}
		if f.IsCreated != false {
			t.Fatalf("Fail :%s", f.IsCreated)
		}
		if f.IsDeleted != false {
			t.Fatalf("Fail :%s", f.IsDeleted)
		}
		if f.IsBin != false {
			t.Fatalf("Fail :%s", f.IsBin)
		}
		for _, s := range f.Sections {
			if len(s.Lines) != 7 {
				t.Fatalf("Fail section line length: %d", len(s.Lines))
			}

		}
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

func saveSomeFiles(t *testing.T, repo *git.Repo) {
	_, err := repo.SaveFile("fileone", "# aaaa\n- a\n- b\n", "first one")
	checkFatal(t, err)
	_, err = repo.SaveFile("filetwo", "# bbbb\n- a\n- b\n", "first two")
	checkFatal(t, err)
	_, err = repo.SaveFile("fileone", "# cccc\n- a\n- b\n", "second one")
	checkFatal(t, err)
	_, err = repo.SaveFile("filetwo", "# dddd\n- a\n- b\n", "second two")
	checkFatal(t, err)
	_, err = repo.SaveFile("fileone", "# AAAAA\n- K\n- b\n", "third one")
	checkFatal(t, err)
	_, err = repo.SaveFile("filetwo", "# ZZZZZ\n- a\n- b\n", "third two")
	checkFatal(t, err)
	_, err = repo.SaveFile("fileone", "# BBBBB\n- A\n- b\n", "firth one")
	checkFatal(t, err)
	_, err = repo.SaveFile("filetwo", "# XXXXX\n- B\n- b\n", "firth two")
	checkFatal(t, err)
}
