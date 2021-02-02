package usecase

import ()

type TransactionUseCase struct {
	TransactionRespository model.TransactionRespository
	PixKeyRepository       model.PixKeyRepository
}

func (t *TransactionUseCase) ( accountId string , account float64 , pixKeyTo string , pixKeyKindTo string , description string ) (*model.Transaction,error){

	//verifiy if account exists
	account, err := t.PixKeyRepository.FindAccount(accountId)

	if err != nil {
		return nil , err
	}

	//verifiy if pixkey exists
	pixKey, err := t.PixKeyRepository.FindKeyByKind(pixKeyTo,pixKeyKindTo)

	if err != nil {
		return nil, err
	}

	transaction, err := model.NewTransaction(account,amount,pixKey,description)

	if err != nil {
		return nil, err
	}

	t.TransactionRespository.Save(transaction)

	if transaction.ID == "" {
		return nil, errors.New("Unable to process this transaction")
	}

	return transaction, nil

}

func (t *Transaction) Confirm( transactionId string ) (*model.Transaction, error) {

	transaction, err := t.TransactionRespository.Find( transactionId )

	if err != nil {
		log.Println("Transaction not found", transactionId )
		return nil, err
	}

	transaction.Status = model.TransactionConfirmed
	err := t.TransactionRespository.Save(transaction)

	if err != nil {
		return nil, err
	}

	return transaction, nil

}

func (t *Transaction) Complete( transactionId string ) (*model.Transaction, error) {

	transaction, err := t.TransactionRespository.Find( transactionId )

	if err != nil {
		log.Println("Transaction not found", transactionId )
		return nil, err
	}

	transaction.Status = model.TransactionCompleted
	err := t.TransactionRespository.Save(transaction)

	if err != nil {
		return nil, err
	}

	return transaction, nil

}

func (t *Transaction) Error( transactionId string , reason string ) (*model.Transaction,error){

	transaction, err := t.TransactionRespository.Find( transactionId )

	if err != nil {
		return nil, err
	}

	transaction.Status = model.TransactionError
	transaction.CancelDescription = reason

	err = t.TransactionRespository.Save(transation)

	if err != nil {
		return nil, err
	}

	return transaction, nil

}