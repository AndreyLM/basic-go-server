package db

import "errors"

// DriverType - driver type
type DriverType string

const (
	// MySQLType - MySQL driver
	MySQLType DriverType = "mysql"
	// MSSQLType  MSSQL driver
	MSSQLType DriverType = "mssql"
	// Postgress - Postgress driver
	Postgress DriverType = "postgress"
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

// GetDriverType - gets driver type by string
func GetDriverType(s string) (DriverType, error) {
	switch s {
	case "mysql":
		return MySQLType, nil
	case "mssql":
		return MySQLType, nil
	case "postgress":
		return Postgress, nil
	default:
		return "", errors.New("Cannot find driver type for such string: " + s)
	}
}
