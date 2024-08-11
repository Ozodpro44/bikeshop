package models

type Bike struct {
	ID    int
	Name  string
	Price float64
	Image string
}

type GetBikesList struct {
	Bikes []*Bike
	Count int32
}
