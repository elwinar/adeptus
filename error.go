package main

import (
	"fmt"
)

// ErrorCode holds the type of error return.
type ErrorCode int

// Here is the list of defined error codes.
const (
	UnusedAptitude    ErrorCode = 000
	UndefinedAptitude ErrorCode = 001

	InvalidCharacterSheet ErrorCode = 100

	InvalidHeaderLine        ErrorCode = 200
	EmptyHeaderKey           ErrorCode = 201
	EmptyHeaderValue         ErrorCode = 202
	DuplicateHeaderLine      ErrorCode = 203
	InvalidBackgroundOptions ErrorCode = 204
	UndefinedBackgroundType  ErrorCode = 205
	UndefinedBackgroundValue ErrorCode = 206

	UndefinedSessionDate     ErrorCode = 300
	InvalidSessionReward     ErrorCode = 301
	DuplicateSessionReward   ErrorCode = 302
	BadSessionRewardPosition ErrorCode = 303

	InvalidUpgrade         ErrorCode = 400
	UndefinedUpgradeName   ErrorCode = 401
	InvalidUpgradeMark     ErrorCode = 402
	InvalidUpgradeCost     ErrorCode = 403
	DuplicateUpgradeCost   ErrorCode = 404
	BadUpgradeCostPosition ErrorCode = 405
	MismatchMarkCost       ErrorCode = 406

	InvalidCharacteristicFormat ErrorCode = 500
	InvalidCharacteristicValue  ErrorCode = 501
	UndefinedCharacteristic     ErrorCode = 502
	DuplicateCharacteristic     ErrorCode = 503
	MissingCharacteristic       ErrorCode = 504

	UndefinedTypeCost  ErrorCode = 600
	UndefinedMatchCost ErrorCode = 601
	UndefinedTierCost  ErrorCode = 602
)

// errorMsgs contains the messages associated to the error codes.
var errorMsgs = map[ErrorCode]string{

	UnusedAptitude:    `the aptitude %s is defined but not used by the universe`,
	UndefinedAptitude: `the aptitude %s is used but not defined by the universe`,

	InvalidCharacterSheet: `the character sheet requires at least a header block and a characteristic block`,

	InvalidHeaderLine:        `line %d: the header line format is invalid`,
	EmptyHeaderKey:           `line %d: the header line key is empty`,
	EmptyHeaderValue:         `line %d: the header line value is empty`,
	DuplicateHeaderLine:      `line %d: the header line is already set`,
	InvalidBackgroundOptions: `line %d: the background options are incorrect`,
	UndefinedBackgroundType:  `line %d: the background %s is not defined by the universe`,
	UndefinedBackgroundValue: `line %d: the background %s of type %s is not defined by the universe`,

	UndefinedSessionDate:     `line %d: the session date is not defined`,
	InvalidSessionReward:     `line %d: the session reward is invalid`,
	DuplicateSessionReward:   `line %d: the session reward is already set`,
	BadSessionRewardPosition: `line %d: bad session reward position`,

	InvalidUpgrade:         `line %d: the upgrade format is invalid`,
	UndefinedUpgradeName:   `line %d: the upgrade name is not defined`,
	InvalidUpgradeMark:     `line %d: the upgrade mark is invalid`,
	InvalidUpgradeCost:     `line %d: the upgrade cost is invalid`,
	DuplicateUpgradeCost:   `line %d: the upgrade cost is already set`,
	BadUpgradeCostPosition: `line %d: bad upgrade cost position`,
	MismatchMarkCost:       `line %d: upgrade with mark "-" expects no cost`,

	InvalidCharacteristicFormat: `line %d: the characteristic format is invalid`,
	InvalidCharacteristicValue:  `line %d: the characteristic value is invalid`,
	UndefinedCharacteristic:     `line %d: the characteristic is not defined`,
	DuplicateCharacteristic:     `line %d: the characteristic is already set`,
	MissingCharacteristic:       `the characteristic %s is not defined for the character`,

	UndefinedTypeCost:  `undefined cost for type %s`,
	UndefinedMatchCost: `undefined cost for type %s with %d matching aptitudes`,
	UndefinedTierCost:  `undefined cost for type %s with %d matching aptitudes on tier %d`,
}

// Error is an error holding a code and variadic printable data.
type Error struct {
	Code ErrorCode
	vars []interface{}
}

// NewError build a new error from an error code.
func NewError(code ErrorCode, v ...interface{}) Error {
	return Error{
		Code: code,
		vars: v,
	}
}

// Error implements the Error interface.
func (e Error) Error() string {
	msg, found := errorMsgs[e.Code]
	if !found {
		panic(fmt.Sprintf("undefined error message for code %d", e.Code))
	}
	return fmt.Sprintf(msg, e.vars...)
}
