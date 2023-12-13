package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

// Model para o exemplo de usuário
type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

var (
	db  *gorm.DB
	rdb *redis.Client
)

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	// Configuração do banco de dados
	dbURL := os.Getenv("DB_URL")
	db, err = gorm.Open("postgres", dbURL)
	if err != nil {
		panic("Failed to connect to database")
	}

	// Configuração do cache Redis
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // sem senha
		DB:       0,  // usar o banco de dados padrão
	})
}

func main() {
	defer db.Close()

	// AutoMigrate cria automaticamente as tabelas no banco de dados
	db.AutoMigrate(&User{})

	// Configuração do servidor Gin
	router := gin.Default()

	router.GET("/users", GetUsers)
	router.GET("/users/:id", GetUser)
	router.POST("/users", CreateUser)
	router.PUT("/users/:id", UpdateUser)
	router.DELETE("/users/:id", DeleteUser)

	router.Run(":8080")
}

// GetUsers busca todos os usuários
func GetUsers(c *gin.Context) {
	var users []User
	db.Find(&users)
	c.JSON(200, users)
}

// GetUser busca um usuário pelo ID
// GetUser busca um usuário pelo ID
func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")

	// Tentar obter o usuário do cache
	cacheKey := fmt.Sprintf("user:%s", id)
	cacheResult, err := rdb.Get(c.Request.Context(), cacheKey).Result()
	if err == nil {
		// Se o usuário estiver no cache, retornar diretamente do cache
		var cachedUser User
		if err := json.Unmarshal([]byte(cacheResult), &cachedUser); err == nil {
			c.JSON(http.StatusOK, cachedUser)
			return
		}
	}

	// Se o usuário não estiver no cache, consultá-lo no banco de dados
	var user User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		fmt.Println(err)
	} else {
		// Armazenar o resultado no cache para futuras solicitações
		userJSON, _ := json.Marshal(user)
		rdb.Set(c.Request.Context(), cacheKey, userJSON, 24*time.Hour) // Cache válido por 24 horas

		c.JSON(http.StatusOK, user)
	}
}

// CreateUser cria um novo usuário
func CreateUser(c *gin.Context) {
	var user User
	c.BindJSON(&user)

	db.Create(&user)
	c.JSON(200, user)
}

// UpdateUser atualiza um usuário pelo ID
func UpdateUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User
	if err := db.Where("id = ?", id).First(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&user)
	db.Save(&user)
	c.JSON(200, user)
}

// DeleteUser exclui um usuário pelo ID
func DeleteUser(c *gin.Context) {
	id := c.Params.ByName("id")
	var user User
	d := db.Where("id = ?", id).Delete(&user)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}
