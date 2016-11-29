package worker

import (
	"github.com/eatbytes/razboynik/services/printer"
	"github.com/eatbytes/razboynik/services/worker/targetwork"
)

func TargetList() error {
	var (
		config *targetwork.Configuration
		err    error
	)

	config, err = targetwork.GetConfiguration()

	if err != nil {
		return err
	}

	printer.PrintTitle("TARGETS")
	for _, item := range config.Targets {
		printer.Println("- " + item.Name)
	}
	printer.Print("\n")

	return nil
}

func TargetAdd() error {
	var (
		config    *targetwork.Configuration
		newTarget *targetwork.Target
		err       error
	)

	config, err = targetwork.GetConfiguration()

	if err != nil {
		return err
	}

	newTarget = targetwork.CreateTarget()
	config.Targets = append(config.Targets, newTarget)

	return targetwork.SaveConfiguration(config)
}
