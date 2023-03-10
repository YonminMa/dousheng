package constants

// 数据库配置
const (
	MySQLDefaultDSN = "root:root@tcp(localhost:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local" //MySQL DSN 	// 架构名
)

// 数据库表名
const (
	VideoTableName    = "video"
	UserTableName     = "user"
	RelationTableName = "relation"
	FavoriteTableName = "favorite"
)

// Jwt
const (
	SecretKey = "secret key"
)
