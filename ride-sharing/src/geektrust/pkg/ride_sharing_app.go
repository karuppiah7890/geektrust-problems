package pkg

type RideSharingApp struct {
	drivers map[string]*Driver
}

func NewRideSharingApp() *RideSharingApp {
	return &RideSharingApp{
		drivers: make(map[string]*Driver),
	}
}
