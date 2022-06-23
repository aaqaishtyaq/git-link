package git

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	git "github.com/go-git/go-git/v5"
)

const (
	filePathSeperator = ":"
	rangeSeperator    = ".."
)

type FileContext struct {
	Filepath string
	Start    int
	End      int
	Repo     *git.Repository
}

func NewFileContext(pathContext string) (*FileContext, error) {
	dir, err := os.Getwd()
	if err != nil {
		return &FileContext{}, err
	}

	repo, err := git.PlainOpen(dir)
	if err != nil {
		return &FileContext{}, err
	}

	contextSlice := strings.Split(pathContext, filePathSeperator)
	path := contextSlice[0]

	var start, end int
	if len(contextSlice) > 1 {
		rangeSlice := strings.Split(contextSlice[1], rangeSeperator)
		if len(rangeSlice) == 1 {
			start, err = strconv.Atoi(rangeSlice[0])
			if err != nil {
				return &FileContext{}, err
			}
		} else if len(rangeSlice) == 2 {
			start, err = strconv.Atoi(rangeSlice[0])
			end, err = strconv.Atoi(rangeSlice[1])
		}
	}

	return &FileContext{
		Filepath: path,
		Start:    start,
		Repo:     repo,
		End:      end,
	}, nil
}

func (f *FileContext) HeadCommitSha() (string, error) {
	h, err := f.Repo.Head()
	if err != nil {
		return "", err
	}

	hash := fmt.Sprintf("%s", h.Hash())
	return hash, nil
}

func (f *FileContext) Remote() (string, error) {
	r, err := f.Repo.Remote("origin")
	if err != nil {
		return "", err
	}

	remote := r.Config().URLs[0]
	if strings.HasPrefix(remote, "git@") {
		remote, err = SSHToHTTPS(remote)
		if err != nil {
			return "", err
		}
	}

	return remote, nil
}

func SSHToHTTPS(remote string) (string, error) {
	// git@github.com:aaqaishtyaq/tools.git
	split1 := strings.Split(remote, "git@")

	var sshSepArr []string
	if len(split1) == 2 {
		httpTrail := strings.TrimSuffix(split1[1], ".git")
		sshSepArr = strings.Split(httpTrail, ":")
	}

	path := fmt.Sprintf("https://%s", strings.Join(sshSepArr, "/"))

	return path, nil
}
