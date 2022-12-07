package algorithm

import (
	"fmt"
	"testing"

	"github.com/willyfrog/intercambio/input"
	"github.com/willyfrog/intercambio/output"
)

func genList(size int) []*input.Row {
	list1 := make([]*input.Row, size)
	list1[0] = &input.Row{Name: "yo", Email: "yo@example.com"}
	if size > 1 {
		list1[1] = &input.Row{Name: "tu", Email: "tu@example.com"}
	}
	if size > 2 {
		list1[2] = &input.Row{Name: "el", Email: "el@example.com"}
	}
	if size > 3 {
		list1[3] = &input.Row{Name: "nosotros", Email: "nosotros@example.com"}
	}
	if size > 4 {
		list1[4] = &input.Row{Name: "vosotros", Email: "vosotros@example.com"}
	}

	return list1
}

func fmtDisplayRow(slice []*input.Row) []string {
	results := make([]string, len(slice))
	for i, e := range slice {
		results[i] = e.Name
	}
	return results
}

func fmtDisplayMatches(slice []*output.Match) []string {
	results := make([]string, len(slice))
	for i, e := range slice {
		results[i] = fmt.Sprintf("%v -> %v", e.SenderEmail, e.ReceiverEmail)
	}
	return results
}

func TestPopElement(t *testing.T) {
	list1 := genList(3)
	for i, testElement := range list1 {
		t.Run(fmt.Sprintf("Pop %v", i), func(t *testing.T) {
			t.Logf("[OK] list1 - pre %v", fmtDisplayRow(list1))
			resultList, element := PopElement(list1, i)
			if len(resultList) != 2 {
				t.Errorf("%v element wasn't poped out", i)
			}
			if element.Name != testElement.Name {
				t.Errorf("%v element wasn't the right one %v != %v", i, element.Name, testElement.Name)
				t.Errorf("resultLists %v", fmtDisplayRow(resultList))
				t.Errorf("list1 %v", fmtDisplayRow(list1))
			} else {
				t.Logf("[OK] resultLists %v", fmtDisplayRow(resultList))
				t.Logf("[OK] list1 %v", fmtDisplayRow(list1))
				t.Logf("[OK] Element %v", element.Name)
			}
		})
	}
}

func TestRun(t *testing.T) {
	list1 := genList(1)
	_, err := Run(list1)
	if err == nil {
		t.Errorf("No error when running with a single element")
	}
	list2 := genList(2)
	res, err := Run(list2)
	if err != nil {
		t.Errorf("Error when running with two  elements")
	}
	if len(res) != 2 {
		t.Errorf("there should be 2 elements in the end result: %v", fmtDisplayMatches(res))
	}
	list5 := genList(5)
	res, err = Run(list5)
	if err != nil {
		t.Errorf("Error when running with two  elements")
	}
	if len(res) != 5 {
		t.Errorf("there should be 5 elements in the end result: %v", fmtDisplayMatches(res))
	}
	for _, e := range list5 {
		foundSender := false
		foundReceiver := false
		for _, eTest := range res {
			foundSender = foundSender || e.Email == eTest.SenderEmail
			foundReceiver = foundReceiver || e.Email == eTest.ReceiverEmail
		}
		if !foundReceiver {
			t.Errorf("there isn't a receiver for %v: %v", e.Email, fmtDisplayMatches(res))
		}
		if !foundSender {
			t.Errorf("there isn't a sender for %v: %v", e.Email, fmtDisplayMatches(res))
		}
	}
}
