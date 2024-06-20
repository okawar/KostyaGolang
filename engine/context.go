package engine

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	Response http.ResponseWriter
	Request  *http.Request
}

type Error struct {
	Message string `json:"message"`
}

func (c *Context) Error(status int, errorMsg string) {
	c.Response.WriteHeader(status)

	em := Error{Message: errorMsg}
	marsh, _ := json.Marshal(em)
	c.Response.Header().Set("Content-Type", "application/json")
	c.Response.Write([]byte(marsh))
}

func (c *Context) Print(data interface{}) {
	c.Response.Header().Set("Content-Type", "application/json")

	marsh, _ := json.Marshal(data)
	c.Response.Write([]byte(marsh))
}

// decoder
func ToStruct[T any](ctx Context) (T, error) {
	decoder := json.NewDecoder(ctx.Request.Body)
	var dt T
	err := decoder.Decode(&dt)
	return dt, err
}

//filter.go
/*package engine

import "net/http"

func StaticFiliter(ctx *Context) bool  {
	static, ok := ctx.CheckStatic(ctx.Request.URL.Path)
	if ok{
		http.ServeFile(ctx.Response, ctx.Request, env.Get[string]("template.url")+static)
		return false
	}
	return true
}
*/
