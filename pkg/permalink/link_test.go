package permalink

import "testing"

func TestGithubPermalink(t *testing.T) {
	want := "https://github.com/example/repo/blob/COMMIT/test.go#L1-L5"

	got := GithubPermaLink(
		"https://github.com/example/repo",
		"COMMIT",
		"test.go",
		1,
		5,
	)

	if want != got {
		t.Errorf("expected %s, got %s", want, got)
	}
}
