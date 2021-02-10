import { CHECK, SpinnerCircular, TextInput } from 'kwc';
import {
  GET_NEW_PROJECT,
  GetNewProject,
} from 'Graphql/client/queries/getNewProject.graphql';
import React, { useEffect } from 'react';

import { CONFIG } from 'index';
import IconLink from '@material-ui/icons/Link';
import styles from './InternalRepository.module.scss';
import useNewProject from 'Graphql/client/hooks/useNewProject';
import { useQuery } from '@apollo/client';

function validateProjectSlug(value: string): string {
  const error = CHECK.getValidationError([
    CHECK.isLowerCase(value),
    CHECK.matches(value, /^[a-z]/, 'Name must start with a lowercase letter'),
    CHECK.matches(value, /.{3,}/, 'Name must contain at least 3 characters'),
    CHECK.isAlphanumeric(
      value.replaceAll('-', ''),
      'Name only can contain lowercase alphanumeric and hyphens'
    ),
    CHECK.isSlug(value),
  ]);
  return error === true ? '' : (error as string);
}
type Props = {
  showErrors: boolean;
};
function InternalRepository({ showErrors }: Props) {
  const { data } = useQuery<GetNewProject>(GET_NEW_PROJECT);
  const { updateValue, updateError, clearError } = useNewProject(
    'internalRepository'
  );
  const { values } = data?.newProject.internalRepository || {};
  const { errors } = data?.newProject.internalRepository || {};

  const slug = values?.slug || '';
  const slugError = errors?.slug;

  useEffect(() => {
    updateValue('url', `${CONFIG.SERVER_URL}.${slug}`);
  }, [updateValue, slug]);

  if (!data) return <SpinnerCircular />;

  const slugOk = validateProjectSlug(slug);

  return (
    <div className={styles.repositoryInternal}>
      <div className={styles.url}>
        <p className={styles.urlTitle}>repository url</p>
        <div className={styles.serverUrlContainer}>
          <IconLink className="icon-regular" />
          <span className={styles.urlContent}>{`${CONFIG.SERVER_URL}/`}</span>
        </div>
      </div>
      <TextInput
        label="repository name"
        customClassname={styles.slug}
        onChange={(value: string) => {
          updateValue('slug', value);
          clearError('slug');
        }}
        onBlur={() => updateError('slug', slugOk)}
        formValue={slug}
        error={showErrors ? slugError : ''}
        helpText="Please write in a URL compatible way"
        showClearButton
      />
    </div>
  );
}

export default InternalRepository;