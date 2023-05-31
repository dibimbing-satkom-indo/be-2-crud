package main

import (
	"errors"
	"fmt"
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

func initDB() {
	dsn := "root:@tcp(localhost:3306)/crud?parseTime=true"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}

func createUser(c *gin.Context) {
	var user User

	// Baca data JSON dari body permintaan
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Simpan data ke database
	err := db.Create(&user).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	fmt.Println("inserted:", user.ID)

	// Tampilkan respons berhasil
	c.JSON(http.StatusOK, gin.H{"message": "User created successfully", "user": user})
}

func getUsers(c *gin.Context) {
	var users []User

	// Dapatkan semua data user dari database
	if err := db.Find(&users).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Tampilkan data user
	c.JSON(http.StatusOK, gin.H{"users": users})
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

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/users", createUser)
	r.GET("/users", getUsers)
	r.GET("/users/:id", getUserById)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)

	return r
}

func main() {
	initDB()
	r := setupRouter()

	// Jalankan server di port 8080
	err = r.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
