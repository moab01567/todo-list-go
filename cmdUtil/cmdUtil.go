package cmdUtil

func CheckIfArgsPresent(args []string) bool {
	if len(args) <= 0 {
		return false
	} else {
		return true
	}
}

func CheckIfHelpFlagPresent(helpFlag string) bool {
	if helpFlag == "--help" || helpFlag == "-h" {
		return true
	}
	return false
}
