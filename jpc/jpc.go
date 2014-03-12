//
// gocommon - Go library to interact with the JoyentCloud
//
//
// Copyright (c) 2013 Joyent Inc.
//
// Written by Daniele Stroppa <daniele.stroppa@joyent.com>
//

package jpc

import (
	"fmt"
	"net/url"
	"os"
	"reflect"
	"runtime"
	"strings"

	"github.com/joyent/gosign/auth"
)

const (
	// Environment variables
	SdcAccount = "SDC_ACCOUNT"
	SdcKeyId   = "SDC_KEY_ID"
	SdcUrl     = "SDC_URL"
	MantaUser  = "MANTA_USER"
	MantaKeyId = "MANTA_KEY_ID"
	MantaUrl   = "MANTA_URL"
)

var Locations = map[string]string{
	"us-east-1": "America/New_York",
	"us-west-1": "America/Los_Angeles",
	"us-sw-1":   "America/Los_Angeles",
	"eu-ams-1":  "Europe/Amsterdam",
}

// getConfig returns the value of the first available environment
// variable, among the given ones.
func getConfig(envVars ...string) (value string) {
	value = ""
	for _, v := range envVars {
		value = os.Getenv(v)
		if value != "" {
			break
		}
	}
	return
}

// getUserHome returns the value of HOME environment
// variable for the user environment.
func getUserHome() string {
	if runtime.GOOS == "windows" {
		return os.Getenv("APPDATA")
	} else {
		return os.Getenv("HOME")
	}
}

// CredentialsFromEnv creates and initializes the credentials from the
// environment variables.
func CredentialsFromEnv(key string) *auth.Credentials {
	var keyName string
	if key == "" {
		keyName = getUserHome() + "/.ssh/id_rsa"
	} else {
		keyName = key
	}
	authentication := auth.Auth{User: getConfig(SdcAccount, MantaUser), KeyFile: keyName, Algorithm: "rsa-sha256"}

	return &auth.Credentials{
		UserAuthentication: authentication,
		SdcKeyId:           getConfig(SdcKeyId),
		SdcEndpoint:        auth.Endpoint{URL: getConfig(SdcUrl)},
		MantaKeyId:         getConfig(MantaKeyId),
		MantaEndpoint:      auth.Endpoint{URL: getConfig(MantaUrl)},
	}
}

// CompleteCredentialsFromEnv gets and verifies all the required
// authentication parameters have values in the environment.
func CompleteCredentialsFromEnv(keyName string) (cred *auth.Credentials, err error) {
	cred = CredentialsFromEnv(keyName)
	v := reflect.ValueOf(cred).Elem()
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		f := v.Field(i)
		if f.String() == "" {
			return nil, fmt.Errorf("Required environment variable not set for credentials attribute: %s", t.Field(i).Name)
		}
	}
	return cred, nil
}

func Region(cred *auth.Credentials) string {
	sdcUrl := cred.SdcEndpoint.URL

	if isLocalhost(sdcUrl) {
		return "some-region"
	}
	return sdcUrl[strings.LastIndex(sdcUrl, "/")+1 : strings.Index(sdcUrl, ".")]
}

func isLocalhost(u string) bool {
	parsedUrl, err := url.Parse(u)
	if err != nil {
		return false
	}
	if strings.HasPrefix(parsedUrl.Host, "localhost") || strings.HasPrefix(parsedUrl.Host, "127.0.0.1") {
		return true
	}

	return false
}
