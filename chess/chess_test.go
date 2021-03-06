package chess

import (
	"testing"
)

func TestBoard_Put(t *testing.T) {
	board := NewBoard()
	_, err := board.Put([2]int{0, 0}, KNIGHT)
	if err != nil {
		t.Fatalf("Expected empty cell %s%d", LETTERS[0], 1)
	}
}

func TestValidPos(t *testing.T) {
	_, err := ValidPos([2]int{-1, 7})
	if err == nil || err.Error() != "invalid position" {
		t.Fatalf("Expected 'invalid posotion' for [%d; %d]", -1, 7)
	}
	_, err = ValidPos([2]int{7, -1})
	if err == nil || err.Error() != "invalid position" {
		t.Fatalf("Expected 'invalid posotion' for [%d; %d]", 7, -1)
	}
	_, err = ValidPos([2]int{-1, -1})
	if err == nil || err.Error() != "invalid position" {
		t.Fatalf("Expected 'invalid posotion' for [%d; %d]", -1, -1)
	}
	_, err = ValidPos([2]int{9, -1})
	if err == nil || err.Error() != "invalid position" {
		t.Fatalf("Expected 'invalid posotion' for [%d; %d]", 9, -1)
	}
	_, err = ValidPos([2]int{9, 1})
	if err == nil || err.Error() != "invalid position" {
		t.Fatalf("Expected 'invalid posotion' for [%d; %d]", 9, -1)
	}
	_, err = ValidPos([2]int{1, 9})
	if err == nil || err.Error() != "invalid position" {
		t.Fatalf("Expected 'invalid posotion' for [%d; %d]", 9, -1)
	}
	_, err = ValidPos([2]int{9, 9})
	if err == nil || err.Error() != "invalid position" {
		t.Fatalf("Expected 'invalid posotion' for [%d; %d]", 9, -1)
	}
	_, err = ValidPos([2]int{1, 1})
	if err != nil {
		t.Fatalf("Not expected err for postion [%d; %d]", 1, 1)
	}
}

func TestBoard_Cell(t *testing.T) {
	board := NewBoard()
	board.Put([2]int{1, 1}, KNIGHT)
	figure, err := board.Cell([2]int{1, 1})
	if err != nil {
		t.Fatalf("Not expected err %s for [%d, %d]", err.Error(), 1, 1)
	}
	if "N" != figure {
		t.Fatalf("Expected 'N', bat actual %s", figure)
	}
}

func TestBoard_EmptyCell(t *testing.T) {
	board := NewBoard()
	_, err := board.EmptyCell([2]int{0, 0})
	if err != nil {
		t.Fatalf("Expected empty cell [%d, %d]", 0, 0)
	}
	board.Put([2]int{0, 0}, KNIGHT)
	_, err = board.EmptyCell([2]int{0, 0})
	if err == nil || err.Error() != "cell not empty" {
		t.Fatalf("Expecte not empty cell [%d, %d]", 0, 0)
	}
}

func TestBoard_FullCell(t *testing.T) {
	board := NewBoard()
	board.Put([2]int{3, 3}, KNIGHT)
	figure, err := board.FullCell([2]int{3, 3})
	if err != nil {
		t.Fatalf("Unexpecter error %s", err.Error())
	}
	if "Nd4" != figure {
		t.Fatalf("Expected 'Nd4', but actual %s", figure)
	}
}

func TestBoard_AllowedMoveKnight(t *testing.T) {
	board := NewBoard()
	moves := board.AllowedMoveKnight([2]int{0, 0})
	if len(moves) != 2 {
		t.Fatalf("Expected 2 step from [%d, %d]", 0, 0)
	}
	a1 := []string{"Nb3", "Nc2"}
	checkSteps(t, a1, moves)
	moves = board.AllowedMoveKnight([2]int{7,7})
	h8 := []string{"Ng6", "Nf7"}
	checkSteps(t, h8, moves)
	moves = board.AllowedMoveKnight([2]int{3,3})
	d4 := []string{"Nc6", "Ne6", "Nf5", "Nf3", "Ne2", "Nc2", "Nb3", "Nb5"}
	checkSteps(t, d4, moves)
	moves = board.AllowedMoveKnight([2]int{0,7})
	a8 := []string{"Nb6", "Nc7"}
	checkSteps(t, a8, moves)
	moves = board.AllowedMoveKnight([2]int{1,1})
	b2 := []string{"Na4", "Nc4", "Nd3", "Nd1"}
	checkSteps(t, b2, moves)
}

func checkSteps(t *testing.T, allowed []string, moves []string){
	for _, s := range moves{
		if !contains(allowed, s){
			t.Fatalf("Allowed steps %s. Unexpected step %s for KNIGHT", allowed, s)
		}
	}
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
