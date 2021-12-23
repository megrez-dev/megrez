package router

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/megrez/internal/entity/vo"
)

func TestCaculatePage(t *testing.T) {
	page := vo.CaculatePage(1, 2, 6)
	b, err := json.Marshal(page)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(string(b))
}