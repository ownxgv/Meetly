package context

import "github.com/gin-gonic/gin"

// HTTPContext — интерфейс для обработки HTTP-запросов
type HTTPContext interface {
	JSON(code int, obj interface{}) error
	BindJSON(obj interface{}) error
	Param(key string) string
}

// GinContextAdapter — адаптер для работы с Gin
type GinContextAdapter struct {
	C *gin.Context
}

func (g *GinContextAdapter) JSON(code int, obj interface{}) error {
	g.C.JSON(code, obj)
	return nil
}

func (g *GinContextAdapter) BindJSON(obj interface{}) error {
	return g.C.ShouldBindJSON(obj)
}

func (g *GinContextAdapter) Param(key string) string {
	return g.C.Param(key)
}
