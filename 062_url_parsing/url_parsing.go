package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {
	// We’ll parse this example URL, which includes a scheme,
	// authentication info, host, port, path, query params, and query fragment.
	s := "postgres://user:pass@host.com:5432/path?k=v#f"

	// Parse the URL and ensure there are no errors.
	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	// Accessing the scheme is straightforward.
	fmt.Println(u.Scheme)

	// User contains all authentication info; call Username and Password
	// on this for individual values.
	fmt.Println(u.User)
	fmt.Println(u.User.Username())
	p, _ := u.User.Password()
	fmt.Println(p)

	// The Host contains both the hostname and the port, if present.
	// Use SplitHostPort to extract them.
	fmt.Println(u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println(host)
	fmt.Println(port)

	//
	fmt.Println(u.Path)
	fmt.Println(u.Fragment)

	fmt.Println(u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println(m)
	fmt.Println(m["k"][0])
}
