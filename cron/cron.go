package cron

import (
	"github.com/robfig/cron/v3"
	"interest-arbitrage/service"
	"log"
)

func RegisterCronFunc(c *cron.Cron) {
	entryId, err := c.AddFunc("*/1 * * * *", service.SyncETHBridgeLockMtsByAll)
	if err != nil {
		log.Fatalf("failed to add job, func: %v, err: %v", "SyncETHBridgeLockMtsByAll", err)
	}
	log.Printf("success to add job, func: %v, EntryID: %v", "SyncETHBridgeLockMtsByAll", entryId)

	entryId, err = c.AddFunc("*/1 * * * *", service.SyncMTSBridgeLockMtsByAll)
	if err != nil {
		log.Fatalf("failed to add job, func: %v, err: %v", "SyncMTSBridgeLockMtsByAll", err)
	}
	log.Printf("success to add job, func: %v, EntryID: %v", "SyncMTSBridgeLockMtsByAll", entryId)
}
