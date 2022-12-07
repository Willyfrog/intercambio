package algorithm

import (
	"fmt"
	"testing"

	"github.com/willyfrog/intercambio/input"
)

func genList() []*input.Row {
	list1 := make([]*input.Row, 3)
	list1[0] = &input.Row{Name: "yo", Email: "yo@example.com"}
	list1[1] = &input.Row{Name: "tu", Email: "tu@example.com"}
	list1[2] = &input.Row{Name: "el", Email: "el@example.com"}

	return list1
}

func fmtDisplayRow(slice []*input.Row) []string {
	results := make([]string, len(slice))
	for i, e := range slice {
		results[i] = e.Name
	}
	return results
}

func TestPopElement(t *testing.T) {
	list1 := genList()
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
