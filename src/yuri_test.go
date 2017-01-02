package main

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateURIMapWithNoUsernameAndPassword(t *testing.T) {
	em := map[string]string{
		"fragment": "FRAG",
		"host":     "stage.example.com:443",
		"opaque":   "",
		"password": "",
		"path":     "/path+to+foo",
		"rawpath":  "/path+to+foo",
		"rawquery": "query1=1&query2=2",
		"scheme":   "https",
		"username": "",
	}
	parsedURI, _ := url.Parse("https://stage.example.com:443/path+to+foo?query1=1&query2=2#FRAG")
	rm := CreateURIMap(parsedURI)
	assert.Equal(t, rm, em)
}

func TestCreateURIMapWithUsername(t *testing.T) {
	em := map[string]string{
		"fragment": "FRAG",
		"host":     "stage.example.com:443",
		"opaque":   "",
		"password": "",
		"path":     "/path+to+foo",
		"rawpath":  "/path+to+foo",
		"rawquery": "query1=1&query2=2",
		"scheme":   "https",
		"username": "username",
	}
	parsedURI, _ := url.Parse("https://username@stage.example.com:443/path+to+foo?query1=1&query2=2#FRAG")
	rm := CreateURIMap(parsedURI)
	assert.Equal(t, rm, em)
}

func TestCreateURIMapWithUsernameAndPassword(t *testing.T) {
	em := map[string]string{
		"fragment": "FRAG",
		"host":     "stage.example.com:443",
		"opaque":   "",
		"password": "password",
		"path":     "/path+to+foo",
		"rawpath":  "/path+to+foo",
		"rawquery": "query1=1&query2=2",
		"scheme":   "https",
		"username": "username",
	}
	parsedURI, _ := url.Parse("https://username:password@stage.example.com:443/path+to+foo?query1=1&query2=2#FRAG")
	rm := CreateURIMap(parsedURI)
	assert.Equal(t, rm, em)
}
