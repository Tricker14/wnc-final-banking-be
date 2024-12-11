package serviceimplement

import (
	"errors"
	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/domain/model"
	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/repository"
	"github.com/21CLC01-WNC-Banking/WNC-Banking-BE/internal/service"
	"github.com/gin-gonic/gin"
)

type AccountService struct {
	accountRepository  repository.AccountRepository
	customerRepository repository.CustomerRepository
	CoreService        service.CoreService
}

func NewAccountService(accountRepo repository.AccountRepository, customerRepo repository.CustomerRepository, coreService service.CoreService) service.AccountService {
	return &AccountService{accountRepository: accountRepo, customerRepository: customerRepo, CoreService: coreService}
}

func (service *AccountService) InternalTransfer(ctx *gin.Context, transferReq model.InternalTransferRequest) error {
	//get customer and check info
	customerId, exists := ctx.Get("customerId")
	if !exists {
		return errors.New("customer not exists")
	}

	//check customerId
	sourceCustomer, err := service.customerRepository.GetOneByIdQuery(ctx, customerId.(int64))
	if err != nil {
		return err
	}
	if sourceCustomer == nil {
		return errors.New("customer not exists")
	}

	//get account by customerId and check sourceNumber
	sourceAccount, err := service.accountRepository.GetOneByCustomerIdQuery(ctx, customerId.(int64))
	if err != nil {
		return err
	}
	if sourceAccount == nil {
		return errors.New("source account not exists")
	}
	if sourceAccount.Number != transferReq.SourceAccountNumber {
		return errors.New("source account not match")
	}

	//check targetNumber
	targetAccount, err := service.accountRepository.GetOneByNumberQuery(ctx, transferReq.TargetAccountNumber)
	if err != nil {
		return err
	}
	if targetAccount == nil {
		return errors.New("target account not exists")
	}
	//estimate fee
	fee, err := service.CoreService.EstimateTransferFee(ctx, transferReq.Amount)
	if err != nil {
		return err
	}

	//check is source fee and change balance
	checkFee := *transferReq.IsSourceFee
	if checkFee {
		totalDeduction := transferReq.Amount + fee
		if sourceAccount.Balance < totalDeduction {
			return errors.New("insufficient balance in source account")
		}
		sourceAccount.Balance -= totalDeduction
		targetAccount.Balance += transferReq.Amount
	} else {
		if sourceAccount.Balance < transferReq.Amount {
			return errors.New("insufficient balance in source account")
		}
		sourceAccount.Balance -= transferReq.Amount
		targetAccount.Balance += transferReq.Amount - fee
	}

	//update to DB
	err = service.accountRepository.UpdateCommand(ctx, *sourceAccount)
	if err != nil {
		return err
	}
	err = service.accountRepository.UpdateCommand(ctx, *targetAccount)
	if err != nil {
		return err
	}

	//Save to history, notify, response history

	return nil
}
