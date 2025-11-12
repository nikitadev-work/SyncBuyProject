package grpcserver

func CheckTitle(in string) error {
	if in != "" {
		return nil
	}
	return ErrIncorrectTitle
}

func CheckDescription(in string) error {
	return nil
}

func CheckCurrency(in string) error {
	supportedCur := []string{"RUB"}
	for _, v := range supportedCur {
		if v == in {
			return nil
		}
	}
	return ErrIncorrectCurrency
}

func CheckToken(in string) error {
	if in != "" {
		return nil
	}
	return ErrIncorrectToken
}

func CheckAmount(in int64) error {
	if in >= 0 {
		return nil
	}
	return ErrIncorrectAmount
}
