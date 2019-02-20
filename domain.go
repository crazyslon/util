package util

import (
	"net/url"
	"regexp"
	"strings"

	"github.com/bobesa/go-domain-util/domainutil"
	"golang.org/x/net/idna"
)

var domainRegexp = regexp.MustCompile("^[0-9\\p{L}-\\.]{0,61}[0-9\\p{L}]\\.[0-9\\p{L}][\\p{L}-]*[0-9\\p{L}]+$")

//GetDomainFromURL return domain from url if it valid,
//in other case return empty string
func GetDomainFromURL(rawurl string) string {

	uri, err := url.ParseRequestURI(rawurl)
	if err != nil {
		return ""
	}

	if uri.Scheme != "http" && uri.Scheme != "https" {
		return ""
	}

	return GetDomain(uri.Hostname())
}

//GetDomain return domain if it valid and return empty string in other case.
func GetDomain(domain string) string {
	if strings.Contains(domain, "_") {
		return ""
	}

	if strings.HasPrefix(domain, "-") {
		return ""
	}

	domain, err := convertToUnicode(domain)
	if err != nil {
		return ""
	}

	if !domainRegexp.MatchString(domain) {
		return ""
	}

	suffix := domainutil.DomainSuffix(domain)
	if len(suffix) == 0 {
		return ""
	}

	return domain
}

func convertToUnicode(domain string) (string, error) {
	if strings.Contains(domain, "xn--") {
		var err error
		domain, err = idna.ToUnicode(domain)
		if err != nil {
			return "", err
		}
	}
	return domain, nil
}

//IsValidDomain check is valid domain or not
func IsValidDomain(domainName string) bool {
	return len(GetDomain(domainName)) > 0
}

//IsValidURL check is string valid url or not
func IsValidURL(rawurl string) bool {

	return len(GetDomainFromURL(rawurl)) > 0
}
