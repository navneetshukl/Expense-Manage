package auth

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/navneetshukl/database"
	"github.com/navneetshukl/helpers"
	"github.com/navneetshukl/models"
	"github.com/navneetshukl/redis"
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

	/* c.JSON(http.StatusOK, gin.H{
		"message": "User Registered Successfully",
	}) */

	c.Redirect(http.StatusSeeOther, "/user/login")

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
		"sub": email,
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

	//? Save this JWT token to Cookie

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, int(time.Hour*24*30), "/", "", false, true)

	//* Storing the name and email to the Redis

	name, err := helpers.GetName(email)

	if err != nil {
		log.Println("Error in Getting the name of current login user ", err)
	}
	data := map[string]interface{}{
		"email": email,
		"name":  name,
	}

	err = redis.StoreUserDetailInRedis(data)
	if err != nil {
		log.Println("Error in storing the user details to redis in Login Handler ", err)
	}

	/* c.JSON(http.StatusOK, gin.H{
		"message": "User Login Successfully",
	}) */

	c.Redirect(http.StatusSeeOther, "/expense")

}

// ! Signup will signup the user
func Signup(c *gin.Context) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", "", -1, "/", "", false, true)

	log.Println("I am on signup page")
	c.Redirect(http.StatusSeeOther, "/user/login")
}
