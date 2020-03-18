package main

import (
	"fmt"
	_ "github.com/GoAdminGroup/go-admin/adapter/gin"               // Import the adapter, it must be imported. If it is not imported, you need to define it yourself.
	_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/sqlite" // sql driver
	"github.com/GoAdminGroup/go-admin/template/types"
	_ "github.com/GoAdminGroup/themes/adminlte" // Import the theme
	//_ "github.com/GoAdminGroup/go-admin/modules/db/drivers/mysql"
	"github.com/GoAdminGroup/go-admin/engine"
	"github.com/GoAdminGroup/go-admin/examples/datamodel"
	//"github.com/GoAdminGroup/go-admin/modules/config"
	//"github.com/GoAdminGroup/go-admin/modules/language"
	"github.com/GoAdminGroup/go-admin/plugins/admin"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Instantiate a GoAdmin engine object.
	eng := engine.Default()

	/*
	   // GoAdmin global configuration, can also be written as a json to be imported.
	   cfg := config.Config{
	       Databases: []config.Database{
	           {
	               Host:         "127.0.0.1",
	               Port:         "3306",
	               User:         "root",
	               Pwd:          "root",
	               Name:         "godmin",
	               MaxIdleCon:   50,
	               MaxOpenCon:   150,
	               Driver:       "mysql",
	           },
	       },
	       UrlPrefix: "admin", // The url prefix of the website.
	       // Store must be set and guaranteed to have write access, otherwise new administrator users cannot be added.
	       Store: config.Store{
	           Path:   "./uploads",
	           Prefix: "uploads",
	       },
	       Language: language.EN,
	   }
	*/

	// Import the business table configuration you need to manage here.
	// About Generators，see: https://github.com/GoAdminGroup/go-admin/blob/master/examples/datamodel/tables.go
	adminPlugin := admin.NewAdmin(datamodel.Generators)

	// Add configuration and plugins, use the Use method to mount to the web framework.
	//_ = eng.AddConfig(cfg).AddPlugins(adminPlugin).Use(r)

	// ***
	if err := eng.AddConfigFromJSON("./config.json").
		AddPlugins(adminPlugin).
		Use(r); err != nil {
		panic(err)
	}

	// ***
	r.Static("/uploads", "./uploads")

	r.GET("/admin", func(ctx *gin.Context) {
		engine.Content(ctx, func(ctx interface{}) (types.Panel, error) {
			return DashboardPage()
		})
	})

	fmt.Println("= = = = = Run Server = = = = =")
	_ = r.Run(":9090")
}
