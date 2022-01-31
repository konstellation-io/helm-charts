import { gql } from '@apollo/client';

export default gql`
  query GetRunningRuntime {
    runningRuntime {
      id
      name
      desc
      labels
      dockerImage
      usertoolsPod
    }
  }
`;
