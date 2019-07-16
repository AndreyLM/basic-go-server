package config

// DriverType - driver type
type DriverType int8

const (
	// MySQLType - MySQL driver
	MySQLType DriverType = iota
	// MSSQLType  MSSQL driver
	MSSQLType
	// Postgress - Postgress driver
	Postgress
)

// ConnectionConfig - connection config
type ConnectionConfig struct {
	User       string
	Password   string
	Host       string
	Port       string
	Database   string
	DriverType DriverType
}
