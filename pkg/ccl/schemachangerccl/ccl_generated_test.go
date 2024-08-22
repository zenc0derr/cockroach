// Copyright 2022 The Cockroach Authors.
//
// Licensed as a CockroachDB Enterprise file under the Cockroach Community
// License (the "License"); you may not use this file except in compliance with
// the License. You may obtain a copy of the License at
//
//     https://github.com/cockroachdb/cockroach/blob/master/licenses/CCL.txt

// Code generated by sctestgen, DO NOT EDIT.

package schemachangerccl

import (
	"testing"

	"github.com/cockroachdb/cockroach/pkg/sql/schemachanger/sctest"
	"github.com/cockroachdb/cockroach/pkg/util/leaktest"
	"github.com/cockroachdb/cockroach/pkg/util/log"
)

func TestBackupRollbacks_ccl_alter_index_configure_zone(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/alter_index_configure_zone"
	sctest.BackupRollbacks(t, path, MultiRegionTestClusterFactory{})
}

func TestBackupRollbacks_ccl_alter_index_configure_zone_multiple(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/alter_index_configure_zone_multiple"
	sctest.BackupRollbacks(t, path, MultiRegionTestClusterFactory{})
}

func TestBackupRollbacks_ccl_create_index(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/create_index"
	sctest.BackupRollbacks(t, path, MultiRegionTestClusterFactory{})
}

func TestBackupRollbacks_ccl_drop_database_multiregion_primary_region(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_database_multiregion_primary_region"
	sctest.BackupRollbacks(t, path, MultiRegionTestClusterFactory{})
}

func TestBackupRollbacks_ccl_drop_table_multiregion(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_table_multiregion"
	sctest.BackupRollbacks(t, path, MultiRegionTestClusterFactory{})
}

func TestBackupRollbacks_ccl_drop_table_multiregion_primary_region(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_table_multiregion_primary_region"
	sctest.BackupRollbacks(t, path, MultiRegionTestClusterFactory{})
}

func TestBackupRollbacksMixedVersion_ccl_alter_index_configure_zone(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/alter_index_configure_zone"
	sctest.BackupRollbacksMixedVersion(t, path, MultiRegionTestClusterFactory{})
}

func TestBackupRollbacksMixedVersion_ccl_alter_index_configure_zone_multiple(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/alter_index_configure_zone_multiple"
	sctest.BackupRollbacksMixedVersion(t, path, MultiRegionTestClusterFactory{})
}

func TestBackupRollbacksMixedVersion_ccl_create_index(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/create_index"
	sctest.BackupRollbacksMixedVersion(t, path, MultiRegionTestClusterFactory{})
}

func TestBackupRollbacksMixedVersion_ccl_drop_database_multiregion_primary_region(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_database_multiregion_primary_region"
	sctest.BackupRollbacksMixedVersion(t, path, MultiRegionTestClusterFactory{})
}

func TestBackupRollbacksMixedVersion_ccl_drop_table_multiregion(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_table_multiregion"
	sctest.BackupRollbacksMixedVersion(t, path, MultiRegionTestClusterFactory{})
}

func TestBackupRollbacksMixedVersion_ccl_drop_table_multiregion_primary_region(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_table_multiregion_primary_region"
	sctest.BackupRollbacksMixedVersion(t, path, MultiRegionTestClusterFactory{})
}

func TestBackupSuccess_ccl_alter_index_configure_zone(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/alter_index_configure_zone"
	sctest.BackupSuccess(t, path, MultiRegionTestClusterFactory{})
}

func TestBackupSuccess_ccl_alter_index_configure_zone_multiple(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/alter_index_configure_zone_multiple"
	sctest.BackupSuccess(t, path, MultiRegionTestClusterFactory{})
}

func TestBackupSuccess_ccl_create_index(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/create_index"
	sctest.BackupSuccess(t, path, MultiRegionTestClusterFactory{})
}

func TestBackupSuccess_ccl_drop_database_multiregion_primary_region(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_database_multiregion_primary_region"
	sctest.BackupSuccess(t, path, MultiRegionTestClusterFactory{})
}

func TestBackupSuccess_ccl_drop_table_multiregion(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_table_multiregion"
	sctest.BackupSuccess(t, path, MultiRegionTestClusterFactory{})
}

func TestBackupSuccess_ccl_drop_table_multiregion_primary_region(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_table_multiregion_primary_region"
	sctest.BackupSuccess(t, path, MultiRegionTestClusterFactory{})
}

