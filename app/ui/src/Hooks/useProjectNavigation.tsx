import ROUTE, { buildRoute } from 'Constants/routes';
import { OverridableComponent } from '@material-ui/core/OverridableComponent';
import { SvgIconTypeMap } from '@material-ui/core';
import { useMemo } from 'react';
import IconHome from '@material-ui/icons/Dashboard';
import IconKG from '@material-ui/icons/EmojiObjects';
import GiteaIcon from 'Components/Icons/GiteaIcon/GiteaIcon';
import MinioIcon from 'Components/Icons/MinioIcon/MinioIcon';
import DroneIcon from 'Components/Icons/DroneIcon/DroneIcon';
import VSIcon from 'Components/Icons/VSIcon/VSIcon';
import JupyterIcon from 'Components/Icons/JupyterIcon/JupyterIcon';
import MlFlowIcon from 'Components/Icons/MlFlowIcon/MlFlowIcon';
import { useQuery } from '@apollo/client';
import { GetMe } from 'Graphql/queries/types/GetMe';
import { loader } from 'graphql.macro';

const GetMeQuery = loader('Graphql/queries/getMe.graphql');

export interface RouteConfiguration {
  id: string;
  label: string;
  Icon: OverridableComponent<SvgIconTypeMap<{}, 'svg'>>;
  canBeDisabled?: boolean;
  disabled?: boolean;
}

export const projectRoutesConfiguration: {
  [key: string]: RouteConfiguration;
} = {
  [ROUTE.PROJECT_OVERVIEW]: {
    id: 'overview',
    label: 'Overview',
    Icon: IconHome,
  },
  [ROUTE.PROJECT_KG]: {
    id: 'knowledge-graph',
    label: 'Knowledge Graph',
    Icon: IconKG,
  },
  [ROUTE.PROJECT_TOOL_GITEA]: {
    id: 'gitea',
    label: 'Gitea',
    Icon: GiteaIcon,
  },
  [ROUTE.PROJECT_TOOL_DRONE]: {
    id: 'drone',
    label: 'Drone',
    Icon: DroneIcon,
  },
  [ROUTE.PROJECT_TOOL_JUPYTER]: {
    id: 'jupyter',
    label: 'Jupyter',
    Icon: JupyterIcon,
    canBeDisabled: true,
    disabled: true,
  },
  [ROUTE.PROJECT_TOOL_MINIO]: {
    id: 'minio',
    label: 'Minio',
    Icon: MinioIcon,
  },
  [ROUTE.PROJECT_TOOL_MLFLOW]: {
    id: 'mlflow',
    label: 'Mlflow',
    Icon: MlFlowIcon,
  },
  [ROUTE.PROJECT_TOOL_VSCODE]: {
    id: 'vscode',
    label: 'Vscode',
    Icon: VSIcon,
    canBeDisabled: true,
    disabled: true,
  },
};

export interface EnhancedRouteConfiguration extends RouteConfiguration {
  to: string;
}

function useProjectNavigation(projectId: string) {
  const { data } = useQuery<GetMe>(GetMeQuery);

  const routesConfigurations: EnhancedRouteConfiguration[] = useMemo(
    () =>
      Object.entries(projectRoutesConfiguration).map(
        ([routeString, { id, label, Icon, canBeDisabled }]) => ({
          to: buildRoute(routeString as ROUTE, projectId),
          id,
          label,
          Icon,
          disabled: canBeDisabled ? !data?.me.areToolsActive : false,
        })
      ),
    [projectId, data?.me.areToolsActive]
  );
  return routesConfigurations;
}

export default useProjectNavigation;
