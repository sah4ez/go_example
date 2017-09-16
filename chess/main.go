package chess

import (
	"github.com/op/go-logging"
	"fmt"
	"errors"
)

const KNIGHT = "N"

var (
	LETTERS = map[int]string{
		0: "a",
		1: "b",
		2: "c",
		3: "d",
		4: "e",
		5: "f",
		6: "g",
		7: "h",
	}
	KNIGHT_MOVES = [8][2]int{
		{1,2},
		{-1,2},
		{2,1},
		{2,-1},
		{1,-2},
		{-1,-2},
		{-2,1},
		{-2,-1},
	}
	LOG = logging.MustGetLogger("chess")
)

type GameBoard interface {
	MoveKnight(from [2]int, to [2]int) string
	AllowedMoveKnight(pos [2]int) []string
	Put(pos [2]int, figure string)
	EmptyCell(pos [2]int) bool
	Cell(pos [2]int) string
}

type Board struct {
	area [8][8]string
}

func NewBoard() *Board {
	return &Board{
		area: [8][8]string{
			{
				"", "", "", "", "", "", "", "",
			},
			{
				"", "", "", "", "", "", "", "",
			},
			{
				"", "", "", "", "", "", "", "",
			},
			{
				"", "", "", "", "", "", "", "",
			},
			{
				"", "", "", "", "", "", "", "",
			},
			{
				"", "", "", "", "", "", "", "",
			},
			{
				"", "", "", "", "", "", "", "",
			},
			{
				"", "", "", "", "", "", "", "",
			},
		},
	}
}

func (b *Board) Put(pos [2]int, figure string) (string, error) {
	res, err := b.EmptyCell(pos)
	if err != nil {
		if !res {
			figure, _ := b.FullCell(pos)
			LOG.Warningf("In cell %s", figure)
		}
		return "", err
	}
	b.area[pos[0]][pos[1]] = figure
	return figure, nil
}

func (b *Board) EmptyCell(pos [2]int) (bool, error) {
	figure, err := b.Cell(pos)
	if err != nil{
		return false, err
	}
	if "" == figure {
		return true, nil
	}
	return false, errors.New("cell not empty")
}

func (b *Board) Cell(pos [2]int) (string, error) {
	_, err := ValidPos(pos)
	if err != nil {
		LOG.Errorf("Invalid position")
		return "", err
	}
	figure := b.area[pos[0]][pos[1]]
	if "" != figure{
		LOG.Debugf("Get [%d;%d] figure %s", pos[0], pos[1], figure)
	}
	return figure, nil
}

func (b *Board) FullCell(pos [2]int) (string, error) {
	_, err := ValidPos(pos)
	if err != nil {
		LOG.Errorf("Invalid position")
		return "", err
	}
	return format(b.area[pos[0]][pos[1]], pos[0] ,pos[1]) , nil
}

func (b *Board) AllowedMoveKnight(pos [2]int) []string{
	moves := []string{}
	for _, p := range KNIGHT_MOVES{
		newPos := moveOffset(pos, p)
		b, _ := ValidPos(newPos)
		if b {
			step :=  format(KNIGHT, newPos[0], newPos[1])
			moves = append(moves, step)
		}
	}
	LOG.Debugf("Allowed steps %s", moves)
	return moves
}

func moveOffset(from [2]int, offset [2]int)[2]int{
	from[0] = from[0] + offset[0]
	from[1] = from[1] + offset[1]
	return from
}


func ValidPos(pos [2]int) (bool, error){
	if (pos[0] < 0 || pos[0] > 7) || (pos[1] < 0 || pos[1] > 7){
		return false, errors.New("invalid position")
	}
	return true, nil
}

func format(figure string, letter int, number int) string{
	return figure + LETTERS[letter] + fmt.Sprintf("%d", number+1)
}
