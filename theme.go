package main

import "github.com/fatih/color"

type Theme struct {
	Title func(...interface{}) string
	Error func(...interface{}) string
	Value func(...interface{}) string
}

var theme Theme

func init() {
	theme.Title = color.New(color.FgGreen, color.Bold).SprintFunc()
	theme.Error = color.New(color.FgRed, color.Bold).SprintFunc()
	theme.Value = color.New(color.FgYellow, color.Bold).SprintFunc()
}
