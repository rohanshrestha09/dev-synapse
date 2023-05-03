package project

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/rohanshrestha09/dev-synapse/dto"
	"github.com/rohanshrestha09/dev-synapse/enums"
	"github.com/rohanshrestha09/dev-synapse/models"
	"github.com/rohanshrestha09/dev-synapse/service/database"
	"github.com/rohanshrestha09/dev-synapse/utils"
)

// Get Project godoc
//
//	@Summary	Get a project
//	@Tags		Project
//	@Accept		json
//	@Produce	json
//	@Param		id	path		int	true	"Project ID"
//	@Success	200	{object}	database.GetResponse[models.Project]
//	@Router		/project/{id} [get]
func Get(ctx *gin.Context) {

	project := ctx.MustGet("project").(models.Project)

	ctx.JSON(
		http.StatusOK,
		database.GetResponse[models.Project]{
			Message: "Project Fetched",
			Data:    project,
		})

}

// Get Project godoc
//
//	@Summary	Get all projects
//	@Tags		Project
//	@Accept		json
//	@Produce	json
//	@Param		page	query		int		false	"Page"
//	@Param		size	query		int		false	"Page size"
//	@Param		sort	query		string	false	"Sort"	Enums(id, created_at, name)
//	@Param		order	query		string	false	"Order"	Enums(asc, desc)
//	@Param		search	query		string	false	"Search"
//	@Success	200		{object}	database.GetAllResponse[models.Project]
//	@Router		/project [get]
func GetAll(ctx *gin.Context) {

	args := database.GetAllArgs[models.Project]{
		Pagination: true,
		Search:     true,
		Include:    map[string][]string{"User": {"Email"}},
		Filter: models.Project{
			Published: true,
			Status:    enums.OPEN,
		},
	}

	response, err := database.GetAll(ctx.BindQuery, args)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Response)
		return
	}

	ctx.JSON(http.StatusOK, response)

}

// Create Project godoc
//
//	@Summary	Create a project
//	@Tags		Project
//	@Accept		mpfd
//	@Produce	json
//	@Param		name				formData	string	true	"Name"
//	@Param		description			formData	string	true	"Description"
//	@Param		published			formData	boolean	false	"Published"
//	@Param		estimatedDuration	formData	int		true	"Estimated Duration"
//	@Param		image				formData	file	true	"File to upload"
//	@Success	201					{object}	database.Response
//	@Router		/project [post]
//	@Security	Bearer
func Create(ctx *gin.Context) {
	authUser := ctx.MustGet("auth").(models.User)

	var projectCreateDto dto.ProjectCreateDTO

	if err := ctx.ShouldBindWith(&projectCreateDto, binding.FormMultipart); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	project := models.Project{
		UserID:            authUser.ID,
		Name:              projectCreateDto.Name,
		Description:       projectCreateDto.Description,
		Published:         projectCreateDto.Published,
		EstimatedDuration: projectCreateDto.EstimatedDuration,
	}

	if file, err := ctx.FormFile("image"); err == nil {
		project.Image, project.ImageName, err = utils.UploadFile(file, enums.PROJECT_DIR, enums.IMAGE)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	response, err := database.Create(&project)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Response)
		return
	}

	ctx.JSON(http.StatusCreated, response("Project Created"))

}

// Update Project godoc
//
//	@Summary	Update project
//	@Tags		Project
//	@Accept		mpfd
//	@Produce	json
//	@Param		id					path		int		true	"Project ID"
//	@Param		name				formData	string	false	"Name"
//	@Param		description			formData	string	false	"Description"
//	@Param		estimatedDuration	formData	int		false	"Estimated Duration"
//	@Param		image				formData	file	false	"File to upload"
//	@Success	201					{object}	database.Response
//	@Router		/project/{id} [patch]
//	@Security	Bearer
func Update(ctx *gin.Context) {

	project := ctx.MustGet("project").(models.Project)

	var projectUpdateDto dto.ProjectUpdateDTO

	if err := ctx.ShouldBindWith(&projectUpdateDto, binding.FormMultipart); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"message": err.Error()})
		return
	}

	var imageUrl, imageName string

	if file, err := ctx.FormFile("image"); err == nil {
		if imageUrl, imageName, err = utils.UploadFile(file, enums.PROJECT_DIR, enums.IMAGE); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		if err := utils.DeleteFile(string(enums.PROJECT_DIR) + project.ImageName); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	data := models.Project{
		Name:              projectUpdateDto.Name,
		Description:       projectUpdateDto.Description,
		EstimatedDuration: projectUpdateDto.EstimatedDuration,
		Image:             imageUrl,
		ImageName:         imageName,
	}

	response, err := database.Update(project, data)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Response)
		return
	}

	ctx.JSON(http.StatusCreated, response("Profile Updated"))

}
