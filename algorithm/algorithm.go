package algorithm

import (
	"errors"
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

func random(slice []*input.Row) int {
	return rand.Intn(len(slice))
}

//TODO: test if some fields are empty
func GenMatch(sender *input.Row, receiver *input.Row) *output.Match {
	return &output.Match{
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

func Run(inputRows []*input.Row) ([]*output.Match, error) {
	if len(inputRows) <= 1 {
		return nil, errors.New("Number of elements is too low, we need at least 2")
	}
	results := []*output.Match{}
	original := inputRows[0]
	sender := original
	list := inputRows[1:]
	var receiver *input.Row
	for len(list) > 1 {
		list, receiver = PopElement(list, random(list))
		results = append(results, GenMatch(sender, receiver))
		sender = receiver
	}
	results = append(results, GenMatch(sender, original))
	return results, nil
}
