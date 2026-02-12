package router

import (
	"github.com/yuricapella/Go-Learning/order_cqrs/src/customer/handlers"
	routeschema "github.com/yuricapella/Go-Learning/order_cqrs/src/router/routeschema"
)

var Routes = []routeschema.Route{
	{
		Path:            "/customers",
		Method:          "POST",
		HandlerFunction: handlers.Create,
	},
	{
		Path:            "/customers/:id",
		Method:          "GET",
		HandlerFunction: handlers.GetByID,
	},
}
