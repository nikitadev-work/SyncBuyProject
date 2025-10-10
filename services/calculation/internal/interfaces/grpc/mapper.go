package grpcserver

import (
	dm "calculation/internal/domain"
	uc "calculation/internal/usecase"
	pb "calculation/proto-codegen/calculation"

	"github.com/google/uuid"
)

func ConvertCalcReqDataToInternalDTO(input *pb.CalcRequest) (*uc.CalculationInputDTO, error) {
	participants := []dm.Participant{}
	for _, v := range input.Participants {
		id, err := uuid.Parse(v.UserId)
		if err != nil {
			return nil, ErrToStatus(err)
		}
		newParticipant := dm.NewParticipant(id)

		participants = append(participants, *newParticipant)
	}

	expenses := []dm.Expense{}
	for _, v := range input.Expenses {
		id, _ := uuid.Parse(v.Id)
		authorID, err := uuid.Parse(v.AuthorId)
		if err != nil {
			return nil, ErrToStatus(err)
		}
		money, err := dm.NewMoney(v.Cost.Amount, v.Cost.Currency)
		if err != nil {
			return nil, err
		}
		newExpense := dm.NewExpense(id, authorID, money)

		expenses = append(expenses, *newExpense)
	}

	return &uc.CalculationInputDTO{
		Participants: participants,
		Expenses:     expenses,
	}, nil
}

func ConvertCalcFromInternalDTOtoResp(input *uc.CalculationOutputDTO) (*pb.CalcResponse, error) {
	intents := []*pb.Intent{}
	for _, v := range input.Intents {
		payerId := v.PayerId
		payeeId := v.PayeeId
		money := pb.Money{
			Currency: v.Money.Currency.Code,
			Amount:   v.Money.Amount,
		}

		newIntent := pb.Intent{
			PayerId: payerId.String(),
			PayeeId: payeeId.String(),
			Amount:  &money,
		}

		intents = append(intents, &newIntent)
	}

	return &pb.CalcResponse{
		Intents: intents,
	}, nil
}
