package project_test

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/konstellation-io/kdl-server/app/api/application/project"
	"github.com/konstellation-io/kdl-server/app/api/entity"
	"github.com/konstellation-io/kdl-server/app/api/infrastructure/logging"
	"github.com/stretchr/testify/require"
	"testing"
)

// TODO move to a helper package
func AddLoggerExpects(logger *logging.MockLogger) {
	logger.EXPECT().Debug(gomock.Any()).Return().AnyTimes()
	logger.EXPECT().Info(gomock.Any()).Return().AnyTimes()
	logger.EXPECT().Warn(gomock.Any()).Return().AnyTimes()
	logger.EXPECT().Error(gomock.Any()).Return().AnyTimes()
	logger.EXPECT().Debugf(gomock.Any(), gomock.Any()).Return().AnyTimes()
	logger.EXPECT().Infof(gomock.Any(), gomock.Any()).Return().AnyTimes()
	logger.EXPECT().Warnf(gomock.Any(), gomock.Any()).Return().AnyTimes()
	logger.EXPECT().Errorf(gomock.Any(), gomock.Any()).Return().AnyTimes()
}

type projectSuite struct {
	ctrl       *gomock.Controller
	interactor project.UseCase
	mocks      projectMocks
}

type projectMocks struct {
	logger *logging.MockLogger
	repo   *project.MockRepository
}

func newProjectSuite(t *testing.T) *projectSuite {
	ctrl := gomock.NewController(t)

	logger := logging.NewMockLogger(ctrl)
	AddLoggerExpects(logger)

	repo := project.NewMockRepository(ctrl)

	interactor := project.NewInteractor(logger, repo)

	return &projectSuite{
		ctrl:       ctrl,
		interactor: interactor,
		mocks: projectMocks{
			logger: logger,
			repo:   repo,
		},
	}
}

func TestInteractor_Create(t *testing.T) {
	s := newProjectSuite(t)
	defer s.ctrl.Finish()

	const projectID = "project.1234"
	ctx := context.Background()
	p := entity.NewProject("project-x", "description")
	expectedProject := entity.Project{
		ID:          projectID,
		Name:        "project-x",
		Description: "description",
	}

	s.mocks.repo.EXPECT().Create(ctx, p).Return(projectID, nil)
	s.mocks.repo.EXPECT().Get(ctx, projectID).Return(expectedProject, nil)

	createdProject, err := s.interactor.Create(ctx, p)

	require.Nil(t, err)
	require.Equal(t, expectedProject, createdProject)
}
