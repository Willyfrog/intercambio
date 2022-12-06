package algorithm

import (
	"math/rand"

	"github.com/willyfrog/intercambio/input"
	"github.com/willyfrog/intercambio/output"
)

// TODO: tests
func PopElement(slice []*input.Row, index int) ([]*input.Row, *input.Row) {
	element := slice[index]
	if index < len(slice)-1 {
		return append(slice[:index], slice[index+1:]...), element
	}
	return slice[:index], element
}

func getOne(slice []*input.Row) int {
	return rand.Intn(len(slice) - 1)
}

//TODO: test if some fields are empty
func Match(sender *input.Row, receiver *input.Row) output.Match {
	return output.Match{
		SenderEmail:   sender.Email,
		SenderName:    sender.Name,
		ReceiverEmail: receiver.Email,
		ReceiverName:  receiver.Name,
		Address1:      receiver.Address1,
		Address2:      receiver.Address2,
		City:          receiver.City,
		Zip:           receiver.Zip,
		Province:      receiver.Province,
		Country:       receiver.Country,
		Phone:         receiver.Phone,
	}
}
