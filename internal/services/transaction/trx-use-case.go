package transaction

import (
	"context"
	"myAPI/internal/adapter/Repository"
	"myAPI/internal/domains"
	"time"
)

type(
	UseCase struct{
		TrxRepo Repository.TrxRepoInterface
	}

	UsecaseInterface interface {
		SendTrx(
			ctx context.Context,
			payload SendTrx,
		) (result UseCaseSendResult, err error)

		GetLogTrx(
			ctx context.Context,
		) ([]domains.Transaction, error)
	}
)

func (uc UseCase) SendTrx(
	ctx context.Context,
	payload SendTrx,
) (result UseCaseSendResult, err error) {
	trx, err := uc.TrxRepo.SendTrx(
		ctx,
		&domains.Transaction{
			Title : payload.Title,
			Ammount : payload.Ammount,
			Description : payload.Description,
			Date : time.Now(),
		})
	result.Trx = SendTrx{
		Title : trx.Title,
		Ammount : trx.Ammount,
		Description : trx.Description,
		Date : trx.Date,
	}
	return result, err
}

func (uc UseCase) GetLogTrx(
	ctx context.Context,
) ([]domains.Transaction, error) {
	return uc.TrxRepo.GetLogTrx(ctx)
}
