package cdb

import (
	"fmt"
	"net/url"
	"path"
)

// NewTCPMySQLURLString generate a mysql url from src url
// src url
//   mysql://<user>:<passwd>@<host>/<db>?<query>
// mysql url
//   <user>:<passwd>@tcp(<host>)/<db>?<query>
func NewTCPMySQLURLString(u *url.URL) string {
	password, exists := u.User.Password()
	if !exists {
		password = ""
	}
	return fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?%s",
		u.User.Username(),
		password,
		u.Host,
		path.Base((u.Path)),
		u.RawQuery,
	)
}

func NewLocalhostMySQLURLString(u *url.URL) string {
	password, exists := u.User.Password()
	if !exists {
		password = ""
	}
	return fmt.Sprintf(
		"%s:%s@tcp/%s?%s",
		u.User.Username(),
		password,
		path.Base((u.Path)),
		u.RawQuery,
	)
}
