package window

type AddSampleFunc func(current, new interface{}) interface{}

type AggregateSingleFunc func(value interface{}, count int) interface{}

type AggregateBlocksFunc func(block []interface{}, size int) interface{}

