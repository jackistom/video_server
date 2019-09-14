package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middleWareHandler struct { //添加中间件
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler { //添加一个新的中间件
	m := middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) { //通过middlewarehandler调用一个serverhttp
	//check session
	validateUserSession(r)

	m.r.ServeHTTP(w, r)
}

func RegisterHandlers() *httprouter.Router { //注册handler
	router := httprouter.New()

	router.POST("/user", CreateUser)

	router.POST("/user/:user_name", Login)

	return router
}

func main() {
	r := RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(":8000", mh)
}
