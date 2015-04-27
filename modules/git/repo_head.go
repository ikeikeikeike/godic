package git

import (
	"fmt"

	git "github.com/libgit2/git2go"
)

type Head struct {
	repo *git.Repository
}

func (h *Head) CommitTree() (tree *git.Tree, err error, noHead bool) {
	commit, err, noHead := h.Commit()
	if err != nil {
		return
	}
	tree, err = commit.Tree()
	return
}

func (h *Head) Commit() (commit *git.Commit, err error, noHead bool) {
	oid, err, noHead := h.CommitId()
	if err != nil {
		return
	}
	commit, err = h.repo.LookupCommit(oid)
	return
}

func (h *Head) CommitId() (oid *git.Oid, err error, noHead bool) {
	headRef, err := h.repo.LookupReference("HEAD")
	if err != nil {
		return
	}
	ref, err := headRef.Resolve()
	if err != nil {
		noHead = true
		return
	}
	oid = ref.Target()
	if oid == nil {
		err = fmt.Errorf("Could not get Target for HEAD(%s)\n", oid.String())
	}
	return
}

func (h *Head) Repo() (*git.Repository) {
	headRef, err := h.repo.LookupReference("HEAD")
	if err != nil {
		headRef, _ = h.repo.Head()
	}

	ref, _ := headRef.Resolve()
	return ref.Owner()
}
