package handlers

import (
	"net/http"
	"strconv"

	"github.com/federus1105/socialmedia/internals/models"
	"github.com/federus1105/socialmedia/internals/repositories"
	"github.com/gin-gonic/gin"
)

type LikeHandler struct {
	lh *repositories.LikeRepository
}

func NewLikeHandler(lh *repositories.LikeRepository) *LikeHandler {
	return &LikeHandler{lh: lh}
}


// LikePost godoc
// @Summary Like a post
// @Description Menyukai sebuah postingan. Pengguna harus login (menggunakan JWT).
// @Tags Likes
// @Accept json
// @Produce json
// @Param post_id path int true "ID Postingan yang ingin di-like"
// @Success 200 {object} models.LikeResponse
// @Failure 400 {object} map[string]interface{} "Invalid post_id"
// @Failure 401 {object} map[string]interface{} "Unauthorized: user not logged in"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Security BearerAuth
// @Router /like/{post_id} [post]
func (lh *LikeHandler) LikePost(ctx *gin.Context) {
	// Ambil user login dari JWT
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
			"error":   "Invalid user ID type in context"})
		return
	}

	// Ambil post_id dari path
	postIDParam := ctx.Param("post_id")
	postID, err := strconv.Atoi(postIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"succes": false,
			"error":  "Invalid post_id",
		})
		return
	}

	// Bangun body Like
	like := models.Like{
		UserID: userID,
		PostID: postID,
	}

	// Simpan ke DB lewat repository
	newLike, err := lh.lh.CreateLike(ctx.Request.Context(), like)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Post liked",
		"data":    newLike,
	})
}
