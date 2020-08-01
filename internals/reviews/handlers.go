package reviews

import (
	"encoding/json"
	"net/http"

	"github.com/Al-un/emprev-api/internals/core"
	"github.com/Al-un/emprev-api/internals/utils"
)

func handleCreateReview(w http.ResponseWriter, r *http.Request, claims core.JwtClaims) {
	var toCreateReview Review
	json.NewDecoder(r.Body).Decode(&toCreateReview)

	result, err := createReview(toCreateReview)
	if err != nil {
		utils.HandleInternalError(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func handleListReviews(w http.ResponseWriter, r *http.Request, claims core.JwtClaims) {
	reviewerUserID := utils.GetVar(r, "reviewerUserID")

	result, err := listReviews(reviewerUserID)
	if err != nil {
		utils.HandleInternalError(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func handleUpdateReview(w http.ResponseWriter, r *http.Request, claims core.JwtClaims) {
	reviewID := utils.GetVar(r, "reviewId")

	var toUpdateReview Review
	json.NewDecoder(r.Body).Decode(&toUpdateReview)

	result, err := updateReview(reviewID, toUpdateReview)
	if err != nil {
		utils.HandleInternalError(w, r, err)
		return
	}

	json.NewEncoder(w).Encode(result)
}
