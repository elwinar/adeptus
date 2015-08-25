type Session struct {
	date 		dateTime
	label		string
	reward		int
	upgrades	[]Upgrade
}

// adds a label and a date to the session
(s Session) func addLabel(raw string) error {
}

// adds a reward to the session
(s Session) func addReward(raw string) error {
}

// adds an upgrade to the session
(s Session) func addUpgrade(raw string) error {
}