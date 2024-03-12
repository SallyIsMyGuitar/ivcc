// Code generated by "stringer -type=EVSEState -output=./state_string.go"; DO NOT EDIT.

package mennekes

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NotInitialized-0]
	_ = x[Idle-1]
	_ = x[EVConnected-2]
	_ = x[PreconditionsValid-3]
	_ = x[ReadyToCharge-4]
	_ = x[Charging-5]
	_ = x[Error-6]
	_ = x[ServiceMode-7]
}

const _EVSEState_name = "NotInitializedIdleEVConnectedPreconditionsValidReadyToChargeChargingErrorServiceMode"

var _EVSEState_index = [...]uint8{0, 14, 18, 29, 47, 60, 68, 73, 84}

func (i EVSEState) String() string {
	if i >= EVSEState(len(_EVSEState_index)-1) {
		return "EVSEState(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _EVSEState_name[_EVSEState_index[i]:_EVSEState_index[i+1]]
}
