package rolling

type SlideWindow struct {
	Success *Number
	Failed  *Number
	Timeout *Number
}

func NewSlideWindow(timeLen int64) *SlideWindow {
	return &SlideWindow{
		Success: NewNumber(timeLen),
		Failed:  NewNumber(timeLen),
		Timeout: NewNumber(timeLen),
	}
}
