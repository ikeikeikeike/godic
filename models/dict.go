package models

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
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

func (m *Dict) GetFile(filename string) (err error) {
	return
}

func (m *Dict) SaveFile(filename, content, message string) (*git.Oid, error) {
	var tip *git.Commit

	branch, err := m.Repo.Head()
	if err == nil {
		tip, _ = m.Repo.LookupCommit(branch.Target())
	}

	err = ioutil.WriteFile(m.Path+"/"+filename, []byte(content), 0644)
	if err != nil {
		return nil, err
	}

	idx, err := m.Repo.Index()
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
		if tip != nil {
			message = fmt.Sprintf("Update: %s", filename)
		}
	}

	if tip == nil {
		return m.Repo.CreateCommit(m.Ref, m.Author, m.Committer, message, tree)
	} else {
		return m.Repo.CreateCommit(m.Ref, m.Author, m.Committer, message, tree, tip)
	}
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

func (m *Dict) DumpRepo() {

	odb, err := m.Repo.Odb()
	if err != nil {
		log.Fatal(err)
	}

	err = odb.ForEach(func(oid *git.Oid) error {
		obj, err := m.Repo.Lookup(oid)
		if err != nil {
			return err
		}

		switch obj := obj.(type) {
		default:
		case *git.Blob:
			break
			fmt.Printf("=================Blob=================\n")
			fmt.Printf("obj:  %s\n", obj)
			fmt.Printf("Type: %s\n", obj.Type())
			fmt.Printf("Id:   %s\n", obj.Id())
			fmt.Printf("Size: %s\n", obj.Size())
		case *git.Commit:
			fmt.Printf("=================Commit=================\n")
			fmt.Printf("obj:  %s\n", obj)
			fmt.Printf("Type: %s\n", obj.Type())
			fmt.Printf("Id:   %s\n", obj.Id())
			author := obj.Author()
			fmt.Printf("    Author:\n        Name:  %s\n        Email: %s\n        Date:  %s\n", author.Name, author.Email, author.When)
			committer := obj.Committer()
			fmt.Printf("    Committer:\n        Name:  %s\n        Email: %s\n        Date:  %s\n", committer.Name, committer.Email, committer.When)
			fmt.Printf("    ParentCount: %d\n", int(obj.ParentCount()))
			fmt.Printf("    TreeId:      %s\n", obj.TreeId())
			fmt.Printf("    Message:\n\n        %s\n\n", strings.Replace(obj.Message(), "\n", "\n        ", -1))
			//fmt.Printf("obj.Parent: %s\n", obj.Parent())
			//fmt.Printf("obj.ParentId: %s\n", obj.ParentId())
			//fmt.Printf("obj.Tree: %s\n", obj.Tree())
		case *git.Tree:
			break
			fmt.Printf("=================Tree=================\n")
			fmt.Printf("obj:  %s\n", obj)
			fmt.Printf("Type: %s\n", obj.Type())
			fmt.Printf("Id:   %s\n", obj.Id())
			fmt.Printf("    EntryCount: %d\n", obj.EntryCount())
		}
		return nil
	})
}
