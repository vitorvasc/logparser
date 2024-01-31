package main

import (
	"bufio"
	"encoding/json"
	"os"

	"logparser/internal/adapter/dto"
	"logparser/internal/adapter/handler"
	"logparser/internal/adapter/repository/memorydb"
	"logparser/internal/core/service"
)

type ProcessFileFunc func(*os.File) *dto.ProcessResult
type GetMatchFunc func(string) (*dto.MatchDetails, error)

const (
	LogFileLocation = "resources/qgames.log"
)

func main() {
	// dependencies
	matchRepository := memorydb.NewMatchRepository()
	createMatchService := service.NewCreateMatchService(matchRepository)
	getMatchService := service.NewGetMatchService(matchRepository)
	logFileHandler := handler.NewLogFileHandler(createMatchService)
	getMatchHandler := handler.NewGetMatchHandler(getMatchService)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		renderMenu()
		println("Please, enter the option number:")

		scanner.Scan()
		switch scanner.Text() {
		case "1":
			_ = processLogFileAndLoadMatches(logFileHandler.CreateMatchesFromLogFile)
			println("Matches loaded successfully. Type any key to continue...")
			scanner.Scan()
		case "2":
			match, err := generateReportByGameNumber(scanner, getMatchHandler.GetMatchByID)
			if err != nil {
				println("An error occurred while generating the report: ", err.Error())
				println("Type any key to continue...")
				scanner.Scan()
			} else {
				result, _ := json.MarshalIndent(match, "", "\t")
				print(string(result))
				println("\nType any key to continue...")
			}
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
	println("Processing log file and loading matches...")
	// open log file
	file, err := os.Open(LogFileLocation)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	return handler(file)
}

func generateReportByGameNumber(scanner *bufio.Scanner, handler GetMatchFunc) (*dto.MatchDetails, error) {
	println("Please, enter the game number:")
	scanner.Scan()
	return handler(scanner.Text())
}
