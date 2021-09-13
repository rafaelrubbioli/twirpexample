package proto

func (c Client) Name ()string {
	switch c {
	case Client_GO:
		return "Go"
	case Client_PYTHON:
		return "Python"
	default:
		return ""
	}
}
