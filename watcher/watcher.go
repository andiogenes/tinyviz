package watcher

import (
	"os"
	"time"
)

// Watcher watches a set of files and perform action on it.
type Watcher struct {
	Errors chan error

	files      []*os.File
	fileQueue  chan *os.File
	quit       chan int
	fileAction func(*Watcher, *os.File) error

	modifyTimes map[string]time.Time
}

// NewWatcher creates new watcher object with specific action.
func NewWatcher(action func(*Watcher, *os.File) error) *Watcher {
	return &Watcher{
		Errors: make(chan error),

		files:      make([]*os.File, 0),
		fileQueue:  make(chan *os.File),
		quit:       make(chan int),
		fileAction: action,

		modifyTimes: make(map[string]time.Time),
	}
}

// NewModTimeWatcher creates new watcher object which perform
// specific action when modification time of file changes.
func NewModTimeWatcher(action func(*os.File) error) *Watcher {
	return NewWatcher(func(w *Watcher, f *os.File) error {
		stat, err := f.Stat()
		if err != nil {
			return err
		}
		curTime := stat.ModTime()

		prevTime := w.modifyTimes[f.Name()]

		if !curTime.Equal(prevTime) {
			if err := action(f); err != nil {
				return err
			}

			w.modifyTimes[f.Name()] = curTime
		}

		return nil
	})
}

// AddFile append file with specific name to the watching list.
func (w *Watcher) AddFile(fileName string) error {
	f, err := os.Open(fileName)
	if err == nil {
		go func() { w.fileQueue <- f }()
	}
	return err
}

// Close stops watching on files.
func (w *Watcher) Close() {
	w.quit <- 0
}

// Run starts watching on files with specific duration.
func (w *Watcher) Run(duration time.Duration) {
	go func() {
		for {
			select {
			case f := <-w.fileQueue:
				w.files = append(w.files, f)
			case <-w.quit:
				for _, file := range w.files {
					file.Close()
				}
				return
			default:
				for _, file := range w.files {
					if err := w.fileAction(w, file); err != nil {
						w.Errors <- err
						break
					}
				}

				time.Sleep(duration)
			}
		}
	}()
}
