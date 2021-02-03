// FIXME: change the urls. Delete the server/
enum ROUTE {
  HOME = '/',
  USERS = '/users',
  USER_SSH_KEY = '/user/ssh-key',
  USER_API_TOKENS = '/user/api-tokens',
  GENERATE_USER_API_TOKEN = '/user/api-tokens/new-token',
  NEW_USER = '/new-user',
  NEW_PROJECT = '/new-project',
  PROJECT = '/project/:projectId',
  PROJECT_OVERVIEW = '/project/:projectId/overview',
  PROJECT_TOOLS = '/project/:projectId/tools',
  PROJECT_KG = '/project/:projectId/knowledge-graph',
  CREATION_PROJECT = '/new-project/create',
}

export type RouteProjectParams = {
  projectId: string;
};

export const buildRoute = {
  project: (route: ROUTE, projectId: string) =>
    route.replace(':projectId', projectId),
};

export default ROUTE;
