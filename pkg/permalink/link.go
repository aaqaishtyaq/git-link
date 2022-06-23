package permalink

import (
	"fmt"
	"log"

	"github.com/aaqaishtyaq/git-link/pkg/git"
)

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

//https://github.com/aaqaishtyaq/home_ops/blob/2fbc4821a35508b477bb6c2e5466ef56cdb3f95a/modules/default.nix#L2

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
