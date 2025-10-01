package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/federus1105/socialmedia/internals/models"
	"github.com/federus1105/socialmedia/internals/repositories"
	"github.com/federus1105/socialmedia/internals/utils"
	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	ph *repositories.PostRepository
}

func NewPostHandler(ph *repositories.PostRepository) *PostHandler {
	return &PostHandler{ph: ph}
}

// CreatePost godoc
// @Summary Create a new post
// @Description Membuat postingan baru dengan caption dan optional image (multipart/form-data). Pengguna harus login (JWT).
// @Tags Posts
// @Accept multipart/form-data
// @Produce json
// @Param caption formData string true "Caption untuk postingan"
// @Param image formData file false "Gambar untuk postingan"
// @Success 201 {object} models.Post
// @Failure 400 {object} map[string]interface{} "Bad request: data tidak valid atau upload gagal"
// @Failure 401 {object} map[string]interface{} "Unauthorized: user not logged in"
// @Failure 500 {object} map[string]interface{} "Internal server error saat menyimpan data"
// @Security BearerAuth
// @Router /postingan [post]
func (ph *PostHandler) CreatePost(ctx *gin.Context) {
	// Ambil user_id dari JWT context
	userIDInterface, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "Unauthorized: user not logged in",
		})
		return
	}

	var userID int
	switch v := userIDInterface.(type) {
	case int:
		userID = v
	case float64:
		userID = int(v)
	default:
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "Invalid user ID type in context",
		})
		return
	}

	// Bind multipart/form-data ke struct
	var body models.PostBody
	if err := ctx.ShouldBind(&body); err != nil {
		log.Println("Gagal bind data:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Format data tidak valid",
		})
		return
	}

	// Upload image
	var filename string
	if body.Image != nil {
		savePath, generatedFilename, err := utils.UploadImageFile(ctx, body.Image, "public", fmt.Sprintf("post_user_%d", userID))
		if err != nil {
			log.Println("Upload gambar gagal:", err)
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}

		if err := ctx.SaveUploadedFile(body.Image, savePath); err != nil {
			log.Println("Gagal menyimpan file gambar:", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"error":   "Gagal menyimpan file gambar",
			})
			return
		}

		filename = generatedFilename
	}

	// Siapkan data post
	post := models.PostBody{
		User:    userID,
		Caption: body.Caption,
	}
	if filename != "" {
		post.ImageStr = filename
	}

	// Simpan ke repository
	newPost, err := ph.ph.CreatePost(ctx.Request.Context(), post)
	if err != nil {
		log.Println("Gagal simpan postingan:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Terjadi kesalahan saat menyimpan data",
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"success": true,
		"data":    newPost,
	})
}
