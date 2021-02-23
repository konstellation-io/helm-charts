import { ErrorMessage, SpinnerCircular } from 'kwc';
import React, { useEffect } from 'react';

import { GetProjects } from 'Graphql/queries/types/GetProjects';
import ProjectContentRoutes from './ProjectContentRoutes';
import ProjectNavigation from './components/ProjectNavigation/ProjectNavigation';
import ProjectPanels from './ProjectPanels';
import { RouteProjectParams } from 'Constants/routes';
import { loader } from 'graphql.macro';
import styles from './Project.module.scss';
import useOpenedProject from 'Graphql/client/hooks/useOpenedProject';
import { useParams } from 'react-router-dom';
import { useQuery } from '@apollo/client';
import {
  GET_OPENED_PROJECT,
  GetOpenedProject,
} from '../../Graphql/client/queries/getOpenedProject.graphql';

const GetProjectsQuery = loader('Graphql/queries/getProjects.graphql');

function Project() {
  const { projectId } = useParams<RouteProjectParams>();
  const { data, error, loading } = useQuery<GetProjects>(GetProjectsQuery);
  const { data: openedProjectData } = useQuery<GetOpenedProject>(
    GET_OPENED_PROJECT
  );
  const { updateOpenedProject } = useOpenedProject();

  useEffect(() => {
    const openedProject = data?.projects.find((p) => p.id === projectId);
    openedProject && updateOpenedProject(openedProject);
  }, [data, projectId, updateOpenedProject]);

  useEffect(
    () => () => updateOpenedProject(null),
    // We want to execute this on on component mount/unmount
    // eslint-disable-next-line react-hooks/exhaustive-deps
    []
  );

  const project = openedProjectData?.openedProject;

  if (loading || !project) return <SpinnerCircular />;
  if (error || !data) return <ErrorMessage />;

  return (
    <div className={styles.container}>
      <ProjectNavigation />
      <div className={styles.contentLayer}>
        <ProjectContentRoutes openedProject={project} />
      </div>
      <div className={styles.panelLayer}>
        <ProjectPanels openedProject={project} />
      </div>
    </div>
  );
}

export default Project;
