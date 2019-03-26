package predictor

import (
	"bytes"
	"context"
	goimage "image"
	"image/png"
	"os"
	"path/filepath"
	"testing"

	"github.com/rai-project/dlframework"

	"github.com/rai-project/dlframework/framework/options"
	"github.com/rai-project/image"
	"github.com/rai-project/image/types"
	nvidiasmi "github.com/rai-project/nvidia-smi"
	tf "github.com/rai-project/tensorflow"
	"github.com/stretchr/testify/assert"
	gotensor "gorgonia.org/tensor"
)

func TestImageEnhancementInference(t *testing.T) {
	tf.Register()
	model, err := tf.FrameworkManifest.FindModel("srgan:1.0")
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

	predictor, err := NewImageEnhancementPredictor(*model, options.WithOptions(opts))
	assert.NoError(t, err)
	assert.NotEmpty(t, predictor)
	defer predictor.Close()

	imgDir, _ := filepath.Abs("./_fixtures")
	imgPath := filepath.Join(imgDir, "penguin.png")
	r, err := os.Open(imgPath)
	if err != nil {
		panic(err)
	}
	img, err := image.Read(r)
	if err != nil {
		panic(err)
	}

	input := make([]*gotensor.Dense, batchSize)
	imgFloats, err := normalizeImageHWC(img.(*types.RGBImage), []float32{127.5, 127.5, 127.5}, 127.5)
	if err != nil {
		panic(err)
	}
	height := img.Bounds().Dy()
	width := img.Bounds().Dx()
	channels := 3
	for ii := 0; ii < batchSize; ii++ {
		input[ii] = gotensor.New(
			gotensor.WithShape(height, width, channels),
			gotensor.WithBacking(imgFloats),
		)
	}

	err = predictor.Predict(ctx, input)
	assert.NoError(t, err)
	if err != nil {
		return
	}

	pred, err := predictor.ReadPredictedFeatures(ctx)
	assert.NoError(t, err)
	if err != nil {
		panic(err)
	}

	f, ok := pred[0][0].Feature.(*dlframework.Feature_Image)
	if !ok {
		panic("expecting an image feature")
	}

	hrimg, _, _ := goimage.Decode(bytes.NewReader(f.Image.GetData()))

	output, _ := os.Create("/tmp/output.png")
	defer output.Close()

	err = png.Encode(output, hrimg)
	if err != nil {
		panic(err)
	}
}
