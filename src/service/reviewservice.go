package service

import (
	"fmt"
	"ginpet/src/db"
	"ginpet/src/helper"
	"ginpet/src/model"
)

var review_connection = db.Pool()
var review_validator = helper.NewValidator()

func CREATE_REVIEW(review model.Review) (model.Review, error) {
	sqlStatement := `INSERT INTO review (Id,appointmentid,rating,reviewtext) VALUES ($1,$2,$3,$4)`

	err := review_validator.ValidateReview(review)
	if err != nil {
		return model.Review{}, fmt.Errorf("validation error: %w", err)
	}

	_, err = review_connection.Exec(sqlStatement, review.Id, review.AppointmentId, review.Rating, review.ReviewText)
	if err != nil {
		return model.Review{}, fmt.Errorf("failed to create review: %w", err)
	}

	reviewResponse := model.Review{
		Id:            review.Id,
		AppointmentId: review.AppointmentId,
		Rating:        review.Rating,
		ReviewText:    review.ReviewText,
	}
	return reviewResponse, nil
}