func TestBackupSuccessMixedVersion_ccl_alter_index_configure_zone(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/alter_index_configure_zone"
	sctest.BackupSuccessMixedVersion(t, path, MultiRegionTestClusterFactory{})
}

func TestBackupSuccessMixedVersion_ccl_alter_index_configure_zone_multiple(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/alter_index_configure_zone_multiple"
	sctest.BackupSuccessMixedVersion(t, path, MultiRegionTestClusterFactory{})
}

func TestBackupSuccessMixedVersion_ccl_create_index(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/create_index"
	sctest.BackupSuccessMixedVersion(t, path, MultiRegionTestClusterFactory{})
}

func TestBackupSuccessMixedVersion_ccl_drop_database_multiregion_primary_region(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_database_multiregion_primary_region"
	sctest.BackupSuccessMixedVersion(t, path, MultiRegionTestClusterFactory{})
}

func TestBackupSuccessMixedVersion_ccl_drop_table_multiregion(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_table_multiregion"
	sctest.BackupSuccessMixedVersion(t, path, MultiRegionTestClusterFactory{})
}

func TestBackupSuccessMixedVersion_ccl_drop_table_multiregion_primary_region(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_table_multiregion_primary_region"
	sctest.BackupSuccessMixedVersion(t, path, MultiRegionTestClusterFactory{})
}

func TestEndToEndSideEffects_ccl_alter_index_configure_zone(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/alter_index_configure_zone"
	sctest.EndToEndSideEffects(t, path, MultiRegionTestClusterFactory{})
}

func TestEndToEndSideEffects_ccl_alter_index_configure_zone_multiple(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/alter_index_configure_zone_multiple"
	sctest.EndToEndSideEffects(t, path, MultiRegionTestClusterFactory{})
}

func TestEndToEndSideEffects_ccl_create_index(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/create_index"
	sctest.EndToEndSideEffects(t, path, MultiRegionTestClusterFactory{})
}

func TestEndToEndSideEffects_ccl_drop_database_multiregion_primary_region(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_database_multiregion_primary_region"
	sctest.EndToEndSideEffects(t, path, MultiRegionTestClusterFactory{})
}

func TestEndToEndSideEffects_ccl_drop_table_multiregion(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_table_multiregion"
	sctest.EndToEndSideEffects(t, path, MultiRegionTestClusterFactory{})
}

func TestEndToEndSideEffects_ccl_drop_table_multiregion_primary_region(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_table_multiregion_primary_region"
	sctest.EndToEndSideEffects(t, path, MultiRegionTestClusterFactory{})
}

func TestExecuteWithDMLInjection_ccl_alter_index_configure_zone(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/alter_index_configure_zone"
	sctest.ExecuteWithDMLInjection(t, path, MultiRegionTestClusterFactory{})
}

func TestExecuteWithDMLInjection_ccl_alter_index_configure_zone_multiple(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/alter_index_configure_zone_multiple"
	sctest.ExecuteWithDMLInjection(t, path, MultiRegionTestClusterFactory{})
}

func TestExecuteWithDMLInjection_ccl_create_index(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/create_index"
	sctest.ExecuteWithDMLInjection(t, path, MultiRegionTestClusterFactory{})
}

func TestExecuteWithDMLInjection_ccl_drop_database_multiregion_primary_region(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_database_multiregion_primary_region"
	sctest.ExecuteWithDMLInjection(t, path, MultiRegionTestClusterFactory{})
}

func TestExecuteWithDMLInjection_ccl_drop_table_multiregion(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_table_multiregion"
	sctest.ExecuteWithDMLInjection(t, path, MultiRegionTestClusterFactory{})
}

func TestExecuteWithDMLInjection_ccl_drop_table_multiregion_primary_region(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_table_multiregion_primary_region"
	sctest.ExecuteWithDMLInjection(t, path, MultiRegionTestClusterFactory{})
}

func TestGenerateSchemaChangeCorpus_ccl_alter_index_configure_zone(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/alter_index_configure_zone"
	sctest.GenerateSchemaChangeCorpus(t, path, MultiRegionTestClusterFactory{})
}

func TestGenerateSchemaChangeCorpus_ccl_alter_index_configure_zone_multiple(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/alter_index_configure_zone_multiple"
	sctest.GenerateSchemaChangeCorpus(t, path, MultiRegionTestClusterFactory{})
}

