package sensitive_word

import (
	"meteor/config/env"
	"meteor/httptrack"
)

const paladinHostDev = "http://eunomia.kt.ziroom.com"
const paladinHostFeature = "http://eunomia.kt.ziroom.com"
const paladinHostPreprod = "http://eunomia.kt.ziroom.com"
const paladinHostProd = "http://eunomia.kp.ziroom.com"

const AccessKeyDev = "8a76951575d51f200175d51f20490001#8a76951575d51f200175d5209d720003"
const AccessKeyFeature = "8a76951575d51f200175d51f20490001#8a76951575d51f200175d5209d720003"
const AccessKeyPreprod = "8a76951575d51f200175d51f20490001#8a76951575d51f200175d5209d720003"
const AccessKeyProd = "8a90b7a4743506b80175d0060d8455f9#8a90b7a4743506b80175d01da00556f2"

var hostName string
var accessKey string

func GetHost() string {
	if hostName == "" {
		setHost()
	}
	return hostName
}

func GetAccessKey() string {
	if accessKey == "" {
		setAccessKey()
	}
	return accessKey
}

func setHost() {
	var httpEnv = httptrack.GetEnv()
	switch httpEnv {
	case env.ZEnvLocal:
		hostName = paladinHostDev
	case env.ZEnvDev:
		hostName = paladinHostDev
	case env.ZEnvFeature:
		hostName = paladinHostFeature
	case env.ZEnvPreprod:
		hostName = paladinHostPreprod
	case env.ZEnvProd:
		hostName = paladinHostProd
	default:
		hostName = paladinHostProd
	}
}

func setAccessKey() {
	var httpEnv = httptrack.GetEnv()
	switch httpEnv {
	case env.ZEnvLocal:
		accessKey = AccessKeyDev
	case env.ZEnvDev:
		accessKey = AccessKeyDev
	case env.ZEnvFeature:
		accessKey = AccessKeyFeature
	case env.ZEnvPreprod:
		accessKey = AccessKeyPreprod
	case env.ZEnvProd:
		accessKey = AccessKeyProd
	default:
		accessKey = AccessKeyProd
	}
}
