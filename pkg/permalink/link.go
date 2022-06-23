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
package permalink

import (
	"fmt"
	"log"

	"github.com/aaqaishtyaq/git-link/pkg/git"
)

// Generate permalink URL
func Generate(args []string) {
	g, err := git.NewFileContext(args[0])
	if err != nil {
		log.Fatal(err)
	}

	r, err := g.Remote()
	if err != nil {
		log.Fatal(err)
	}

	sha, err := g.HeadCommitSha()
	if err != nil {
		log.Fatal(err)
	}

	url := GithubPermaLink(r, sha, g.Filepath, g.Start, g.End)

	log.Default().Println(url)
}

// GithubPermaLink Github permalink
func GithubPermaLink(remote, commit, path string, start, end int) string {
	url := fmt.Sprintf("%s/blob/%s/%s", remote, commit, path)

	if start != 0 {
		url = fmt.Sprintf("%s#L%d", url, start)
		if end != 0 {
			url = fmt.Sprintf("%s-L%d", url, end)
		}
	}

	return url
}
