package usecase

type PurchaseUsecaseInterface interface {
	// Purchase
	CreatePurchase(CreateInviteInputDTO) (*CreateInviteOutputDTO, error)
	GetPurchase(GetPurchaseInputDTO) (*GetPurchaseOutputDTO, error)

	// Participant
	CreateInvite(CreateInviteInputDTO) (*CreateInviteOutputDTO, error)
	JoinByInvite(JoinByInviteInputDTO) error
	RemoveParticipant(RemoveParticipantInputDTO) error
	ListParticipantsByPurchaseId(ListParticipantsByPurchaseIdInputDTO) (*ListParticipantsByPurchaseIdOutputDTO, error)

	// Task
	CreateTask(CreateInviteInputDTO) (*CreateInviteOutputDTO, error)
	TakeTask(TakeTaskInputDTO) error
	DeleteTask(DeleteTaskInputDTO) error
	ListTasksByPurchaseId(ListTasksByPurchaseIdInputDTO) (*ListTasksByPurchaseIdOutputDTO, error)
	MarkTaskAsDone(MarkTaskAsDoneInputDTO) error

	// Status
	LockPurchase(LockPurchaseInputDTO) error
	UnlockPurchase(UnlockPurchaseInputDTO) error
	GetSnapshot(GetSnapshotInputDTO) (*GetSnapshotOutputDTO, error)
	MarkSettlementInitiated(MarkSettlementInitiatedInputDTO) error
	FinishPurchase(FinishPurchaseInputDTO) error
}

type PurchaseUsecase struct {
	Repository PurchaseUsecaseInterface
	TxManager  TxManagerInterface
}

func NewPurchaseUsecase(repository PurchaseRepositoryInterface, txManager TxManagerInterface) *PurchaseUsecase {
	return &PurchaseUsecase{}
}

// Purchase
func (uc *PurchaseUsecase) CreatePurchase(input CreateInviteInputDTO) (*CreateInviteInputDTO, error) {
	//TODO
	return nil, nil
}

func (uc *PurchaseUsecase) GetPurchase(input GetPurchaseInputDTO) (*GetPurchaseOutputDTO, error) {
	//TODO
	return nil, nil
}

// Participant
func (uc *PurchaseUsecase) CreateInvite(input CreateInviteInputDTO) (*CreateInviteOutputDTO, error) {
	//TODO
	return nil, nil
}

func (uc *PurchaseUsecase) JoinByInvite(input JoinByInviteInputDTO) error {
	//TODO
	return nil
}

func (uc *PurchaseUsecase) RemoveParticipant(input RemoveParticipantInputDTO) error {
	//TODO
	return nil
}

func (uc *PurchaseUsecase) ListParticipantsByPurchaseId(input ListParticipantsByPurchaseIdInputDTO) (*ListParticipantsByPurchaseIdOutputDTO, error) {
	//TODO
	return nil, nil
}

// Task
func (uc *PurchaseUsecase) CreateTask(input CreateTaskInputDTO) (*CreateInviteOutputDTO, error) {
	//TODO
	return nil, nil
}

func (uc *PurchaseUsecase) TakeTask(input TakeTaskInputDTO) error {
	//TODO
	return nil
}

func (uc *PurchaseUsecase) DeleteTask(input DeleteTaskInputDTO) error {
	//TODO
	return nil
}

func (uc *PurchaseUsecase) ListTasksByPurchaseId(input ListTasksByPurchaseIdInputDTO) (*ListTasksByPurchaseIdOutputDTO, error) {
	//TODO
	return nil, nil
}

func (uc *PurchaseUsecase) MarkTaskAsDone(input MarkTaskAsDoneInputDTO) error {
	//TODO
	return nil
}

// Status
func (uc *PurchaseUsecase) LockPurchase(input LockPurchaseInputDTO) error {
	//TODO
	return nil
}

func (uc *PurchaseUsecase) UnlockPurchase(input LockPurchaseInputDTO) error {
	//TODO
	return nil
}

func (uc *PurchaseUsecase) GetSnapshot(input GetSnapshotInputDTO) (*GetSnapshotOutputDTO, error) {
	//TODO
	return nil, nil
}

func (uc *PurchaseUsecase) MarkSettlementInitiated(input MarkSettlementInitiatedInputDTO) error {
	//TODO
	return nil
}

func (uc *PurchaseUsecase) FinishPurchase(input FinishPurchaseInputDTO) error {
	//TODO
	return nil
}
