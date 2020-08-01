package reviews

import (
	"github.com/Al-un/emprev-api/internals/core"
	"github.com/gorilla/mux"
)

func LoadEndpoints(router *mux.Router) {
	router.Handle("/reviews", core.DoIfAdmin(handleCreateReview)).Methods("POST")
	router.Handle("/reviews/{reviewerUserId}", core.DoIfLogged(handleListReviews)).Methods("GET")
	router.Handle("/reviews/{reviewerUserId}/{reviewId}", core.DoIfLogged(handleUpdateReview)).Methods("PATCH")
}
