package controllers

import (
	"example/go-jwt/initializers"
	"example/go-jwt/models"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context){

	var body struct {
		Email string
		Password string
	}


	 if c.Bind(&body)!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Failed to get data",
		})

		return
	 }

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password),10)

	if err!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Failed to hash password",
		})
		return 
	}

	user := models.User{Email: body.Email,Password: string(hash)}

	result := initializers.DB.Create(&user)

	if result.Error !=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Failed to create user",
		})
		return 
	}

	c.JSON(http.StatusOK, gin.H{})

}

func Login(c *gin.Context){

	var body struct {
		Email string
		Password string
	}


	 if c.Bind(&body)!=nil{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Failed to get data",
		})

		return
	 }

	 var user  models.User
	 initializers.DB.First(&user, "email = ?", body.Email)

	 if user.ID==0{
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Invalid Email or Password",
		})

		return
	 }

	 err := bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(body.Password))

	 if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{
			"error":"Invalid Email or Password",
		})

		return
	 }


token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	"sub": user.ID,
	"exp":time.Now().Add(time.Hour*24*30).Unix(),
})


tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

if err!=nil{
	c.JSON(http.StatusBadRequest,gin.H{
			"error":"Failed to create token",
		})

		return
}

		c.SetSameSite(http.SameSiteLaxMode)
		c.SetCookie("Authorization",tokenString,3600*24*30,"","",false,true)
c.JSON(http.StatusOK, gin.H{
})

}

func Validate(c *gin.Context){

	user,_:=c.Get("user")

		c.JSON(http.StatusOK,gin.H{
			"message":user,
		})

}