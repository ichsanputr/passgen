package passgen

import "errors"

var (
	count int
)

func ValidatePass(input int, l bool, u bool, n bool, s bool) error {

	if l {
		count++
	}
	if u {
		count++
	}
	if n {
		count++
	}
	if s {
		count++
	}

	if input == 0 || count == 0 || input < count {
		// return error with erros.New
		// errors.New for return error with string
		// 2 option for type error = nil or error.New
		// except use error.New can use fmt.Errorf()
		return errors.New("length must not be less than the number of requirements || length and terms cannot be empty")
	}
	return nil
}
