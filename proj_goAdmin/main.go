package main

import (
	"fmt"

	_ "github.com/GoAdminGroup/go-admin/adapter/gin"
	"github.com/GoAdminGroup/go-admin/engine"

	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/sqlite" // sql driver
	"github.com/GoAdminGroup/go-admin/plugins/admin"
	"github.com/GoAdminGroup/go-admin/plugins/example"
	_ "github.com/GoAdminGroup/themes/adminlte"
	"github.com/gin-gonic/gin"
	"proj_goadmin/tables"
	//"github.com/GoAdminGroup/go-admin/modules/config"
	//_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
)

func main() {
	r := gin.Default()

	//gin.SetMode(gin.ReleaseMode)
	gin.SetMode(gin.DebugMode)

	eng := engine.Default()

	adminPlugin := admin.NewAdmin(tables.Generators)
	//adminPlugin := admin.NewAdmin(tables.Generators)

	//adminPlugin.AddGenerator("posts", tables.GetPostsTable)

	examplePlugin := example.NewExample()

	err := eng.AddConfigFromJSON("./config.json").AddPlugins(adminPlugin, examplePlugin).Use(r)
	if err != nil {
		panic(err)
	}

	r.Static("/uploads", "./uploads")

	r.GET("/admin", func(ctx *gin.Context) {
		fmt.Println("### CALLED~~ \"/admin\"")
		return
	})

	fmt.Println("= = = = = Run Server = = = = =")
	r.Run(":9090")
}
