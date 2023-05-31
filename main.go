package main

import (
	"crud/modules/users"
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `json:"name"`
}

var db *gorm.DB
var err error

func initDB() (*gorm.DB, error) {
	dsn := "root:@tcp(localhost:3306)/crud?parseTime=true"
	return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}

func getUserById(c *gin.Context) {
	var user User
	userID := c.Param("id")

	// Dapatkan data user dari database berdasarkan ID
	if err := db.First(&user, userID).Error; err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Tampilkan data user
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func updateUser(c *gin.Context) {
	var user User
	userID := c.Param("id")

	// Dapatkan data user dari database berdasarkan ID
	if err := db.First(&user, userID).Error; err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Baca data JSON dari body permintaan
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simpan perubahan ke database
	if err := db.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Tampilkan respons berhasil
	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "user": user})
}

func deleteUser(c *gin.Context) {
	var user User
	userID := c.Param("id")

	// Dapatkan data user dari database berdasarkan ID
	if err := db.First(&user, userID).Error; err != nil {
		if errors.Is(gorm.ErrRecordNotFound, err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Hapus data user dari database
	if err := db.Delete(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Tampilkan respons berhasil
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}

func main() {
	db, err := initDB()
	if err != nil {
		log.Fatalln("initDB:", err)
	}

	r := gin.Default()
	usersHandler := users.DefaultRequestHandler(db)

	r.POST("/users", usersHandler.Create)
	r.GET("/users", usersHandler.Read)

	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}

	// request handler: menerima request, mengirim response
	// controller: validasi dan transformasi data
	// use case: pemrosesan data
	// repository: persistensi data
}
