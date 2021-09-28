package http

import (
	"github.com/ArtalkJS/ArtalkGo/lib"
	"github.com/ArtalkJS/ArtalkGo/model"
	"github.com/labstack/echo/v4"
)

type ParamsAdminSiteAdd struct {
	Name string `mapstructure:"name" param:"required"`
	Url  string `mapstructure:"url" param:"required"`
}

func ActionAdminSiteAdd(c echo.Context) error {
	if isOK, resp := AdminOnly(c); !isOK {
		return resp
	}

	var p ParamsAdminSiteAdd
	if isOK, resp := ParamsDecode(c, ParamsAdminSiteAdd{}, &p); !isOK {
		return resp
	}

	if !model.FindSite(p.Name).IsEmpty() {
		return RespError(c, "site 已存在")
	}

	site := model.Site{}
	site.Name = p.Name
	site.Url = p.Url
	err := lib.DB.Create(&site).Error
	if err != nil {
		return RespError(c, "site 创建失败")
	}

	return RespSuccess(c)
}
