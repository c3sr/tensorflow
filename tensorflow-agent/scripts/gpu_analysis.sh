#!/usr/bin/env bash

DATABASE_ADDRESS=$1
BATCHSIZE=$2
MODELNAME=$3
OUTPUTFOLDER=output_gpu
DATABASE_NAME=carml

cd ..

if [ ! -d $OUTPUTFOLDER ]; then
  mkdir $OUTPUTFOLDER
fi

if [ -f tensorflow-agent ]; then
  rm tensorflow-agent
fi
go build -tags=nolibjpeg

echo "Start to run model analysis"

./tensorflow-agent evaluation model info --database_address=$DATABASE_ADDRESS --database_name=$DATABASE_NAME --model_name=$MODELNAME --sort_output --format=csv,table --plot_all --output="$OUTPUTFOLDER/$MODELNAME/model_info"

echo "Start to run layer analysis"

./tensorflow-agent evaluation layer info --database_address=$DATABASE_ADDRESS --database_name=$DATABASE_NAME --model_name=$MODELNAME --batch_size=$BATCHSIZE --sort_output --format=csv,table --plot_all --output="$OUTPUTFOLDER/$MODELNAME/$BATCHSIZE/layer_info"

./tensorflow-agent evaluation layer aggre_info --database_address=$DATABASE_ADDRESS --database_name=$DATABASE_NAME --model_name=$MODELNAME --batch_size=$BATCHSIZE --sort_output --format=csv,table --plot_all --output="$OUTPUTFOLDER/$MODELNAME/$BATCHSIZE/layer_aggre_info"

echo "Start to run gpu analysis"

./tensorflow-agent evaluation gpu_kernel info --database_address=$DATABASE_ADDRESS --database_name=$DATABASE_NAME --model_name=$MODELNAME --batch_size=$BATCHSIZE --sort_output --format=csv,table --output="$OUTPUTFOLDER/$MODELNAME/$BATCHSIZE/gpu_kernel_info"

./tensorflow-agent evaluation gpu_kernel name_aggre_info --database_address=$DATABASE_ADDRESS --database_name=$DATABASE_NAME --model_name=$MODELNAME --batch_size=$BATCHSIZE --sort_output --format=csv,table --output="$OUTPUTFOLDER/$MODELNAME/$BATCHSIZE/gpu_kernel_name_aggre_info"

./tensorflow-agent evaluation gpu_kernel model_aggre_info --database_address=$DATABASE_ADDRESS --database_name=$DATABASE_NAME --model_name=$MODELNAME --batch_size=$BATCHSIZE --sort_output --format=csv,table --output="$OUTPUTFOLDER/$MODELNAME/$BATCHSIZE/gpu_kernel_model_aggre_info"

./tensorflow-agent evaluation gpu_kernel layer_aggre_info --database_address=$DATABASE_ADDRESS --database_name=$DATABASE_NAME --model_name=$MODELNAME --batch_size=$BATCHSIZE --sort_output --format=csv,table --plot_all --output="$OUTPUTFOLDER/$MODELNAME/$BATCHSIZE/gpu_kernel_layer_aggre_info"
