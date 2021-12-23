package vo

type Pagination struct {
	Current int
	Pre     int
	Next    int
	Range   []int
	First   int
	Last    int
}

func CaculatePagination(current, size, count int) *Pagination {
	max := count / size
	if count % size > 0 {
		max++
	}
	
	pagination := &Pagination{}
	pagination.Current = current
	if current > 0 {
		pagination.Pre = current - 1
	}
	if current < max {
		pagination.Next = current + 1
	}
	rangeStart := 1
	rangeEnd := max
	if current > 3 {
		pagination.First = 1
		rangeStart = current - 2
	}
	if current < max-2 {
		pagination.Last = max
		rangeEnd = current + 2
	}
	pageRange := []int{}
	for i := rangeStart; i <= rangeEnd; i++ {
		pageRange = append(pageRange, i)
	}
	pagination.Range = pageRange
	return pagination
}
