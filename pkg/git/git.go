/*
Copyright © 2022 Aaqa Ishtyaq <aaqaishtyaq@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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

// FileContext about the file for which generate URL
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

// HeadCommitSha returns commit Sha for git HEAD
func (f *FileContext) HeadCommitSha() (string, error) {
	h, err := f.Repo.Head()
	if err != nil {
		return "", err
	}

	hash := fmt.Sprintf("%s", h.Hash())
	return hash, nil
}

// Remote returns git remote url
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

// SSHToHTTPS converts Github SSH path to HTTPS
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
