package configs

import (
	"database/sql"
	"log"

	"github.com/soulwill1/estudos/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Init() {
	var err error
	dsn := "root:abc123@tcp(127.0.0.1:3306)/"

	// conecte-sa ao MySQL sem especificar um banco de dados
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed to connect to MySQL: %v", err)
	}
	defer sqlDB.Close()

	// Verifica se o banco de dados existe e cria se necessário
	_, err = sqlDB.Exec("CREATE DATABASE IF NOT EXISTS polls CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;")
	if err != nil {
		log.Fatalf("failed to create database: %v", err)
	}

	// conecta ao banco de dados polls
	dsn = "root:abc123@tcp(127.0.0.1:3306)/polls?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Executa a migração
	db.AutoMigrate(&models.Poll{}, &models.PollOption{})
}


func GetMySQL() *gorm.DB {
	return db
}
