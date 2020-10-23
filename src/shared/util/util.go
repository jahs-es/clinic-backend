package util

func GetExpresion(column string, value string) string {
	if value != "" {
		return column + " like '%" + value + "%' OR "
	} else {
		return ""
	}
}

