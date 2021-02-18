package project_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"github.com/konstellation-io/kdl-server/app/api/entity"
	"github.com/konstellation-io/kdl-server/app/api/infrastructure/droneservice"
	"github.com/konstellation-io/kdl-server/app/api/infrastructure/giteaservice"
	"github.com/konstellation-io/kdl-server/app/api/infrastructure/minioservice"
	"github.com/konstellation-io/kdl-server/app/api/pkg/clock"
	"github.com/konstellation-io/kdl-server/app/api/pkg/logging"
	"github.com/konstellation-io/kdl-server/app/api/usecase/project"
)

type projectSuite struct {
	ctrl       *gomock.Controller
	interactor project.UseCase
	mocks      projectMocks
}

type projectMocks struct {
	logger       *logging.MockLogger
	repo         *project.MockRepository
	clock        *clock.MockClock
	giteaService *giteaservice.MockGiteaClient
	minioService *minioservice.MockMinioService
	droneService *droneservice.MockDroneService
}

func newProjectSuite(t *testing.T) *projectSuite {
	ctrl := gomock.NewController(t)

	logger := logging.NewMockLogger(ctrl)
	logging.AddLoggerExpects(logger)

	repo := project.NewMockRepository(ctrl)

	clockMock := clock.NewMockClock(ctrl)

	giteaService := giteaservice.NewMockGiteaClient(ctrl)

	minioService := minioservice.NewMockMinioService(ctrl)

	droneService := droneservice.NewMockDroneService(ctrl)

	interactor := project.NewInteractor(logger, repo, clockMock, giteaService, minioService, droneService)

	return &projectSuite{
		ctrl:       ctrl,
		interactor: interactor,
		mocks: projectMocks{
			logger:       logger,
			repo:         repo,
			clock:        clockMock,
			giteaService: giteaService,
			minioService: minioService,
			droneService: droneService,
		},
	}
}

func TestInteractor_Create(t *testing.T) {
	s := newProjectSuite(t)
	defer s.ctrl.Finish()

	const (
		projectID   = "project.1234"
		projectName = "The Project X"
		projectDesc = "The Project X Description"
	)

	var internalRepoName = "project-x"

	ctx := context.Background()
	now := time.Now().UTC()

	createProject := entity.NewProject(projectName, projectDesc)
	createProject.CreationDate = now
	createProject.Repository = entity.Repository{
		Type:             entity.RepositoryTypeInternal,
		InternalRepoName: internalRepoName,
	}

	expectedProject := entity.Project{
		ID:           projectID,
		Name:         projectName,
		Description:  projectDesc,
		CreationDate: now,
		Repository: entity.Repository{
			Type:             entity.RepositoryTypeInternal,
			InternalRepoName: internalRepoName,
		},
	}

	s.mocks.giteaService.EXPECT().CreateRepo(internalRepoName).Return(nil)
	s.mocks.minioService.EXPECT().CreateBucket(internalRepoName).Return(nil)
	s.mocks.droneService.EXPECT().ActivateRepository(internalRepoName).Return(nil)
	s.mocks.clock.EXPECT().Now().Return(now)
	s.mocks.repo.EXPECT().Create(ctx, createProject).Return(projectID, nil)
	s.mocks.repo.EXPECT().Get(ctx, projectID).Return(expectedProject, nil)

	createdProject, err := s.interactor.Create(ctx, project.CreateProjectOption{
		Name:             projectName,
		Description:      projectDesc,
		RepoType:         entity.RepositoryTypeInternal,
		InternalRepoName: &internalRepoName,
		ExternalRepoURL:  nil,
	})

	require.NoError(t, err)
	require.Equal(t, expectedProject, createdProject)
}

func TestInteractor_FindAll(t *testing.T) {
	s := newProjectSuite(t)
	defer s.ctrl.Finish()

	ctx := context.Background()
	expectedProjects := []entity.Project{
		entity.NewProject("project-x", "Project X"),
	}

	s.mocks.repo.EXPECT().FindAll(ctx).Return(expectedProjects, nil)

	p, err := s.interactor.FindAll(ctx)

	require.NoError(t, err)
	require.Equal(t, p, expectedProjects)
}

func TestInteractor_GetByID(t *testing.T) {
	s := newProjectSuite(t)
	defer s.ctrl.Finish()

	const projectID = "project.1234"

	ctx := context.Background()
	expectedProject := entity.NewProject("project-x", "Project X")

	s.mocks.repo.EXPECT().Get(ctx, projectID).Return(expectedProject, nil)

	p, err := s.interactor.GetByID(ctx, projectID)

	require.NoError(t, err)
	require.Equal(t, p, expectedProject)
}
