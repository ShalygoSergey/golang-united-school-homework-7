package coverage

import (
	"fmt"
	"os"
	"testing"
	"time"
)

// DO NOT EDIT THIS FUNCTION
func init() {
	content, err := os.ReadFile("students_test.go")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile("autocode/students_test", content, 0644)
	if err != nil {
		panic(err)
	}
}

// WRITE YOUR CODE BELOW

func TestPersonLen(t *testing.T) {
	expected := 1
	people := People([]Person{{
		firstName: "firstName",
		lastName:  "lastName",
		birthDay:  time.Time{},
	}})

	length := People.Len(people)
	if length != expected {
		t.Errorf("Expected: %d, got %d", expected, length)
	}
}

func TestPersonLess(t *testing.T) {
	birthDay1 := time.Date(2021, 12, 12, 1, 1, 1, 1, time.UTC)
	birthDay2 := time.Date(2020, 10, 11, 6, 5, 4, 2, time.UTC)
	people := People([]Person{{
		firstName: "firstName1",
		lastName:  "lastName1",
		birthDay:  birthDay1,
	}, {
		firstName: "firstName1",
		lastName:  "lastName1",
		birthDay:  birthDay1,
	}, {
		firstName: "firstName2",
		lastName:  "lastName1",
		birthDay:  birthDay1,
	}, {
		firstName: "firstName1",
		lastName:  "lastName2",
		birthDay:  birthDay2,
	}})

	tData := map[string]struct {
		People   People
		I        int
		J        int
		Expected bool
	}{
		"1": {People: people, I: 0, J: 1, Expected: false},
		"2": {People: people, I: 0, J: 2, Expected: true},
		"3": {People: people, I: 0, J: 3, Expected: true}}
	for name, tcase := range tData {
		t.Run(name, func(t *testing.T) {
			got := tcase.People.Less(tcase.I, tcase.J)
			if got != tcase.Expected {
				t.Errorf("[%s] expected: %t, got %t", name, tcase.Expected, got)
			}
		})
	}
}

func TestPersonSwap(t *testing.T) {
	people := People([]Person{{
		firstName: "firstName1",
		lastName:  "lastName",
		birthDay:  time.Time{},
	}, {
		firstName: "firstName2",
		lastName:  "lastName",
		birthDay:  time.Time{},
	}})

	if people.Less(1, 0) {
		t.Error("Swap error 1")
	}

	people.Swap(0, 1)

	if people.Less(0, 1) {
		t.Error("Swap error 2")
	}
}

func TestMatrixNew(t *testing.T) {
	tData := map[string]struct {
		Str  string
		Rows int
		Cols int
		Err  error
	}{
		"1": {Str: "1 2\n2 3", Rows: 2, Cols: 2, Err: nil},
		"2": {Str: "1\n2 3", Rows: 0, Cols: 0, Err: fmt.Errorf("")},
		"3": {Str: "1 A\n2 3", Rows: 0, Cols: 0, Err: fmt.Errorf("")}}
	for name, tcase := range tData {
		t.Run(name, func(t *testing.T) {
			got, err := New(tcase.Str)

			if err != nil && tcase.Err == nil {
				t.Errorf("[%s] error happend while not expected: %s", name, err.Error())
			}

			if got != nil && got.rows != tcase.Rows {
				t.Errorf("[%s] rows expected: %d, got %d", name, tcase.Rows, got.rows)
			}

			if got != nil && got.cols != tcase.Cols {
				t.Errorf("[%s] cols expected: %d, got %d", name, tcase.Cols, got.cols)
			}
		})
	}
}

func TestMatrixRows(t *testing.T) {

	got, _ := New("1 2\n3 4")
	rows := got.Rows()

	if rows[0][0] != 1 ||
		rows[0][1] != 2 ||
		rows[1][0] != 3 ||
		rows[1][1] != 4 {
		t.Error("Rows error")
	}
}

func TestMatrixCols(t *testing.T) {

	got, _ := New("1 2\n3 4")
	cols := got.Cols()

	if cols[0][0] != 1 ||
		cols[0][1] != 3 ||
		cols[1][0] != 2 ||
		cols[1][1] != 4 {
		t.Error("Cols error")
	}
}

func TestMatrixSet(t *testing.T) {

	got, _ := New("1 2\n3 4")

	if !got.Set(0, 0, 0) {
		t.Error("Set error 1")
	}

	if got.Set(8, 8, 0) {
		t.Error("Set error 2")
	}
}
