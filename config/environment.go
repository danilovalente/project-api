package config

import (
	"github.com/spf13/viper"
)

//Values stores the current configuration values
var Values Config

//Config contains the application's configuration values. Add here your own variables and bind it on init() function
type Config struct {
//DBConnectionString to connect to Mongo
DBConnectionString string
//DBConnectionCertificateFileName defines the TLS Certificate for DB Connections. If not set, no TLS is configured
DBConnectionCertificateFileName string
	//Port contains the port in which the application listens
	Port string
	//AppName for displaying in Monitoring
	AppName string
	//LogLevel - DEBUG or INFO or WARNING or ERROR or PANIC or FATAL
	LogLevel string
	//TestRun state if the current execution is a test execution
	TestRun bool
	//UsePrometheus to enable prometheus metrics endpoint
	UsePrometheus bool
}

func init() {
_ = viper.BindEnv("DBConnectionString", "DB_CONNECTION_STRING")
_ = viper.BindEnv("DBConnectionCertificateFileName", "DB_CONNECTION_CERTIFICATE_FILE_NAME")
	_ = viper.BindEnv("TestRun", "TESTRUN")
	viper.SetDefault("TestRun", false)
	_ = viper.BindEnv("UsePrometheus", "USEPROMETHEUS")
	viper.SetDefault("UsePrometheus", false)
	_ = viper.BindEnv("Port", "PORT")
	viper.SetDefault("Port", "8080")
	_ = viper.BindEnv("AppName", "APP_NAME")
	viper.SetDefault("AppName", "project-api")
	_ = viper.BindEnv("LogLevel", "LOG_LEVEL")
	viper.SetDefault("LogLevel", "INFO")
	_ = viper.Unmarshal(&Values)
}
