package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkGetDomainFromURL(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetDomainFromURL("https://spark-public.s3.amazonaws.com/dataanalysis/loansData.csv")
	}
}

func BenchmarkGetDomain(b *testing.B) {
	for n := 0; n < b.N; n++ {
		GetDomain("test.domain.org")
	}
}

func TestGetDomainFromURL(t *testing.T) {

	assert.Equal(t,
		"",
		GetDomainFromURL("http://http://google.com/?test=1"),
	)

	require.Equal(t,
		"spark-public.s3.amazonaws.com",
		GetDomainFromURL("https://spark-public.s3.amazonaws.com/dataanalysis/loansData.csv"),
	)

	assert.Equal(t,
		"",
		GetDomainFromURL("gopher://domain.unknown/"),
	)

	assert.Equal(t,
		"",
		GetDomainFromURL("https://192.168.0.0"),
	)

	assert.Equal(t,
		"google.com",
		GetDomainFromURL("http://google.com"),
	)

	assert.Equal(t,
		"",
		GetDomainFromURL("http://google.local"),
	)

	assert.Equal(t,
		"",
		GetDomainFromURL("com"),
	)

	assert.Equal(t,
		"amazon.fancy.uk",
		GetDomainFromURL("http://amazon.fancy.uk/test_123?p1=1&P_2=2"),
	)

	assert.Equal(t,
		"example.co.uk",
		GetDomainFromURL("https://user:password@example.co.uk:8080/some/path?and&query#hash"),
	)

	assert.Equal(t,
		"",
		GetDomainFromURL("http://ama_zon.fancy.uk/test123"),
	)

	assert.Equal(t,
		"кто.рф",
		GetDomainFromURL("https://кто.рф"),
	)

	assert.Equal(t,
		"кто.рф",
		GetDomainFromURL("http://xn--j1ail.xn--p1ai/"),
	)
}

func TestGetDomain(t *testing.T) {

	assert.Equal(t, "mk.com", GetDomain("mk.com"))
	assert.Equal(t, "m.com", GetDomain("m.com"))
	assert.Equal(t, "google.com", GetDomain("google.com"))
	assert.Equal(t, "test-domain.ru", GetDomain("test-domain.ru"))
	assert.Equal(t, "test.domain.org", GetDomain("test.domain.org"))
	assert.Equal(t, "кто.рф", GetDomain("кто.рф"))
	assert.Equal(t, "кто.рф", GetDomain("xn--j1ail.xn--p1ai"))
	assert.Equal(
		t,
		"中国互联网络信息中心.中国",
		GetDomain("中国互联网络信息中心.中国"),
	)
	assert.Equal(
		t,
		"中国互联网络信息中心.中国",
		GetDomain("xn--fiqa61au8b7zsevnm8ak20mc4a87e.xn--fiqs8s"),
	)

	invalidDomains := []string{
		//invalid tld
		"google.test", "google.unknown", "google.s1",
		"test", "test.", "mkyong.t.t.c", ".com", "com",
		"test,com",

		//contains dot at the end
		"google.com.",

		//url or with params
		"http://google.com", "google.com/?param=1",
		"google.com/?param=1&otherparam=2",

		//with underscore
		"test.go_ogle.com",

		//hyphen -
		"-test.google.com", "test.google.com-",
		"sub.mkyong-.com",
		//"sub.-mkyong.com", todo

		//invalid characters
		"test.goog!le.com", "test.goog*le.com", "test.goog|le.com",
	}

	for _, d := range invalidDomains {
		assert.Empty(t, GetDomain(d))
	}
}
