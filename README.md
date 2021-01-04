# MLModelScope TensorFlow Agent

[![Build Status](https://dev.azure.com/yhchang/c3sr/_apis/build/status/c3sr.tensorflow?branchName=master)](https://dev.azure.com/yhchang/c3sr/_build/latest?definitionId=5&branchName=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/c3sr/tensorflow)](https://goreportcard.com/report/github.com/c3sr/tensorflow)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

This is the TensorFlow agent for [MLModelScope](mlmodelscope.org), an open-source framework and hardware agnostic, extensible and customizable platform for evaluating and profiling ML models across datasets / frameworks / systems, and within AI application pipelines.

Currently it has most of the models from TensorFlow Model Zoo built in, including Image Classification, Object Detection, Instance Segmentation, Semantic Segmentation, Image Enhancement. More built-in models are coming. 
One can evaluate the **~80** models on any systems of insterest with either local TensorFlow installation or TensorFlow docker images. 

Check out [MLModelScope](mlmodelscope.org) and welcome to contribute.

# Bare Minimum Installation

## Prerequsite System Library Installation
We first discuss a bare minimum tensorflow-agent installation without the tracing and profiling capabilities. To make this work, you will need to have the following system libraries preinstalled in your system. 

- The CUDA library (required)
- The CUPTI library (required)
- The Tensorflow library (required)
- The libjpeg-turbo library (optional, but preferred)

### The CUDA Library

Please refer to Nvidia CUDA library installation on this. Find the location of your local CUDA installation, which is typically at `/usr/local/cuda/`, and setup the path to the `libcublas.so` library. Place the following in either your `~/.bashrc` or `~/.zshrc` file:

```
export LIBRARY_PATH=$LIBRARY_PATH:/usr/local/cuda/lib64
export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/usr/local/cuda/lib64
```

### The CUPTI Library

Please refer to Nvidia CUPTI library installation on this. Find the location of your local CUPTI installation, which is typically at `/usr/local/cuda/extras/CUPTI`, and setup the path to the `libcupti.so` library. Place the following in either your `~/.bashrc` or `~/.zshrc` file:

```
export LIBRARY_PATH=$LIBRARY_PATH:/usr/local/cuda/extras/CUPTI/lib64
export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/usr/local/cuda/extras/CUPTI/lib64
```

### The TensorFlow C Library

The TensorFlow C library is required for the TensorFlow Go package. If you want to use TensorFlow Docker Images (e.g. NVIDIA GPU CLOUD (NGC)) instead, skip this step for now and refer to our later section on this.

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

You can test the installed TensorFlow C library using an [example C program](https://www.tensorflow.org/install/lang_c#build).

To build the TensorFlow C library from source, refer to [TensorFlow in Go](https://github.com/tensorflow/tensorflow/tree/master/tensorflow/go#building-the-tensorflow-c-library-from-source) .


### Use libjpeg-turbo for Image Preprocessing

[libjpeg-turbo](https://github.com/libjpeg-turbo/libjpeg-turbo) is a JPEG image codec that uses SIMD instructions (MMX, SSE2, AVX2, NEON, AltiVec) to accelerate baseline JPEG compression and decompression. It outperforms libjpeg by a significant amount.

You need libjpeg installed.  
```
sudo apt-get install libjpeg-dev  
```
The default is to use libjpeg-turbo, to opt-out, use build tag `nolibjpeg`.

To install libjpeg-turbo, refer to [libjpeg-turbo](https://github.com/libjpeg-turbo/libjpeg-turbo/releases).

Linux

```
  export TURBO_VER=2.0.2
  cd /tmp
  wget https://cfhcable.dl.sourceforge.net/project/libjpeg-turbo/${TURBO_VER}/libjpeg-turbo-official_${TURBO_VER}_amd64.deb
  sudo dpkg -i libjpeg-turbo-official_${TURBO_VER}_amd64.deb
```

macOS

```
brew install jpeg-turbo
```


## Installation of Go for Compilation

Since we use `go` for MLModelScope development, it's required to have `go` installed in your system before proceed.

Please follow [Installing Go Compiler](https://github.com/c3sr/rai/blob/master/docs/developer_guide.md) to have `go` installed.


## Bare Minimum Tensorflow-agent Installation

Download and install the MLModelScope TensorFlow Agent by running the following command in any location, assuming you have installed `go` following the above instruction.

```
go get -v github.com/c3sr/tensorflow

```

You can then install the dependency packages through `go get`.

```
cd $GOPATH/src/github.com/c3sr/tensorflow
go get -u -v ./...
```

An alternative to install the dependency packages is to use [Dep](https://github.com/golang/dep).

```
dep ensure -v
```

This installs the dependency in `vendor/`.

The CGO interface passes go pointers to the C API. There is an error in the CGO runtime. We can disable the error by placing

```
export GODEBUG=cgocheck=0
```

in your `~/.bashrc` or `~/.zshrc` file and then run either `source ~/.bashrc` or `source ~/.zshrc`


Build the TensorFlow agent with GPU enabled
```
cd $GOPATH/src/github.com/c3sr/tensorflow/tensorflow-agent
go build 
```

Build the TensorFlow agent without GPU or libjpeg-turbo
```
cd $GOPATH/src/github.com/c3sr/tensorflow/tensorflow-agent
go build -tags="nogpu nolibjpeg" 
```

If everything is successful, you should have an executable `tensorflow-agent` binary in the current directory.

### Configuration Setup

To run the agent, you need to setup the correct configuration file for the agent. Some of the information may not make perfect sense for all testing scenarios, but they are required and will be needed for later stage testing. Some of the port numbers as specified below can be changed depending on your later setup for those service. 

So let's just set them up as is, and worry about the detailed configuration parameter values later. 

You must have a `carml` config file called `.carml_config.yml` under your home directory. An example config file `carml_config.yml.example` is in [github.com/c3sr/MLModelScope](https://github.com/c3sr/MLModelScope) . You can move it to `~/.carml_config.yml`.

The following configuration file can be placed in `$HOME/.carml_config.yml` or can be specified via the `--config="path"` option.

```yaml
app:
  name: carml
  debug: true
  verbose: true
  tempdir: ~/data/carml
registry:
  provider: consul
  endpoints:
    - localhost:8500
  timeout: 20s
  serializer: jsonpb
database:
  provider: mongodb
  endpoints:
    - localhost
tracer:
  enabled: true
  provider: jaeger
  endpoints:
    - localhost:9411
  level: FULL_TRACE
logger:
  hooks:
    - syslog
```

## Test Installation

With the configuration and the above  bare minimumn installation, you should be ready to test the installation and see how things work. 

Here are a few examples. First, make sure we are in the right location
```
cd $GOPATH/src/github.com/c3sr/tensorflow/tensorflow-agent
```

To see a list of help
```
./tensorflow-agent -h
```

To see a list of models that we can run with this agent
```
./tensorflow-agent info models
```

To run an inference using the default DNN model `mobilenet_v1_1.0_224` with a default input image.

```
./tensorflow-agent predict urls --profile=false --publish=false
```

The above `--profile=false --publish=false` command parameters tell the agent that we do not want to use profiling capability and publish the results, as we haven't installed the MongoDB database to store profiling data and the tracer service to accept tracing information.

# External Service Installation to Enable Tracing and Profiling

We now discuss how to install a few external services that make the agent fully useful in terms of collecting tracing and profiling data.

## External Srvices

MLModelScope relies on a few external services. These services provide tracing, registry, and database servers.

These services can be installed and enabled in different ways. We discuss how we use `docker` below to show how this can be done. You can also not use `docker` but install those services from either binaries or source codes directly.

### Installing Docker

Refer to [Install Docker](https://docs.docker.com/install/).

On Ubuntu, an easy way is using

```
curl -fsSL get.docker.com -o get-docker.sh | sudo sh
sudo usermod -aG docker $USER
```

On macOS, [intsall Docker Destop](https://docs.docker.com/docker-for-mac/install/)


### Starting Trace Server

This service is required.

- On x86 (e.g. intel) machines, start [jaeger](http://jaeger.readthedocs.io/en/latest/getting_started/) by

```
docker run -d -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 -p5775:5775/udp -p6831:6831/udp -p6832:6832/udp \
  -p5778:5778 -p16686:16686 -p14268:14268 -p9411:9411 jaegertracing/all-in-one:latest
```

- On ppc64le (e.g. minsky) machines, start [jaeger](http://jaeger.readthedocs.io/en/latest/getting_started/) machine by

```
docker run -d -e COLLECTOR_ZIPKIN_HTTP_PORT=9411 -p5775:5775/udp -p6831:6831/udp -p6832:6832/udp \
  -p5778:5778 -p16686:16686 -p14268:14268 -p9411:9411 carml/jaeger:ppc64le-latest
```

The trace server runs on http://localhost:16686

### Starting Registry Server

This service is not required if using TensorFlow-agent for local evaluation.

- On x86 (e.g. intel) machines, start [consul](https://hub.docker.com/_/consul/) by

```
docker run -p 8500:8500 -p 8600:8600 -d consul
```

- On ppc64le (e.g. minsky) machines, start [consul](https://hub.docker.com/_/consul/) by

```
docker run -p 8500:8500 -p 8600:8600 -d carml/consul:ppc64le-latest
```

The registry server runs on http://localhost:8500

### Starting Database Server

This service is not required if not using database to publish evaluation results.

- On x86 (e.g. intel) machines, start [mongodb](https://hub.docker.com/_/mongo/) by

```
docker run -p 27017:27017 --restart always -d mongo:3.0
```

You can also mount the database volume to a local directory using

```
docker run -p 27017:27017 --restart always -d  -v $HOME/data/carml/mongo:/data/db mongo:3.0
```

### Configuration

You must have a `carml` config file called `.carml_config.yml` under your home directory. An example config file `~/.carml_config.yml` is already discussed above. Please update the port numbers for the above external services accordingly if you decide to choose a different ports above.


### Testing

The testing steps are very similar to those testing we discussed above, except that you can now safely use both the profiling and publishing services.


# Alternative ways to Install Tensorflow-agent using a Tensorflow Docker Image

Instead of using a local Tensorflow library to install the MLModelScope `tensorflow-agent`, we can also use a Tensorflow docker image to start this process. 

## Local Setup

You need to follow the above similar procedures to setup `go` and get all the related `rai-project` projects in your local go development environment.

```
go get -v github.com/c3sr/tensorflow
cd $GOPATH/src/github.com/c3sr/tensorflow
go get -u -v ./...
```

You also need to have the `.carml_config.yml` configuraiton file as discussed above to be placed under $HOME as `.carml_config.yml`

You can also setup all the external services as discussed above in your local host machine where you plan to use the Tensorflow Docker container.

## Setup within TensorFlow Docker Image

Continue if you have

* installed all the dependencies
* downloaded carml_config_example.yml to $HOME as .carml_config.yml
* launched docker external services on the host machine of the docker container you are going to use

, otherwise read above

Assuming you want to use the NGC TensorFlow docker image. Here is an example on how to do this: 

```
docker run --gpus=all --shm-size=1g --ulimit memlock=-1 --ulimit stack=67108864 -it --privileged=true --network host \
-v $GOPATH:/workspace/go1.12/global \
-v $GOROOT:/workspace/go1.12_root \
-v ~/.carml_config.yml:/root/.carml_config.yml \
nvcr.io/nvidia/tensorflow:19.06-py3
```

NOTE: The SHMEM allocation limit is set to the default of 64MB.  This may be insufficient for TensorFlow. NVIDIA recommends the use of the following flags:
   ```--shm-size=1g --ulimit memlock=-1 --ulimit stack=67108864 ...```

Within the container, set up the environment so that the agent can find the TensorFlow C library.

```
export GOPATH=/workspace/go1.12/global
export GOROOT=/workspace/go1.12_root
export PATH=$GOROOT/bin:$PATH

ln -s /usr/local/lib/tensorflow/libtensorflow_cc.so /usr/local/lib/tensorflow/libtensorflow.so
export CGO_LDFLAGS="${CGO_LDFLAGS} -L /usr/local/lib/tensorflow/"
export CGO_CFLAGS="${CGO_CFLAGS} -I /opt/tensorflow/tensorflow-source"
export CGO_CXXFLAGS="${CGO_CXXFLAGS} -I /opt/tensorflow/tensorflow-source"

export PATH=$PATH:$(go env GOPATH)/bin  
export GODEBUG=cgocheck=0  
```

Build the TensorFlow agent with GPU enabled
```
cd $GOPATH/src/github.com/c3sr/tensorflow/tensorflow-agent
go build 
```

Build the TensorFlow agent without GPU or libjpeg-turbo
```
cd $GOPATH/src/github.com/c3sr/tensorflow/tensorflow-agent
go build -tags="nogpu nolibjpeg" 
```


# Use the Agent with the [MLModelScope Web UI](https://github.com/c3sr/mlmodelscope)

```
./tensorflow-agent serve -l -d -v
```

Refer to [TODO] to run the web UI to interact with the agent.

# Use the Agent through Command Line

Run ```./tensorflow-agent -h``` to list the available commands.

Run ```./tensorflow-agent info models``` to list the available models.

Run ```./tensorflow-agent predict``` to evaluate a model. This runs the default evuation. 
```./tensorflow-agent predict -h``` shows the available flags you can set.

An example run is

```
./tensorflow-agent predict urls --trace_level=FRAMEWORK_TRACE --model_name=Inception_v3
```

Refer to [TODO] to run the web UI to interact with the agent.

# Use the Agent through Pre-built Docker Images

We have [pre-built docker images](https://hub.docker.com/r/carml/tensorflow/tags) on Dockerhub. The images are `carml/tensorflow:amd64-cpu-latest`, `carml/tensorflow:amd64-gpu-latest` and `carml/tensorflow:amd64-gpu-ngc-latest`. The entrypoint is set as `tensorflow-agent` thus these images act similar as the command line above.

An example run is

```
docker run --gpus=all --shm-size=1g --ulimit memlock=-1 --ulimit stack=67108864 --privileged=true \
    --network host \
    -v ~/.carml_config.yml:/root/.carml_config.yml \ 
    -v ~/results:/go/src/github.com/c3sr/tensorflow/results \
    carml/tensorflow:amd64-gpu-latest predict urls --trace_level=FRAMEWORK_TRACE --model_name=Inception_v3
```
NOTE: The SHMEM allocation limit is set to the default of 64MB.  This may be insufficient for TensorFlow.  NVIDIA recommends the use of the following flags:
   ```--shm-size=1g --ulimit memlock=-1 --ulimit stack=67108864 ...```

NOTE: To run with GPU, you need to meet following requirements:

- Docker >= 19.03 with nvidia-container-toolkit (otherwise need to use nvidia-docker)
- CUDA >= 10.1
- NVIDIA Driver >= 418.39 

# Notes on installing TensorFlow from source (ignore this if you are a user)

## Install Bazel

- [Installing Bazel on Ubuntu](https://docs.bazel.build/versions/master/install-ubuntu.html)

- [Installing Bazel on macOS](https://docs.bazel.build/versions/master/install-os-x.html#install-on-mac-os-x-homebrew)

## Build

Build TensorFlow 1.14.0 with the following scripts.

```sh
go get -d github.com/tensorflow/tensorflow/tensorflow/go
cd ${GOPATH}/src/github.com/tensorflow/tensorflow
git fetch --all
git checkout v1.14.0
./configure
```

Configure the build and then run 

```
bazel build -c opt //tensorflow:libtensorflow.so
cp ${GOPATH}/src/github.com/tensorflow/tensorflow/bazel-bin/tensorflow/libtensorflow.so /opt/tensorflow/lib
```

Need to put the directory that contains `libtensorflow_framework.so` and `libtensorflow.so` into `$PATH`.

## PowerPC

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
