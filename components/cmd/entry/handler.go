package main

import (
	"github.com/alexeyco/simpletable"
	"github.com/autoai-org/aiflow/components/cmd/pkg/entities"
	"github.com/autoai-org/aiflow/components/cmd/pkg/runtime"
	"github.com/autoai-org/aiflow/components/cmd/pkg/utilities"
	"path/filepath"
	"strconv"
	"strings"
)

func build(solverName string) {
	// Read Dockerfile
	dockerClient := runtime.NewDockerRuntime()
	// Read Base Image Name
	tomlString := utilities.ReadFileContent("./cvpm.toml")
	packageInfo := entities.LoadPackageFromConfig(tomlString)
	if solverName == "*" {
		for _, solver := range packageInfo.Solvers {
			imageName := packageInfo.Package.Vendor+"-"+packageInfo.Package.Name +"-" + solver.Name
			dockerClient.Build(strings.ToLower(imageName), "./docker_"+solver.Name)
		}
	} else {
		imageName := packageInfo.Package.Vendor+"-"+packageInfo.Package.Name +"-" + solverName
		dockerClient.Build(strings.ToLower(imageName), "./docker_"+ solverName)
	}
}

func printImages() {
	dockerClient := runtime.NewDockerRuntime()
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Image ID"},
			{Align: simpletable.AlignCenter, Text: "Repo Tags"},
			{Align: simpletable.AlignCenter, Text: "Size"},
		},
	}
	images := dockerClient.ListImages()
	for _, image := range images {
		r := []*simpletable.Cell{
			{Text: image.ID},
			{Text: strings.Join(image.RepoTags, ",")},
			{Text: strconv.FormatInt(image.Size, 10)},
		}
		table.Body.Cells = append(table.Body.Cells, r)
	}
	table.Println()
}

func printContainers() {
	dockerClient := runtime.NewDockerRuntime()
	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "Container ID"},
			{Align: simpletable.AlignCenter, Text: "Status"},
			{Align: simpletable.AlignCenter, Text: "Image"},
		},
	}
	containers := dockerClient.ListContainers()
	for _, container := range containers {
		r := []*simpletable.Cell{
			{Text: container.ID},
			{Text: container.Status},
			{Text: container.Image},
		}
		table.Body.Cells = append(table.Body.Cells, r)
	}
	table.Println()
}

func generate() {
	tomlFilePath := filepath.Join("./", "cvpm.toml")
	cvpmToml := utilities.ReadFileContent(tomlFilePath)
	solvers := entities.LoadSolversFromConfig(cvpmToml)
	runtime.RenderRunnerTpl("./", solvers)
	runtime.GenerateDockerFiles("./")
}
