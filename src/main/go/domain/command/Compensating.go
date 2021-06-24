package command

type (
	BaseCompensating interface {
	}
	RegisterCommandCompensating struct {
		AggregatId string
		BaseCompensating
	}
)