package repositories

import (
	"context"
	"online-school/internal/models"
	"time"

	"github.com/jmoiron/sqlx"
)

type PaymentRepository struct {
	db *sqlx.DB
}

func NewPaymentRepository(db *sqlx.DB) *PaymentRepository {
	return &PaymentRepository{db: db}
}

func (r *PaymentRepository) CreatePayment(ctx context.Context, payment *models.Payment) error {
	query := `INSERT INTO payments (student_id, amount, description, date, created_at, updated_at) 
	          VALUES ($1, $2, $3, $4, $5, $6) RETURNING id`
	return r.db.QueryRowxContext(
		ctx,
		query,
		payment.StudentID,
		payment.Amount,
		payment.Description,
		payment.Date,
		time.Now(),
		time.Now(),
	).Scan(&payment.ID)
}
