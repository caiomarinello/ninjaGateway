package main

import (
	"encoding/gob"
	"log"
	"net/http"
	"os"

	"github.com/caiomarinello/ninjaGateway/components"
	"github.com/caiomarinello/ninjaGateway/db"
	hdl "github.com/caiomarinello/ninjaGateway/handlers"
	mdw "github.com/caiomarinello/ninjaGateway/middleware"
	repos "github.com/caiomarinello/ninjaGateway/repositories"
	"github.com/caiomarinello/ninjaGateway/utils"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/srinathgs/mysqlstore"
)

var mysqlSessionStore *mysqlstore.MySQLStore

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var secretKey = []byte(os.Getenv("SECRET_KEY"))
	var tableName = "sessions"
	var dbConnString = os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWD") + "@tcp(127.0.0.1:3306)/" + os.Getenv("DB_NAME") + "?parseTime=true&loc=Local"
	mysqlSessionStore, err = mysqlstore.NewMySQLStore(dbConnString, tableName, "/", 3600, secretKey)
	if err != nil {
		log.Println(err)
	}
	mysqlSessionStore.Options.HttpOnly = true
	mysqlSessionStore.Options.SameSite = http.SameSiteNoneMode
	mysqlSessionStore.Options.Secure = false // FALSE ONLY FOR DEVELOPMENT. CHANGE IT TO TRUE OTHERWISE
	mysqlSessionStore.Options.MaxAge = 30 * 24 * 60 * 60 //1 month
	mysqlSessionStore.Options.Path = "/"
	gob.Register(components.User{})
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := gin.Default()

	dbConn := db.OpenSqlConnection()
	defer dbConn.Close()

	userAuthorized := r.Group("/", mdw.AuthenticateSession(mysqlSessionStore))
	admin := r.Group("/", mdw.AuthenticateAdminSession(mysqlSessionStore))

	r.POST("/login", hdl.HandleLogin(mysqlSessionStore, repos.NewUserRepository(dbConn))) 
	r.POST("/register", hdl.HandleRegisterUser(repos.NewUserRepository(dbConn)))
	r.POST("/logout", hdl.HandleLogout(mysqlSessionStore))

	userAuthorized.GET("/products", func(c *gin.Context) {
		utils.ForwardRequest(c, os.Getenv("UPSTREAM_URL") + "/products")
	})
	userAuthorized.GET("/product/:productId", func(c *gin.Context) {
		utils.ForwardRequest(c, os.Getenv("UPSTREAM_URL") + "/product/" + c.Param("productId"))
	})
	admin.POST("/product", func(c *gin.Context) {
		utils.ForwardRequest(c, os.Getenv("UPSTREAM_URL") + "/product")
	})
	admin.PUT("/product/:productId", func(c *gin.Context) {
		utils.ForwardRequest(c, os.Getenv("UPSTREAM_URL") + "/product/" + c.Param("productId"))
	})
	admin.DELETE("/product/:productId", func(c *gin.Context) {
		utils.ForwardRequest(c, os.Getenv("UPSTREAM_URL") + "/product/" + c.Param("productId"))
	})

	r.Run(":8085")
}
