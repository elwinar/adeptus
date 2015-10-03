package main

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
	
	// UnusedAptitude is returned when an aptitude is defined in the universe but not used.
	UnusedAptitude ErrorCode = 500
	
	// UndefinedAptitude is returned when an non-defined aptitude is used by the universe.
	UndefinedAptitude ErrorCode = 501
	
	// UndefinedBackground is returned when a character uses a non defined backgroud.
	UndefinedBackground ErrorCode = 502
	
	// UndefinedCharacteristic is returned when a character uses a non defined characteristic.
	UndefinedCharacteristic = 503
	
	// DuplicateCharacteristic is returned when a characteristic is found twice for a character
	DuplicateCharacteristic = 504
	
	// MissingCharacteristic is returned when a characteristic is missing from a character's sheet.
	MissingCharacteristic = 505
	
)

// errorMsgs contains the messages associated to the error codes.
var errorMsgs = map[ErrorCode]string{
    
        UnusedAptitude:                "the aptitude %s is defined but not used"
        UndefinedAptitude:             "the aptitude %s is used but not defined"
    
	InsuficientData:               "the sheet requires at least a header block and a characteristic block",
        
	InvalidPairKeyValue:           "line %d: the header line is not in the proper key:value format",
	EmptyKey:                      "line %d: the header line's key is empty",
	EmptyValue:                    "line %d: the header line's value is empty",
	DuplicateMeta:                 "line %d: the header's background %s is already defined",
	InvalidOptions:                "line %d: the header's options are incorrect",
        
	InvalidCharacteristicFormat:   "line %d: the characteristic is not properly formated",
	NotIntegerCharacteristicValue: "line %d: the characteristic value is not an integer",
        
	UndefinedDate:                 "line %d: the session's header contains no date",
        
	InvalidReward:                 "line %d: the session's reward is not properly set",
	RewardAlreadyFound:            "line %d: the session has more than one reward",
	WrongRewardPosition:           "line %d: the session's reward should be in second or last position",
        
	InvalidUpgrade:                "line %d: the upgrade's format is invalid",
	InvalidMark:                   "line %d: upgrade's mark should be \"+\", \"-\" or \"=\"",
	InvalidCost:                   "line %d: the upgrade's cost is not properly formated",
	CostAlreadyFound:              "line %d: the upgrade has more than one cost",
	WrongCostPosition:             "line %d: the upgrade's cost should be in second or last position",
	EmptyName:                     "line %d: the upgrade has no name",
        
        InvalidCharacteristicCase:     "line %d: the characteristic name must be upper case"
        InvalidCharacteristicValue:    "line %d: the characteristic value must be an integer"
        UndefinedCharacteristic:       "line %d: the characteristic is not defined"
        DuplicateCharacteristic:       "line %d: the characteristic is already set"
        MissingCharacteristic:         "the characteristic %s is not defined for the character"
        
        UndefinedBackgroundType:       "line %d: the background type %s is not defined"
        UndefinedBackgroundValue:      "line %d: the background %s of type %s is not defined"
        
        UndefinedTypeCost:             "undefined cost for type %s"
        UndefinedMatchCost:            "undefined cost for type %s with %d matching aptitudes"
        UndefinedTierCost:             "undefined cost for type %s with %d matching aptitudes on tier %d"
        
        MismatchMarkCost:              "line %d: mark "-" expects no cost value"
        
}

// Error is an error holding a code and variadic printable data.
type Error struct {
	Code ErrorCode
	vars []interface{}
}

// NewError build a new error from an error code.
func NewError(code ErrorCode, ...v interface{}) Error {
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