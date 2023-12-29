package sso

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/exp/slog"
	"io"
	"net/http"
	"strings"
	"time"
)

var (
	ClientID     = "100027"
	ClientSecret = "LZ92e_efkiRJFbPGQOTUyQ=="
	SSOURL       = "https://sso.dev.seayoo.com/v2/oidc"
	RedirectURL  = "http://localhost:8888/login"
)

type Response struct {
	Code    int
	Message string
	Data    any
}

type ShiYouSSOToken struct {
	AccessToken      string `json:"access_token"`
	TokenType        string `json:"token_type"`
	Scope            string `json:"scope"`
	ExpiresIn        int64  `json:"expires_in"`
	ResponseError    string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type UserInfo struct {
	UserName string `json:"username"` // KOA 账号, 如 liuyi1
	Nickname string `json:"nickname"` // 中文姓名
	Email    string `json:"email"`
}

func beforeLogin(c *gin.Context) {
	loginUrl := fmt.Sprintf("%s/authorize/?response_type=code&client_id=%s&redirect_uri=%s&&scope=read&t=%d",
		SSOURL, ClientID, RedirectURL, time.Now().Unix())
	slog.Info("get redirect uri : ", "login_url", loginUrl)
	c.Redirect(301, loginUrl)
}

func login(c *gin.Context) {
	code := c.Query("code")
	slog.Info("login code auth code", "code", code)

	userInfo, err := doShiYouSSOLogin(code)
	if err != nil {
		slog.Error("do shiYouSSOLogin error", "err", err)
		return
	}

	c.JSON(http.StatusOK, &Response{
		Code:    200,
		Message: "success",
		Data:    gin.H{"token": "kkkk", "user_info": userInfo},
	})
}

func doShiYouSSOLogin(code string) (string, error) {
	tokenUrl := fmt.Sprintf("client_id=%s&client_secret=%s&grant_type=authorization_code&code=%s&redirect_uri=%s",
		ClientID, ClientSecret, code, RedirectURL)
	ssoUrl := fmt.Sprintf("%s/oidc/token", SSOURL)

	resp, err := http.Post(ssoUrl, "application/x-www-form-urlencoded", strings.NewReader(tokenUrl))
	if err != nil {
		slog.Error("get response have failed", "err", err)
		return "", err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		slog.Error("read response failed", "err", err)
		return "", err
	}

	var tokenInfo ShiYouSSOToken
	if err := json.Unmarshal(body, &tokenInfo); err != nil {
		slog.Error("unmarshal token failed", "err", err)
		return "", err
	}

	slog.Info("get token info ", "tokenInfo", tokenInfo)
	userInfoUrl := fmt.Sprintf("%s/oidc/info", SSOURL)
	getRequest, err := http.NewRequest(http.MethodGet, userInfoUrl, nil)
	if err != nil {
		slog.Error("get request failed", "err", err)
		return "", err
	}

	jwt := fmt.Sprintf("%s %s", tokenInfo.TokenType, tokenInfo.AccessToken)
	getRequest.Header.Add("Authorization", jwt)

	// get user info
	client := &http.Client{Timeout: 60 * time.Second}
	userResp, err := client.Do(getRequest)
	defer userResp.Body.Close()

	infoByte, _ := io.ReadAll(userResp.Body)
	var userInfo UserInfo
	if err = json.Unmarshal(infoByte, &userInfo); err != nil {
		slog.Error("json.Unmarshal userInfo failed", "err", err)
		return "", err
	}

	slog.Info("get userInfo success", "userInfo", userInfo)
	return userInfo.Email, nil
}
