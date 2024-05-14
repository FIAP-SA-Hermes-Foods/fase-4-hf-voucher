package valueobject

import (
	"errors"
	"fmt"
	"regexp"
)

type Id struct {
	Value string `json:"value,omitempty"`
}

var regId = regexp.MustCompile(`[0-9-.]*`)

func (c Id) Validate() error {

	if len(c.Value) == 0 {
		return errors.New("id value is null or invalid")
	}

	idMatch := regId.FindStringSubmatch(c.Value)

	if idMatch == nil {
		return fmt.Errorf("id %s is not valid", c.Value)
	}

	return nil
}
