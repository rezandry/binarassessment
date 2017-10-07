package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func InitDb() *gorm.DB {
	var err error
	db, err := gorm.Open("postgres", "host=127.0.0.1 user=postgres dbname=profileapi sslmode=disable password=profileapi")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Connection Success!")
	}
	return db
}

type Profile struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Token string `json:"token"`
}

func main() {
	migrate()
	r := gin.Default()
	r.POST("/profile", CreateProfile)
	r.Use(Middleware)
	r.GET("/profile/:id", ReadProfile)
	r.GET("/profile", ReadProfiles)
	r.PUT("/profile/:id", UpdateProfile)
	r.DELETE("/profile/:id", DeleteProfile)
	r.Run(":8080")
}

func migrate() {
	var profile Profile
	db := InitDb()
	defer db.Close()
	db.AutoMigrate(&profile)
}

func ReadProfile(c *gin.Context) {
	var profile Profile
	db := InitDb()
	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&profile).Error; err != nil {
		fmt.Println(err)
	}
	profile.Token = ""
	c.JSON(200, profile)
}

func ReadProfiles(c *gin.Context) {
	var profile []Profile
	db := InitDb()
	defer db.Close()
	if err := db.Find(&profile).Error; err != nil {
		fmt.Println(err)
	}
	for _, p := range profile {
		p.Token = ""
	}
	c.JSON(200, profile)
}

func CreateProfile(c *gin.Context) {
	var profile Profile
	db := InitDb()
	defer db.Close()
	c.BindJSON(&profile)
	profile.Token = GenerateToken("Ini adalah combination string yang akan dibuat token serta harusnya ditambahi dengan salt, bisa diambil dari combinasi data nama atau apapun")
	db.Create(&profile)
	c.JSON(200, profile)
}

func UpdateProfile(c *gin.Context) {
	var profile Profile
	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).First(&profile).Error; err != nil {
		fmt.Println(err)
	}
	c.BindJSON(&profile)
	db.Save(&profile)
	profile.Token = ""
	c.JSON(200, profile)
}

func DeleteProfile(c *gin.Context) {
	var profile Profile
	db := InitDb()
	defer db.Close()

	id := c.Params.ByName("id")
	if err := db.Where("id = ?", id).Delete(&profile).Error; err != nil {
		fmt.Println(err)
	}
	c.JSON(200, gin.H{"Account ID : " + id: " deleted."})
}

func GenerateToken(text string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(text))
	h := sha256.New()
	h.Write([]byte(encoded))
	var sha = h.Sum(nil)
	var token = base64.StdEncoding.EncodeToString([]byte(sha))
	return token
}

func Middleware(c *gin.Context) {
	var profile Profile
	db := InitDb()
	defer db.Close()

	token := c.Request.Header.Get("token")

	if token == "" {
		data := map[string]interface{}{
			"Message": "Token Required",
		}
		c.AbortWithStatusJSON(400, data)
		return
	}
	if err := db.Where("token = ?", token).First(&profile).Error; err != nil {
		data := map[string]interface{}{
			"Message": "You have no authorization",
		}
		c.AbortWithStatusJSON(400, data)
		return
	}
	c.Next()
}
