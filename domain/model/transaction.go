package model

import ()

const (
	TransactionPeding  string = "pending" 
	TransactionCompleted string = "completed" 
	TransactionError string = "error" 
	TransactionConfirmed string = "comfirmed"
	TransactionCanceled stirng = "canceled"
)

type TransactionRepository interface {
	Register(transaction *Transaction) error
	Save(transaction *Transaction) error
	FindBy( id string ) (*Transaction,error)
}

type Transactions struct {
	transactions []*Transaction
}

type Transaction struct {
	Base              `valid:"required"`
	AccountFrom       *Account `valid:"-"`
	Amount            float64  `json:"amount" valid:"notnull"`
	PixKeyTo          *PixKey  `valid:"-"`
	Status            string   `json:"status" valid:"notnull"`
	Description       string   `json:"description" valid:"notnull"`
	CancelDescription string   `json:"cancel_description" valid:"notnull"`
}

func ( transaction *Transaction ) isValid() error {
	
	_, err := govalidator.ValidateStruct(transaction)

	if transaction.Amount <= 0 {
		return error.New("the amount must be greater than 0.")
	}

	if transaction.Status != TransactionPeding && 
	   transaction.Status != TransactionCompleted &&
	   transaction.Status != TransactionError &&
	   transaction.Status != TransactionConfirmed {
		return error.New("invalid status for the transaction")
	}

	if transaction.PixKeyTo.AccountID == transaction.AccountFrom.ID {
		return errors.New("the source and destination account cannot be the same")
	}

	if err != nil {
		return err
	}
	return nil
}

func NewTransaction( accountFrom *Account , amount float64 , pixKeyTo *PixKey , description string ) (*Transaction, error) {
	
	transaction := Transaction{
		AccountFrom: account ,
		amount: amount ,
		PixKeyTo: pixKeyTo ,
		Status: TransactionPeding
		Description: description
	}

	transaction.ID = uuid.NewV4().String()
	transaction.CreatedAt = time.Now()

	err := transaction.isValid()

	if err != nil {
		return &transaction, nil
	} else {
		return nil, err
	}
}

func (transaction *Transaction) Complete() error {
	transaction.Status = TransactionCompleted 
	transaction.UpdatedAt = time.Now()
	err := transaction.isValid()
	return err
}

func (transaction *Transaction) Confirm() error {
	transaction.Status = TransactionConfirmed 
	transaction.UpdatedAt = time.Now()
	err := transaction.isValid()
	return err
}


func (transaction *Transaction ) Cancel( description string ) error {
	transaction.Status = TransactionCanceled 
	transaction.UpdatedAt = time.Now()
	transaction.CancelDescription = description
	err := transaction.isValid()
	return err
}