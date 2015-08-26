package main

import (
	"time"
)

type Session struct {
	Date       time.Time
	Label      string
	Experience int
	Upgrades   []Upgrade
}

type rawSession struct {
	date       string
	label      string
	experience string
}
