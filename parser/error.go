package parser

import "fmt"

type ErrorCode int

const (
	EmptySheet                    = 000 // EmptySheet is returned when the sheet doesn't contain anything other than empty lines (or no line)
	InvalidKeyValuePair ErrorCode = 100 // InvalidKeyValuePair is returned when a header line isn't of the right format
	UnknownKey                    = 101 // UnknownKey is returned when the header key isn't a valid one
	NoDate                        = 200 // NoDate is returned when no date of the right format is found in the headline of a session
	InvalidReward                 = 201 // InvalidReward is returned when the reward of a session headline isn't correctly formed
	RewardAlreadyFound            = 202 // RewardAlreadyFound is returned when two or more rewards are found for the same session headline
	WrongRewardPosition           = 203 // WrongRewardPosition is returned when the reward isn't on the second or last position of a session headline
	InvalidUpgrade                = 300 // InvalidUpgrade is returned when the upgrade format is invalid
	InvalidMark                   = 301 // InvalidMark is returned when the mark of an upgrade isn't recognized
	InvalidCost                   = 302 // InvalidCost is returned when the cost of an upgrade isn't correctly formed
	CostAlreadyFound              = 303 // CostAlreadyFound is returned when two or more costs are found for the same upgrade
	WrongCostPosition             = 304 // WrongCostPosition is returned when the cost isn't on the second or last position of an upgrade
	EmptyName                     = 305 // EmptyName is returned when the name of an upgrade is empty
)

type Error struct {
	Line int
	Code ErrorCode
}

func NewError(line int, code ErrorCode) Error {
	return Error{
		Line: line,
		Code: code,
	}
}

func (e Error) Error() string {
	return fmt.Sprintf("line %d: %d", e.Line, e.Code)
}
