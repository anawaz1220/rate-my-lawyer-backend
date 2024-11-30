package handlers

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent"
	"github.com/mikestefanello/pagoda/ent/jurisdiction"
	"github.com/mikestefanello/pagoda/ent/lawyer"
	"github.com/mikestefanello/pagoda/ent/predicate"
	"github.com/mikestefanello/pagoda/ent/review"
	"github.com/mikestefanello/pagoda/ent/rmluser"
	"github.com/mikestefanello/pagoda/pkg/services"
)

type LawyerAPI struct {
	services.Container
}

func (h *LawyerAPI) Init(c *services.Container) error {
	h.Container = *c
	return nil
}

type SubmitReviewRequest struct {
	Rating  int    `json:"rating"`
	Comment string `json:"comment"`
	UserID  *int   `json:"user_id"` // Optional user ID
}

func (h *LawyerAPI) Routes(g *echo.Group) {
	g.GET("/api/lawyers", h.GetLawyers)
	g.GET("/api/lawyers/:id", h.GetLawyerDetails)
	g.POST("/api/lawyers/:id/reviews", h.SubmitReview)
	g.GET("/api/lawyers/:id/reviews", h.GetLawyerReviews)
}

type ReviewResponse struct {
	ID        int       `json:"id"`
	Rating    int       `json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
	UserName  string    `json:"user_name,omitempty"`
}

type UserInfo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (h *LawyerAPI) GetLawyerReviews(c echo.Context) error {
	lawyerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid lawyer ID"})
	}

	ctx := c.Request().Context()

	// Check if the lawyer exists
	exists, err := h.Container.ORM.Lawyer.Query().Where(lawyer.ID(lawyerID)).Exist(ctx)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to check lawyer existence"})
	}
	if !exists {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Lawyer not found"})
	}

	reviews, err := h.Container.ORM.Review.Query().
		Where(review.HasLawyerWith(lawyer.ID(lawyerID))).
		Order(ent.Desc(review.FieldCreatedAt)).
		WithUser().
		All(ctx)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch reviews"})
	}

	var response []ReviewResponse
	for _, r := range reviews {
		reviewResponse := ReviewResponse{
			ID:        r.ID,
			Rating:    r.Rating,
			Comment:   r.Comment,
			CreatedAt: r.CreatedAt,
		}
		if r.Edges.User != nil {
			reviewResponse.UserName = r.Edges.User.Name
		}
		response = append(response, reviewResponse)
	}

	return c.JSON(http.StatusOK, response)
}

func (h *LawyerAPI) SubmitReview(c echo.Context) error {
	lawyerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid lawyer ID"})
	}

	var req SubmitReviewRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request body"})
	}

	if req.Rating < 1 || req.Rating > 5 {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Rating must be between 1 and 5"})
	}

	ctx := c.Request().Context()

	// Use a retry loop for optimistic locking
	maxRetries := 5
	var review *ent.Review
	for i := 0; i < maxRetries; i++ {
		review, err = h.submitReviewTx(ctx, lawyerID, &req)
		if err == nil {
			break
		}
		if !ent.IsConstraintError(err) {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
		}
		log.Printf("Retry %d: Constraint error occurred, retrying...", i+1)
	}

	if review == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to submit review after multiple attempts"})
	}

	response := ReviewResponse{
		ID:        review.ID,
		Rating:    review.Rating,
		Comment:   review.Comment,
		CreatedAt: review.CreatedAt,
	}

	// Fetch the associated user if it exists
	user, err := review.QueryUser().Only(ctx)
	if err == nil && user != nil {
		response.UserName = user.Name
	} else if !ent.IsNotFound(err) {
		log.Printf("Error fetching associated user: %v", err)
	}

	return c.JSON(http.StatusCreated, response)
}

func (h *LawyerAPI) submitReviewTx(ctx context.Context, lawyerID int, req *SubmitReviewRequest) (*ent.Review, error) {
	tx, err := h.Container.ORM.Tx(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	lawyer, err := tx.Lawyer.Query().Where(lawyer.ID(lawyerID)).Only(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, echo.NewHTTPError(http.StatusNotFound, "Lawyer not found")
		}
		return nil, err
	}

	reviewCreate := tx.Review.Create().
		SetRating(req.Rating).
		SetComment(req.Comment).
		SetCreatedAt(time.Now()).
		SetLawyer(lawyer)

	if req.UserID != nil {
		user, err := tx.RMLUser.Query().Where(rmluser.ID(*req.UserID)).Only(ctx)
		if err != nil {
			if ent.IsNotFound(err) {
				return nil, echo.NewHTTPError(http.StatusNotFound, "User not found")
			}
			return nil, err
		}
		reviewCreate.SetUser(user)
		_, err = user.Update().AddReviewCount(1).Save(ctx)
		if err != nil {
			return nil, err
		}
	}

	review, err := reviewCreate.Save(ctx)
	if err != nil {
		return nil, err
	}

	newReviewCount := lawyer.ReviewCount + 1
	newAverageRating := (lawyer.AverageRating*float64(lawyer.ReviewCount) + float64(req.Rating)) / float64(newReviewCount)

	_, err = lawyer.Update().
		SetAverageRating(newAverageRating).
		SetReviewCount(newReviewCount).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return review, tx.Commit()
}

type LawyerListResponse struct {
	ID               int     `json:"id"`
	FullName         string  `json:"full_name"`
	Address          string  `json:"address"`
	PracticingStatus string  `json:"practicing_status"`
	AverageRating    float64 `json:"average_rating"`
	ReviewCount      int     `json:"review_count"`
}

func (h *LawyerAPI) GetLawyers(c echo.Context) error {
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page < 1 {
		page = 1
	}
	limit := 6
	offset := (page - 1) * limit

	fullName := c.QueryParam("full_name")
	practicingStatus := c.QueryParam("practicing_status")
	jurisdictionName := c.QueryParam("jurisdiction")

	ctx := c.Request().Context()

	predicates := []predicate.Lawyer{}

	if fullName != "" {
		predicates = append(predicates, lawyer.FullNameContainsFold(fullName))
	}

	if practicingStatus != "" {
		predicates = append(predicates, lawyer.PracticingStatus(practicingStatus))
	}

	if jurisdictionName != "" {
		predicates = append(predicates, lawyer.HasJurisdictionsWith(jurisdiction.NameEQ(jurisdictionName)))
	}

	lawyers, err := h.Container.ORM.Lawyer.Query().
		Where(predicates...).
		Limit(limit).
		Offset(offset).
		Order(ent.Asc(lawyer.FieldID)).
		All(ctx)

	if err != nil {
		log.Printf("Error fetching lawyers: %v", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch lawyers"})
	}

	var response []LawyerListResponse
	for _, l := range lawyers {
		log.Printf("Lawyer ID: %d, Average Rating: %f, Review Count: %d", l.ID, l.AverageRating, l.ReviewCount)
		response = append(response, LawyerListResponse{
			ID:               l.ID,
			FullName:         l.FullName,
			Address:          l.Address,
			PracticingStatus: l.PracticingStatus,
			AverageRating:    l.AverageRating,
			ReviewCount:      l.ReviewCount,
		})
	}

	return c.JSON(http.StatusOK, response)
}

type LawyerDetailsResponse struct {
	ID               int      `json:"id"`
	FullName         string   `json:"full_name"`
	FirstName        string   `json:"first_name"`
	MiddleName       string   `json:"middle_name"`
	LastName         string   `json:"last_name"`
	Gender           string   `json:"gender"`
	Address          string   `json:"address"`
	Phone            string   `json:"phone"`
	PracticingStatus string   `json:"practicing_status"`
	ProfileURL       string   `json:"profile_url"`
	AverageRating    float64  `json:"average_rating"`
	ReviewCount      int      `json:"review_count"`
	Jurisdictions    []string `json:"jurisdictions"`
	Decisions        []string `json:"decisions"`
}

func (h *LawyerAPI) GetLawyerDetails(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid lawyer ID"})
	}

	ctx := c.Request().Context()

	l, err := h.Container.ORM.Lawyer.Query().
		Where(lawyer.ID(id)).
		WithJurisdictions().
		WithDecisions().
		First(ctx)

	if err != nil {
		if ent.IsNotFound(err) {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "Lawyer not found"})
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to fetch lawyer details"})
	}

	response := LawyerDetailsResponse{
		ID:               l.ID,
		FullName:         l.FullName,
		FirstName:        l.FirstName,
		MiddleName:       l.MiddleName,
		LastName:         l.LastName,
		Gender:           l.Gender,
		Address:          l.Address,
		Phone:            l.Phone,
		PracticingStatus: l.PracticingStatus,
		ProfileURL:       l.ProfileURL,
		AverageRating:    l.AverageRating,
		ReviewCount:      l.ReviewCount,
	}

	for _, j := range l.Edges.Jurisdictions {
		response.Jurisdictions = append(response.Jurisdictions, j.Name)
	}

	for _, d := range l.Edges.Decisions {
		response.Decisions = append(response.Decisions, d.URL)
	}

	return c.JSON(http.StatusOK, response)
}

func init() {
	Register(new(LawyerAPI))
}
