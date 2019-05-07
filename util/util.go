package util

import (
	"bytes"
	"encoding/json"
)

func PrettyFormatJSON(raw []byte) string {
	var tmp bytes.Buffer
	json.Indent(&tmp, raw, "", "    ")
	return string(tmp.Bytes())
}
