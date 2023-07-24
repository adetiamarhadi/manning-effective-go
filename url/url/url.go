package url

import (
	"errors"
	"fmt"
	"strings"
)

type URL struct {
	Scheme string
	Host   string
	Path   string
}

// Parse parses rawurl into a URL reference
func Parse(rawurl string) (*URL, error) {
	scheme, rest, ok := parseScheme(rawurl)
	if !ok {
		return nil, errors.New("missing scheme")
	}
	host, path := parseHostPath(rest)
	return &URL{scheme, host, path}, nil
}

func parseScheme(rawurl string) (scheme, rest string, ok bool) {
	return split(rawurl, "://", 1)
}

func parseHostPath(hostpath string) (host, path string) {
	host, path, ok := split(hostpath, "/", 0)
	if !ok {
		host = hostpath
	}
	return host, path
}

func split(s, sep string, n int) (a, b string, ok bool) {
	i := strings.Index(s, sep)
	if i < n {
		return "", "", false
	}
	return s[:i], s[i+len(sep):], true
}

func (u *URL) Port() string {
	_, port, _ := split(u.Host, ":", 0)
	return port
}

func (u *URL) Hostname() string {
	host, _, ok := split(u.Host, ":", 0)
	if !ok {
		host = u.Host
	}
	return host
}

func (u *URL) String() string {
	if u == nil {
		return ""
	}
	var s strings.Builder
	if sc := u.Scheme; sc != "" {
		s.WriteString(sc)
		s.WriteString("://")
	}
	if h := u.Host; h != "" {
		s.WriteString(h)
	}
	if p := u.Path; p != "" {
		s.WriteByte('/')
		s.WriteString(p)
	}
	return s.String()
}

func (u *URL) testString() string {
	return fmt.Sprintf("scheme=%q, host=%q, path=%q", u.Scheme, u.Host, u.Path)
}
