package controllers

import (
	"go-template/models"
	"go-template/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTaskHandler(c *gin.Context) {
	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido: " + err.Error()})
		return
	}

	if len(task.Title) < 3 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "El título debe tener al menos 3 caracteres"})
		return
	}

	if len(task.Description) < 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "La descripción debe tener al menos 5 caracteres"})
		return
	}

	result, err := services.CreateTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear la tarea"})
		return
	}

	c.JSON(http.StatusOK, result)
}

// Listar todas las tareas
func GetTasksHandler(c *gin.Context) {
	tasks, err := services.GetAllTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener las tareas"})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// Marcar una tarea como completada
func CompleteTaskHandler(c *gin.Context) {
	id := c.Param("id")

	if _, err := primitive.ObjectIDFromHex(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := services.CompleteTask(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar la tarea"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Tarea marcada como completada"})
}

// Obtener una tarea por ID
func GetTaskByIDHandler(c *gin.Context) {
	id := c.Param("id")

	if _, err := primitive.ObjectIDFromHex(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	task, err := services.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Tarea no encontrada"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// Eliminar una tarea por ID
func DeleteTaskHandler(c *gin.Context) {
	id := c.Param("id")

	if _, err := primitive.ObjectIDFromHex(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	err := services.DeleteTaskByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar la tarea"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Tarea eliminada"})
}

// Eliminar tareas completadas
func DeleteCompletedTasksHandler(c *gin.Context) {
	err := services.DeleteCompletedTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron eliminar las tareas completadas"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Tareas completadas eliminadas"})
}
