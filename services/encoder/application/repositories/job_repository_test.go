package repositories_test

import (
	"encoder/application/repositories"
	"encoder/domain"
	"encoder/framework/database"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestJobRepositoryInsert(t *testing.T) {

	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.New().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	job, err := domain.NewJob("output_path", "Pending", video)
	require.Nil(t, err)

	repoJob := repositories.NewJobRepository(db)
	repoJob.Insert(job)

	j, err := repoJob.Find(job.ID)

	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, j.ID, job.ID)
	require.Equal(t, j.VideoID, video.ID)

}

func TestJobRepositoryUpdate(t *testing.T) {

	db := database.NewDbTest()
	defer db.Close()

	video := domain.NewVideo()
	video.ID = uuid.New().String()
	video.FilePath = "path"
	video.CreatedAt = time.Now()

	repo := repositories.VideoRepositoryDb{Db: db}
	repo.Insert(video)

	job, err := domain.NewJob("output_path", "Pending", video)
	require.Nil(t, err)

	repoJob := repositories.NewJobRepository(db)
	repoJob.Insert(job)

	job.Status = "Complete"

	repoJob.Update(job)

	j, err := repoJob.Find(job.ID)

	require.NotEmpty(t, j.ID)
	require.Nil(t, err)
	require.Equal(t, j.Status, job.Status)

}
