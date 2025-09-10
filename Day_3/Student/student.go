package student

type Student struct {
	Name  string
	Roll  int
	Marks map[string]int
	Avg   float64
	Grade string
}

func (s *Student) Calculate() {
	total := 0
	for _, m := range s.Marks {
		total += m
	}
	s.Avg = float64(total) / float64(len(s.Marks))

	if s.Avg <= 50 {
		s.Grade = "Fail"
	} else if s.Avg < 70 {
		s.Grade = "B"
	} else {
		s.Grade = "A"
	}
}
