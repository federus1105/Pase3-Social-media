package handlers

import (
	"net/http"
	"strconv"

	"github.com/federus1105/socialmedia/internals/repositories"
	"github.com/gin-gonic/gin"
)

type FollowHandler struct {
	fh *repositories.FollowRepository
}

func NewFollowHandler(fh *repositories.FollowRepository) *FollowHandler {
	return &FollowHandler{fh: fh}
}

// Follow godoc
// @Summary Follow a user
// @Description Mengikuti pengguna lain. Pengguna harus login (menggunakan JWT).
// @Tags Follow
// @Accept json
// @Produce json
// @Param id_user path int true "ID pengguna yang akan di-follow"
// @Success 200 {object} map[string]interface{} "Follow berhasil"
// @Failure 400 {object} map[string]interface{} "ID user tidak valid"
// @Failure 401 {object} map[string]interface{} "Unauthorized / tidak login"
// @Failure 500 {object} map[string]interface{} "Kesalahan server saat memproses follow"
// @Security BearerAuth
// @Router /follow/{id_user} [post]
func (h *FollowHandler) Follow(ctx *gin.Context) {
	// Ambil user_id dari JWT context (follower)
	userIDInterface, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "Unauthorized: user not logged in",
		})
		return
	}

	var followerID int
	switch v := userIDInterface.(type) {
	case int:
		followerID = v
	case float64:
		followerID = int(v)
	default:
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"success": false,
			"error":   "Invalid user ID type in context",
		})
		return
	}

	// Ambil following_id dari path param
	followingIDParam := ctx.Param("id_user")
	followingID, err := strconv.Atoi(followingIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid following_id",
		})
		return
	}

	// Simpan ke DB via repository
	newFollow, err := h.fh.Follow(ctx.Request.Context(), followerID, followingID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "followed successfully",
		"data":    newFollow,
	})
}