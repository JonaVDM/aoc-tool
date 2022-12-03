package gen

import (
	"errors"
	"fmt"
	"os"
	"text/template"
)

func GenerateTemplates(year, day int) error {
	data := DayTempl{
		Day:  fmt.Sprintf("day%02d", day),
		Year: year,
	}
	pwd, err := os.Getwd()
	if err != nil {
		return err
	}

	_, err = os.Stat(data.Day)

	if errors.Is(err, os.ErrNotExist) {
		err = os.Mkdir(data.Day, os.ModePerm)
	}
	if err != nil {
		return err
	}

	err = generateDayFile(pwd, data)
	if err != nil {
		return err
	}

	err = generateDayTestFile(pwd, data)
	return err
}

func generateDayFile(pwd string, data DayTempl) error {
	dayFile, err := os.Create(fmt.Sprintf("%s/%s/%s.go", pwd, data.Day, data.Day))
	if err != nil {
		return err
	}
	defer dayFile.Close()

	dayTemplate := template.Must(template.New("day").Parse(string(DayTemplate())))
	return dayTemplate.Execute(dayFile, data)
}

func generateDayTestFile(pwd string, data DayTempl) error {
	testFile, err := os.Create(fmt.Sprintf("%s/%s/%s_test.go", pwd, data.Day, data.Day))
	if err != nil {
		return err
	}
	defer testFile.Close()

	testTemplate := template.Must(template.New("test").Parse(string(DayTestTemplate())))
	return testTemplate.Execute(testFile, data)
}
