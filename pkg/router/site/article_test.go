package site

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/megrez/pkg/entity/vo"
)

func TestCalculatePage(t *testing.T) {
	page := vo.CalculatePagination(1, 2, 6)
	b, err := json.Marshal(page)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	fmt.Println(string(b))
}
