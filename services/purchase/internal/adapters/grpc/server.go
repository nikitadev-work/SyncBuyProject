package gprcserver

import (
	uc "purchase/internal/usecase"
	pb "purchase/proto-codegen"

	"github.com/nikitadev-work/SyncBuyProject/common/kit/logger"
)

type PurchaseServer struct {
	pb.UnimplementedPurchaseServiceServer
	usecase uc.PurchaseUsecase
	logger  logger.LoggerInterface
}

func New(usecase uc.PurchaseUsecase, logger logger.LoggerInterface) *PurchaseServer {
	return &PurchaseServer{
		usecase: usecase,
		logger:  logger,
	}
}
