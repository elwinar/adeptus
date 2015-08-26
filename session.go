package adeptus

type Session struct {
	Date time.Time
	Title string
	Reward int
	Upgrades []Upgrade
}

func ParseSession(raw string) (Session, error) {
	return Session{}, nil
}
