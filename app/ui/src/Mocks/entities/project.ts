import { member1, member2 } from './member';

const toolUrls = {
  drone: 'drone',
  gitea: 'drone',
  jupyter: 'gitea',
  minio: 'jupyter',
  mlflow: 'minio',
  vscode: 'mlflow',
};

export const project1 = {
  id: 'projectId1',
  name: 'projectName1',
  description: 'projectDescription1',
  favorite: false,
  creationDate: '2020-02-02',
  lastActivationDate: '2020-02-02',
  repository: null,
  needAccess: false,
  archived: false,
  error: null,
  toolUrls,
  members: [member1, member2],
};

export const project2 = {
  id: 'projectId2',
  name: 'projectName2',
  description: 'projectDescription2',
  favorite: true,
  creationDate: '2020-03-03',
  lastActivationDate: '2020-03-03',
  repository: null,
  needAccess: false,
  archived: false,
  error: null,
  toolUrls,
  members: [member1, member2],
};

export const projectNoAccess = {
  id: 'projectId3',
  name: 'projectName3',
  description: 'projectDescription3',
  favorite: false,
  creationDate: '2020-02-02',
  lastActivationDate: '2020-02-02',
  repository: null,
  needAccess: true,
  archived: false,
  error: null,
  toolUrls,
  members: [member1, member2],
};

export const projectArchived = {
  id: 'projectId4',
  name: 'projectName4',
  description: 'projectDescription4',
  favorite: true,
  creationDate: '2020-03-03',
  lastActivationDate: '2020-03-03',
  repository: null,
  needAccess: false,
  archived: true,
  error: null,
  toolUrls,
  members: [member1, member2],
};
