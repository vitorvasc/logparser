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
type GetReportByMatchID func(string) (map[string]*dto.MatchDetails, error)
type GetReportForAllMatches func() (map[string]*dto.MatchDetails, error)

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
		print("Please, enter the option number: ")

		scanner.Scan()
		switch scanner.Text() {
		case "1":
			result := processLogFileAndLoadMatches(logFileHandler.CreateMatchesFromLogFile)
			println()
			print(result.TotalProcessedMatches, " matches processed successfully.")
			typeAnyKeyToContinue(scanner)
		case "2":
			match, err := generateReportByGameNumber(scanner, getMatchHandler.GetSimpleReportByMatchID)
			if err != nil {
				print("An error occurred while generating the report: ", err.Error())
				typeAnyKeyToContinue(scanner)
				continue
			}

			result, _ := json.MarshalIndent(match, "", "\t")
			print(string(result))
			typeAnyKeyToContinue(scanner)
		case "3":
			match, err := generateReportByGameNumber(scanner, getMatchHandler.GetCompleteReportByMatchID)
			if err != nil {
				print("An error occurred while generating the report: ", err.Error())
				typeAnyKeyToContinue(scanner)
				continue
			}

			result, _ := json.MarshalIndent(match, "", "\t")
			print(string(result))
			typeAnyKeyToContinue(scanner)
		case "4":
			matches, err := generateReportForAllGames(getMatchHandler.GetCompleteReportForAllMatches)
			if err != nil {
				print("An error occurred while generating the report: ", err.Error())
				typeAnyKeyToContinue(scanner)
				continue
			}

			result, _ := json.MarshalIndent(matches, "", "\t")
			print(string(result))
			typeAnyKeyToContinue(scanner)
		case "9":
			println("Exiting...")
			os.Exit(0)
		default:
			println("Invalid option")
		}
	}
}

func renderMenu() {
	for i := 0; i < 10; i++ {
		println()
	}
	println("[MENU] The following options are available:")
	println("1 - Process log file and load matches")
	println("2 - Generate report by game number (simple report)")
	println("3 - Generate report by game number, including deaths by cause (complete report)")
	println("4 - Generate report for all games (complete report, including deaths by cause)")
	println("9 - Exit")
}

func processLogFileAndLoadMatches(handler ProcessFileFunc) *dto.ProcessResult {
	println()
	println()
	println("Processing log file and loading matches...")
	// open log file
	file, err := os.Open(LogFileLocation)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	return handler(file)
}

func generateReportByGameNumber(scanner *bufio.Scanner, handler GetReportByMatchID) (map[string]*dto.MatchDetails, error) {
	print("Please, enter the game number: ")
	scanner.Scan()
	return handler(scanner.Text())
}

func generateReportForAllGames(handler GetReportForAllMatches) (map[string]*dto.MatchDetails, error) {
	return handler()
}

func typeAnyKeyToContinue(scanner *bufio.Scanner) {
	println()
	print("\nType any key to continue...")
	scanner.Scan()
}
