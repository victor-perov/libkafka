package Fetch

type Response struct {
	ThrottleTimeMs int32
	TopicResponses []TopicResponse
}

func (r *Response) PartitionResponse() *PartitionResponse {
	defer func() { recover() }()
	return &(r.TopicResponses[0].PartitionResponses[0])
}

type TopicResponse struct {
	Topic              string
	PartitionResponses []PartitionResponse
}

type PartitionResponse struct {
	Partition           int32
	ErrorCode           int16
	HighWatermark       int64
	LastStableOffset    int64
	LogStartOffset      int64
	AbortedTransactions []AbortedTransaction
	//
	RecordSet []byte // NULLABLE_BYTES
}

type AbortedTransaction struct {
	ProducerId  int64
	FirstOffset int64
}
