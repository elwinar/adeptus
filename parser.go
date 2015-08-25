const(	
	REWARD_REGEX = ^( )*\+ (\d)* xp$
	DATE_LABEL_REGEX = ^(\d){2}[/-](\d){2}[/-](\d){4} (.)*$
)

func Parse(filename string) (Character, error) {
	f, err := os.Open(filename)
	if err != nil {
		// ... error
	}
	scanner := bufio.NewScanner(f)
	
	Header, err = scanHeader(&scanner)
	Session, err = scanSession(&scanner)
	
	err = scanner.Err()
	if err != nil {
		// ... error
	}
}

// Reads line until the header_end token is reached and returns the header
func scanHeader(scanner *bufio.Scanner) (h Header, err error) {
	h := Header{}
	err := nil
	
	for scanner.Scan() {
		text := scanner.Text()
		if text == "\n" {
			return
		}
		err = h.addMetadata(text)
		if err != nil {
			// ... error
		}
	}
}

// Reads line until the session_end token is reached and returns the session
func scanSession(scanner *bufio.Scanner) (s Session, err error) {
	s := Session{}
	err := nil
	
	// scan label
	text := scanner.Text()
	if regexp.Match(DATE_LABEL_REGEX, text) {
		err = s.addLabel(text)
		if err != nil {
			// ... error
		}
	}
	else {
		err = errors.New("A label is required")
		// ... error (no label, generate error)
	}
	
	// scan potential reward
	text = scanner.Text()
	if regexp.Match(REWARD_REGEX, text) {
		err = s.addReward(text)
		if err != nil {
			// ... error
		}
		text = scanner.Text()
	}
	
	// scan upgrades
	for scanner.Scan() {
		text = scanner.Text()
		if text == "\n" {
			return s, nil
		}
		err = s.addUpgrade(text)
		if err != nil {
			// ... error
		}
	}
}