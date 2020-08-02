package reviews

import (
	"net/http"

	"github.com/Al-un/emprev-api/internals/core"
	"github.com/gorilla/mux"
)

func LoadEndpoints(router *mux.Router) {
	router.Handle("/reviews/", core.DoIfAdmin(handleListAllReviews)).Methods(http.MethodGet, http.MethodOptions)
	router.Handle("/reviews/", core.DoIfAdmin(handleCreateReview)).Methods(http.MethodPost, http.MethodOptions)
	router.Handle("/reviews/{reviewerUserID}/", core.DoIfLogged(handleListReviews)).Methods(http.MethodGet, http.MethodOptions)
	router.Handle("/reviews/{reviewerUserID}/{reviewID}/", core.DoIfLogged(handleUpdateReview)).Methods(http.MethodPatch, http.MethodOptions)
}
