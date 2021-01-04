package predictor

import (
	"context"
	"encoding/json"
	"errors"
	"regexp"
	"strings"
	"time"

	proto "github.com/c3sr/tensorflow"
	graph "github.com/c3sr/tensorflow/graph"
	"github.com/c3sr/tracer"
	opentracing "github.com/opentracing/opentracing-go"
)

type traceNode struct {
	device         string
	node           *proto.NodeExecStats
	timelineOpName string
	graphOpName    string
}

type Trace struct {
	nodes []traceNode
}

// https://github.com/tensorflow/tensorflow/blob/master/tensorflow/python/client/timeline.py
func NewTrace(data *proto.StepStats, graphPath string) (*Trace, error) {
	if len(data.GetDevStats()) == 0 {
		return nil, errors.New("no device stats available")
	}
	nodes := []traceNode{}
	g, err := graph.New(graphPath)
	if err != nil {
		return nil, err
	}
	for _, dev := range data.GetDevStats() {
		for _, nd := range dev.GetNodeStats() {
			timelineOpName, _, _ := parseOpLabel(nd.GetTimelineLabel())
			if timelineOpName == "Const" {
				continue
			}
			traceNode := traceNode{
				device:         dev.GetDevice(),
				node:           nd,
				timelineOpName: timelineOpName,
			}
			graphOpName, err := g.LookUpNodeOperatorName(nd.GetNodeName())
			if err == nil {
				traceNode.graphOpName = graphOpName
			}
			nodes = append(nodes, traceNode)
		}
	}
	return &Trace{
		nodes: nodes,
	}, nil
}

func parseOpLabel(name string) (op string, inputs []string, output string) {
	re := regexp.MustCompile(`(?m)(.*) = (.*)\((.*)\)`)
	allMatches := re.FindAllStringSubmatch(name, -1)
	if len(allMatches) != 1 {
		return
	}
	matches := allMatches[0]
	if len(matches) != 4 {
		return
	}
	output, op, sinputs := matches[1], matches[2], matches[3]
	inputs = strings.Split(sinputs, ",")
	return
}

// Notes about start and end time from the NodeExecStats proto:
// For GPU, there is no difference between op_end_rel_micros and
// all_end_rel_micros. All are kernel times.
// For CPU, op_end_rel is the kernel time, while all_end_rel_micros includes
// some post-processing. Besides, currently, there is no way to measure
// the execution time of async ops accurately.
func (t *Trace) Publish(ctx context.Context, opts ...opentracing.StartSpanOption) error {
	for layerSequenceIndex, tr := range t.nodes {
		device := tr.device
		graphOpName := tr.graphOpName
		timelineOpName := tr.timelineOpName
		node := tr.node
		startTime := time.Unix(0, node.GetAllStartMicros()*int64(time.Microsecond))
		endTime := time.Unix(0, (node.GetAllStartMicros()+node.GetAllEndRelMicros())*int64(time.Microsecond))

		// skip if the span is too small
		// if endTime.Sub(startTime) <= 10*time.Microsecond {
		// 	continue
		// }

		tags := opentracing.Tags{
			"trace_source":         "framework",
			"framework_name":       "tensorflow",
			"layer_sequence_index": layerSequenceIndex,
			"op_name":              timelineOpName,
			"static_op_name":       graphOpName,
			"device":               device,
			// "all_start_micros":    node.GetAllStartMicros(),
			// "all_end_rel_micros":  node.GetAllEndRelMicros(),
			// "op_start_rel_micros": node.GetOpStartRelMicros(),
			// "op_end_rel_micros":   node.GetOpEndRelMicros(),
			"timeline_label":   node.GetTimelineLabel(),
			"scheduled_micros": node.GetScheduledMicros(),
			"thread_id":        node.GetThreadId(),
			// "start_time":          startTime,
			// "end_time":            endTime,
		}

		if len(node.GetOutput()) != 0 {
			shapes := make([][]int64, len(node.GetOutput()))
			for jj, o := range node.GetOutput() {
				ss := o.GetTensorDescription().GetShape().GetDim()
				if len(ss) == 0 {
					shapes[jj] = []int64{}
					continue
				}
				shape := make([]int64, len(ss))
				for ii, s := range ss {
					shape[ii] = s.GetSize_()
				}
				shapes[jj] = shape
			}
			tags["shape"] = shapes
		}

		memStatsTags := opentracing.Tags{}
		if node.GetMemoryStats() != nil {
			stats := node.GetMemoryStats()
			memStatsTags = opentracing.Tags{
				"temp_memory_size":                   stats.GetTempMemorySize(),
				"persistent_memory_size":             stats.GetPersistentMemorySize(),
				"device_temp_memory_size":            stats.GetDeviceTempMemorySize(),
				"device_persistent_memory_size":      stats.GetDevicePersistentMemorySize(),
				"host_persistent_tensor_alloc_ids":   stats.GetPersistentTensorAllocIds(),
				"device_persistent_tensor_alloc_ids": stats.GetDevicePersistentTensorAllocIds(),
			}
		}

		outputTags := opentracing.Tags{}
		if node.GetOutput() != nil {
			stats := node.GetOutput()
			bts, err := json.Marshal(stats)
			if err != nil {
				break
			}
			outputTags = opentracing.Tags{
				"output": string(bts),
			}
		}

		memAllocTags := opentracing.Tags{}
		if node.GetMemory() != nil {
			stats := node.GetMemory()
			bts, err := json.Marshal(stats)
			if err != nil {
				break
			}
			memAllocTags = opentracing.Tags{
				"memory": string(bts),
			}
		}

		// NOT USED, SAME INFO IN OUTPUT
		// tensorTags := opentracing.Tags{}
		// if node.GetReferencedTensor() != nil {
		// 	stats := node.GetReferencedTensor()
		// 	bts, err := json.Marshal(stats)
		// 	if err != nil {
		// 		break
		// 	}
		// 	tensorTags = opentracing.Tags{
		// 		"tensor": string(bts),
		// 	}
		// }

		s, _ := tracer.StartSpanFromContext(
			ctx,
			tracer.FRAMEWORK_TRACE,
			node.GetNodeName(),
			opentracing.StartTime(startTime),
			tags,
			memStatsTags,
			outputTags,
			memAllocTags,
			// tensorTags,
		)
		if s == nil {
			log.WithField("node_name", node.GetNodeName()).
				WithField("tags", tags).
				Error("failed to create span from context")
			continue
		}
		s.FinishWithOptions(opentracing.FinishOptions{
			FinishTime: endTime,
		})
	}

	return nil
}
