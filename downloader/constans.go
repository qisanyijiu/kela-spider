package downloader

type State int

const (
	waiting State = 10
	running State = 20
	finish  State = 30
)

func (s *State) Desc () string {
	switch *s {
	case waiting:
		return "waiting"
	case running:
		return "running"
	case finish:
		return "finish"
	default:
		return "waiting"
	}
}
