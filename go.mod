module github.com/ijasMohamad/negt

go 1.19

require (
	github.com/aymerick/raymond v2.0.2+incompatible
	github.com/gertd/go-pluralize v0.2.1
	github.com/iancoleman/strcase v0.2.0
	github.com/manifoldco/promptui v0.9.0
	github.com/spf13/cobra v1.6.1
	golang.org/x/text v0.4.0
)

require (
	github.com/chzyer/readline v1.5.1 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/sys v0.2.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

retract [v0.1.0, v0.1.1]

retract v0.1.2

retract [v1.0.0, v1.0.1]
