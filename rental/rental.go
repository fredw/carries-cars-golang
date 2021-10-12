package rental

type Rental struct {
	duration int
	distance int
}

func (r Rental) Duration() int {
	return r.duration
}

func NewRental(duration, distance int) Rental {
	return Rental{
		duration: duration,
		distance: distance,
	}
}
