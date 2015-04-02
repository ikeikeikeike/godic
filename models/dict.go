package models

import git "github.com/libgit2/git2go"

type Dict struct {
	Path string
	Repo *git.Repository

	Ref            string
	CommitterName  string
	CommitterEmail string
}

func NewDict() *Dict {
	return &Dict{
		Ref:            "master",
		CommitterName:  "Unk",
		CommitterEmail: "unk@example.com",
	}
}

func (m *Dict) Init(path string) (err error) {
	if m.Repo, err = git.OpenRepository(path); err != nil {
		if m.Repo, err = git.InitRepository(path, false); err != nil {
			return
		}
	}
	m.Path = path
	return
}

func (m *Dict) Stats(opts *git.StatusOptions) (entries []git.StatusEntry, err error) {
	var stats *git.StatusList
	if stats, err = m.Repo.StatusList(opts); err != nil {
		return
	}
	var statscnt int
	if statscnt, err = stats.EntryCount(); err != nil {
		return
	}

	var entry git.StatusEntry
	for i := 0; i < statscnt; i++ {
		entry, err = stats.ByIndex(i)
		if err != nil {
			return
		}
		entries = append(entries, entry)
	}
	return
}

func (m *Dict) StatsModified() ([]git.StatusEntry, error) {
	return m.Stats(&git.StatusOptions{})
}

func (m *Dict) StatsUntracked() ([]git.StatusEntry, error) {
	opts := &git.StatusOptions{Flags: git.StatusOptIncludeUntracked}
	return m.Stats(opts)
}
