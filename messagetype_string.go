// Code generated by "stringer -type=MessageType"; DO NOT EDIT.

package websocket

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[MessageText-1]
	_ = x[MessageBinary-2]
}

const _MessageType_name = "MessageTextMessageBinary"

var _MessageType_index = [...]uint8{0, 11, 24}

func (i MessageType) String() string {
	i -= 1
	if i < 0 || i >= MessageType(len(_MessageType_index)-1) {
		return "MessageType(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _MessageType_name[_MessageType_index[i]:_MessageType_index[i+1]]
}