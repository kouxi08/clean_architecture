package driver

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	// "github.com/kouxi08/clean_architecture/tree/main/adapter/controlleru"
	_ "github.com/go-sql-driver/mysql"
	"github.com/kouxi08/clean_architecture/tree/main/adapter/gateway"
	"github.com/kouxi08/clean_architecture/tree/main/adapter/presenter"
	"github.com/kouxi08/clean_architecture/tree/main/usecase/interactor"
)

func Serve(addr string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
		return
	}
	user := controller.User{
		OutputFactory: presenter.NewUserOutputPort,
		InputFactory:  interactor.NewUserInputPort,
		RepoFactory:   gateway.NewUserRepository,
		Conn:          conn,
	}
	http.HandleFunc("/", user.GetUserByID)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}
