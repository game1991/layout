package env

import "os"

var (
	zonst_env string
)

func init() {
	zonst_env = os.Getenv("ZONST_ENV")
}

func IsDevelopment() bool {
	return GetZonstEnv() == "development"
}

func GetZonstEnv() string {
	return zonst_env
}
