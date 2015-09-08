package parser

import "fmt"

// ErrorCode holds the type of error return
type ErrorCode int

const (
	// EmptySheet is returned when the sheet is empty
	EmptySheet = 000

	// InvalidKeyValuePair is returned when a header line isn't of the right format
	InvalidKeyValuePair = 100

	// UnknownKey is returned when the header key isn't a valid one
	UnknownKey = 101

	// NoDate is returned when no date of the right format is found in the headline
	NoDate = 200

	// InvalidReward is returned when the reward of a headline isn't correctly formed
	InvalidReward = 201

	// RewardAlreadyFound is returned when two or more rewards are found for the same headline
	RewardAlreadyFound = 202

	// WrongRewardPosition is returned when the reward isn't on the second or last position of a headline
	WrongRewardPosition = 203

	// InvalidUpgrade is returned when the upgrade format is invalid
	InvalidUpgrade = 300

	// InvalidMark is returned when the mark of an upgrade isn't recognized
	InvalidMark = 301

	// InvalidCost is returned when the cost of an upgrade isn't correctly formed
	InvalidCost = 302

	// CostAlreadyFound is returned when two or more costs are found for the same upgrade
	CostAlreadyFound = 303

	// WrongCostPosition is returned when the cost isn't on the second or last position of an upgrade
	WrongCostPosition = 304

	// EmptyName is returned when the name of an upgrade is empty
	EmptyName = 305

	// InvalidCharacteristicFormat is returned when the characteristic format is incorrect
	InvalidCharacteristicFormat = 400

	// NotIntegerCharacteristicValue is returned when the characteristic value is not a positive integer
	NotIntegerCharacteristicValue = 401
)

// Error is an error encountered when parsing the sheet
type Error struct {
	Line int
	Code ErrorCode
}

// NewError build a new error from the line and error code
func NewError(line int, code ErrorCode) Error {
	return Error{
		Line: line,
		Code: code,
	}
}

// Error implements the Error interface
func (e Error) Error() string {
	return fmt.Sprintf("line %d: %d", e.Line, e.Code)
}
