package types

type Batch uint8

const (
	BatchStageA Batch = iota
	BatchStageB Batch = iota
)

func (T *Batch) Toggle() {
	if *T == BatchStageA {
		*T = BatchStageB
	} else if *T == BatchStageB {
		*T = BatchStageA
	}
}
