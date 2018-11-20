package store

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var controller = &Controller{Repository: Repository{}}

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes of our API
type Routes []Route

var routes = Routes{
	Route{
		"Authentication",
		"POST",
		"/get-token",
		controller.GetToken,
	},
	Route{
		"AddMacAddress",
		"POST",
		"/AddMacAddress",
		AuthenticationMiddleware(controller.AddMacAddress),
	},
	Route{
		"UpdateMacAddress",
		"PUT",
		"/UpdateMacAddress",
		AuthenticationMiddleware(controller.UpdateMacAddress),
	},
	// Get MacAddress by {id}
	Route{
		"GetMacAddress",
		"GET",
		"/macAddresses/{id}",
		controller.GetMacAddress,
	},
	// Get MacAddresses list
	Route{
		"GetMacAddresses",
		"GET",
		"/macAddresses",
		controller.GetMacAddresses,
	},
	// Delete MacAddress by {id}
	Route{
		"DeleteMacAddress",
		"DELETE",
		"/deleteMacAddress/{id}",
		AuthenticationMiddleware(controller.DeleteMacAddress),
	},
	// Search macAddress with string
	Route{
		"SearchMacAddress",
		"GET",
		"/Search/{query}",
		controller.SearchMacAddress,
	},
	Route{
		"AddScan",
		"POST",
		"/AddScan",
		AuthenticationMiddleware(controller.AddScan),
	},
	Route{
		"UpdateScan",
		"PUT",
		"/UpdateScan",
		AuthenticationMiddleware(controller.UpdateScan),
	},
	// Get Scan by {id}
	Route{
		"GetScan",
		"GET",
		"/scans/{id}",
		controller.GetScan,
	},
	// Get Scans by {id}
	Route{
		"GetScans",
		"GET",
		"/scans",
		controller.GetScans,
	},
	// Delete Scan by {id}
	Route{
		"DeleteScan",
		"DELETE",
		"/deleteScan/{id}",
		AuthenticationMiddleware(controller.DeleteScan),
	},
	// Search scan with string
	Route{
		"SearchScan",
		"GET",
		"/Search/{query}",
		controller.SearchScan,
	},
}

// NewRouter configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		log.Println(route.Name)
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
