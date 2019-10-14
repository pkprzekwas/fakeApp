package todo

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/pkprzekwas/fakeApp/common"
	"net/http"
)

type (
	// TodoModel describes a TodoModel type
	TodoModel struct {
		gorm.Model
		Title     string `json:"title"`
		Completed int    `json:"completed"`
	}
	// TransformedTodo represents a formatted todo
	TransformedTodo struct {
		ID        uint   `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}
)

func AutoMigrate() {
	db := common.GetDB()
	db.AutoMigrate(&TodoModel{})
}

func Get(c *gin.Context) {
	id := c.Param("id")
	var todo TodoModel

	db := common.GetDB()
	db.First(&todo, id)

	if todo.ID == 0 {
		response(http.StatusNotFound,"No todo of id", c)
		return
	}

	_todo := TransformedTodo{
		ID:        todo.ID,
		Title:     todo.Title,
		Completed: common.IntToBool(todo.Completed),
	}

	response(http.StatusOK,  _todo, c)
}

func List(c *gin.Context)  {
	var todos []TodoModel
	var _todos []TransformedTodo

	db := common.GetDB()
	db.Find(&todos)

	if len(todos) == 0 {
		response(http.StatusNotFound,"No todos in the DB", c)
		return
	}

	for _, todo := range todos {
		_todos = append(_todos, TransformedTodo{
			ID: todo.ID,
			Title: todo.Title,
			Completed: common.IntToBool(todo.Completed),
		})
	}

	response(http.StatusOK, _todos, c)
}


func Create(c *gin.Context) {
	// TODO: change it to post form data
	var json TransformedTodo
	if err := c.ShouldBindJSON(&json); err != nil {
		response(http.StatusBadRequest, err.Error(), c)
		return
	}

	todo := TodoModel{
		Title: json.Title,
		Completed: common.BoolToInt(json.Completed),
	}

	db := common.GetDB()
	db.Save(&todo)

	response(http.StatusOK, fmt.Sprintf("Todo item saved successfully: %d", todo.ID), c)
}

func Update(c *gin.Context) {
	id := c.Param("id")
	var todo TodoModel

	db := common.GetDB()
	db.First(&todo, id)

	if todo.ID == 0 {
		response(http.StatusNotFound, "Todo item was not found.", c)
		return
	}

	var json TransformedTodo
	if err := c.ShouldBindJSON(&json); err != nil {
		response(http.StatusBadRequest, err.Error(), c)
		return
	}

	todo.Title = json.Title
	todo.Completed = common.BoolToInt(json.Completed)
	db.Save(&todo)

	response(http.StatusOK, todo, c)
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	var todo TodoModel

	db := common.GetDB()
	db.First(&todo, id)

	if todo.ID == 0 {
		response(http.StatusNotFound,"Todo item was not found.", c)
		return
	}

	db.Delete(&todo)

	response(http.StatusOK, "Successfully deleted", c)
}

func response(status int, data interface{}, c *gin.Context) {
	c.JSON(status, gin.H{
		"status": status,
		"data": data,
	})
}
