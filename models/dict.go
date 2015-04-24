package models

import (
	"fmt"
	"io/ioutil"
	"time"

	git "github.com/libgit2/git2go"
)

type Dict struct {
	Path string
	Repo *git.Repository

	Ref       string
	Author    *git.Signature
	Committer *git.Signature
}

func NewDict() *Dict {
	return &Dict{
		Ref:       "HEAD",
		Author:    &git.Signature{Name: "Author", Email: "author@example.com", When: time.Now()},
		Committer: &git.Signature{Name: "Committer", Email: "committer@example.com", When: time.Now()},
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

func (m *Dict) Save(name, content, message string, create bool) {
	if create {
		m.Create(name, content, message)
	}
	m.Update(name, content, message)
}

func (m *Dict) Create(filename, content, message string) (*git.Oid, error) {
	idx, err := m.Repo.Index()
	if err != nil {
		return nil, err
	}
	err = ioutil.WriteFile(m.Path+"/"+filename, []byte(content), 0644)
	if err != nil {
		return nil, err
	}
	if err = idx.AddByPath(filename); err != nil {
		return nil, err
	}

	treeID, err := idx.WriteTree()
	if err != nil {
		return nil, err
	}
	tree, err := m.Repo.LookupTree(treeID)
	if err != nil {
		return nil, err
	}

	if message == "" {
		message = fmt.Sprintf("Create: %s", filename)
	}
	return m.Repo.CreateCommit(m.Ref, m.Author, m.Committer, message, tree)
}

func (m *Dict) Update(filename, content, message string) (*git.Oid, error) {
	branch, err := m.Repo.Head()
	if err != nil {
		return nil, err
	}
	tip, err := m.Repo.LookupCommit(branch.Target())
	if err != nil {
		return nil, err
	}

	idx, err := m.Repo.Index()
	if err != nil {
		return nil, err
	}
	err = ioutil.WriteFile(m.Path+"/"+filename, []byte(content), 0644)
	if err != nil {
		return nil, err
	}
	if err = idx.AddByPath(filename); err != nil {
		return nil, err
	}

	treeID, err := idx.WriteTree()
	if err != nil {
		return nil, err
	}
	tree, err := m.Repo.LookupTree(treeID)
	if err != nil {
		return nil, err
	}

	if message == "" {
		message = fmt.Sprintf("Update: %s", filename)
	}
	return m.Repo.CreateCommit(m.Ref, m.Author, m.Committer, message, tree, tip)
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

func (m *Dict) ModifiedStats() ([]git.StatusEntry, error) {
	return m.Stats(&git.StatusOptions{})
}

func (m *Dict) UntrackedStats() ([]git.StatusEntry, error) {
	opts := &git.StatusOptions{Flags: git.StatusOptIncludeUntracked}
	return m.Stats(opts)
}
