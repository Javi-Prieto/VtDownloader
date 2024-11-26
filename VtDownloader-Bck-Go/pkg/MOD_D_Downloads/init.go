package modddownloads

import (
	"VtDownloader.Back.Go/pkg/MOD_D_Downloads/handler"
	"VtDownloader.Back.Go/pkg/MOD_D_Downloads/service"
	"github.com/gorilla/mux"
)

/*
Init

Description:

	Method used to initialize the Module D for the Downloads

Parameters:

	*r, mux.Router: Used to manage all the endpoints
*/
func Init(r *mux.Router) {
	serviceToUse := service.New()
	handler.New(r, serviceToUse)
}
