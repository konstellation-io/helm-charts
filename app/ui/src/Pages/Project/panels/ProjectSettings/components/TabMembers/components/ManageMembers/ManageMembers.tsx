import {
  Button,
  BUTTON_ALIGN,
  ModalContainer,
  ModalLayoutConfirmList,
  Select,
} from 'kwc';
import React, { useRef } from 'react';
import styles from './ManageMembers.module.scss';
import useMembers from 'Graphql/hooks/useMembers';
import { AccessLevel } from 'Graphql/types/globalTypes';
import {
  defaultModalInfo,
  getModalInfo,
  ModalInfo,
} from '../../confirmationModals';
import { GetProjectMembers_project_members } from 'Graphql/queries/types/GetProjectMembers';
import MemberItem from './components/MemberItem/MemberItem';
import DeleteIcon from '@material-ui/icons/Delete';
import useBoolState from 'Hooks/useBoolState';
import capitalize from 'lodash.capitalize';

const accessLevelSeparator = () => (
  <Button
    key="separator"
    label="CHANGE MEMBERS LEVEL TO"
    className={styles.separator}
    align={BUTTON_ALIGN.LEFT}
    disabled
  />
);

type Props = {
  projectId: string;
  selectedMembers: GetProjectMembers_project_members[];
  onCompleteManage: () => void;
};
function ManageMembers({
  projectId,
  selectedMembers,
  onCompleteManage,
}: Props) {
  const {
    activate: showModal,
    deactivate: closeModal,
    value: isModalVisible,
  } = useBoolState();
  const modalInfo = useRef<ModalInfo>(defaultModalInfo);

  function onComplete() {
    closeModal();
    onCompleteManage();
  }

  const { updateMembersAccessLevel, removeMembersById } = useMembers(
    projectId,
    {
      onCompleteUpdate: onComplete,
      onCompleteRemove: onComplete,
    }
  );

  const nMembers = selectedMembers.length;
  const hasNotMembersSelected = nMembers === 0;
  const membersIds = () => selectedMembers.map(({ user }) => user.id);

  function showAccessLevelModal(accessLevel: AccessLevel) {
    modalInfo.current = getModalInfo({
      nMembers,
      type: 'update',
      action: () => updateMembersAccessLevel(membersIds(), accessLevel),
      accessLevel: accessLevel,
    });
    showModal();
  }

  function showRemoveModal() {
    modalInfo.current = getModalInfo({
      nMembers,
      type: 'remove',
      action: () => removeMembersById(membersIds()),
    });
    showModal();
  }

  const removeMembersButton = () => (
    <Button
      key="remove"
      label="Remove members"
      onClick={showRemoveModal}
      className={styles.manageMemberButton}
      Icon={DeleteIcon}
      align={BUTTON_ALIGN.LEFT}
    />
  );

  const AccessLevelButton = (accessLevel: AccessLevel) => (
    <Button
      key={accessLevel}
      label={capitalize(accessLevel)}
      onClick={() => showAccessLevelModal(accessLevel)}
      className={styles.manageMemberButton}
      align={BUTTON_ALIGN.LEFT}
    />
  );

  const optionToButton = {
    'remove members': removeMembersButton,
    'access levels': accessLevelSeparator,
    admin: () => AccessLevelButton(AccessLevel.ADMIN),
    manager: () => AccessLevelButton(AccessLevel.MANAGER),
    viewer: () => AccessLevelButton(AccessLevel.VIEWER),
  };

  const selectionText =
    nMembers === 1 ? `${nMembers} User selected` : `${nMembers} Users selected`;

  return (
    <div className={styles.container}>
      <span className={styles.actionsLabel}>ACTIONS</span>
      <Select
        placeholder={selectionText}
        className={styles.accessLevelSelector}
        options={Object.keys(optionToButton)}
        CustomOptions={optionToButton}
        disabledOptions={[Object.keys(optionToButton)[1]]}
        showSelectAllOption={false}
        shouldSort={false}
        height={30}
        disabled={hasNotMembersSelected}
        hideError
      />
      {isModalVisible && (
        <ModalContainer
          title={modalInfo.current.title}
          onAccept={modalInfo.current.action}
          onCancel={closeModal}
          actionButtonLabel={modalInfo.current.acceptLabel}
          actionButtonCancel="Cancel"
          warning={modalInfo.current.warning}
          error={modalInfo.current.error}
          blocking
        >
          <ModalLayoutConfirmList message={modalInfo.current.message}>
            {selectedMembers.map((m) => (
              <MemberItem key={m.user.id} member={m} />
            ))}
          </ModalLayoutConfirmList>
        </ModalContainer>
      )}
    </div>
  );
}

export default ManageMembers;