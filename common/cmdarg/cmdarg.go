package cmdarg

import "strings"

// Arg is used by flag to accept multiple argument.
type Arg []string

func (c *Arg) String() string {
	return strings.Join([]string(*c), " ")
}


func (c *Arg) Set(value string) error {
	*c = append(*c, value)
	return nil
}
