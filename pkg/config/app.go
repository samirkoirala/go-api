package config

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	// Configure Viper
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AddConfigPath("../..")

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: No config file found: %v", err)
	}

	viper.AutomaticEnv()

	// dsn := "user:userpassword@tcp(127.0.0.1:3306)/simplerest?charset=utf8mb4&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("DB_USER"),
		viper.GetString("DB_PASSWORD"),
		viper.GetString("DB_HOST"),
		viper.GetString("DB_PORT"),
		viper.GetString("DB_NAME"),
	)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to the database: %v", err))
	}
	db = d
	fmt.Println("Database connection established")
}

func GetDB() *gorm.DB {
	return db
}
