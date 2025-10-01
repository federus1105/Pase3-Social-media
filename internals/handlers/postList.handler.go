package handlers

import (
	"net/http"
	"strconv"

	"github.com/federus1105/socialmedia/internals/repositories"
	"github.com/gin-gonic/gin"
)

type PostListHandler struct {
	plh *repositories.PostListRepository
}

func NewPostListHandler(plh *repositories.PostListRepository) *PostListHandler {
	return &PostListHandler{plh: plh}
}

// GetUserPosts godoc
// @Summary Get posts of a user you follow
// @Description Mengambil semua postingan dari user yang diikuti. Hanya bisa melihat postingan jika sudah follow.
// @Tags Posts
// @Produce json
// @Param following_id path int true "ID user yang ingin dilihat postingannya"
// @Success 200 {object} models.GetUserPostsResponse
// @Failure 400 {object} map[string]interface{} "ID tidak valid"
// @Failure 401 {object} map[string]interface{} "Unauthorized: user not logged in"
// @Failure 403 {object} map[string]interface{} "Forbidden: belum mengikuti user tersebut"
// @Failure 500 {object} map[string]interface{} "Kesalahan server"
// @Security BearerAuth
// @Router /listpostingan/{following_id} [get]
func (h *PostListHandler) GetUserPosts(ctx *gin.Context) {
	// Ambil user login dari JWT
	userIDInterface, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "Unauthorized",
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
			"error":   "Invalid user ID type",
		})
		return
	}

	// Ambil user yang mau dilihat
	followingIDParam := ctx.Param("following_id")
	followingID, err := strconv.Atoi(followingIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid user ID",
		})
		return
	}

	// Ambil postingan
	posts, err := h.plh.GetUserPostsIfFollow(ctx, userID, followingID)
	if err != nil {
		if err.Error() == "Anda belum mengikuti user ini" {
			ctx.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"error":   err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    posts,
	})
}
