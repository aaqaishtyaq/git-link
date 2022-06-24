package git

import "testing"

func TestSshtoHttpConversion(t *testing.T) {
	sshurl := "git@github.com:example/repo.git"
	want := "https://github.com/example/repo"

	got := SSHToHTTPS(sshurl)

	if want != got {
		t.Errorf("expected %s, got %s", want, got)
	}
}
