package parser

import "fmt"

// ErrorCode holds the type of error return
type ErrorCode int

const (
	// InsuficientData is returned when the sheet does not contain the minimal two mandatory blocks
	InsuficientData = 000

	// InvalidKeyValuePair is returned when a header line isn't of the right format
	InvalidKeyValuePair = 100

	// EmptyMetaKey is returned when the meta has an empty key
	EmptyMetaKey = 101

	// EmptyMetaValue is returned when the meta has an empty value
	EmptyMetaValue = 102

	// InvalidOptions is returned when the options provided are not allowed
	InvalidOptions = 103

	// DuplicateMeta is returned when a meta is specified more than once
	DuplicateMeta = 104

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

// errorMsgs contains the messages associated to the error codes.
var errorMsgs = make(map[ErrorCode]string)

// init sets the errorMsg map
func init() {
	errorMsgs[InsuficientData] = "insufficient data: the sheet requires at least a header block and a characteristic block"
	errorMsgs[InvalidKeyValuePair] = "invalid pair key:value: the header line is not in the proper format"
	errorMsgs[EmptyMetaKey] = "empty key: the header line's key is empty"
	errorMsgs[EmptyMetaValue] = "empty value: the header line's value is empty"
	errorMsgs[InvalidOptions] = "invalid options: the header's options are incorrect"
	errorMsgs[DuplicateMeta] = "duplicate meta: the header's history is already defined"
	errorMsgs[NoDate] = "empty date: the session's header contains no date"
	errorMsgs[InvalidReward] = "invalid reward: the session's reward is not properly set"
	errorMsgs[RewardAlreadyFound] = "duplicate reward: the session has more than one reward"
	errorMsgs[WrongRewardPosition] = "wrong reward position: the session's reward should be in second or last position"
	errorMsgs[InvalidUpgrade] = "invalid upgrade: the upgrade's format is invalid"
	errorMsgs[InvalidMark] = "invalid mark: upgrade's mark should be \"+\", \"-\" or \"=\""
	errorMsgs[InvalidCost] = "invalid cost: the upgrade's cost is not properly formated"
	errorMsgs[CostAlreadyFound] = "duplicate cost:  the upgrade has more than one cost"
	errorMsgs[WrongCostPosition] = "wrong cost position: the upgrade's cost should be in second or last position"
	errorMsgs[EmptyName] = "empty name: the character's name is empty"
	errorMsgs[InvalidCharacteristicFormat] = "invalid characteristic format: the characteristic is not properly formated"
	errorMsgs[NotIntegerCharacteristicValue] = "invalid characteristic value: the characteristic value is not an integer"
}

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
	msg, found := errorMsgs[e.Code]
	if !found {
		panic(fmt.Sprintf("undefined error message for code %d", e.Code))
	}
	return fmt.Sprintf("line %d: %s", e.Line, msg)
}
