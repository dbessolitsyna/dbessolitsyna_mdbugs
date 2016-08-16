package controllers

import (
	"github.com/mindoktor/mdbugs/app/models"
	"github.com/mindoktor/mdbugs/app/routes"
	"github.com/revel/revel"
	"github.com/ottob/go-semver/semver"
	"fmt"
)

type App struct {
	GorpController
	PageTitle string
}

func (c App) PageLoad() revel.Result {
	c.PageTitle = "Correct page title"
	return nil
}

func (c App) Index() revel.Result {
	c.RenderArgs["pageTitle"] = c.PageTitle
	return c.Render()
}

func (c App) IndexPost(message string) revel.Result {
	vA, err := semver.NewVersion("1.2.3")
        if err != nil {
                fmt.Println(err.Error())
        }
        
	newCase := models.Case{Message: message, GuideVersion: vA}
	
	if err := c.Txn.Insert(&newCase); err != nil {
		panic(err)
	}
	
	c.Flash.Success("Success")
	return c.Redirect(routes.App.Index())
}
