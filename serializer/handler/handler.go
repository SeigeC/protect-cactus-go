package handler

import (
	"github.com/gin-gonic/gin"
)

// Method HTTP method
type Method int

// methods
const (
	GET  Method = 1 << iota
	POST Method = 1 << iota
)

// Content-Type
const (
	ContentTypeJSON           = "application/json"
	ContentTypeFormUrlencoded = "application/x-www-form-urlencoded"
)

// M for map
type M map[string]interface{}

// S for slice
type S []interface{}

// ActionFunc handle the requests
type ActionFunc func(c *Context) (ActionResponse, error)

// ActionResponse for action response
type ActionResponse interface{}

// Handler contains all routers' info
type Handler struct {
	Name        string
	Middlewares gin.HandlersChain
	Actions     map[string]*Action
	SubHandlers []*Handler
}

func NewHandler(str string) Handler {
	return Handler{Name: str}
}

func (h *Handler) Use(middleware ...gin.HandlerFunc) {
	h.Middlewares = append(h.Middlewares, middleware...)
}

func (h *Handler) pushAction(method Method, relativePath string, f ActionFunc, login bool) {
	if h.Actions == nil{
		h.Actions = make(map[string]*Action)
	}
	if _, ok := h.Actions[relativePath]; ok {
		panic("Handler 已存在 " + relativePath)
	}
	h.Actions[relativePath] = NewAction(method, f, login)
}

// POST append a post func to actions
func (h *Handler) POST(relativePath string, f ActionFunc) {
	h.pushAction(POST, relativePath, f, false)
}

// POST append a need login post func to actions
func (h *Handler) LoginPOST(relativePath string, f ActionFunc) {
	h.pushAction(POST, relativePath, f, true)
}

// GET append a get func to actions
func (h *Handler) GET(relativePath string, f ActionFunc) {
	h.pushAction(GET, relativePath, f, false)
}

// GET append a get func to actions
func (h *Handler) LoginGet(relativePath string, f ActionFunc) {
	h.pushAction(GET, relativePath, f, true)
}

// Handler create a new handler
func (h *Handler) Handler(str string) *Handler {
	handler := &Handler{Name: str}
	h.SubHandlers = append(h.SubHandlers, handler)
	return handler
}

// Mount mount handler
func (h *Handler) Mount(r *gin.Engine) {
	h.mount(&r.RouterGroup)
}

func (h *Handler) mount(g *gin.RouterGroup) {
	g = g.Group(h.Name)
	g.Use(h.Middlewares...)
	for name, action := range h.Actions {
		actionHandler := action.GetHandler()
		if action.Method&GET != 0 {
			g.GET(name, actionHandler)
		}
		if action.Method&POST != 0 {
			g.POST(name, actionHandler)
		}
	}
	for _, sub := range h.SubHandlers {
		sub.mount(g)
	}
}
