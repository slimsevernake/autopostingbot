// Code generated by "stringer -type=MediaType"; DO NOT EDIT.

package command

import "strconv"

const _MediaType_name = "PhotoVideo"

var _MediaType_index = [...]uint8{0, 5, 10}

func (i MediaType) String() string {
	if i < 0 || i >= MediaType(len(_MediaType_index)-1) {
		return "MediaType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _MediaType_name[_MediaType_index[i]:_MediaType_index[i+1]]
}