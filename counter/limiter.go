package counter

type Limiter struct{
	Size int
	Rate int
	Allowed bool
}

func NewLimiter(size int, rate int) *Limiter{
	return &Limiter{Size: size,Rate: rate}
}


