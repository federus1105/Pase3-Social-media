package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/federus1105/socialmedia/internals/models"
	"github.com/federus1105/socialmedia/internals/pkg"
	"github.com/federus1105/socialmedia/internals/repositories"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	ar *repositories.AuthRepository
}

func NewAutHandler(ar *repositories.AuthRepository) *AuthHandler {
	return &AuthHandler{ar: ar}
}

// Register godoc
// @Summary Register
// @Tags Authentication
// @Accept json
// @Produce json
// @Param order body models.UserRegister true "Register"
// @Success 201 {object} models.UserRegister
// @Router /auth/register [post]
func (a *AuthHandler) Register(ctx *gin.Context) {
	var body models.UserRegister
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request body",
		})
		return
	}
	newOrder, err := a.ar.Register(ctx.Request.Context(), body)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"order":   newOrder,
	})
}

// Login godoc
// @Summary Login
// @Tags Authentication
// @Accept json
// @Produce json
// @Param order body models.UserAuth true "Login"
// @Success 201 {object} models.UserAuth
// @Router /auth/login [post]
func (a *AuthHandler) Login(ctx *gin.Context) {
	// menerima body dan validasi
	var body models.UserAuth
	if err := ctx.ShouldBind(&body); err != nil {
		if strings.Contains(err.Error(), "required") {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "Email dan Password harus diisi",
			})
			return
		}

		if strings.Contains(err.Error(), "min") {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "Password minimum 8 karakter",
			})
			return
		}
		log.Println("Internal Server Error.\nCause: ", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "internal server error",
		})
		return
	}
	// ambil data user
	user, err := a.ar.Login(ctx.Request.Context(), body.Email)
	if err != nil {
		if strings.Contains(err.Error(), "user not found") {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   "Email atau Password salah",
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err,
			"status":  http.StatusInternalServerError,
		})
		return
	}

	// bandingkan password
	hc := pkg.NewHashConfig()
	isMatched, err := hc.CompareHashAndPassword(body.Password, user.Password)
	if err != nil {
		log.Println("Internal Server Error.\nCause: ", err)
		re := regexp.MustCompile("hash|crypto|argon2id|format")
		if re.Match([]byte(err.Error())) {
			log.Println("Error during Hashing")
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "internal server error",
		})
		return
	}
	if !isMatched {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Email atau Password salah",
		})
		return
	}

	// jika match, maka buatkan jwt dan kirim via response
	claims := pkg.NewJWTClaims(user.Id, user.Email)
	jwtToken, err := claims.GenToken()
	if err != nil {
		log.Println("Internal Server Error: Cause:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "internal server errorrr",
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"msg":     "Login SuccesFully",
		"token":   jwtToken,
	})
}
