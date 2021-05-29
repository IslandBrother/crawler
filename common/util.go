package common

func GetErrorString(err error) string {
	if err != nil {
		return err.Error()
	}
	return ""
}
