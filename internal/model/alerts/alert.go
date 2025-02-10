package alerts

type Alert struct {
	id        int
	alertType string
}

func NewAlert(id int, alertType string) *Alert {
	return &Alert{id, alertType}
}

func (a *Alert) GetType() string {
	return a.alertType
}
