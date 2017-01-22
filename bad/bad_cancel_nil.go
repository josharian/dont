// Package bad contains bad code for templates to match, during development.
// It'll all get deleted when I understand better the code structure and how to write tests.
package bad

import "context"

func f() {
	c, cancel := context.WithCancel(nil)
	defer cancel()
	_ = c
}
