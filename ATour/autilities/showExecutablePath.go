package autilities

import "os"

func ShowExecutablePath() {
	exePath, err := os.Executable()

	if err != nil {
		PLine("Failed to get executable path: ", err)
	} else {
		PLine("Executable Path : ", exePath)
	}
}
