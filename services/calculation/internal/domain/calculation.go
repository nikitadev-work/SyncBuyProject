package domain

// Participants - all of the users for current purchase
func CalculateIntents(participants []Participant, expenses []Expense) ([]Intent, error) {
	for i := 1; i < len(expenses); i++ {
		if expenses[i].Money.Currency.Code != expenses[i-1].Money.Currency.Code {
			return nil, ErrCurrenciesInListDiffers
		}
	}

	for i := 0; i < len(expenses); i++ {
		if expenses[i].Money.Amount < 0 {
			return nil, ErrNegativeAmount
		}
	}

	if len(participants) == 0 {
		return nil, ErrEmptyParticipantsList
	}

	if len(expenses) == 0 {
		return nil, ErrEmptyExpensesList
	}

	if len(participants) == 1 {
		return nil, ErrOnlyOneParticipant
	}

	intents := []Intent{}
	amountOfParticipants := len(participants)

	for _, expense := range expenses {
		curCost := expense.Money.Amount
		base, remainder, err := DivideIntoNParts(curCost, amountOfParticipants-1) //Base - the amount for each participant, r - amount of participants, who gets 1 from reminder
		if err != nil {
			return nil, err
		}

		for _, participant := range participants {
			if participant.UserId == expense.AuthorId {
				continue
			}

			remAdd := 0
			if remainder > 0 {
				remAdd += 1
				remainder--
			}

			money, err := NewMoney(base+int64(remAdd), expense.Money.Currency.Code)
			if err != nil {
				return nil, err
			}

			newIntent, err := NewIntent(participant.UserId, expense.AuthorId, money)
			if err != nil {
				return nil, err
			}

			intents = append(intents, newIntent)
		}
	}

	return intents, nil
}
