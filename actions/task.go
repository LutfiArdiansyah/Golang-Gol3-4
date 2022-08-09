package actions

import (
	"digitask/models"
	"github.com/gobuffalo/buffalo"
	uuid2 "github.com/gofrs/uuid"
	"net/http"
)

// TaskGet default implementation.
func TaskGet(c buffalo.Context) error {
	tasks := models.Tasks{}

	err := models.DB.All(&tasks)

	if err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(err))
	}

	return c.Render(http.StatusOK, r.JSON(tasks))
}

// TaskGetByID default implementation.
func TaskGetByID(c buffalo.Context) error {
	id := c.Param("id")

	task := models.Task{}

	err := models.DB.Find(&task, id)

	if err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(err))
	}

	return c.Render(http.StatusOK, r.JSON(task))
}

// TaskCreate default implementation.
func TaskCreate(c buffalo.Context) error {
	payload := &models.Task{}

	if err := c.Bind(payload); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(err))
	}

	id, _ := uuid2.NewV4()

	task := models.Task{ID: id, Name: payload.Name, Assignee: payload.Assignee, Status: payload.Status, DueDate: payload.DueDate}

	err := models.DB.Create(&task)

	if err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(err))
	}

	return c.Render(http.StatusOK, r.JSON(task))
}

// TaskUpdate default implementation.
func TaskUpdate(c buffalo.Context) error {
	payload := &models.Task{}

	id := c.Param("id")

	if err := c.Bind(payload); err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(err))
	}

	task := models.Task{}

	errFind := models.DB.Find(&task, id)

	if errFind != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(errFind))
	}

	task.Name = payload.Name

	task.Assignee = payload.Assignee

	task.Status = payload.Status

	task.DueDate = payload.DueDate

	err := models.DB.Save(&task)

	if err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(err))
	}

	return c.Render(http.StatusOK, r.JSON(task))
}

// TaskDelete default implementation.
func TaskDelete(c buffalo.Context) error {
	id := c.Param("id")

	task := models.Task{}

	errFind := models.DB.Find(&task, id)

	if errFind != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(errFind))
	}

	err := models.DB.Destroy(&task)

	if err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(err))
	}

	return c.Render(http.StatusNoContent, nil)
}
