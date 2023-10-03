package pkg

type RideSharingApp struct {
	drivers map[string]*Driver
	riders  map[string]*Rider
}

func NewRideSharingApp() *RideSharingApp {
	return &RideSharingApp{
		drivers: make(map[string]*Driver),
		riders:  make(map[string]*Rider),
	}
}
