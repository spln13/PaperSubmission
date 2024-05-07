package enum

type ResponseType int

const (
	OperateOK   ResponseType = 0
	OperateFail ResponseType = 1
)

func (p ResponseType) String() string {
	switch p {
	case OperateOK:
		return "OK"
	case OperateFail:
		return "Fail"
	default:
		return "Unknown"
	}
}
