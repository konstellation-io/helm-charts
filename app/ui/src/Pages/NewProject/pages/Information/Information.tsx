import { SpinnerCircular, TextInput } from 'kwc';
import { generateSlug, getErrorMsg } from 'Utils/string';
import {
  validateProjectDescription,
  validateProjectName,
} from './InformationUtils';

import DescriptionScore from 'Components/DescriptionScore/DescriptionScore';
import React from 'react';
import { newProject } from 'Graphql/client/cache';
import styles from './Information.module.scss';
import useNewProject from 'Graphql/client/hooks/useNewProject';
import useQualityDescription from 'Hooks/useQualityDescription/useQualityDescription';
import { useReactiveVar } from '@apollo/client';

const limits = {
  maxHeight: 400,
  minHeight: 375,
};

type Props = {
  showErrors: boolean;
};
function Information({ showErrors }: Props) {
  const project = useReactiveVar(newProject);
  const { updateValue, updateError, clearError } = useNewProject('information');
  const { updateValue: updateInternalRepositoryValue } = useNewProject(
    'internalRepository'
  );

  const { values, errors } = project.information || {};
  const { name, description } = values;
  const { name: errorName, description: errorDescription } = errors;

  const { descriptionScore, loading } = useQualityDescription(description);

  if (!project) return <SpinnerCircular />;

  return (
    <div className={styles.container}>
      <TextInput
        label="project name"
        onChange={(v: string) => {
          updateValue('name', v);
          clearError('name');
        }}
        onBlur={() => {
          updateInternalRepositoryValue('slug', generateSlug(name));
          const isValidName = validateProjectName(name);
          updateError('name', getErrorMsg(isValidName));
        }}
        formValue={name}
        autoFocus
        showClearButton
        error={showErrors ? errorName : ''}
      />
      <TextInput
        label="project description"
        formValue={description}
        onChange={(v: string) => {
          updateValue('description', v);
          clearError('description');
        }}
        onBlur={() => {
          const isValidDescription = validateProjectDescription(description);
          updateError('description', getErrorMsg(isValidDescription));
        }}
        limits={limits}
        error={showErrors ? errorDescription : ''}
        helpText={`A minimum of 50 words is required to get a valid score. Words: ${
          description.split(' ').length
        }`}
        showClearButton
        textArea
        lockHorizontalGrowth
      />
      <DescriptionScore score={descriptionScore} loading={loading} />
    </div>
  );
}

export default Information;
