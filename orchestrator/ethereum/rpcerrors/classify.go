package rpcerrors

import (
	"context"
	"errors"
	"io"
	"net"
	"strings"
)

type Kind int

const (
	Unknown Kind = iota
	RateLimit
	RangeTooLarge
	Timeout
	Network
	UnknownBlock
	NodeIssue
	HistoryPruned
	Fatal
)

func (k Kind) String() string {
	switch k {
	case RateLimit:
		return "rate_limit"
	case RangeTooLarge:
		return "range_too_large"
	case Timeout:
		return "timeout"
	case Network:
		return "network"
	case UnknownBlock:
		return "unknown_block"
	case NodeIssue:
		return "node_issue"
	case HistoryPruned:
		return "history_pruned"
	case Fatal:
		return "fatal"
	default:
		return "unknown"
	}
}

var (
	rateLimitMarkers = []string{
		"rate limit",
		"rate-limit",
		"too many requests",
		"429",
		"throttl",
		"quota",
		"credits",
		"max query rate",
		"exceeded the quota",
		"request limit",
		"capacity",
	}

	rangeMarkers = []string{
		"limit exceeded",
		"block range",
		"range is too",
		"response size",
		"response is too large",
		"exceed maximum",
		"query returned more than",
		"more than 10000 results",
		"logs matched by query exceeds",
		"returned more than",
		"too many logs",
		"getlogs is limited",
		"log response size",
		"max results",
		"requested too many",
		"batch size is too",
		"payload too large",
	}

	timeoutMarkers = []string{
		"timeout",
		"deadline exceeded",
		"timed out",
		"i/o timeout",
	}

	networkMarkers = []string{
		"eof",
		"connection reset",
		"connection refused",
		"broken pipe",
		"no such host",
		"tls handshake",
		"bad gateway",
		"gateway timeout",
		"service unavailable",
		"502",
		"503",
		"504",
		"no contract code at given address",
		"websocket",
		"dial tcp",
	}

	unknownBlockMarkers = []string{
		"unknown block",
		"one of the blocks specified in filter",
		"header not found",
		"block not found",
	}

	historyPrunedMarkers = []string{
		"history has been pruned",
		"missing trie node",
		"historical state not available",
		"archive node",
		"pruned state",
		"state not found",
		"requested block is too old",
		"block is pruned",
	}

	fatalMarkers = []string{
		"no rpc clients available",
		"invalid sender",
		"insufficient funds",
	}
)

// Classify inspects an error string and returns the most specific Kind it
// matches. Network providers rarely expose stable error codes, so we rely on
// substring matching on the error text. Order matters: range limits often come
// back as generic "limit exceeded" which would otherwise be swallowed by the
// rate-limit bucket.
func Classify(err error) Kind {
	if err == nil {
		return Unknown
	}
	if errors.Is(err, context.DeadlineExceeded) {
		return Timeout
	}
	if errors.Is(err, context.Canceled) {
		return Fatal
	}
	if errors.Is(err, io.EOF) || errors.Is(err, io.ErrUnexpectedEOF) {
		return Network
	}
	var netErr net.Error
	if errors.As(err, &netErr) && netErr.Timeout() {
		return Timeout
	}

	msg := strings.ToLower(err.Error())

	if containsAny(msg, fatalMarkers) {
		return Fatal
	}
	if containsAny(msg, historyPrunedMarkers) {
		return HistoryPruned
	}
	if containsAny(msg, unknownBlockMarkers) {
		return UnknownBlock
	}
	if containsAny(msg, rangeMarkers) {
		return RangeTooLarge
	}
	if containsAny(msg, rateLimitMarkers) {
		return RateLimit
	}
	if containsAny(msg, timeoutMarkers) {
		return Timeout
	}
	if containsAny(msg, networkMarkers) {
		return Network
	}
	return Unknown
}

// IsRetryable reports whether the caller should retry (on the same or another
// RPC). Non-retryable kinds require the caller to take a different action
// (range resize, abort).
func IsRetryable(k Kind) bool {
	switch k {
	case RateLimit, Timeout, Network, NodeIssue, HistoryPruned, Unknown:
		return true
	default:
		return false
	}
}

// ShouldRotate reports whether the current RPC is likely the problem and the
// caller should rotate to a different one before retrying.
func ShouldRotate(k Kind) bool {
	switch k {
	case RateLimit, Timeout, Network, NodeIssue, HistoryPruned:
		return true
	default:
		return false
	}
}

func containsAny(s string, markers []string) bool {
	for _, m := range markers {
		if strings.Contains(s, m) {
			return true
		}
	}
	return false
}
