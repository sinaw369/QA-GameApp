package entity

type Question struct {
	ID              uint
	Text            string
	PossibleAnswer  []PossibleAnswer
	CorrectAnswerID uint
	Difficulty      QuestionDifficulty
	CategoryID      uint
}
type PossibleAnswer struct {
	ID     uint
	Text   string
	Choice PossibleAnswerChoice
}
type QuestionDifficulty uint8
type PossibleAnswerChoice uint8

const (
	QuestionDifficultyEasy QuestionDifficulty = iota + 1
	QuestionDifficultyMedium
	QuestionDifficultyHard
)
const (
	PossibleAnswerA PossibleAnswerChoice = iota + 1
	PossibleAnswerB
	PossibleAnswerC
	PossibleAnswerD
)

func (q QuestionDifficulty) IsValid() bool {
	if q >= QuestionDifficultyEasy && q <= QuestionDifficultyHard {
		return true
	}
	return false

}

func (p PossibleAnswerChoice) IsValid() bool {
	if p >= PossibleAnswerA && p <= PossibleAnswerD {
		return true
	}
	return false

}
