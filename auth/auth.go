package auth

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/navneetshukl/database"
	"github.com/navneetshukl/models"
	"golang.org/x/crypto/bcrypt"
)

// !Register will create a user in our database
func Register(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")
	limit := c.PostForm("limit")
	password := c.PostForm("password")

	newpassword, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Println("Error in encrypting the password of the user ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Some Error Occured.Please try again",
		})
		return
	}

	user := models.User{
		Name:     name,
		Email:    email,
		Password: string(newpassword),
		Limit:    limit,
	}

	db, err := database.ConnectToDatabase()
	if err != nil {
		log.Println("Error in Connecting to database in Register function ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Some Error Occured,Please try again",
		})
		return

	}

	succ := db.Create(&user)

	if succ.Error != nil {
		log.Println("Error in storing to database in Register function", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "There is some internal error.Please try after sometime",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User Registered Successfully",
	})

}

func Home(c *gin.Context) {
	c.HTML(http.StatusOK, "register.page.tmpl", nil)
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.page.tmpl", nil)
}

// ! Login function will login the user
func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	db, err := database.ConnectToDatabase()
	if err != nil {
		log.Println("Error in Connecting to database in Register function ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Some Error Occured,Please try again",
		})
		return

	}

	var loginData models.User

	db.Where("email=?", email).First(&loginData)
	if loginData.ID == 0 {
		log.Println("Email Does not exist")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email Does not Exist",
		})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(loginData.Password), []byte(password))

	if err != nil {
		log.Println("Password does not exist")
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Password does not match",
		})
		return
	}

	//* Implement JWT authentication method

	secret := os.Getenv("SECRET")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": loginData.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})
	tokenString, err := token.SignedString([]byte(secret))

	if err != nil {
		log.Println("Error in signing the token ", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Some Error Occured.Please try again",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "User Login Successfully",
		"token":   tokenString,
	})
}
