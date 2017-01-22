// Package examples contains example templates.
// They are being used for development.
// Long term plan is to have a builtin package,
// with subdirs corresponding to other packages,
// e.g. builtin/os, builtin/context, and builtin/net/http,
// which should probably also suffice for testing and documentation purposes.
package examples

import (
	"bufio"
	"bytes"
	"context"
	"io"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"testing"

	"github.com/josharian/dont/match" // it says dont match, but it means match. oh the irony.
)

// All usages of context.WithCancel must call cancel on all exit paths.
func DontLoseCancelContext(parent context.Context) {
	_, cancel := context.WithCancel(parent)
	match.PossiblyUnused(cancel)
}

// TODO: variants with other context constructors
func DontCallWithCancelNil() {
	context.WithCancel(nil)
}

// All uses of bufio.Scanner should check Err after scanning,
// except for types that are known to never return errors.
func DontForgetToCheckScannerErr(r io.Reader) {
	match.NotType(r, strings.NewReader(""))
	match.NotType(r, bytes.NewReader(nil))
	s := bufio.NewScanner(r)
	for s.Scan() {
	}
	match.Unused(s.Err)
}

// If err == nil, then resp will be nil, and resp.Body.Close will panic.
// TODO: same for opening files, etc.
// https://github.com/golang/go/issues/17780
func DontCloseRespBodyOnError(url string) {
	resp, err := http.Get(url)
	defer resp.Body.Close()
	if err != nil {
		return
	}
}

// describe, find go issue
func DontRaceOnParallelTestData(t *testing.T, names []string) {
	for _, name := range names {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			match.Used(name)
		})
	}
}

// TODO: variants sync.WaitGroup
// TODO: match just parts of functions.
// TODO: also match go func(wg) ...
// TODO: ref go issue
func DontCallWGAddInsideGoFuncLiteral(wg *sync.WaitGroup) {
	go func() {
		match.NoInterveningStatements()
		wg.Add(1)
	}()
}

// borrowed from staticcheck
func DontCallFindAllWithZeroResults(r *regexp.Regexp, b []byte) {
	r.FindAll(b, 0)
}
