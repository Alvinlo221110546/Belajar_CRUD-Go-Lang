package routes

import (
	"database/sql"
	"net/http"

	"github.com/Alvinlo221110546/crud-employee-go/controller"
)

func MapRoutes(server *http.ServeMux, db *sql.DB) {
	server.HandleFunc("/", controller.NewHelloWorldController())
	server.HandleFunc("/employee", controller.NewIndexEmployee(db))
	server.HandleFunc("/employee/create", controller.NewCreateEmployeeContoller(db))
	server.HandleFunc("/employee/update", controller.NewUpdateEmployeeController(db))
	server.HandleFunc("/employee/delete", controller.NewDeleteEmployeeController(db))
}
