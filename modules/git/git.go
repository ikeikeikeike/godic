package git

import (
	"bytes"
	"container/list"
)

const prettyLogFormat = `--pretty=format:%H`

func parsePrettyFormatLog(r *Repo, logByts []byte) (*list.List, error) {
	l := list.New()
	if len(logByts) == 0 {
		return l, nil
	}

	for _, commitId := range bytes.Split(logByts, []byte{'\n'}) {
		commit, err := r.GetCommitByHash(string(commitId))
		if err != nil {
			return nil, err
		}
		l.PushBack(commit)
	}

	return l, nil
}
