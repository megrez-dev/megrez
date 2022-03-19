package view

import (
	"testing"
	"time"
)

func TestTimeFormat(t *testing.T) {
	blogBirthStr := "2022-03-19 17:31:11"
	blogBirth, err := time.Parse("2006-01-02 15:04:05", blogBirthStr)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(blogBirth.Format("2006-01-02 15:04:05"))
}
