// Copyright Project Harbor Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package dao

import (
	"context"
	"testing"
	"time"

	"github.com/goharbor/harbor/src/common/dao"
	"github.com/goharbor/harbor/src/common/utils"
	"github.com/goharbor/harbor/src/jobservice/job"
	"github.com/goharbor/harbor/src/lib/errors"
	"github.com/goharbor/harbor/src/lib/orm"
	"github.com/goharbor/harbor/src/lib/q"
	"github.com/stretchr/testify/suite"
)

type taskMysqlDAOTestSuite struct {
	suite.Suite
	ctx          context.Context
	taskDAO      *taskMysqlDAO
	executionDAO *executionMysqlDAO
	executionID  int64
	taskID       int64
}

func (t *taskMysqlDAOTestSuite) SetupSuite() {
	dao.PrepareTestForDB()
	t.ctx = orm.Context()
	t.taskDAO = &taskMysqlDAO{}
	t.executionDAO = &executionMysqlDAO{}
}

func (t *taskMysqlDAOTestSuite) SetupTest() {
	id, err := t.executionDAO.Create(t.ctx, &Execution{
		VendorType: "test",
		Trigger:    "test",
		ExtraAttrs: "{}",
	})
	t.Require().Nil(err)
	t.executionID = id
	id, err = t.taskDAO.Create(t.ctx, &Task{
		ExecutionID: t.executionID,
		Status:      "success",
		StatusCode:  1,
		ExtraAttrs:  `{"key":"value"}`,
	})
	t.Require().Nil(err)
	t.taskID = id
}

func (t *taskMysqlDAOTestSuite) TearDownTest() {
	err := t.taskDAO.Delete(t.ctx, t.taskID)
	t.Nil(err)

	err = t.executionDAO.Delete(t.ctx, t.executionID)
	t.Nil(err)
}

func (t *taskMysqlDAOTestSuite) TestCount() {
	count, err := t.taskDAO.Count(t.ctx, &q.Query{
		Keywords: map[string]interface{}{
			"ExecutionID":    t.executionID,
			"ExtraAttrs.key": "value",
		},
	})
	t.Require().Nil(err)
	t.Equal(int64(1), count)

	count, err = t.taskDAO.Count(t.ctx, &q.Query{
		Keywords: map[string]interface{}{
			"ExecutionID":    t.executionID,
			"ExtraAttrs.key": "incorrect-value",
		},
	})
	t.Require().Nil(err)
	t.Equal(int64(0), count)
}

func (t *taskMysqlDAOTestSuite) TestList() {
	tasks, err := t.taskDAO.List(t.ctx, &q.Query{
		Keywords: map[string]interface{}{
			"ExecutionID":    t.executionID,
			"ExtraAttrs.key": "value",
		},
	})
	t.Require().Nil(err)
	t.Require().Len(tasks, 1)
	t.Equal(t.taskID, tasks[0].ID)

	tasks, err = t.taskDAO.List(t.ctx, &q.Query{
		Keywords: map[string]interface{}{
			"ExecutionID":    t.executionID,
			"ExtraAttrs.key": "incorrect-value",
		},
	})
	t.Require().Nil(err)
	t.Require().Len(tasks, 0)
}

func (t *taskMysqlDAOTestSuite) TestGet() {
	// not exist
	_, err := t.taskDAO.Get(t.ctx, 10000)
	t.Require().NotNil(err)
	t.True(errors.IsNotFoundErr(err))

	// exist
	task, err := t.taskDAO.Get(t.ctx, t.taskID)
	t.Require().Nil(err)
	t.NotNil(task)
}

func (t *taskMysqlDAOTestSuite) TestCreate() {
	// reference the non-existing execution
	_, err := t.taskDAO.Create(t.ctx, &Task{
		ExecutionID: 10000,
		Status:      "success",
		StatusCode:  1,
		ExtraAttrs:  "{}",
	})
	t.Require().NotNil(err)
	t.True(errors.IsErr(err, errors.ViolateForeignKeyConstraintCode))

	// reference the existing execution is covered by SetupTest
}

