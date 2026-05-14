package parsing

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dharmab/goacmi/objects"
)

// ParseTimeFrame attemps to parse a line of ACMI data describing a time frame into a [time.Duration].
func ParseTimeFrame(line string) (time.Duration, error) {
	if !strings.HasPrefix(line, "#") {
		return 0, fmt.Errorf("line does not contain TimeFrame: %s", line)
	}
	seconds, err := strconv.ParseFloat(strings.TrimPrefix(line, "#"), 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing duration: %w", err)
	}
	duration := time.Duration(seconds*1000) * time.Millisecond
	return duration, nil
}

// ParseObjectUpdate attemps to parse a line of ACMI data describing an object update into an [objects.Update].
func ParseObjectUpdate(line string) (*objects.Update, error) {
	update := &objects.Update{}

	if strings.HasPrefix(line, "-") {
		update.IsRemoval = true
		line = line[1:]
	}

	idStr, propertiesStr, _ := strings.Cut(line, ",")

	id, err := parseID(idStr)
	if err != nil {
		return nil, err
	}
	update.ID = id

	properties, err := parseProperties(propertiesStr)
	if err != nil {
		return nil, err
	}
	update.Properties = properties

	return update, nil
}

func parseID(idStr string) (uint64, error) {
	id, err := strconv.ParseUint(idStr, 16, 64)
	if err != nil {
		return 0, fmt.Errorf("error parsing object ID: %w", err)
	}
	return id, nil
}

func parseProperties(propertiesStr string) (map[string]string, error) {
	properties := make(map[string]string)
	if propertiesStr == "" {
		return properties, nil
	}
	for _, prop := range splitOnUnescapedCommas(propertiesStr) {
		key, value, ok := strings.Cut(prop, "=")
		if !ok {
			return nil, fmt.Errorf("error parsing property: %s", prop)
		}
		properties[key] = strings.TrimSpace(value)
	}
	return properties, nil
}

// splitOnUnescapedCommas splits s on commas not preceded by a backslash,
// unescaping \, to , in each resulting segment.
func splitOnUnescapedCommas(s string) []string {
	var parts []string
	var current strings.Builder
	for i := 0; i < len(s); i++ {
		if s[i] == '\\' && i+1 < len(s) && s[i+1] == ',' {
			current.WriteByte(',')
			i++
		} else if s[i] == ',' {
			parts = append(parts, current.String())
			current.Reset()
		} else {
			current.WriteByte(s[i])
		}
	}
	return append(parts, current.String())
}
