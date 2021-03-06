/*
Copyright 2016 The TensorFlow Authors. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package predictor

// #cgo LDFLAGS: -L/opt/tensorflow/lib -ltensorflow -lstdc++
// #cgo !nogpu LDFLAGS: -L/usr/local/cuda/lib64 -lcublas -lcudart -lcudnn -lcurand -lcusparse -lcufft
// #cgo CFLAGS: -I${SRCDIR} -O3 -Wall -g
// #cgo !nogpu CFLAGS: -I/usr/local/cuda/include
import "C"
