package pkg

type RideIdAlreadyExists string

func (r RideIdAlreadyExists) Error() string {
	return string(r)
}
