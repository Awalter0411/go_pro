package main

type OpenClose interface {
	open()
	close()
}

type Door struct {
}

func (d Door) open() {

}

func (d Door) close() {

}

func main() {
	var d1 OpenClose
	d1 = Door{}
	d1.open()
}
