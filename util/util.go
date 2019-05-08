package util

import (
	"bytes"
	"encoding/json"
	"net/url"
	"path"
)

func PrettyFormatJSON(raw []byte) string {
	var tmp bytes.Buffer
	json.Indent(&tmp, raw, "", "    ")
	return string(tmp.Bytes())
}

func JoinBaseAndPath(base, endpoint string) (string, error) {
	baseURL, err := url.Parse(base)
	if err != nil {
		return "", err
	}
	tmp := path.Join(baseURL.Path, endpoint)
	baseURL.Path = tmp
	return baseURL.String(), nil
}