func TestGenerateSchemaChangeCorpus_ccl_create_index(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/create_index"
	sctest.GenerateSchemaChangeCorpus(t, path, MultiRegionTestClusterFactory{})
}

func TestGenerateSchemaChangeCorpus_ccl_drop_database_multiregion_primary_region(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_database_multiregion_primary_region"
	sctest.GenerateSchemaChangeCorpus(t, path, MultiRegionTestClusterFactory{})
}

func TestGenerateSchemaChangeCorpus_ccl_drop_table_multiregion(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_table_multiregion"
	sctest.GenerateSchemaChangeCorpus(t, path, MultiRegionTestClusterFactory{})
}

func TestGenerateSchemaChangeCorpus_ccl_drop_table_multiregion_primary_region(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_table_multiregion_primary_region"
	sctest.GenerateSchemaChangeCorpus(t, path, MultiRegionTestClusterFactory{})
}

func TestPause_ccl_alter_index_configure_zone(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/alter_index_configure_zone"
	sctest.Pause(t, path, MultiRegionTestClusterFactory{})
}

func TestPause_ccl_alter_index_configure_zone_multiple(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/alter_index_configure_zone_multiple"
	sctest.Pause(t, path, MultiRegionTestClusterFactory{})
}

func TestPause_ccl_create_index(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/create_index"
	sctest.Pause(t, path, MultiRegionTestClusterFactory{})
}

func TestPause_ccl_drop_database_multiregion_primary_region(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_database_multiregion_primary_region"
	sctest.Pause(t, path, MultiRegionTestClusterFactory{})
}

func TestPause_ccl_drop_table_multiregion(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_table_multiregion"
	sctest.Pause(t, path, MultiRegionTestClusterFactory{})
}

func TestPause_ccl_drop_table_multiregion_primary_region(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_table_multiregion_primary_region"
	sctest.Pause(t, path, MultiRegionTestClusterFactory{})
}

func TestPauseMixedVersion_ccl_alter_index_configure_zone(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/alter_index_configure_zone"
	sctest.PauseMixedVersion(t, path, MultiRegionTestClusterFactory{})
}

func TestPauseMixedVersion_ccl_alter_index_configure_zone_multiple(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/alter_index_configure_zone_multiple"
	sctest.PauseMixedVersion(t, path, MultiRegionTestClusterFactory{})
}

func TestPauseMixedVersion_ccl_create_index(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/create_index"
	sctest.PauseMixedVersion(t, path, MultiRegionTestClusterFactory{})
}

func TestPauseMixedVersion_ccl_drop_database_multiregion_primary_region(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_database_multiregion_primary_region"
	sctest.PauseMixedVersion(t, path, MultiRegionTestClusterFactory{})
}

func TestPauseMixedVersion_ccl_drop_table_multiregion(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_table_multiregion"
	sctest.PauseMixedVersion(t, path, MultiRegionTestClusterFactory{})
}

func TestPauseMixedVersion_ccl_drop_table_multiregion_primary_region(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_table_multiregion_primary_region"
	sctest.PauseMixedVersion(t, path, MultiRegionTestClusterFactory{})
}

func TestRollback_ccl_alter_index_configure_zone(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/alter_index_configure_zone"
	sctest.Rollback(t, path, MultiRegionTestClusterFactory{})
}

func TestRollback_ccl_alter_index_configure_zone_multiple(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/alter_index_configure_zone_multiple"
	sctest.Rollback(t, path, MultiRegionTestClusterFactory{})
}

func TestRollback_ccl_create_index(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/create_index"
	sctest.Rollback(t, path, MultiRegionTestClusterFactory{})
}

func TestRollback_ccl_drop_database_multiregion_primary_region(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_database_multiregion_primary_region"
	sctest.Rollback(t, path, MultiRegionTestClusterFactory{})
}

func TestRollback_ccl_drop_table_multiregion(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_table_multiregion"
	sctest.Rollback(t, path, MultiRegionTestClusterFactory{})
}

func TestRollback_ccl_drop_table_multiregion_primary_region(t *testing.T) {
	defer leaktest.AfterTest(t)()
	defer log.Scope(t).Close(t)
	const path = "pkg/ccl/schemachangerccl/testdata/end_to_end/drop_table_multiregion_primary_region"
	sctest.Rollback(t, path, MultiRegionTestClusterFactory{})
}
