package components

// The funcs bellow are return a colorized string
func Gray (letter string) string {
	return "\033[30m"+letter+"\033[0m"
}

func Red (letter string) string {
	return "\033[31m"+letter+"\033[0m"
}

func Green (letter string) string {
	return "\033[32m"+letter+"\033[0m"
}

func Yellow (letter string) string {
	return "\033[33m"+letter+"\033[0m"
}

func Blue (letter string) string {
	return "\033[34m"+letter+"\033[0m"
}

func Pink (letter string) string {
	return "\033[35m"+letter+"\033[0m"
}

func Sky (letter string) string {
	return "\033[36m"+letter+"\033[0m"
}