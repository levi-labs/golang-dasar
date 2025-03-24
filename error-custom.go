package main

import "fmt"

type validationError struct {
	Message string
}
type notFoundError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"id is empty validation error"}
	}
	if id != "john" {
		return &notFoundError{"data is already in use"}
	}

	return nil
}
func main() {
	err := SaveData("", nil)
	if err != nil {
		//akan dilakukan type assertion mengecek apakah error tersebut merupakan instance Type of validation error? jika ya
		//maka variable ok akan bersifat boolean true, dan hasil dari type assertion yang merupakan objek dari interface Error
		//akan di tampung pada variable validationErr
		if validationErr, ok := err.(*validationError); ok {
			fmt.Println("validation error", validationErr.Message)
		} else if notFoundErr, ok := err.(*notFoundError); ok {
			fmt.Println("notFoundError", notFoundErr.Message)
		} else {
			fmt.Println("unknown error", err.Error())
		}

		switch finalError := err.(type) {
		case *validationError:
			fmt.Println("validation error", finalError.Message)
		case *notFoundError:
			fmt.Println("notFoundError", finalError.Message)
		default:
			fmt.Println("unknown error", err.Error())
		}
	}
}
