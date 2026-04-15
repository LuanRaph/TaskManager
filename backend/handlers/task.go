package handlers

import (
	"log"
	"task-manager/database"
	"task-manager/models"

	"github.com/gin-gonic/gin"
)

func GetTask(c *gin.Context) {
	rows, err := database.DB.Query("SELECT id, title, description, done, created_at FROM tasks")
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()
	var tasks []models.Task
	for rows.Next() {
		var task models.Task
		rows.Scan(&task.ID, &task.Title, &task.Description, &task.Done, &task.CreatedAt)
		tasks = append(tasks, task)
	}
	c.JSON(200, tasks)
}

func CreateTask(c *gin.Context) {
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := database.DB.QueryRow(
		"INSERT INTO tasks (title, description) VALUES ($1, $2) RETURNING id", task.Title, task.Description,
	).Scan(&task.ID)
	if err != nil {
		log.Println("Erro ao criar task:", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, task)
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	_, err := database.DB.Exec(
		"UPDATE tasks SET title=$1, description=$2, done=$3 WHERE id=$4",
		task.Title, task.Description, task.Done, id,
	)
	if err != nil {
		log.Println("Erro ao atualizar task:", err)
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Task atualizada"})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id")

	_, err := database.DB.Exec("DELETE FROM tasks WHERE id=$1", id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Task deletada!"})
}
