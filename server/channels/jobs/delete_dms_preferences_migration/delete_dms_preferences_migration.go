// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package delete_dms_preferences_migration

import (
	"time"

	"github.com/mattermost/mattermost/server/public/model"
	"github.com/mattermost/mattermost/server/v8/channels/jobs"
	"github.com/mattermost/mattermost/server/v8/channels/store"
	"github.com/pkg/errors"
)

const (
	timeBetweenBatches = 1 * time.Second
)

// MakeWorker creates a batch migration worker to delete empty drafts.
func MakeWorker(jobServer *jobs.JobServer, store store.Store, app jobs.BatchMigrationWorkerAppIFace) model.Worker {
	return jobs.MakeBatchMigrationWorker(
		jobServer,
		store,
		app,
		model.MigrationKeyDeleteDmsPreferences,
		timeBetweenBatches,
		doDeleteDmsPreferencesMigrationBatch,
	)
}

// doDeleteDmsPreferencesMigrationBatch deletes any limit_visible_dms_gms preferences with a value > 40
func doDeleteDmsPreferencesMigrationBatch(data model.StringMap, store store.Store) (model.StringMap, bool, error) {
	rowAffected, err := store.Preference().DeleteVisibleDmsGms()
	if err != nil {
		return nil, false, errors.Wrapf(err, "failed to delete limit_visible_dms_gms with a value > 40")
	}

	if rowAffected == 0 {
		return nil, true, nil
	}

	return nil, false, nil
}