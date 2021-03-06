package predictor

import (
	"bytes"
	"context"
	"io/ioutil"
	"strings"

	"github.com/c3sr/config"
	"github.com/c3sr/dlframework"
	"github.com/c3sr/dlframework/framework/agent"
	"github.com/c3sr/dlframework/framework/options"
	common "github.com/c3sr/dlframework/framework/predictor"
	"github.com/c3sr/tensorflow"
	"github.com/c3sr/tracer"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	tf "github.com/tensorflow/tensorflow/tensorflow/go"
	gotensor "gorgonia.org/tensor"
)

type ImageClassificationPredictor struct {
	*ImagePredictor
	inputLayer         string
	probabilitiesLayer string
	probabilities      interface{}
}

func NewImageClassificationPredictor(model dlframework.ModelManifest, opts ...options.Option) (common.Predictor, error) {
	ctx := context.Background()
	span, ctx := tracer.StartSpanFromContext(ctx, tracer.APPLICATION_TRACE, "new_predictor")
	defer span.Finish()

	modelInputs := model.GetInputs()
	if len(modelInputs) != 1 {
		return nil, errors.New("number of inputs not supported")
	}
	firstInputType := modelInputs[0].GetType()
	if strings.ToLower(firstInputType) != "image" {
		return nil, errors.New("input type not supported")
	}

	predictor := new(ImageClassificationPredictor)
	return predictor.Load(ctx, model, opts...)
}

func (self *ImageClassificationPredictor) Load(ctx context.Context, modelManifest dlframework.ModelManifest, opts ...options.Option) (common.Predictor, error) {
	pred, err := self.ImagePredictor.Load(ctx, modelManifest, opts...)
	if err != nil {
		return nil, err
	}

	p := &ImageClassificationPredictor{
		ImagePredictor: pred,
	}

	model, err := ioutil.ReadFile(p.GetGraphPath())
	if err != nil {
		return nil, errors.Wrapf(err, "cannot read %s", p.GetGraphPath())
	}
	modelReader := bytes.NewReader(model)

	p.inputLayer, err = p.GetInputLayerName(modelReader, "input_layer")
	if err != nil {
		return nil, errors.Wrap(err, "failed to get the input layer name")
	}
	p.probabilitiesLayer, err = p.GetOutputLayerName(modelReader, "probabilities_layer")
	if err != nil {
		return nil, errors.Wrap(err, "failed to get the probabilities layer name")
	}

	return p, nil
}

// Predict ...
func (p *ImageClassificationPredictor) Predict(ctx context.Context, data interface{}, opts ...options.Option) error {
	if data == nil {
		return errors.New("input data nil")
	}
	input, ok := data.([]*gotensor.Dense)
	if !ok {
		return errors.New("input data is not slice of dense tensors")
	}

	session := p.tfSession
	graph := p.tfGraph

	tensor, err := makeTensorFromGoTensors(input)
	if err != nil {
		return err
	}

	sessionSpan, ctx := tracer.StartSpanFromContext(ctx, tracer.MODEL_TRACE, "c_predict",
		opentracing.Tags{
			"evaluation_trace_level": p.TraceLevel(),
		})

	err = p.cuptiStart(ctx)
	if err != nil {
		return err
	}

	fetches, err := session.Run(ctx,
		map[tf.Output]*tf.Tensor{
			graph.Operation(p.inputLayer).Output(0): tensor,
		},
		[]tf.Output{
			graph.Operation(p.probabilitiesLayer).Output(0),
		},
		nil,
		p.runOptions(),
		p.GetGraphPath(),
	)

	p.cuptiClose()

	sessionSpan.Finish()

	if err != nil {
		return errors.Wrapf(err, "failed to perform session.Run")
	}

	p.probabilities = fetches[0].Value()

	return nil
}

// ReadPredictedFeatures ...
func (p *ImageClassificationPredictor) ReadPredictedFeatures(ctx context.Context) ([]dlframework.Features, error) {
	span, ctx := tracer.StartSpanFromContext(ctx, tracer.APPLICATION_TRACE, "read_predicted_features")
	defer span.Finish()

	e, ok := p.probabilities.([][]float32)
	if !ok {
		return nil, errors.New("output is not of type [][]float32")
	}

	labels, err := p.GetLabels()
	if err != nil {
		return nil, errors.New("cannot get the labels")
	}

	return p.CreateClassificationFeatures(ctx, e, labels)
}

func (p ImageClassificationPredictor) Modality() (dlframework.Modality, error) {
	return dlframework.ImageClassificationModality, nil
}

func init() {
	config.AfterInit(func() {
		framework := tensorflow.FrameworkManifest
		agent.AddPredictor(framework, &ImageClassificationPredictor{
			ImagePredictor: &ImagePredictor{
				ImagePredictor: common.ImagePredictor{
					Base: common.Base{
						Framework: framework,
					},
				},
			},
		})
	})
}
