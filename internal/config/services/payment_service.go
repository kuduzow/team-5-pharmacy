package services

import (
	"errors"

	"github.com/kuduzow/team-5-pharmacy/internal/config/repository"
	"github.com/kuduzow/team-5-pharmacy/internal/models"
	"gorm.io/gorm"
)

var (
	ErrPaymentNotFound = errors.New("платеж не найден")
)

type PaymentService interface {
	CreatePayment(req models.CreatePaymentRequest) (*models.Payment, error)
	UpdatePayment(id uint, req *models.UpdatePaymentRequest) (*models.Payment, error)
	GetPaymentByID(id uint) (*models.Payment, error)
	DeletePayment(id uint) error
}

type paymentService struct {
	repo repository.PaymentRepository
}

func NewPaymentService(paymentRepo repository.PaymentRepository) PaymentService {
	return &paymentService{repo: paymentRepo}
}

func (r *paymentService) CreatePayment(req models.CreatePaymentRequest) (*models.Payment, error) {
	if err := r.validatePaymentCreate(&req); err != nil {
		return nil, err
	}

	payment := models.Payment{Amount: req.Amount, Status: req.Status, Method: req.Method}

	if err := r.repo.Create(&payment); err != nil {
		return nil, err
	}
	return &payment, nil
}

func (r *paymentService) validatePaymentCreate(req *models.CreatePaymentRequest) error {
	if req.Amount <= 0 {
		return errors.New("сумма не может быть отрицательной")
	}

	switch req.Status {
	case models.PaymentStatusFailed, models.PaymentStatusPending, models.PaymentStatusSuccess:
	default:
		return errors.New("некорректный статус")
	}

	switch req.Method {
	case models.PaymentMethodCard, models.PaymentMethodCash, models.PaymentMethodBankTransfer:
	default:
		return errors.New("некорректный метод")
	}

	return nil

}

func (r *paymentService) applyPaymentValidate(req *models.UpdatePaymentRequest) error {
	if req.Amount != nil {
		if *req.Amount <= 0 {
			return errors.New("сумма не может меньше или равна нулю")
		}
	}

	if req.Status != nil {
		switch *req.Status {
		case models.PaymentStatusFailed, models.PaymentStatusPending, models.PaymentStatusSuccess:
		default:
			return errors.New("некорректный статус")
		}

	}

	if req.Method != nil {
		switch *req.Method {
		case models.PaymentMethodBankTransfer, models.PaymentMethodCard, models.PaymentMethodCash:
		default:
			return errors.New("некорректный метод")
		}
	}
	return nil
}

func (r *paymentService) UpdatePayment(id uint, req *models.UpdatePaymentRequest) (*models.Payment, error) {
	if err := r.applyPaymentValidate(req); err != nil {
		return nil, err
	}

	payment, err := r.repo.GetById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPaymentNotFound
		}
		return nil, err
	}

	if req.Amount != nil {
		payment.Amount = *req.Amount
	}

	if req.Method != nil {
		payment.Method = *req.Method
	}

	if req.Status != nil {
		payment.Status = *req.Status
	}

	if err := r.repo.Update(payment); err != nil {
		return nil, err
	}

	return payment, nil

}

func (r *paymentService) DeletePayment(id uint) error {
	if _, err := r.repo.GetById(id); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ErrPaymentNotFound
		}
		return err
	}

	return r.repo.Delete(id)
}

func (r *paymentService) GetPaymentByID(id uint) (*models.Payment, error) {
	payment, err := r.repo.GetById(id)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrPaymentNotFound
		}
		return nil, err
	}

	return payment, nil
}
