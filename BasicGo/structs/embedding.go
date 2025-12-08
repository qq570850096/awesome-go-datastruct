package structdemo

// Logger 用于演示通过组合/嵌入实现“带日志”的组件。
type Logger struct {
	Prefix string
}

func (l *Logger) Log(msg string) string {
	return l.Prefix + ": " + msg
}

// Service 通过嵌入 Logger 复用日志能力，而不是继承。
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

// Info 直接复用嵌入的 Logger 上的方法。
func (s *Service) Info(msg string) string {
	return s.Log(msg)
}

