package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

func main() {
	mutexExample()
}

func dirWalk() {
	filepath.Walk("/Volumes/Workspace/GoWorkspace", func(path string, info os.FileInfo, err error) error {
		fmt.Println(path, "|", info.Name(), "Size", info.Size(), "IsDir", info.IsDir(), "file mode", info.Mode())
		return nil
	})
}

func mutexExample() {
	m := new(sync.Mutex)

	for i := 0; i < 10; i++ {
		go func(i int) {
			m.Lock()
			fmt.Println("Start lock", i)
			time.Sleep(time.Millisecond * 500)
			fmt.Println("End lock", i)
			m.Unlock()
		}(i)
	}

	var input string
	fmt.Scanln(&input)
}
