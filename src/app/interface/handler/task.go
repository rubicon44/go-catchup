package handler

import (
	"net/http"
	"strconv"

	"app/usecase"

	"github.com/gin-gonic/gin"
)

// TaskHandler task handlerのinterface
type TaskHandler interface {
	Post(c *gin.Context)
	Get(c *gin.Context)
	Put(c *gin.Context)
	Delete(c *gin.Context)
}

type taskHandler struct {
	taskUsecase usecase.TaskUsecase
}

// NewTaskHandler task handlerのコンストラクタ
func NewTaskHandler(taskUsecase usecase.TaskUsecase) TaskHandler {
	return &taskHandler{taskUsecase: taskUsecase}
}

type requestTask struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type responseTask struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// Post taskを保存するときのハンドラー
func (th *taskHandler) Post(c *gin.Context) {
	var req requestTask
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTask, err := th.taskUsecase.Create(req.Title, req.Content)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := responseTask{
		ID:      createdTask.ID,
		Title:   createdTask.Title,
		Content: createdTask.Content,
	}

	c.JSON(http.StatusCreated, res)
}

// Get taskを取得するときのハンドラー
func (th *taskHandler) Get(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	foundTask, err := th.taskUsecase.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := responseTask{
		ID:      foundTask.ID,
		Title:   foundTask.Title,
		Content: foundTask.Content,
	}

	c.JSON(http.StatusOK, res)
}

// Put taskを更新するときのハンドラー
func (th *taskHandler) Put(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req requestTask
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTask, err := th.taskUsecase.Update(id, req.Title, req.Content)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := responseTask{
		ID:      updatedTask.ID,
		Title:   updatedTask.Title,
		Content: updatedTask.Content,
	}

	c.JSON(http.StatusOK, res)
}

// Delete taskを削除するときのハンドラー
func (th *taskHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = th.taskUsecase.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
