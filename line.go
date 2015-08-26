package main

import(
	"regexp"
)

var formats []string{
	"2006/02/03",
	"2006-02-03",
	"2006.02.03",
}

const(
	regex_date = `\d{4}[-/\.]\d{2}[-/\.]\d{2}`
	regex_xp = `\(?\d*xp\)?`
	regex_mark = `[\*\+\-]`
	regex_value = `([\+\-])?\d+`
	regex_blank = `[\t ]*`
	regex_date_separator = `[/-\.]`
	date_format = `2006/02/03`
)

type Line struct {
	raw	string
}

// creates a new line given a raw value
func NewLine(raw string) *Line {
	return &Line{
		raw: raw,
	}
}

// returns true if the line contains a date
func (l *line) HasDate() (m bool) {
	m, _ regexp.MatchString(`^` + regex_blank + regex_date, l.raw)
	return
}

// returns the date within the line and removes it
func (l *line) GetDate() (s string) {
	if !l.HasDate() {
		return
	}
	r := regexp.MustCompile(`^` + regex_blank + regex_date)
	s = r.FindString(l.raw)
	l.raw = strings.Replace(l.raw, s, "", 1)
	s = r.ReplaceAllString(s, "/")
	s = strings.TrimSpace(s)
}

// returns true if the line contains xp
func (l *line) HasXp() (m bool) {
	m, _ regexp.MatchString(regex_xp, l.raw)
	return
}

// returns the xp within the line and removes it
func (l *line) GetXp() (s string) {
	if !l.HasXp() {
		return
	}
	r := regexp.MustCompile(regex_xp)
	s = r.FindString(l.raw)
	l.raw = strings.Replace(l.raw, s, "", 1)
	s = regexp.MustCompile(`\d*`).FindString(s)
}

// returns true if the line contains mark
func (l *line) HasMark() (m bool) {
	m, _ regexp.MatchString(regex_mark, l.raw)
	return
}

// returns the mark within the line and removes it
func (l *line) GetMark() (s string) {
	if !l.HasMark() {
		return
	}
	r := regexp.MustCompile(regex_mark)
	s = r.FindString(l.raw)
	l.raw = strings.Replace(l.raw, s, "", 1)
	s = strings.TrimSpace(s)
}

// returns true if the line contains value
func (l *line) HasValue() (m bool) {
	m, _ regexp.MatchString(regex_value, l.raw)
	return
}

// returns the value within the line and removes it
func (l *line) GetValue() (s string) {
	if !l.HasValue() {
		return
	}
	r = regexp.MustCompile(regex_value)
	s = r.FindString(l.raw)
	l.raw = strings.Replace(l.raw, s, "", 1)
}

// returns true if the line contains an label
func (l *line) HasKey() bool {
	return strings.Contains(l.raw, ":")
}

// returns the key within the line and removes it
func (l *line) GetKey() (s string) {
	if !l.HasKey() {
		return
	}
	split := strings.SplitN(l.raw, ":", 2)
	s = split[0]
	l.raw = split[1]
	s = strings.TrimSpace(s)
	return
}

// returns the label within the line and removes it
func (l *line) GetLabel() (s string) {
	return strings.TrimSpace(l.raw)
}