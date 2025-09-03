package gitinfo

import "testing"

func TestGetCommitHash(t *testing.T) {
	hash, err := GetCommitHash()
	if err != nil {
		t.Errorf("Erro ao obter hash do commit: %v", err)
	}
	if hash == "" {
		t.Error("Hash do commit está vazio")
	}
}
