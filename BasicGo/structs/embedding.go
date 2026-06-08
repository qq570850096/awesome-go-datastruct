package structdemo

type Logger struct {
	Prefix string
}

func (l *Logger) Log(msg string) string {
	return l.Prefix + ": " + msg
}

type Service struct {
	Logger
	Name string
}

func NewService(name string) *Service {
	return &Service{
		Logger: Logger{Prefix: name},
		Name:   name,
	}
}

func (s *Service) Info(msg string) string {
	return s.Log(msg)
}
