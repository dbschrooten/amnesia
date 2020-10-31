package dashboard

import (
	"html/template"
	"log"
)

var (
	OverviewTpl *template.Template
	AboutTpl    *template.Template
	UserTpl     *template.Template
)

func LoadTemplate(name string, files []string) (*template.Template, error) {
	result, err := template.New(name).
		ParseFiles(files...)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func Setup() error {
	var err error

	OverviewTpl, err = LoadTemplate(
		"overview",
		[]string{
			"src/dashboard/views/layout.html",
			"src/dashboard/views/overview.html",
		},
	)

	if err != nil {
		return err
	}

	log.Print("Setup dashboard")

	return nil
}
