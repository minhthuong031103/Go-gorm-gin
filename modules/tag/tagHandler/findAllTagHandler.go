package taghandler

import (
	appctx "go-gorm-gin/components/appCtx"
	tagrepository "go-gorm-gin/modules/tag/tagRepository"
	tagservice "go-gorm-gin/modules/tag/tagService"
	"net/http"

	"github.com/gin-gonic/gin"
)

// FindAllTag godoc
// @Summary Retrieve all tags
// @Description Retrieve a list of all tags with pagination support
// @Tags tags
// @Accept json
// @Produce json
// @Param page query integer false "Page number (default is 1)"
// @Param limit query integer false "Number of items per page (default is 10)"
// @Success 200 {object} tagresponse.FindAllTagResponse "Paginated list of tags"
// @Router /tags [get
func FindAllTag(appCtx appctx.AppContext) gin.HandlerFunc {

	tagRepository := tagrepository.NewTagRepositoryImpl(appCtx.GetDb())
	tagService := tagservice.NewTagService(tagRepository)

	return func(c *gin.Context) {

		tags, err := tagService.FindAllTags()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, tags)

	}
}
