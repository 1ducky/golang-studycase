package main

import (
	"log"
	"net/http"
	"restApi/config"
	"restApi/internal/auth"
	"restApi/internal/auth/jwt"
	"restApi/internal/database"
	"restApi/internal/http/middleware"
	"restApi/internal/image"
	"restApi/internal/todos"
	todolog "restApi/internal/todos/todo-log"
	"restApi/internal/user"
	"time"
)

func main() {
	config.LoadEnv()
	conf := config.NewConfig()
	log.Print(conf.Database)
	db, err := database.NewDatabase(conf.Database)
	if err != nil {
		panic(err)
	}
	dbtx := database.NewDBTransaction(db)

	todologRepo := todolog.NewTodoLogRepository(db)

	todoRepo := todos.NewTodoMemory(db)
	todoService := todos.NewService(todoRepo, todologRepo, dbtx)
	todoHandler := todos.NewTodoHandler(todoService)

	userRepo := user.NewUserMemory(db)
	userService := user.NewService(userRepo)
	userHandler := user.NewHandler(userService)

	jwtService := jwt.NewJWT(conf.Auth)

	authRepo := auth.NewRepository(db)
	authService := auth.NewAuthService(authRepo, jwtService)
	authHandler := auth.NewHandler(authService)

	imageRepo := image.NewLocalStorage(conf.Image)
	imageService := image.NewService(imageRepo)
	imageHandler := image.NewHandler(imageService)

	mux := http.NewServeMux()
	todos.Register(todoHandler, mux)
	auth.Register(authHandler, mux)
	image.Register(imageHandler, mux)
	user.Register(userHandler, mux)

	middlewares := middleware.Chain(mux, middleware.Logging, middleware.NewAuthMiddleware(authService))

	server := &http.Server{
		Addr:              ":8000",
		Handler:           middlewares,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       10 * time.Second,
	}

	log.Fatal(server.ListenAndServe())

}
