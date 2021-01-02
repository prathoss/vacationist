package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	DbHost     string
	DbUser     string
	DbPassword string
	DbPort     int
	SslMode    string
)

func init() {
	var badVars []string

	host, isHostSet := os.LookupEnv("DB_HOST")
	if isHostSet {
		DbHost = host
	} else {
		badVars = append(badVars, "DB_HOST is not set")
	}

	user, isUserSet := os.LookupEnv("DB_USER")
	if isUserSet {
		DbUser = user
	} else {
		badVars = append(badVars, "DB_USER is not set")
	}

	passwd, isPasswdSet := os.LookupEnv("DB_PASSWORD")
	if isPasswdSet {
		DbPassword = passwd
	} else {
		badVars = append(badVars, "DB_PASSWORD is not set")
	}

	port, isPortSet := os.LookupEnv("DB_PORT")
	if isPortSet {
		if iPort, err := strconv.Atoi(port); err != nil {
			badVars = append(badVars, "DB_PORT is not integer")
		} else {
			DbPort = iPort
		}
	} else {
		badVars = append(badVars, "DB_PORT is not set")
	}

	sslMode, isSslModeSet := os.LookupEnv("SSL_MODE")
	if isSslModeSet {
		if sslMode == "enable" || sslMode == "disable" {
			SslMode = sslMode
		} else {
			badVars = append(badVars, "SSL_MODE has invalid value")
		}
	} else {
		SslMode = "disbale"
	}

	if len(badVars) != 0 {
		panic(fmt.Errorf("Errors with env variables:\n\t%s", strings.Join(badVars, "\n\t")))
	}
}

func GetDbConnection() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=vacationist sslmode=%s", DbHost, DbPort, DbUser, DbPassword, SslMode)
}
