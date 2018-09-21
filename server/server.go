package server

import (
	"fmt"
	"log"

	"github.com/paradigm-network/paradigm-fn2/config"
	"github.com/paradigm-network/paradigm-fn2/pkg/docker"
	"github.com/paradigm-network/paradigm-fn2/pkg/server"
)

type PullTask struct {
	ImageName string
	Err       error
}

func newPullTask(imageName string, result error) PullTask {
	return PullTask{
		ImageName: imageName,
		Err:       result,
	}
}

//PullBaseDockerImage fetch base images from the registry
func PullBaseDockerImage(verbose bool) []PullTask {
	baseImages := []string{
		"metrue/fn2-java-base",
		"metrue/fn2-julia-base",
		"metrue/fn2-python-base",
		"metrue/fn2-node-base",
		"metrue/fn2-d-base",
	}

	count := len(baseImages)
	results := make(chan PullTask, count)

	task := func(image string, verbose bool) error {
		return docker.Pull(image, verbose)
	}

	fmt.Println("fn2 is pulling some basic resources")
	for _, image := range baseImages {
		go func(img string) {
			err := task(img, verbose)
			results <- newPullTask(img, err)
		}(image)
	}

	var pullResutls []PullTask
	for result := range results {
		pullResutls = append(pullResutls, result)

		if len(pullResutls) == count {
			close(results)
		}
	}

	return pullResutls
}

// Start parses input and launches the fn2 server in a blocking process
func Start(verbose bool) error {
	if !docker.IsRunning() {
		panic("make sure docker is running on your host")
	} else {
		go PullBaseDockerImage(true)
	}

	grpcEndpoint := fmt.Sprintf("0.0.0.0:%d", config.GRPC_PORT)
	httpEndpoint := fmt.Sprintf("0.0.0.0:%d", config.HTTP_PORT)
	go func() {
		s := server.NewFn2ServiceServer(grpcEndpoint)
		err := s.Start()
		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Printf("fn2 serves on %d", config.HTTP_PORT)
	return Run(grpcEndpoint, httpEndpoint)
}
