package parallel

import "log"

func Go(f func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println("service panic:", err)
			}
		}()
		f()
	}()
}
