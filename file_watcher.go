package main

import (
	"github.com/fsnotify/fsnotify"
	"github.com/ns-cn/keeper/env"
	"github.com/robfig/cron/v3"
	"log"
)

func watchForUpdate(cron *cron.Cron) {
	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Start listening for events.
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) {
					cron.Stop()
					for _, entry := range cron.Entries() {
						cron.Remove(entry.ID)
					}
					updateInFiles(cron)
					cron.Start()
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()
	err = watcher.Add(env.CfgFile.Value)
	if err != nil {
		log.Fatal(err)
	}
	select {}
}
