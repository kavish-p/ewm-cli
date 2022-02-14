package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func AppendLine(fileName string, line string) {
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	if _, err := f.WriteString(line); err != nil {
		log.Println(err)
	}
}

func GenerateDefectHTMLReport(fileName string) {

	_, error := os.Stat(fileName)
	if os.IsNotExist(error) {
		log.Println("Specified file does not exist: " + fileName)
		log.Println("No defects were previously created.")
	} else {
		file, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		rows := ``

		rowTemplate := `
			<tr>
				<td>%s</td>
				<td>%s</td>
				<td>%s</td>
				<td>%s</td>
			</tr>
		`

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			elements := strings.Split(scanner.Text(), ",")
			row := fmt.Sprintf(rowTemplate, elements[0], elements[1], elements[2], elements[3])
			rows = rows + row
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		report := DefectBaseReportStart + rows + DefectBaseReportEnd

		AppendLine("DefectsReport.html", report)
		log.Println("Defects Report Generated Successfully")
	}

}
