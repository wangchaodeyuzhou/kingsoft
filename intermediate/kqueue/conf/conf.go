package conf

var bootstrap *Bootstrap
var services = make(map[string][]WorkerNode)

type Bootstrap struct {
	Services *WorkerServices
}

type WorkerServices struct {
	Workers []Worker
}

// Worker 一个 worker 有多个 workerNode
type Worker struct {
	WorkerId string
	Node     []WorkerNode
}

type WorkerNode struct {
	Id  string
	URL string
}

func DefaultConfig() any {
	return &Bootstrap{
		Services: defaultWorkerServices(),
	}
}

func defaultWorkerServices() *WorkerServices {
	return &WorkerServices{
		Workers: make([]Worker, 0),
	}
}
