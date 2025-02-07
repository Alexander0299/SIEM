package logs

type Log struct {
	id       int
	message  string
	severity string
}

func NewLog(id int, message, severity string) *Log {
	return &Log{id, message, severity}
}

func (l *Log) GetSeverity() string {
	return l.severity
}

func (l *Log) SetSeverity(newSeverity string) {
	l.severity = newSeverity
}