func (t *taskMysqlDAOTestSuite) TestUpdate() {
	// not exist
	err := t.taskDAO.Update(t.ctx, &Task{ID: 10000}, "Status")
	t.Require().NotNil(err)
	t.True(errors.IsNotFoundErr(err))

	// exist
	err = t.taskDAO.Update(t.ctx, &Task{
		ID:     t.taskID,
		Status: "failed",
	}, "Status")
	t.Require().Nil(err)
	task, err := t.taskDAO.Get(t.ctx, t.taskID)
	t.Require().Nil(err)
	t.Equal("failed", task.Status)
}

func (t *taskMysqlDAOTestSuite) TestUpdateStatus() {
	// update status to running
	status := job.RunningStatus.String()
	statusRevision := time.Now().Unix()
	err := t.taskDAO.UpdateStatus(t.ctx, t.taskID, status, statusRevision)
	t.Require().Nil(err)

	task, err := t.taskDAO.Get(t.ctx, t.taskID)
	t.Require().Nil(err)
	t.Equal(int32(1), task.RunCount)
	t.True(time.Unix(statusRevision, 0).Equal(task.StartTime))
	t.Equal(status, task.Status)
	t.Equal(job.RunningStatus.Code(), task.StatusCode)
	t.Equal(statusRevision, task.StatusRevision)
	t.NotEqual(time.Time{}, task.UpdateTime)
	t.Equal(time.Time{}, task.EndTime)

	// update status to success
	status = job.SuccessStatus.String()
	err = t.taskDAO.UpdateStatus(t.ctx, t.taskID, status, statusRevision)
	t.Require().Nil(err)

	task, err = t.taskDAO.Get(t.ctx, t.taskID)
	t.Require().Nil(err)
	t.Equal(int32(1), task.RunCount)
	t.True(time.Unix(statusRevision, 0).Equal(task.StartTime))
	t.Equal(status, task.Status)
	t.Equal(job.SuccessStatus.Code(), task.StatusCode)
	t.Equal(statusRevision, task.StatusRevision)
	t.NotEqual(time.Time{}, task.EndTime)

	// update status to running again with different revision
	status = job.RunningStatus.String()
	statusRevision = time.Now().Add(1 * time.Second).Unix()
	err = t.taskDAO.UpdateStatus(t.ctx, t.taskID, status, statusRevision)
	t.Require().Nil(err)

	task, err = t.taskDAO.Get(t.ctx, t.taskID)
	t.Require().Nil(err)
	t.Equal(int32(2), task.RunCount)
	t.True(time.Unix(statusRevision, 0).Equal(task.StartTime))
	t.Equal(status, task.Status)
	t.Equal(job.RunningStatus.Code(), task.StatusCode)
	t.Equal(statusRevision, task.StatusRevision)
	t.Equal(time.Time{}, task.EndTime)
}

func (t *taskMysqlDAOTestSuite) TestDelete() {
	// not exist
	err := t.taskDAO.Delete(t.ctx, 10000)
	t.Require().NotNil(err)
	t.True(errors.IsNotFoundErr(err))

	// happy pass is covered by TearDownTest
}

func (t *taskMysqlDAOTestSuite) TestListStatusCount() {
	scs, err := t.taskDAO.ListStatusCount(t.ctx, t.executionID)
	t.Require().Nil(err)
	t.Require().Len(scs, 1)
	t.Equal("success", scs[0].Status)
	t.Equal(int64(1), scs[0].Count)
}

func (t *taskMysqlDAOTestSuite) TestGetMaxEndTime() {
	now := time.Now()
	err := t.taskDAO.Update(t.ctx, &Task{
		ID:      t.taskID,
		EndTime: now,
	}, "EndTime")
	t.Require().Nil(err)
	endTime, err := t.taskDAO.GetMaxEndTime(t.ctx, t.executionID)
	t.Require().Nil(err)
	t.Equal(now.Unix(), endTime.Unix())
}

func TestTaskMysqlDAOSuite(t *testing.T) {
	if !utils.IsDBMysql() {
		return
	}
	suite.Run(t, &taskMysqlDAOTestSuite{})
}
