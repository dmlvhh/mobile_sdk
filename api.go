package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func GetToken(c *Client, r *GetTokenReq) (res *Response, err error) {
	c.config.APIBaseURL = "http://token.dctxiot.com/api"
	path := "/token/getToken"
	params := url.Values{}
	params.Add("appid", r.Appid)
	params.Add("app_secret", r.AppSecret)
	params.Add("channel_id", r.ChannelId)
	params.Add("times", r.Times)
	params.Add("sign", r.Sign)
	req, err := c.BuildRequest(http.MethodGet, path, nil, params)
	if err != nil {
		log.Printf("Failed to build request: %v", err)
		return nil, err
	}
	var getRes GetTokenRes
	res.Data = getRes
	fmt.Println(res)
	err = c.Do(req, &res)
	if err != nil {
		return nil, err
	}
	return
}
