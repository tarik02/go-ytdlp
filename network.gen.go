// Copyright (c) Liam Stanley <me@liamstanley.io>. All rights reserved. Use
// of this source code is governed by the MIT license that can be found in
// the LICENSE file.
//
// Code generated by cmd/codegen. DO NOT EDIT.
//
// Network Option Group

package ytdlp

import (
	"strconv"
)

// Use the specified HTTP/HTTPS/SOCKS proxy. To enable SOCKS proxy, specify a
// proper scheme, e.g. socks5://user:pass@127.0.0.1:1080/. Pass in an empty string
// (--proxy "") for direct connection
//
// Proxy maps to cli flags: --proxy=URL.
func (c *Command) Proxy(url string) *Command {
	c.addFlag(&Flag{
		ID:   "proxy",
		Flag: "--proxy",
		Args: []string{url},
	})
	return c
}

// UnsetProxy unsets any flags that were previously set by
// Proxy().
func (c *Command) UnsetProxy() *Command {
	c.removeFlagByID("proxy")
	return c
}

// Time to wait before giving up, in seconds
//
// SocketTimeout maps to cli flags: --socket-timeout=SECONDS.
func (c *Command) SocketTimeout(seconds float64) *Command {
	c.addFlag(&Flag{
		ID:   "socket_timeout",
		Flag: "--socket-timeout",
		Args: []string{
			strconv.FormatFloat(seconds, 'g', -1, 64),
		},
	})
	return c
}

// UnsetSocketTimeout unsets any flags that were previously set by
// SocketTimeout().
func (c *Command) UnsetSocketTimeout() *Command {
	c.removeFlagByID("socket_timeout")
	return c
}

// Client-side IP address to bind to
//
// SourceAddress maps to cli flags: --source-address=IP.
func (c *Command) SourceAddress(ip string) *Command {
	c.addFlag(&Flag{
		ID:   "source_address",
		Flag: "--source-address",
		Args: []string{ip},
	})
	return c
}

// UnsetSourceAddress unsets any flags that were previously set by
// SourceAddress().
func (c *Command) UnsetSourceAddress() *Command {
	c.removeFlagByID("source_address")
	return c
}

// Enable file:// URLs. This is disabled by default for security reasons.
//
// EnableFileUrls maps to cli flags: --enable-file-urls.
func (c *Command) EnableFileUrls() *Command {
	c.addFlag(&Flag{
		ID:   "enable_file_urls",
		Flag: "--enable-file-urls",
		Args: nil,
	})
	return c
}
