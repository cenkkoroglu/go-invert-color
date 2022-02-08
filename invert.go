package invert

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Invert(color string) (string, error) {
	color = strings.Trim(color, " ")
	color = strings.ToLower(color)

	startWithHash := false
	if strings.HasPrefix(color, "#") {
		startWithHash = true
		color = strings.Replace(color, "#", "", -1)
	}

	if len(color) == 3 || len(color) == 6 {
		if len(color) == 3 {
			var regex = regexp.MustCompile(`(.)(.)(.)`)
			color = regex.ReplaceAllString(color, `$1$1$2$2$3$3`)
		}
	} else {
		return "", errors.New(fmt.Sprintf("invalid hex length (%d). Length must be 3 or 6 characters", len(color)))
	}

	regex, _ := regexp.Compile(`^[a-f0-9]{6}$`)

	if !regex.MatchString(color) {
		return "", errors.New(fmt.Sprintf("invalid hex string #%s", color))
	}

	rHexDec, _ := hexDec(color[0:2])
	r := decHex(255 - rHexDec)
	if !(len(r) > 1) {
		r = "0" + r
	}

	gHexDec, _ := hexDec(color[2:4])
	g := decHex(255 - gHexDec)
	if !(len(g) > 1) {
		g = "0" + g
	}

	bHexDec, _ := hexDec(color[4:6])
	b := decHex(255 - bHexDec)
	if !(len(b) > 1) {
		b = "0" + b
	}

	result := r + g + b
	if startWithHash {
		result = "#" + r + g + b
	}

	return result, nil
}

func decHex(number int64) string {
	return strconv.FormatInt(number, 16)
}

func hexDec(str string) (int64, error) {
	return strconv.ParseInt(str, 16, 0)
}
