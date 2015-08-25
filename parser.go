const(	
	REWARD_REGEX = ^( )*\+( )*(\d)*( )*xp$
)

func Parse(filename string) (c Character, err error) {
	var c Character
	var err error

	f, err := os.Open(filename)
	if err != nil {
		err = fmt.Errorf("Unable to open character: \"%s.\"", filename)
		return
	}
	scanner := bufio.NewScanner(f)
	
	h, err = scanHeader(scanner)
	if err != nil {
		err = fmt.Errorf("Incorrect formating in header: \"%s.\"", filename)
		return
	}
	c.AddHeader(h)
	
	it := 0;
	for scanner.Scan() {
		it++;
		s, err := scanSession(scanner)
		if err != nil {
			err = fmt.Errorf("Incorrect formating in session %d: \"%s.\"", it, filename)
			return
		}
		c.AddSession(s)
	}
	
	err = scanner.Err()
	if err != nil {
		err = fmt.Errorf("Error during scan: \"%s.\"", filename)
		return
	}
}

// Scans the pairs key:value in the lines and returns the header
func scanHeader(scanner *bufio.Scanner) (h Header, err error) {
	var h Header
	var err error
	
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "\n" {
			return
		}
		err = h.addMetadata(text)
		if err != nil {
			return
		}
	}
}

// Reads line until the session_end token is reached and returns the session
func scanSession(scanner *bufio.Scanner) (s Session, err error) {
	var s Session
	var err error
	
	// scan label
	text := strings.TrimSpace(scanner.Text())
	err = s.addLabel(text)
	if err != nil {
		return
	}
	
	// scan potential reward
	text := strings.TrimSpace(scanner.Text())
	if regexp.Match(REWARD_REGEX, text) {
		s.addReward(text)
		text := nextLine(scanner)
	}
	
	// scan upgrades
	for scanner.Scan() {
		text := strings.TrimSpace(scanner.Text())
		if text == "\n" {
			return 
		}
		err = s.addUpgrade(text)
		if err != nil {
			return
		}
	}
}