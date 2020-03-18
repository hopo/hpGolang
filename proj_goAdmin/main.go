package main

import (
	"fmt"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin"
	"github.com/GoAdminGroup/go-admin/engine"

	"proj_goadmin/tables"

	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/sqlite" // sql driver
	"github.com/GoAdminGroup/go-admin/plugins/admin"
	"github.com/GoAdminGroup/go-admin/plugins/example"
	_ "github.com/GoAdminGroup/themes/adminlte"
	"github.com/gin-gonic/gin"
	//"github.com/GoAdminGroup/go-admin/modules/config"
	//_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
)

func goadminInit(ctx *gin.Engine) {
	eng := engine.Default()

	adminPlugin := admin.NewAdmin(tables.Generators)
	//adminPlugin := admin.NewAdmin(tables.Generators)

	adminPlugin.AddGenerator("books", tables.GetBooksTable)

	examplePlugin := example.NewExample()

	err := eng.AddConfigFromJSON("./config.json").
		AddPlugins(adminPlugin, examplePlugin).
		Use(ctx)
	if err != nil {
		panic(err)
	}
}

func main() {
	r := gin.Default()

	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)

	goadminInit(r)

	r.Static("/uploads", "./uploads")

	r.GET("/admin", EndpointGetAdmin)

	fmt.Println("= = = = = Run Server = = = = =")
	//r.Run(":9090")
	r.Run(":8080")
}

// EndpointGetAdmin "/admin"
func EndpointGetAdmin(ctx *gin.Context) {
	fmt.Println("### CALLED~~ \"/admin\"")
	return
}
