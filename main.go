package main

import (
	"context"
	"errors"
	"flag"
	"github.com/genuinetools/pkg/cli"
	"github.com/sirupsen/logrus"
	"google.golang.org/api/cloudbuild/v1"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	build   *cloudbuild.Build
	config  string
	verbose bool
)

func main() {
	p := cli.NewProgram()
	p.Name = "cloudbuild-ci"
	p.Description = "A bot that fires a JSON blob down range to my cloud function"

	p.FlagSet = flag.NewFlagSet("global", flag.ExitOnError)
	p.FlagSet.StringVar(&config, "config", "cloudbuild.yaml", "the filename to build")
	p.FlagSet.BoolVar(&verbose, "verbose", false, "enable debug logging")

	p.Before = func(ctx context.Context) error {
		if verbose {
			logrus.SetLevel(logrus.DebugLevel)
		}
		return nil
	}

	p.Action = func(ctx context.Context, args []string) error {
		logrus.Debugf("Trying to read file: %s", config)
		yf, err := ioutil.ReadFile(config)
		if err != nil {
			logrus.Fatalf("Could not parse %s: #%v", config, err.Error())
			return err
		}

		logrus.Debugf("Trying to create build from file")
		if err := yaml.Unmarshal(yf, &build); err != nil {
			logrus.Fatalf("Could not create build from file %s: %v", config, err.Error())
			return err
		}

		if build == nil {
			logrus.Fatalf("Could not unmarshal yaml into build: %v", build)
			return errors.New("Could not unmarshal yaml")
		}

		logrus.WithFields(logrus.Fields{
			"project": build.ProjectId,
			"images":  build.Images,
			"tags":    build.Tags,
		}).Debugf("Got build")

		if len(build.Steps) == 0 {
			logrus.Fatalf("No steps found in build: %s", config)
			return errors.New("No build steps found")
		}

		logrus.Infof("ok")
		return nil
	}
	p.Run()
}
