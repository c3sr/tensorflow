# MLModelScope TensorFlow Agent

[![Build Status](https://travis-ci.org/rai-project/tensorflow.svg?branch=master)](https://travis-ci.org/rai-project/tensorflow)
[![Build Status](https://dev.azure.com/dakkak/rai/_apis/build/status/tensorflow)](https://dev.azure.com/dakkak/rai/_build/latest?definitionId=14)
[![Go Report Card](https://goreportcard.com/badge/github.com/rai-project/tensorflow)](https://goreportcard.com/report/github.com/rai-project/tensorflow)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

[![](https://images.microbadger.com/badges/version/carml/tensorflow:ppc64le-gpu-latest.svg)](https://microbadger.com/images/carml/tensorflow:ppc64le-gpu-latest> 'Get your own version badge on microbadger.com') [![](https://images.microbadger.com/badges/version/carml/tensorflow:ppc64le-cpu-latest.svg)](https://microbadger.com/images/carml/tensorflow:ppc64le-cpu-latest 'Get your own version badge on microbadger.com') [![](https://images.microbadger.com/badges/version/carml/tensorflow:amd64-cpu-latest.svg)](https://microbadger.com/images/carml/tensorflow:amd64-cpu-latest 'Get your own version badge on microbadger.com') [![](https://images.microbadger.com/badges/version/carml/tensorflow:amd64-gpu-latest.svg)](https://microbadger.com/images/carml/tensorflow:amd64-gpu-latest 'Get your own version badge on microbadger.com')

## Installation

Download and install the MLModelScope TensorFlow Agent:

```
go get -v github.com/rai-project/tensorflow

```

The agent requires The TensorFlow C library and other Go packages.

### The TensorFlow C library

The TensorFlow C library is required for the TensorFlow Go package.
You can download pre-built TensorFlow C library from [Install TensorFlow for C](https://www.tensorflow.org/install/lang_c).

Extract the downloaded archive to `/opt/tensorflow/`.

```
tar -C /opt/tensorflow -xzf (downloaded file)
```

Configure the linker environmental variables since the TensorFlow C library is extracted to a non-system directory. Place the following in either your `~/.bashrc` or `~/.zshrc` file

Linux

```
export LIBRARY_PATH=$LIBRARY_PATH:/opt/tensorflow/lib
export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/opt/tensorflow/lib
```

macOS

```
export LIBRARY_PATH=$LIBRARY_PATH:/opt/tensorflow/lib
export DYLD_LIBRARY_PATH=$DYLD_LIBRARY_PATH:/opt/tensorflow/lib
```

You can test the installed TensorFlow C library using an [examle C program](https://www.tensorflow.org/install/lang_c#build).

To build the TensorFlow C library from source, refer to [TensorFlow in Go](https://github.com/tensorflow/tensorflow/tree/master/tensorflow/go#building-the-tensorflow-c-library-from-source) .

### Go packages

You can install the dependency through `go get`.

```
cd $GOPATH/src/github.com/rai-project/tensorflow
go get -u -v ./...
```

Or use [Dep](https://github.com/golang/dep).

```
dep ensure -v
```

This installs the dependency in `vendor/`.

### libjpeg-turbo

[libjpeg-turbo](https://github.com/libjpeg-turbo/libjpeg-turbo) is a JPEG image codec that uses SIMD instructions (MMX, SSE2, AVX2, NEON, AltiVec) to accelerate baseline JPEG compression and decompression. It outperforms libjpeg by a significant amount.

The default is to use libjpeg-turb, to opt-out, use build tag `nolibjpeg`.

To install libjpeg-turbo, refer to [libjpeg-turbo](https://github.com/libjpeg-turbo/libjpeg-turbo).

## Run

With GPU enabled

```
cd $GOPATH/src/github.com/rai-project/tensorflow
go run tensorflow-agent/main.go -l -d -v
```

Without GPU or libjpeg-turbo

```
cd $GOPATH/src/github.com/rai-project/tensorflow
go run -tags="nogpu nolibjpeg" tensorflow-agent/main.go -l -d -v
```

## Notes on installing TensorFlow from source

### Install Bazel

{{% notice note %}}
Currently there's issue using bazel 0.19.1 to build TensorFlow 1.12 with CUDA 10.0
{{% /notice %}}

- [Installing Bazel on Ubuntu](https://docs.bazel.build/versions/master/install-ubuntu.html)

- [Installing Bazel on macOS](https://docs.bazel.build/versions/master/install-os-x.html#install-on-mac-os-x-homebrew)

### Build

Build TensorFlow 1.12 with the following scripts.

```sh
go get -d github.com/tensorflow/tensorflow/tensorflow/go
cd ${GOPATH}/src/github.com/tensorflow/tensorflow
git fetch --all
git checkout r1.12
./configure
```

For linux with gpu, an example `.tf_configure.bazelrc` is

```
build --action_env PYTHON_BIN_PATH="/usr/bin/python"
build --action_env PYTHON_LIB_PATH="/usr/lib/python3/dist-packages"
build --python_path="/usr/bin/python"
build:ignite --define with_ignite_support=true
build --define with_xla_support=true
build --action_env TF_NEED_OPENCL_SYCL="0"
build --action_env TF_NEED_ROCM="0"
build --action_env TF_NEED_CUDA="1"
build --action_env CUDA_TOOLKIT_PATH="/usr/local/cuda"
build --action_env TF_CUDA_VERSION="10.0"
build --action_env CUDNN_INSTALL_PATH="/usr/lib/x86_64-linux-gnu"
build --action_env TF_CUDNN_VERSION="7"
build --action_env TENSORRT_INSTALL_PATH="/usr/lib/x86_64-linux-gnu"
build --action_env TF_TENSORRT_VERSION="5.0.0"
build --action_env NCCL_INSTALL_PATH="/usr/lib/x86_64-linux-gnu"
build --action_env NCCL_HDR_PATH="/usr/include"
build --action_env TF_NCCL_VERSION="2"
build --action_env TF_CUDA_COMPUTE_CAPABILITIES="3.5,7.0"
build --action_env LD_LIBRARY_PATH="/usr/local/cuda/extras/CUPTI/lib64:/home/abduld/.gvm/pkgsets/go1.11/global/overlay/lib"
build --action_env TF_CUDA_CLANG="0"
build --action_env GCC_HOST_COMPILER_PATH="/usr/bin/gcc"
build --config=cuda
test --config=cuda
build:opt --copt=-march=native
build:opt --host_copt=-march=native
build:opt --define with_default_optimizations=true
build:v2 --define=tf_api_version=2
```

For macos without gpu, an example `.tf_configure.bazelrc` is

```
build --action_env PYTHON_BIN_PATH="/usr/local/opt/python/bin/python3.7"
build --action_env PYTHON_LIB_PATH="/usr/local/Cellar/python/3.7.0/Frameworks/Python.framework/Versions/3.7/lib/python3.7/site-packages"
build --python_path="/usr/local/opt/python/bin/python3.7"
build:ignite --define with_ignite_support=true
build --define with_xla_support=true
build --action_env TF_NEED_OPENCL_SYCL="0"
build --action_env TF_NEED_ROCM="0"
build --action_env TF_NEED_CUDA="0"
build --action_env TF_DOWNLOAD_CLANG="0"
build:opt --copt=-march=native
build:opt --host_copt=-march=native
build:opt --define with_default_optimizations=true
build:v2 --define=tf_api_version=2
```

Then run

```bash
bazel build -c opt //tensorflow:libtensorflow.so
cp ${GOPATH}/src/github.com/tensorflow/tensorflow/bazel-bin/tensorflow/libtensorflow.so /opt/tensorflow/lib
```

Need to put the directory that contains `libtensorflow_framework.so` and `libtensorflow.so` into `$PATH`.

### PowerPC

For TensorFlow compilation, here are the recommended tensorflow-configure settings:

```
export CC_OPT_FLAGS="-mcpu=power8 -mtune=power8"
export GCC_HOST_COMPILER_PATH=/usr/bin/gcc

ANACONDA_HOME=$(conda info --json | python -c "import sys, json; print json.load(sys.stdin)['default_prefix']")
export PYTHON_BIN_PATH=$ANACONDA_HOME/bin/python
export PYTHON_LIB_PATH=$ANACONDA_HOME/lib/python2.7/site-packages

export USE_DEFAULT_PYTHON_LIB_PATH=0
export TF_NEED_CUDA=1
export TF_CUDA_VERSION=9.0
export CUDA_TOOLKIT_PATH=/usr/local/cuda-9.0
export TF_CUDA_COMPUTE_CAPABILITIES=3.5,3.7,5.2,6.0,7.0
export CUDNN_INSTALL_PATH=/usr/local/cuda-9.0
export TF_CUDNN_VERSION=7
export TF_NEED_GCP=1
export TF_NEED_OPENCL=0
export TF_NEED_HDFS=1
export TF_NEED_JEMALLOC=1
export TF_ENABLE_XLA=1
export TF_CUDA_CLANG=0
export TF_NEED_MKL=0
export TF_NEED_MPI=0
export TF_NEED_VERBS=0
export TF_NEED_GDR=0
export TF_NEED_S3=0
```

### Issues

- Install tensorflow 1.12.0 with CUDA 10.0

Build from source -> build the pip package -> GPU support -> bazel build -> ERROR: Config value cuda is not defined in any .rc file https://github.com/tensorflow/tensorflow/issues/23401
