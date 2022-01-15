package Register

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

var jwtKeys = []byte("my_secret_key")

var users = map[string]string{
	"lwgg": "12345678",
	"cjgg": "90abcdef",
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Hash(pwd string) string  {
	hash,err := bcrypt.GenerateFromPassword([]byte(pwd),bcrypt.DefaultCost)
	if err != nil{
		panic(err.Error())
	}
	encodingPwd := string(hash)
	return encodingPwd
}

func Sign(c *gin.Context)  {
	var cred Credentials
	cred.Username = c.PostForm("username")
	cred.Password = c.PostForm("password")
	err := bcrypt.CompareHashAndPassword([]byte(Hash(users[cred.Username])),[]byte(cred.Password))
	if err != nil  {
		c.JSON(http.StatusOK,gin.H{
			"status": http.StatusOK,
			"msg": "登陆失败",
		})
		return
	}
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: cred.Username,
		StandardClaims:jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKeys)
	if err != nil{
		c.JSON(http.StatusInternalServerError,gin.H{
			"status": http.StatusInternalServerError,
			"msg": "服务器错误",
		})
		return
	}
	cookieName := &http.Cookie{
		Name: "token",
		Value: tokenString,
		Expires: expirationTime,
	}
	http.SetCookie(c.Writer,cookieName)
}

func Welcome(c *gin.Context)  {
	cookies,err := c.Request.Cookie("token")
	if err == nil{
		c.Next()
	}else {
		c.Abort()
		c.JSON(http.StatusNotFound,gin.H{
			"status" : http.StatusNotFound,
			"msg": "无法访问页面",
			"error": "",
		})
	}
	tokenString := cookies.Value
	claims := &Claims{}
	token,err := jwt.ParseWithClaims(tokenString,claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKeys,nil
	})
	if err != nil{
		c.JSON(http.StatusUnauthorized,gin.H{})
		return
	}
	if !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"user": claims.Username,
		"msg": "登陆成功",
	})
}

func Refresh(c *gin.Context)  {
	cookie,err := c.Request.Cookie("token")
	if err == nil{
		c.Next()
	}else {
		c.Abort()
		c.JSON(http.StatusNotFound,gin.H{
			"status" : http.StatusNotFound,
			"msg": "无法访问页面",
			"error": "",
		})
		return
	}
	tokenStr := cookie.Value
	claim := &Claims{}
	token,err := jwt.ParseWithClaims(tokenStr,claim, func(token *jwt.Token) (interface{}, error) {
		return jwtKeys,err
	})
	if err != nil{
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized,gin.H{
				"status": http.StatusUnauthorized,
			})
			return
		}
		c.JSON(http.StatusBadRequest,gin.H{
			"status": http.StatusBadRequest,
		})
		return
	}
	if !token.Valid{
		c.JSON(http.StatusUnauthorized,gin.H{
			"status": http.StatusUnauthorized,
		})
	}
	if time.Unix(claim.ExpiresAt,0).Sub(time.Now()) > 30 * time.Second{
		c.JSON(http.StatusBadRequest,gin.H{
			"status": http.StatusBadRequest,
		})
	}
	expirationTime := time.Now().Add(5* time.Minute)
	claim.ExpiresAt = expirationTime.Unix()
	tkn := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	tokenString, err := tkn.SignedString(jwtKeys)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{
			"status": http.StatusInternalServerError,
		})
	}
	cookieName := &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	}
	http.SetCookie(c.Writer,cookieName)
}
