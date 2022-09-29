// Code generated by "enumer -type Feature"; DO NOT EDIT.

package api

import (
	"fmt"
	"strings"
)

const _FeatureName = "Offline"

var _FeatureIndex = [...]uint8{0, 7}

const _FeatureLowerName = "offline"

func (i Feature) String() string {
	i -= 1
	if i < 0 || i >= Feature(len(_FeatureIndex)-1) {
		return fmt.Sprintf("Feature(%d)", i+1)
	}
	return _FeatureName[_FeatureIndex[i]:_FeatureIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _FeatureNoOp() {
	var x [1]struct{}
	_ = x[Offline-(1)]
}

var _FeatureValues = []Feature{Offline}

var _FeatureNameToValueMap = map[string]Feature{
	_FeatureName[0:7]:      Offline,
	_FeatureLowerName[0:7]: Offline,
}

var _FeatureNames = []string{
	_FeatureName[0:7],
}

// FeatureString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func FeatureString(s string) (Feature, error) {
	if val, ok := _FeatureNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _FeatureNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Feature values", s)
}

// FeatureValues returns all values of the enum
func FeatureValues() []Feature {
	return _FeatureValues
}

// FeatureStrings returns a slice of all String values of the enum
func FeatureStrings() []string {
	strs := make([]string, len(_FeatureNames))
	copy(strs, _FeatureNames)
	return strs
}

// IsAFeature returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Feature) IsAFeature() bool {
	for _, v := range _FeatureValues {
		if i == v {
			return true
		}
	}
	return false
}
