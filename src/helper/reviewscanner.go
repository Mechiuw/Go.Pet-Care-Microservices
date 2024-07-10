package helper

import (
	"database/sql"
	"ginpet/src/model"
)

func ScanReview(rows *sql.Rows) []model.Review {
	// clients instance, error instance
	reviews := []model.Review{}
	var err error

	// scan each property of the struct
	for rows.Next() {
		review := model.Review{}
		err = rows.Scan(
			&review.Id,
			&review.AppointmentId,
			&review.Rating,
			&review.ReviewText,
		)

		//handle the error
		if err != nil {
			panic("failed to scan row: " + err.Error())
		}

		reviews = append(reviews, review)
	}

	// handle the error and the exceptions for row iterations
	err = rows.Err()
	if err != nil {
		panic("error occurred during row iteration: " + err.Error())
	}

	//return the value
	return reviews
}
