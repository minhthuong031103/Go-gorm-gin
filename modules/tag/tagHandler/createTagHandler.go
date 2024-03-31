package taghandler

import (
	appctx "go-gorm-gin/components/appCtx"
	tagrequest "go-gorm-gin/modules/tag/tagData/tagRequest"
	tagrepository "go-gorm-gin/modules/tag/tagRepository"
	tagservice "go-gorm-gin/modules/tag/tagService"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateTag godoc
//
//	@Summary		Create a new tag
//	@Description	Create a new tag based on the provided JSON data
//	@Tags			tags
//	@Accept			json
//	@Produce		json
//	@Param			tag	body		tagrequest.CreateTagRequest	true	"Tag data"
//	@Success		200	{object}	model.Tag					"Tag created successfully"
//	@Router			/tags [post]
func CreateTag(appCtx appctx.AppContext) gin.HandlerFunc {
	tagRepository := tagrepository.NewTagRepositoryImpl(appCtx.GetDb())
	tagService := tagservice.NewTagService(tagRepository)

	return func(c *gin.Context) {

		var request tagrequest.CreateTagRequest

		if err := c.ShouldBindJSON(&request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := request.Validate(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		tag, err := tagService.CreateTag(c, request.ToModel())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, *tag)

	}
}
