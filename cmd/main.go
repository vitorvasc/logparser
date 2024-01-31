package main

import (
	"bufio"
	"os"

	"logparser/internal/adapter/dto"
	"logparser/internal/adapter/handler"
	"logparser/internal/adapter/repository/memorydb"
	"logparser/internal/core/service"
)

type ProcessFileFunc func(*os.File) *dto.ProcessResult

const (
	LogFileLocation = "resources/qgames.log"
)

func main() {
	// dependencies
	matchRepository := memorydb.NewMatchRepository()
	matchService := service.NewCreateMatchService(matchRepository)
	logHandler := handler.NewLogFileHandler(matchService)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		renderMenu()
		println("Please, enter the option number:")

		scanner.Scan()
		switch scanner.Text() {
		case "1":
			_ = processLogFileAndLoadMatches(logHandler.CreateMatchesFromLogFile)
			println("Matches loaded successfully. Type any key to continue...")
			scanner.Scan()
		case "2":
			println("Generating report by game number...")
		case "3":
			println("Generating report by weapon type...")
		case "9":
			println("Exiting...")
			os.Exit(0)
		default:
			println("Invalid option")
		}
	}
}

func renderMenu() {
	println("[MENU] The following options are available:")
	println("1 - Process log file and load matches")
	println("2 - Generate report by game number")
	println("9 - Exit")
}

func processLogFileAndLoadMatches(handler ProcessFileFunc) *dto.ProcessResult {
	// open log file
	file, err := os.Open(LogFileLocation)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	return handler(file)
}

func generateReportByGameNumber(scanner *bufio.Scanner) {
	scanner.Scan()
}
