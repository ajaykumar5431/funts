package array

type Dayss struct {
	Monday  int
	Tuesday int
}

func (z Dayss) Day1() int {
	return z.Monday * z.Tuesday

}
