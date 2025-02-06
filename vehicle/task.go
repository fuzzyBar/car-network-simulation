package vehicle

type Task struct {
	ID          int
	Size        int // 任务大小
	ResourceReq int // 任务需要的计算资源
	Remaining   int // 剩余处理时间
}
