package main

import (
	"fmt"
)

// ErrorCode holds the type of error return.
type ErrorCode int

// Here is the list of defined error codes.
const (
	// Parse sheet errors
	InvalidCharacterSheet ErrorCode = iota

	InvalidHeaderLine
	EmptyHeaderKey
	EmptyHeaderValue
	DuplicateHeaderLine
	InvalidHeaderOptions

	UndefinedSessionDate
	InvalidSessionReward
	DuplicateSessionReward
	ForbidenRewardPosition

	EmptyUpgrade
	InvalidUpgradeFormat
	InvalidUpgradeValue
	InvalidUpgradeMark
	InvalidUpgradeCost
	DuplicateUpgradeCost
	ForbidenUpgradeMark
	ForbidenCostPosition
	ForbidenUpgradeLoss
	ForbidenUpgradeValue
	DuplicateUpgrade

	UndefinedTypeCost
	UndefinedMatchCost
	UndefinedTierCost

	UndefinedCharacteristic
	UndefinedBackground

	UnitTest
)

// errorMsgs contains the messages associated to the error codes.
var errorMsgs = map[ErrorCode]string{

	InvalidCharacterSheet: `the character sheet requires at least a header block and a characteristic block`,

	InvalidHeaderLine:    `line %d: the header line format is invalid`,
	EmptyHeaderKey:       `line %d: the header line key is empty`,
	EmptyHeaderValue:     `line %d: the header line value is empty`,
	DuplicateHeaderLine:  `line %d: the header line is already set`,
	InvalidHeaderOptions: `line %d: the background options are incorrect`,

	UndefinedSessionDate:   `line %d: the session date is not defined`,
	InvalidSessionReward:   `line %d: the session reward is invalid`,
	DuplicateSessionReward: `line %d: the session reward is already set`,
	ForbidenRewardPosition: `line %d: bad session reward position`,

	EmptyUpgrade:         `line %d: the upgrade name is not defined`,
	InvalidUpgradeFormat: `line %d: the upgrade format is invalid`,
	InvalidUpgradeValue:  `line %d: the upgrade value is invalid`,
	InvalidUpgradeMark:   `line %d: the upgrade mark is invalid`,
	InvalidUpgradeCost:   `line %d: the upgrade cost is invalid`,
	DuplicateUpgradeCost: `line %d: the upgrade cost is already set`,
	ForbidenUpgradeMark:  `line %d: the upgrade mark is forbiden`,
	ForbidenCostPosition: `line %d: bad upgrade cost position`,
	ForbidenUpgradeLoss:  `line %d: the upgrade is absent from sheet`,
	ForbidenUpgradeValue: `line %d: the upgrade value is forbiden`,
	DuplicateUpgrade:     `line %d: the upgrade is already set`,

	UndefinedTypeCost:  `undefined cost for type %s`,
	UndefinedMatchCost: `undefined cost for type %s with %d matching aptitudes`,
	UndefinedTierCost:  `undefined cost for type %s with %d matching aptitudes on tier %d`,

	UndefinedCharacteristic: `line %d: the characteristic is not defined`,
	UndefinedBackground:     `line %d: the background is not defined`,

	UnitTest: `should not be seen outside unit testing`,
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
