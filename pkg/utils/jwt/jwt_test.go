package jwt

import "testing"

func TestToken(t *testing.T) {
	var id uint = 1
	token, err := GenerateToken(id)
	if err != nil {
		t.Errorf("GenerateToken() error = %v", err)
		return
	}
	c, err := ParseToken(token)
	if err != nil {
		t.Errorf("ParseToken() error = %v", err)
		return
	}
	if c.ID != id {
		t.Errorf("ParseToken() error = %v", err)
		return
	}
}
