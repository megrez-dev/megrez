package vo

type Page struct {
	Current int
	Pre     int
	Next    int
	Range   []int
	First   int
	Last    int
}

func CaculatePage(current, size, count int) *Page {
	max := count / size
	if count % size > 0 {
		max++
	}
	
	page := &Page{}
	page.Current = current
	if current > 0 {
		page.Pre = current - 1
	}
	if current < max {
		page.Next = current + 1
	}
	rangeStart := 1
	rangeEnd := max
	if current > 3 {
		page.First = 1
		rangeStart = current - 2
	}
	if current < max-2 {
		page.Last = max
		rangeEnd = current + 2
	}
	pageRange := []int{}
	for i := rangeStart; i <= rangeEnd; i++ {
		pageRange = append(pageRange, i)
	}
	page.Range = pageRange
	return page
}
