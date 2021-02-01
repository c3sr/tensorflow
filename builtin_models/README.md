# builtin_models

This folder contains all the built-in model manifests.
The models are collected from [TensorFlow Models](https://github.com/tensorflow/models).

## How to Add a Model

No code is required to add a new model.

## Update the Built-in Model Catalog

After updating model manifests or adding new model manifests, run `make generate` in the root `tensorflow` (`..`).

# model test

### Image Object Detection

Note: Only record first three detections, excluding background.
| Name                        | Image                                  | Label | Xmin  | Xmax  | Ymin  | Ymax  | Probability |
|:---------------------------:|:--------------------------------------:|:-----:|:-----:|:-----:|:-----:|:-----:|:-----------:|
| SSD_VGG16                   | Post-processing not implemented        |       |       |       |       |       |             |
| SSD_MobileNet_v2_Quantized_300x300_COCO     | Require TFLite         |       |       |       |       |       |             |
| RFCN_ResNet101_COCO         | ./_fixtures/lane_control.jpg           | person| 0.447 | 0.500 | 0.483 | 0.698 | 0.998       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.030 | 0.332 | 0.564 | 0.815 | 0.998       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.697 | 1.000 | 0.589 | 0.993 | 0.995       |
| Faster_RCNN_Resnet50_Lowproposals_COCO| ./_fixtures/lane_control.jpg | car   | 0.018 | 0.341 | 0.570 | 0.809 | 0.998       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.705 | 0.988 | 0.575 | 0.994 | 0.993       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.290 | 0.431 | 0.506 | 0.601 | 0.990       |
| SSDLite_MobileNet_v2_COCO   | ./_fixtures/lane_control.jpg           | car   | 0.722 | 0.996 | 0.594 | 0.991 | 0.918       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.634 | 0.885 | 0.572 | 0.938 | 0.891       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.012 | 0.342 | 0.585 | 0.808 | 0.878       |
| Faster_RCNN_Inception_v2_COCO | ./_fixtures/lane_control.jpg         | car   | 0.022 | 0.335 | 0.582 | 0.814 | 0.993       |
|                             | ./_fixtures/lane_control.jpg           | person| 0.451 | 0.502 | 0.486 | 0.698 | 0.989       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.124 | 0.294 | 0.504 | 0.601 | 0.980       |
| MLPerf_SSD_ResNet34_1200x1200 | ./_fixtures/lane_control.jpg         | car   | 0.019 | 0.330 | 0.585 | 0.795 | 0.993       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.288 | 0.437 | 0.501 | 0.603 | 0.991       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.121 | 0.286 | 0.497 | 0.598 | 0.970       |
| Faster_RCNN_ResNet50_COCO   | ./_fixtures/lane_control.jpg           | car   | 0.018 | 0.341 | 0.570 | 0.809 | 0.998       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.705 | 0.988 | 0.575 | 0.994 | 0.993       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.290 | 0.431 | 0.506 | 0.601 | 0.990       |
| Faster_RCNN_ResNet101_Lowproposals_COCO| ./_fixtures/lane_control.jpg| person| 0.445 | 0.499 | 0.488 | 0.697 | 0.994       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.714 | 0.994 | 0.585 | 0.988 | 0.994       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.017 | 0.340 | 0.572 | 0.812 | 0.994       |
| SSD_MobileNet_v1_COCO       | ./_fixtures/lane_control.jpg           | car   | 0.611 | 0.866 | 0.574 | 0.937 | 0.936       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.010 | 0.334 | 0.589 | 0.817 | 0.911       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.709 | 0.994 | 0.581 | 0.989 | 0.893       |
| SSD_MobileNet_v1_0.75_Depth_300x300_COCO14_Sync| ./_fixtures/lane_control.jpg| car   | 0.001 | 0.334 | 0.576 | 0.811 | 0.740       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.692 | 0.989 | 0.605 | 0.990 | 0.714       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.296 | 0.414 | 0.502 | 0.589 | 0.681       |
| SSD_Inception_v2_COCO       | ./_fixtures/lane_control.jpg           | car   | 0.013 | 0.336 | 0.571 | 0.810 | 0.985       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.615 | 0.869 | 0.561 | 0.936 | 0.940       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.288 | 0.434 | 0.500 | 0.602 | 0.885       |
| Faster_RCNN_ResNet101_Lowproposals_COCO| ./_fixtures/lane_control.jpg| car   | 0.705 | 0.993 | 0.592 | 0.984 | 0.997       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.614 | 0.848 | 0.573 | 0.935 | 0.997       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.099 | 0.288 | 0.502 | 0.598 | 0.997       |
| Faster_RCNN_ResNet101_COCO  | ./_fixtures/lane_control.jpg           | car   | 0.017 | 0.339 | 0.574 | 0.816 | 0.997       |
|                             | ./_fixtures/lane_control.jpg           | person| 0.445 | 0.499 | 0.488 | 0.697 | 0.994       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.714 | 0.994 | 0.585 | 0.988 | 0.994       |
| Faster_RCNN_NAS_COCO        | ./_fixtures/lane_control.jpg           | car   | 0.024 | 0.338 | 0.577 | 0.806 | 0.999       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.614 | 0.838 | 0.571 | 0.933 | 0.998       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.705 | 0.993 | 0.592 | 0.984 | 0.997       |
| Faster_RCNN_Inception_ResNet_v2_Atrous_COCO| ./_fixtures/lane_control.jpg| person| 0.453 | 0.501 | 0.486 | 0.683 | 0.997       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.022 | 0.344 | 0.575 | 0.815 | 0.996       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.717 | 0.998 | 0.586 | 0.989 | 0.990       |
| Faster_RCNN_Inception_ResNet_v2_Atrous_Lowproposals_COCO| ./_fixtures/lane_control.jpg | person   | 0.453 | 0.501 | 0.486 | 0.683 | 0.997       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.022 | 0.344 | 0.575 | 0.815 | 0.996       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.717 | 0.998 | 0.586 | 0.989 | 0.990       |
| SSD_MobileNet_v2_COCO       | ./_fixtures/lane_control.jpg           | car   | 0.620 | 0.868 | 0.563 | 0.923 | 0.958       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.022 | 0.343 | 0.577 | 0.812 | 0.935       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.723 | 0.996 | 0.593 | 0.997 | 0.834       |
| MLPerf_SSD_MobileNet_v1_300x300| ./_fixtures/lane_control.jpg        | car   | 0.010 | 0.333 | 0.587 | 0.813 | 0.897       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.610 | 0.888 | 0.576 | 0.943 | 0.882       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.289 | 0.430 | 0.499 | 0.593 | 0.804       |
| SSD_ResNet50_FPN_Shared_Box_Predictor_640x640_COCO14_Sync| ./_fixtures/lane_control.jpg| car   | 0.015 | 0.342 | 0.578 | 0.809 | 0.858       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.608 | 0.854 | 0.563 | 0.940 | 0.807       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.289 | 0.437 | 0.502 | 0.604 | 0.805       |
| SSD_MobileNet_v1_FPN_Shared_Box_Predictor_640x640_COCO14_Sync| ./_fixtures/lane_control.jpg           | car   | 0.014 | 0.341 | 0.569 | 0.815 | 0.858       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.286 | 0.436 | 0.503 | 0.601 | 0.757       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.497 | 0.560 | 0.538 | 0.598 | 0.738       |
| SSD_MobileNet_v1_PPN_Shared_Box_Predictor_300x300_COCO14_Sync| ./_fixtures/lane_control.jpg  | car   | 0.016 | 0.333 | 0.584 | 0.808 | 0.805       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.136 | 0.301 | 0.500 | 0.595 | 0.793       |
|                             | ./_fixtures/lane_control.jpg           | car   | 0.630 | 0.960 | 0.574 | 0.972 | 0.754       |
