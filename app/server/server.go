package server

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	// Database driver ...
	_ "github.com/go-sql-driver/mysql"

	"github.com/MrWebUzb/apiserver/app/middleware"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

// Server ...
type Server struct {
	Log    *logrus.Logger
	Router *mux.Router
	DB     *sql.DB
}

// NewServer ...
func NewServer() *Server {
	return &Server{
		Log:    logrus.New(),
		Router: mux.NewRouter(),
		DB:     nil,
	}
}

// Start ...
func (s *Server) Start(ctx context.Context, level logrus.Level) (err error) {
	s.DB, err = s.ConfigureDB()

	if err != nil {
		s.Log.Fatalf("Error with connection database: %+v", err)
	} else {
		s.Log.Infof("Successfully connected to database")
	}

	// Configurations
	s.ConfigureLogger(level)
	s.ConfigureErrorRouters()
	s.ConfigureRouters()

	wrappedRouter := middleware.RequestLog(s.Log, s.Router)

	server := &http.Server{
		Addr:    ":8080",
		Handler: wrappedRouter,
	}

	go func() {
		if err = server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			s.Log.Fatalf("Error listening: %s\n", err)
		}
	}()

	s.Log.Infof("Server started: http://localhost:8080\n")
	<-ctx.Done()

	s.Log.Warnf("Server shutting down ...")

	ctxShutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer func() {
		cancel()

		err = s.DB.Close()

		if err != nil {
			s.Log.Fatalf("Error while closing database connection: %+v", err)
		}

		s.Log.Warnf("Successfully closed database")
	}()

	if err = server.Shutdown(ctxShutdown); err != nil {
		s.Log.Fatalf("Error when shutdown server: %s", err)
	}

	s.Log.Warnf("Server properly shut down")

	if err == http.ErrServerClosed {
		err = nil
	}

	return
}

// ConfigureLogger ...
func (s *Server) ConfigureLogger(level logrus.Level) {
	s.Log.SetLevel(level)

	// multiWriter := io.MultiWriter(os.Stdout, s.log.Writer())
	// s.log.SetOutput(multiWriter)
}

// ConfigureDB ...
func (s *Server) ConfigureDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "asliddin:asliddin@/todo_app?parseTime=true&charset=utf8")

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
