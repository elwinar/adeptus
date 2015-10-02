package main

import "fmt"

// ErrorCode holds the type of error return.
type ErrorCode int

const (
	// InsuficientData is returned when the sheet does not contain the minimal two mandatory blocks.
	InsuficientData ErrorCode = 000

	// InvalidKeyValuePair is returned when a header line isn't of the right format.
	InvalidKeyValuePair ErrorCode = 100

	// EmptyMetaKey is returned when the meta has an empty key.
	EmptyMetaKey ErrorCode = 101

	// EmptyMetaValue is returned when the meta has an empty value.
	EmptyMetaValue ErrorCode = 102

	// InvalidOptions is returned when the options provided are not allowed.
	InvalidOptions ErrorCode = 103

	// DuplicateMeta is returned when a meta is specified more than once.
	DuplicateMeta ErrorCode = 104

	// NoDate is returned when no date of the right format is found in the headline.
	NoDate ErrorCode = 200

	// InvalidReward is returned when the reward of a headline isn't correctly formed.
	InvalidReward ErrorCode = 201

	// RewardAlreadyFound is returned when two or more rewards are found for the same headline.
	RewardAlreadyFound ErrorCode = 202

	// WrongRewardPosition is returned when the reward isn't on the second or last position of a headline.
	WrongRewardPosition ErrorCode = 203

	// InvalidUpgrade is returned when the upgrade format is invalid.
	InvalidUpgrade ErrorCode = 300

	// InvalidMark is returned when the mark of an upgrade isn't recognized.
	InvalidMark ErrorCode = 301

	// InvalidCost is returned when the cost of an upgrade isn't correctly formed.
	InvalidCost ErrorCode = 302

	// CostAlreadyFound is returned when two or more costs are found for the same upgrade.
	CostAlreadyFound ErrorCode = 303

	// WrongCostPosition is returned when the cost isn't on the second or last position of an upgrade.
	WrongCostPosition ErrorCode = 304

	// EmptyName is returned when the name of an upgrade is empty.
	EmptyName ErrorCode = 305

	// InvalidCharacteristicFormat is returned when the characteristic format is incorrect.
	InvalidCharacteristicFormat ErrorCode = 400

	// NotIntegerCharacteristicValue is returned when the characteristic value is not a positive integer.
	NotIntegerCharacteristicValue ErrorCode = 401
)

// errorMsgs contains the messages associated to the error codes.
var errorMsgs = map[ErrorCode]string{
	InsuficientData:               "insufficient data: the sheet requires at least a header block and a characteristic block",
	InvalidKeyValuePair:           "invalid pair key:value: the header line is not in the proper format",
	EmptyMetaKey:                  "empty key: the header line's key is empty",
	EmptyMetaValue:                "empty value: the header line's value is empty",
	InvalidOptions:                "invalid options: the header's options are incorrect",
	DuplicateMeta:                 "duplicate meta: the header's history is already defined",
	NoDate:                        "empty date: the session's header contains no date",
	InvalidReward:                 "invalid reward: the session's reward is not properly set",
	RewardAlreadyFound:            "duplicate reward: the session has more than one reward",
	WrongRewardPosition:           "wrong reward position: the session's reward should be in second or last position",
	InvalidUpgrade:                "invalid upgrade: the upgrade's format is invalid",
	InvalidMark:                   "invalid mark: upgrade's mark should be \"+\", \"-\" or \"=\"",
	InvalidCost:                   "invalid cost: the upgrade's cost is not properly formated",
	CostAlreadyFound:              "duplicate cost:  the upgrade has more than one cost",
	WrongCostPosition:             "wrong cost position: the upgrade's cost should be in second or last position",
	EmptyName:                     "empty name: the character's name is empty",
	InvalidCharacteristicFormat:   "invalid characteristic format: the characteristic is not properly formated",
	NotIntegerCharacteristicValue: "invalid characteristic value: the characteristic value is not an integer",
}

// Error is an error encountered when running the application.
type Error struct {
	Line int
	Code ErrorCode
}

// NewError build a new error from the line and error code.
func NewError(line int, code ErrorCode) Error {
	return Error{
		Line: line,
		Code: code,
	}
}

// Error implements the Error interface.
func (e Error) Error() string {
	msg, found := errorMsgs[e.Code]
	if !found {
		panic(fmt.Sprintf("undefined error message for code %d", e.Code))
	}
	return fmt.Sprintf("line %d: %s", e.Line, msg)
}
