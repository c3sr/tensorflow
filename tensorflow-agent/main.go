package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/Unknwon/com"
	"github.com/c3sr/config"
	dlcmd "github.com/c3sr/dlframework/framework/cmd"
	cmd "github.com/c3sr/dlframework/framework/cmd/server"
	common "github.com/c3sr/dlframework/framework/predictor"
	"github.com/c3sr/logger"
	"github.com/c3sr/tensorflow"
	"github.com/c3sr/tensorflow/graph"
	"github.com/c3sr/tensorflow/predictor"
	"github.com/c3sr/tracer"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	modelName    string
	modelVersion string
	hostName, _  = os.Hostname()
	framework    = tensorflow.FrameworkManifest
	log          *logrus.Entry
)

func graphConvert(c *cobra.Command, args []string) error {
	ctx := context.Background()

	model, err := framework.FindModel(modelName + ":" + modelVersion)
	if err != nil {
		return err
	}

	workDir, err := model.WorkDir()
	if err != nil {
		return err
	}

	predictorFramework := &predictor.ImagePredictor{
		ImagePredictor: common.ImagePredictor{
			Base: common.Base{
				Framework: model.MustResolveFramework(),
				Model:     *model,
				WorkDir:   workDir,
			},
		},
	}

	err = predictorFramework.Download(ctx, *model)
	if err != nil {
		return errors.Wrapf(err, "failed to download %s model", model.MustCanonicalName())
	}

	g, err := graph.New(predictorFramework.GetGraphPath())
	if err != nil {
		return err
	}

	bts, err := g.MarshalJSON()
	if err != nil {
		return err
	}

	baseDir := filepath.Join("experiments", framework.Name, framework.Version, model.Name, model.Version)
	if !com.IsDir(baseDir) {
		os.MkdirAll(baseDir, os.ModePerm)
	}

	ioutil.WriteFile(filepath.Join(baseDir, "model_info.json"), bts, 0644)

	return err
}

var graphCmd = &cobra.Command{
	Use:   "graph",
	Short: "Converts the frozen graph into a model graph summary",
	RunE: func(c *cobra.Command, args []string) error {
		tracer.SetLevel(tracer.NO_TRACE)
		if modelName == "all" {
			models := framework.Models()
			pb := dlcmd.NewProgress("download models", len(models))
			for _, model := range models {
				modelName = model.Name
				modelVersion = model.Version
				err := graphConvert(c, args)
				if err != nil {
					log.WithField("model_name", modelName).
						WithField("model_version", modelVersion).
						Error("failed to convert graph")
				}
				pb.Increment()
			}
			pb.Finish()
			return nil
		}
		return graphConvert(c, args)
	},
}

func init() {
	config.AfterInit(func() {
		log = logger.New().WithField("pkg", "tensorflow-agent")
	})
	graphCmd.PersistentFlags().StringVar(&modelName, "model_name", "MobileNet_v1_1.0_224", "the name of the model to use for conversion")
	graphCmd.PersistentFlags().StringVar(&modelVersion, "model_version", "1.0", "the version of the model to use for conversion")
}

func main() {
	rootCmd, err := cmd.NewRootCommand(tensorflow.Register, framework)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	rootCmd.AddCommand(graphCmd)

	defer tracer.Close()
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
