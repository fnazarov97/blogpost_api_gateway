package handlers

import (
	"blogpost/genprotos/author"
	"blogpost/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CreateAuthor godoc
// @Summary     Create author
// @Description create a new author
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       author        body     models.CreateAuthorModel true  "author body"
// @Param       Authorization header   string                   false "Authorization"
// @Success     201           {object} models.JSONResponse{data=models.Author}
// @Failure     400           {object} models.JSONErrorResponse
// @Router      /v1/author [post]
func (h Handler) CreateAuthor(c *gin.Context) {
	var body models.CreateAuthorModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}

	// TODO - validation should be here

	author, err := h.grpcClients.Author.AddAuthor(c.Request.Context(), &author.CreateAuthorReq{
		Fullname: body.Fullname,
	})

	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, models.JSONResponse{
		Message: "Author | Created",
		Data:    author,
	})
}

// GetAuthorByID godoc
// @Summary     get author by id
// @Description get an author by id
// @Tags        authors
// @Accept      json
// @Param       id            path   string true  "Author ID"
// @Param       Authorization header string false "Authorization"
// @Produce     json
// @Success     200 {object} models.JSONResponse{data=models.Author}
// @Failure     400 {object} models.JSONErrorResponse
// @Router      /v1/author/{id} [get]
func (h Handler) GetAuthorByID(c *gin.Context) {
	idStr := c.Param("id")

	// TODO - validation
	author, err := h.grpcClients.Author.GetAuthorByID(c.Request.Context(), &author.Id{
		Id: idStr,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, models.JSONResponse{
		Message: "Author | Created",
		Data:    author,
	})
}

// GetAuthorList godoc
// @Summary     List author
// @Description get author
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       offset        query    int    false "0"
// @Param       limit         query    int    false "10"
// @Param       search        query    string false "search"
// @Param       Authorization header   string false "Authorization"
// @Success     200           {object} models.JSONResponse{data=[]models.Author}
// @Router      /v1/author [get]
func (h Handler) GetAuthorList(c *gin.Context) {
	offsetStr := c.DefaultQuery("offset", h.Conf.DefaultOffset)
	limitStr := c.DefaultQuery("limit", h.Conf.DefaultLimit)
	search := c.DefaultQuery("search", "")
	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: "offset error",
		})
		return
	}
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: "limit error",
		})
		return
	}
	authorList, err := h.grpcClients.Author.GetAuthorList(c.Request.Context(), &author.GetAuthorListReq{
		Offset: int64(offset),
		Limit:  int64(limit),
		Search: search,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.JSONErrorResponse{
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.JSONResponse{
		Message: "OK",
		Data:    authorList,
	})
}

// UpdateAuthor godoc
// @Summary     update author
// @Description update a new author
// @Tags        authors
// @Accept      json
// @Produce     json
// @Param       author        body     models.UpdateAuthorModel true  "author body"
// @Param       Authorization header   string                   false "Authorization"
// @Success     200           {object} models.JSONResponse{data=models.Author}
// @Response    400           {object} models.JSONErrorResponse
// @Router      /v1/author [put]
func (h Handler) UpdateAuthor(c *gin.Context) {
	var body models.UpdateAuthorModel
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{Error: err.Error()})
		return
	}
	updated, err := h.grpcClients.Author.UpdateAuthor(c.Request.Context(), &author.UpdateAuthorReq{
		Id:       body.ID,
		Fullname: body.Fullname,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, models.JSONErrorResponse{
			Error: "Faild update!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data":    updated,
		"message": "author | Update",
	})
}

// DeleteAuthor godoc
// @Summary     delete author by id
// @Description delete an author by id
// @Tags        authors
// @Accept      json
// @Param       id            path   string true  "author ID"
// @Param       Authorization header string false "Authorization"
// @Produce     json
// @Success     200 {object} models.JSONResponse{data=models.Author}
// @Failure     404 {object} models.JSONErrorResponse
// @Router      /v1/author/{id} [delete]
func (h Handler) DeleteAuthor(c *gin.Context) {
	idStr := c.Param("id")
	deleted, err := h.grpcClients.Author.DeleteAuthor(c.Request.Context(), &author.Id{
		Id: idStr,
	})
	if err != nil {
		c.JSON(http.StatusNotFound, models.JSONErrorResponse{
			Error: "Author have been deleted already!"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "author deleted",
		"data":    deleted,
	})
}
