package order

import (
	"context"
	orderDomain "github.com/poster-keisuke/sample-clearn-architecture/app/domain/order"
	productDomain "github.com/poster-keisuke/sample-clearn-architecture/app/domain/product"
)

type CreteOrderUseCase struct {
	orderRepo   orderDomain.OrderRepository
	productRepo productDomain.ProductRepository
}

func NewCreteOrderUseCase(
	orderRepo orderDomain.OrderRepository,
	productRepo productDomain.ProductRepository,
) *CreteOrderUseCase {
	return &CreteOrderUseCase{
		orderRepo:   orderRepo,
		productRepo: productRepo,
	}
}

type ProductIDAndAmount struct {
	ID     string `json:"id"`
	Amount int    `json:"amount"`
}

type CreteOrderUseCaseInputDto struct {
	Products []*ProductIDAndAmount `json:"products"`
}

func (uc *CreteOrderUseCase) Run(ctx context.Context, input CreteOrderUseCaseInputDto) error {
	//newOrder := orderDomain.NewOrder()
	//
	//ids := make([]string, 0, len(input.Products))
	//for _, p := range input.Products {
	//	ids = append(ids, p.ID)
	//}
	//products, err := uc.productRepo.GetMultiByIDs(ctx, ids)
	//if err != nil {
	//	return xerrors.Errorf(": %w", err)
	//}

	return nil
}
