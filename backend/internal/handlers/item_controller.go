package handlers

import (
	"net/http"
	"strconv"

	"project/internal/dto"
	"project/internal/ent"

	"project/internal/services"

	"github.com/gin-gonic/gin"
)

type ItemController interface {
	FindAll(ctx *gin.Context)
	FindById(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type itemController struct {
	service services.IItemService
}

func NewItemController(service services.IItemService) ItemController {
	return &itemController{service: service}
}

// FindAll godoc
//
//	@Summary		Get all items
//	@Description	Get all items from the database
//	@Tags			items
//	@Produce		json
//	@Success		200	{object}	ent.Item
//	@Failure		500
//	@Router			/items [get]
func (c *itemController) FindAll(ctx *gin.Context) {
	// サービスのFindAll()メソッドを呼び出し、結果をJSON形式で返す
	items, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": items})
}

// FindById godoc
//
//	@Summary		Get item by ID
//	@Description	Get item from the database by ID
//	@Tags			items
//	@Produce		json
//	@Param			id	path	int	true	"Item ID"
//	@Success		200
//	@Failure		400
//	@Failure		404
//	@Failure		500
//	@Router			/api/v1/items/{id} [get]\
func (c *itemController) FindById(ctx *gin.Context) {
	// IDで商品を取得
	user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userId := user.(*ent.User).ID

	// strconv.ParseUint()文字列を整数に変換, 10進数, 64ビット
	itemId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	item, err := c.service.FindById(int(itemId), int(userId))
	if err != nil {
		if err.Error() == "item not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": item})
}

// Create godoc
//
//	@Summary		Create a new item
//	@Description	Create a new item in the database
//	@Tags			items
//	@Accept			json
//	@Produce		json
//	@Param			item	body		dto.CreateItemInput	true	"Item information"	example({"name": "item1", "price": 100, "description": "This is item1"})
//	@Success		201		{object}	ent.Item
//	@Failure		400
//	@Failure		500
//	@Router			/items [post]
func (c *itemController) Create(ctx *gin.Context) {
	// ユーザー情報を取得
	user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userId := user.(*ent.User).ID

	var input dto.CreateItemInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newItem, err := c.service.Create(input, int(userId))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": newItem})
}

// Update godoc
//
//	@Summary		Update an item
//	@Description	Update an item in the database
//	@Tags			items
//	@Accept			json
//	@Produce		json
//	@Param			id		path	int					true	"Item ID"	default(1)
//	@Param			item	body	dto.UpdateItemInput	true	"Item information"
//	@Success		200
//	@Failure		400
//	@Failure		404
//	@Failure		500
//	@Router			/items/{id} [put]
func (c *itemController) Update(ctx *gin.Context) {
	//ユーザーを取得
	user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userId := user.(*ent.User).ID

	itemId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	// UpdateItemInput構造体を定義
	var input dto.UpdateItemInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedItem, err := c.service.Update(int(itemId), input, int(userId))
	if err != nil {
		if err.Error() == "item not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": updatedItem})
}

// Delete godoc
//
//	@Summary		Delete an item
//	@Description	Delete an item in the database
//	@Tags			items
//	@Produce		json
//	@Param			id	path		int	true	"Item ID"
//	@Success		204	{object}	string
//	@Failure		400
//	@Failure		404
//	@Failure		500
//	@Router			/items/{id} [delete]
func (c *itemController) Delete(ctx *gin.Context) {
	user, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	userId := user.(*ent.User).ID

	// idを取得
	itemId, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid id"})
		return
	}

	err = c.service.Delete(int(itemId), int(userId))
	if err != nil {
		if err.Error() == "item not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
