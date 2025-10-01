package handlers

import (
	"net/http"
	"strconv"

	"github.com/federus1105/socialmedia/internals/models"
	"github.com/federus1105/socialmedia/internals/repositories"
	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	ch *repositories.CommentRepository
}

func NewCommentHandler(ch *repositories.CommentRepository) *CommentHandler {
	return &CommentHandler{ch: ch}
}

// CreateComment godoc
// @Summary Create a new comment on a post
// @Description Membuat komentar pada postingan tertentu, hanya untuk pengguna yang telah login (JWT).
// @Tags Comments
// @Accept json
// @Produce json
// @Param postingan_id path int true "ID Postingan"
// @Param comment body models.Comment true "Comment Body"
// @Success 200 {object} models.Comment
// @Failure 400 {object} map[string]interface{} "Invalid request body or parameter"
// @Failure 401 {object} map[string]interface{} "Unauthorized access"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Security BearerAuth
// @Router /comment/{postingan_id} [post]
func (ch *CommentHandler) CreateComment(ctx *gin.Context) {
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

	// Ambil postingan_id dari path param
	postinganIDParam := ctx.Param("postingan_id")
	postinganID, err := strconv.Atoi(postinganIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid postingan_id",
		})
		return
	}

	// Ambil data body komentar
	var commentReq models.Comment
	if err := ctx.ShouldBindJSON(&commentReq); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"error":   "Invalid request body",
		})
		return
	}

	// Bangun CommentBody lengkap untuk dikirim ke repository
	comment := models.Comment{
		PostinganId: postinganID, // dari path param
		UserID:      userID,      // dari JWT
		Teks:        commentReq.Teks,
	}

	// Simpan ke DB
	newComment, err := ch.ch.CreateComment(ctx, comment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   "Failed to create comment",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"data":    newComment,
	})
}
