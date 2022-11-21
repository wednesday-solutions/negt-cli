package hbs

import (
	"fmt"
	"path/filepath"

	"github.com/ijasMohamad/cliApp/gqlgenUtils/fileUtils"
)

func TestingFunction() {

	fields := []string{"name", "class", "subject"}
	ctx := map[string]interface{}{
		"fields": fields,
	}
	path, _ := filepath.Abs(".")
	path = fmt.Sprintf("%s/%s", path, "gql/models")
	file := "testFile.js"

	source := `Hello world
	
	"Hello world"
	
	"{{test fields}}"

	`

	tpl, err := GenerateTemplate(source, ctx)
	if err != nil {
		fmt.Println("Error generating template", err)
	}
	fileUtils.WriteToFile(path, file, tpl)
}