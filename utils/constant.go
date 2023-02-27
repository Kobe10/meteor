package utils

import (
	"meteor/config/env"
)

const (
	ConfigEnv  = "GVA_CONFIG"
	ConfigFile = "config"
)

const configLocal = "config-local"
const configDev = "config-dev"
const configFeature = "config-feature"
const configPreprod = "config-preprod"
const configProd = "config-prod"

var EnvConfig string

func GetEnvConfig(env_ string) string {
	if EnvConfig == "" {
		setEnvConfig(env_)
	}
	return EnvConfig
}

func setEnvConfig(env_ string) {
	switch env_ {
	case env.ZEnvLocal:
		EnvConfig = configLocal
	case env.ZEnvDev:
		EnvConfig = configDev
	case env.ZEnvFeature:
		EnvConfig = configFeature
	case env.ZEnvPreprod:
		EnvConfig = configPreprod
	case env.ZEnvProd:
		EnvConfig = configProd
	default:
		EnvConfig = configProd
	}
}
