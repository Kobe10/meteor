package utils

import (
	"meteor/config/env"
	"meteor/httptrack"
)

const domainDev = "http://127.0.0.1:8888/url/"
const domainFeature = "http://meteor.kt.ziroom.com/url/"
const domainPreprod = "http://meteor.kq.ziroom.com/url/"
const domainProd = "http://meteor.kp.ziroom.com/url/"

var urlDomain string

func getUrlDomain() string {
	if urlDomain == "" {
		setUrlDomain()
	}
	return urlDomain
}

func setUrlDomain() {
	var httpEnv = httptrack.GetEnv()
	switch httpEnv {
	case env.ZEnvLocal:
		urlDomain = domainDev
	case env.ZEnvDev:
		urlDomain = domainDev
	case env.ZEnvFeature:
		urlDomain = domainFeature
	case env.ZEnvPreprod:
		urlDomain = domainPreprod
	case env.ZEnvProd:
		urlDomain = domainProd
	default:
		urlDomain = domainProd
	}
}
