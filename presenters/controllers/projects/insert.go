package attachment

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
	entity "github.com/the-mug-codes/service-manager-api/entities"
	repository "github.com/the-mug-codes/service-manager-api/repositories"
	project "github.com/the-mug-codes/service-manager-api/use_cases/project"
)

type insertProject struct {
	Kind   string `json:"kind" binding:"required"`
	Title  string `json:"title" binding:"required"`
	Labels []*struct {
		ID uuid.UUID `json:"id" binding:"required"`
	} `json:"labels"`
	PersonID     uuid.UUID  `json:"person_id"`
	Requirements []string   `json:"requirements" binding:"required"`
	Content      string     `json:"content" binding:"required"`
	Comments     *string    `json:"comments"`
	Deliveries   []string   `json:"deliveries"`
	DueDate      *time.Time `json:"due_date"`
}

func (data insertProject) dataToInsert() (dataToInsert *entity.Project) {
	dataToInsert = &entity.Project{
		Kind:         data.Kind,
		Title:        data.Title,
		PersonID:     data.PersonID,
		Requirements: data.Requirements,
		Content:      data.Content,
		Comments:     data.Comments,
		Deliveries:   data.Deliveries,
		DueDate:      data.DueDate,
	}
	for _, label := range data.Labels {
		dataToInsert.Labels = append(dataToInsert.Labels, &entity.Label{
			ID: label.ID,
		})
	}
	return dataToInsert
}

// @Summary		Create a Project
// @Description	Creates a new project in database.
// @Tags			Projects
// @Accept			json
// @Produce		json
// @Param			payload	body		insertProject	true	"payload"
// @Success		201		{object}	helper.ResponseOne[entities.Project]
// @Failure		400		{object}	helper.Error
// @Failure		401		{object}	helper.Error
// @Failure		404		{object}	helper.Error
// @Router			/projects [post]
func Insert(context *gin.Context) {
	var requestBody *insertProject
	err := context.ShouldBindBodyWith(&requestBody, binding.JSON)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot bind data", err.Error())
		return
	}
	response, err := project.Insert(repository.Project(context), *requestBody.dataToInsert())
	if err != nil {
		helper.ErrorResponse(context, 401, "cannot insert", err.Error())
		return
	}
	helper.SuccessResponseOne(context, 201, response)
}
