package main
import (
    "github.com/cristian-fioravanti/ecommerceGo/middleware"
    "github.com/cristian-fioravanti/ecommerceGo/routes"
    "github.com/cristian-fioravanti/ecommerceGo/database"
    "github.com/gin-gonic/gin"
    "log"
)

func main() {
    port := ":8088"
    
    router := gin.New() //create new router
    router.Use(gin.Logger()) //use logger middleware

    routes.UserRoutes(router,database.Client) //call user routes
    router.Use(middleware.Authentication()) //use authentication middleware

    log.Fatal(router.Run(port)) //run the server on port 8080
}