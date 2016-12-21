package worker

import (
	"strconv"

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

func TargetEdit(name string) error {
	var (
		config *targetwork.Configuration
		target *targetwork.Target
		err    error
	)

	config, err = targetwork.GetConfiguration()

	if err != nil {
		return err
	}

	target, _, err = targetwork.FindTarget(config, name)

	if err != nil {
		return err
	}

	targetwork.EditTarget(target)

	return targetwork.SaveConfiguration(config)
}

func TargetDetail(name string) error {
	var (
		config *targetwork.Configuration
		target *targetwork.Target
		err    error
	)

	config, err = targetwork.GetConfiguration()

	if err != nil {
		return err
	}

	target, _, err = targetwork.FindTarget(config, name)

	if err != nil {
		return err
	}

	printer.Println("Name: " + target.Name)
	printer.Println("Url: " + target.Config.Url)
	printer.Println("Method: " + target.Config.Method)
	printer.Println("Parameter: " + target.Config.Parameter)
	printer.Println("Shellmethod: " + target.Config.Shellmethod)
	printer.Println("Shellscope: " + target.Config.Shellscope)
	printer.Println("NoExtra: " + strconv.FormatBool(target.Config.NoExtra))

	return nil
}

func TargetRun(name string) error {
	var (
		config *targetwork.Configuration
		target *targetwork.Target
		err    error
	)

	config, err = targetwork.GetConfiguration()

	if err != nil {
		return err
	}

	target, _, err = targetwork.FindTarget(config, name)

	if err != nil {
		return err
	}

	return Run(target.Config)
}

func TargetRemove(name string) error {
	var (
		config *targetwork.Configuration
		index  int
		err    error
	)

	config, err = targetwork.GetConfiguration()

	if err != nil {
		return err
	}

	_, index, err = targetwork.FindTarget(config, name)

	if err != nil {
		return err
	}

	config.Targets = append(config.Targets[:index], config.Targets[index+1:]...)

	return targetwork.SaveConfiguration(config)
}
