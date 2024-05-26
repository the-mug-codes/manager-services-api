package attachment

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	entity "github.com/kodit-tecnologia/service-manager/entities"
	repository "github.com/kodit-tecnologia/service-manager/repositories"
	project "github.com/kodit-tecnologia/service-manager/use_cases/project"
	helper "github.com/the-mug-codes/adapters-service-api/server/helpers"
)

type updateProject struct {
	Kind         string      `json:"kind" binding:"required"`
	Title        string      `json:"title" binding:"required"`
	Labels       []uuid.UUID `json:"labels"`
	PersonID     uuid.UUID   `json:"person_id"`
	Requirements []string    `json:"requirements" binding:"required"`
	Content      string      `json:"content" binding:"required"`
	Comments     *string     `json:"comments"`
	Deliveries   []string    `json:"deliveries"`
	DueDate      *time.Time  `json:"due_date"`
}

func (data updateProject) dataToUpdate(id uuid.UUID) (dataToUpdate *entity.Project) {
	dataToUpdate = &entity.Project{
		ID:           id,
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
		dataToUpdate.Labels = append(dataToUpdate.Labels, &entity.Label{
			ID: label,
		})
	}
	return dataToUpdate
}

// @Summary		Change a Project
// @Description	Updates a project by idin database.
// @Tags			Projects
// @Accept			json
// @Produce		json
// @Param			id		path		uuid.UUID		true	"ID"
// @Param			payload	body		updateProject	true	"payload"
// @Success		200		{object}	helper.ResponseOne[entities.Project]
// @Failure		400		{object}	helper.Error
// @Failure		401		{object}	helper.Error
// @Failure		404		{object}	helper.Error
// @Router			/projects/{id} [put]
func Update(context *gin.Context) {
	var requestBody *updateProject
	id, haveId := context.Params.Get("id")
	err := context.ShouldBindBodyWith(&requestBody, binding.JSON)
	if err != nil || !haveId {
		helper.ErrorResponse(context, 400, "cannot bind data", err.Error())
		return
	}
	uuid, err := uuid.Parse(id)
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot update", err.Error())
		return
	}
	response, err := project.Update(repository.Project(context), *requestBody.dataToUpdate(uuid))
	if err != nil {
		helper.ErrorResponse(context, 400, "cannot update", err.Error())
		return
	}
	helper.SuccessResponseOne(context, 200, response)
}
