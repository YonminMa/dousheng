package constants

// 数据库配置
const (
	MySQLDefaultDSN = "root:root@tcp(123.60.170.254:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local" //MySQL DSN 	// 架构名
)

// 数据库表名
const (
	VideoTableName = "video"
	UserTableName  = "user"
)

// Jwt
const (
	SecretKey = "secret key"
)
