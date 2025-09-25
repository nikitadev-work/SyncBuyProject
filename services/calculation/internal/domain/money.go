package domain

type Money struct {
	Amount   int64
	Currency Currency
}

func NewMoney(amount int64, code string) (Money, error) {
	if amount < 0 {
		return Money{}, ErrNegativeAmount
	}

	if !IsSupportedCurrencyFormat(code) {
		return Money{}, ErrIncorrectCurrencyFormat
	}

	newCurr, err := NewCurrency(code)
	if err != nil {
		return Money{}, err
	}

	return Money{Amount: amount, Currency: newCurr}, nil
}

func (m1 *Money) Add(m2 Money) (Money, error) {
	if m1.Currency.Code != m2.Currency.Code {
		return Money{}, ErrCurrencyMismatch
	}

	return Money{Amount: m1.Amount + m2.Amount, Currency: m1.Currency}, nil
}

func (m1 *Money) Sub(m2 Money) (Money, error) {
	if m1.Currency.Code != m2.Currency.Code {
		return Money{}, ErrCurrencyMismatch
	}

	res := m1.Amount - m2.Amount

	if res < 0 {
		return Money{}, ErrNegativeSubResult
	}

	return Money{Amount: res, Currency: m1.Currency}, nil
}

func (m1 *Money) Cmp(m2 Money) (int, error) {
	if m1.Currency.Code != m2.Currency.Code {
		return 0, ErrCurrencyMismatch
	}

	if m1.Amount == m2.Amount {
		return 0, nil
	} else if m1.Amount > m2.Amount {
		return 1, nil
	} else {
		return -1, nil
	}
}

func DivideIntoNParts(amount int64, n int) (int64, int64, error) {
	if n <= 0 {
		return 0, 0, ErrIncorrectPartsAmount
	}

	base := amount / int64(n)
	r := amount % int64(n)
	return base, r, nil
}
