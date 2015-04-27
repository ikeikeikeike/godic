package git

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"

	git "github.com/libgit2/git2go"
)

type Repo struct {
	Path string

	Repo *git.Repository
	Head *Head

	Author    *git.Signature
	Committer *git.Signature
}

func NewRepo() *Repo {
	return &Repo{
		Author:    &git.Signature{Name: "Author", Email: "author@example.com", When: time.Now()},
		Committer: &git.Signature{Name: "Committer", Email: "committer@example.com", When: time.Now()},
	}
}

func (r *Repo) Init(path string) (err error) {
	if r.Repo, err = git.OpenRepository(path); err != nil {
		if r.Repo, err = git.InitRepository(path, false); err != nil {
			return
		}
	}
	r.Path = path
	r.Head = &Head{r.Repo}
	return
}

func (r *Repo) FolderInfo(path string) (entries []*git.TreeEntry, err error) {
	tree, err, noHead := r.Head.CommitTree()
	if err != nil {
		if noHead {
			err = nil
		}
		return
	}

	var i uint64
	for i = 0; i < tree.EntryCount(); i++ {
		entries = append(entries, tree.EntryByIndex(i))
	}
	return
}

func (r *Repo) FolderFileNames(path string) (names []string, err error) {
	info, err := r.FolderInfo(path)
	if err != nil {
		return
	}

	for _, e := range info {
		if e.Type != git.ObjectTree {
			names = append(names, e.Name)
		}
	}
	return
}

func (r *Repo) GetFileBlob(filename string) (*git.Blob, error) {
	tree, err, _ := r.Head.CommitTree()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	entry, err := tree.EntryByPath(filename)
	if err != nil {
		return nil, err
	}
	return r.Repo.LookupBlob(entry.Id)
}

func (r *Repo) GetCommit(filename string) (*git.Commit, error) {
	commit, err, _ := r.Head.Commit()
	if err != nil {
		return nil, err
	}
	return commit, nil
}

func (r *Repo) SaveFile(filename, content, message string) (*git.Oid, error) {
	var tip *git.Commit

	branch, err := r.Repo.Head()
	if err == nil {
		tip, _ = r.Repo.LookupCommit(branch.Target())
	}

	err = ioutil.WriteFile(r.Path+"/"+filename, []byte(content), 0644)
	if err != nil {
		log.Error(err)
		return nil, err
	}

	idx, err := r.Repo.Index()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if err = idx.AddByPath(filename); err != nil {
		return nil, err
	}

	treeID, err := idx.WriteTree()
	if err != nil {
		return nil, err
	}
	tree, err := r.Repo.LookupTree(treeID)
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
		return r.Repo.CreateCommit("HEAD", r.Author, r.Committer, message, tree)
	} else {
		return r.Repo.CreateCommit("HEAD", r.Author, r.Committer, message, tree, tip)
	}
}

func (r *Repo) Stats(opts *git.StatusOptions) (entries []git.StatusEntry, err error) {
	var stats *git.StatusList
	if stats, err = r.Repo.StatusList(opts); err != nil {
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

func (r *Repo) ModifiedStats() ([]git.StatusEntry, error) {
	return r.Stats(&git.StatusOptions{})
}

func (r *Repo) UntrackedStats() ([]git.StatusEntry, error) {
	opts := &git.StatusOptions{Flags: git.StatusOptIncludeUntracked}
	return r.Stats(opts)
}

func (r *Repo) DumpRepo() error {

	odb, err := r.Repo.Odb()
	if err != nil {
		log.Error(err)
		return err
	}

	err = odb.ForEach(func(oid *git.Oid) error {
		obj, err := r.Repo.Lookup(oid)
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
	return nil
}
