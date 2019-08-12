package predictor

import (
	"context"
	"testing"

	"github.com/rai-project/dlframework/framework/options"
	nvidiasmi "github.com/rai-project/nvidia-smi"
	tf "github.com/rai-project/tensorflow"
	"github.com/stretchr/testify/assert"
)

func TestRaw(t *testing.T) {
	tf.Register()
	// model, err := tf.FrameworkManifest.FindModel("AI_Matrix_ResNet50:1.0")
	model, err := tf.FrameworkManifest.FindModel("AI_Matrix_DIEN:1.0")
	assert.NoError(t, err)
	assert.NotEmpty(t, model)

	device := options.CPU_DEVICE
	if nvidiasmi.HasGPU {
		device = options.CUDA_DEVICE
	}

	batchSize := 1
	ctx := context.Background()
	opts := options.New(options.Context(ctx),
		options.Device(device, 0),
		options.BatchSize(batchSize))

	predictor, err := NewRawPredictor(*model, options.WithOptions(opts))
	assert.NoError(t, err)
	assert.NotEmpty(t, predictor)
	defer predictor.Close()

	err = predictor.Predict(ctx, nil)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	_, err = predictor.ReadPredictedFeatures(ctx)
	assert.NoError(t, err)
	if err != nil {
		return
	}
	// assert.InDelta(t, float32(0.998212), pred[0][0].GetProbability(), 0.001)
	// assert.Equal(t, int32(104), pred[0][0].GetClassification().GetIndex())
}
