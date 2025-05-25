package main

type CatDetail interface {
	PortionToEat(height float64, width float64) float64
	someUnexported()
}

func (*Cat) PortionToEat(height float64, width float64) float64 {
	return height * width
}

func (*Cat) someUnexported() {
}

func (Cat) GrowUp(weight float64) float64 {
	return weight * 2
}
