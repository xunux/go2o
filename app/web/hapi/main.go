/**
 * Copyright 2015 @ z3q.net.
 * name : main.go
 * author : jarryliu
 * date : 2016-09-09 17:41
 * description :
 * history :
 */
package hapi

import (
	"fmt"
	"github.com/jsix/goex/echox"
	"github.com/jsix/gof"
	"go2o/core/variable"
	"net/http"
	"net/url"
)

type mainC struct {
	gof.App
}

func (m *mainC) Info(c *echox.Context) error {
	return c.String(http.StatusOK, `
        release : 2016-09-10
    `)
}

// 测试HAPI
func (m *mainC) Test(c *echox.Context) error {
	memberId := getMemberId(c)
	if memberId <= 0 {
		return requestLogin(c)
	}
	d := gof.Message{
		ErrCode: 0,
		Result:  true,
		Data:    memberId,
	}
	return c.JSONP(http.StatusOK, c.QueryParam("callback"), d)
}

// 请求登录
func (m *mainC) RequestLogin(c *echox.Context) error {
	referrer := c.QueryParam("return_url")
	if referrer == "" {
		referrer = c.Request().Referer()
	}
	target := fmt.Sprintf("%s://%s%s/auth?return_url=%s",
		variable.DOMAIN_PASSPORT_PROTO, variable.DOMAIN_PREFIX_PASSPORT,
		variable.Domain, url.QueryEscape(referrer))
	return c.Redirect(http.StatusFound, target)
}

// 跳转到用户中心
func (m *mainC) RedirectUc(c *echox.Context) error {
	returnUrl := c.QueryParam("url")
	if len(returnUrl) > 0 && returnUrl[0] != '/' {
		returnUrl = "/" + returnUrl
	}
	target := fmt.Sprintf("http://%s%s%s", variable.DOMAIN_PREFIX_MEMBER,
		variable.Domain, returnUrl)
	return c.Redirect(302, target)
}
