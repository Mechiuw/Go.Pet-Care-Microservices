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

	fmt.Println("successfully create review")
	return reviewResponse, nil
}

func UPDATE_REVIEW(id string, review model.Review) (model.Review, error) {
	sqlStatement := `UPDATE review SET appointmentid=$1 , rating=$2, reviewText=$3 WHERE id=$4`

	err := validator.ValidateReview(review)
	if err != nil {
		return model.Review{}, fmt.Errorf("validation error: %w", err)
	}

	_, err = review_connection.Exec(sqlStatement, review.Id, review.Rating, review.ReviewText, id)
	if err != nil {
		return model.Review{}, fmt.Errorf("failed to update review: %w", err)
	}

	reviewResponse := model.Review{
		Id:            review.Id,
		AppointmentId: review.AppointmentId,
		Rating:        review.Rating,
		ReviewText:    review.ReviewText,
	}
	fmt.Println("successfully update review")
	return reviewResponse, nil
}

func DELETE_REVIEW(id string) (model.Review, error) {
	sqlStatement := `DELETE FROM review WHERE id=$1`

	_, err := review_connection.Exec(sqlStatement, id)
	if err != nil {
		return model.Review{}, fmt.Errorf("failed to delete review: %w", err)
	}

	fmt.Println("successfully delete review")
	return model.Review{}, nil
}

func GET_BY_ID_REVIEW(id string) (model.Review, error) {
	sqlStatement := `SELECT * FROM review WHERE id=$1`
	fetchedReview := model.Review{}

	err := connection.QueryRow(sqlStatement, id).Scan(&fetchedReview.Id, &fetchedReview.AppointmentId, &fetchedReview.Rating, &fetchedReview.ReviewText)
	if err != nil {
		return model.Review{}, fmt.Errorf("failed to fetch review: %w", err)
	}

	fmt.Println("successfully delete review")
	return fetchedReview, nil
}

func GET_ALL_REVIEW() ([]model.Review, error) {
	sqlStatement := `SELECT * FROM review`

	rows, err := review_connection.Query(sqlStatement)

	if err != nil {
		return []model.Review{}, fmt.Errorf("failed to fetch review: %w", err)
	}

	reviews := helper.ScanReview(rows)

	defer rows.Close()
	return reviews, nil
}
