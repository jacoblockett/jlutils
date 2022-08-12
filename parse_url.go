package jlutils

import (
	"regexp"
	"strconv"
	"strings"
)

type QueryToMap func(s string) Query

type Query map[string]string

type URL struct {
	Protocol   string
	Hostname   string
	Subdomain  string
	Domain     string
	Port       int
	Path       string
	Query      string
	Fragment   string
	Input      string
	QueryToMap QueryToMap
}

func ParseURL(s string) URL {
	rx := regexp.MustCompile(`(?:(?P<protocol>[a-z]+):\/\/)?(?P<hostname>(?:\.?[a-z0-9]+)+)(?::(?P<port>\d+))?(?:\/(?P<path>(?:[^\s\"\#\'\<\>\?]+)))?(?:\?(?P<query>[^\s\"\#\'\<\>\?]+))?(?:#(?P<fragment>(?:[^\s\"\#\'\<\>\?]+)?)$)?`)
	match := rx.FindStringSubmatch(s)
	u := URL{}

	for i, name := range rx.SubexpNames() {
		rm := match[i]
		if i != 0 && len(rm) != 0 {
			switch x := name; x {
			case "protocol":
				u.Protocol = rm
			case "hostname":
				u.Hostname = rm
				s, d := parseHostname(rm)
				u.Subdomain = s
				u.Domain = d
			case "port":
				i, _ := strconv.Atoi(rm)
				u.Port = i
			case "path":
				u.Path = rm
			case "query":
				u.Query = rm
			case "fragment":
				u.Fragment = rm
			}
		}
	}

	u.Input = s
	u.QueryToMap = func(s string) (q Query) {
		sp := strings.Split(s, "&")

		for _, p := range sp {
			kv := strings.Split(p, "=")
			m := make(map[string]string)
			m[kv[0]] = kv[1]
		}

		return q
	}

	return u
}

func parseHostname(s string) (subdomain string, domain string) {
	sp := strings.Split(s, ".")

	if len(sp) >= 3 && sp[len(sp)-2] == "co" {
		// .co.*
		domain = sp[len(sp)-3] + "." + sp[len(sp)-2] + "." + sp[len(sp)-1]
		sp = sp[:len(sp)-3]
		subdomain = strings.Join(sp, ".")
	} else if len(sp) >= 3 {
		// .com et al. w/ sub
		domain = sp[len(sp)-2] + "." + sp[len(sp)-1]
		sp = sp[:len(sp)-2]
		subdomain = strings.Join(sp, ".")
	} else {
		// .com et al. w/o sub
		subdomain = ""
		domain = strings.Join(sp, ".")
	}

	return subdomain, domain
}
