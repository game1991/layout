package env

import "os"

var (
	zonst_env string
)

func init() {
	zonst_env = os.Getenv("DEVELOP_ENV")
}

// IsDevelopment judge env development
func IsDevelopment() bool {
	return GetZonstEnv() == "development"
}

// GetZonstEnv ...
func GetZonstEnv() string {
	return zonst_env
}
