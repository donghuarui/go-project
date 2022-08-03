package service

type Apple struct {
}
type Banana struct {
}

func (apple Apple) Eat(taste string) string {
	return "苹果" + taste
}

func (banana Banana) Eat(taste string) string {
	return "香蕉" + taste
}
