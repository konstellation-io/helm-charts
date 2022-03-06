import AnimateHeight from 'react-animate-height';
import CircularProgress from '@material-ui/core/CircularProgress';
import { NavButtonLink } from '../../ProjectNavigation';
import NavigationButton from '../NavigationButton/NavigationButton';
import IconPause from '@material-ui/icons/Pause';
import IconPlay from '@material-ui/icons/PlayArrow';
import IconSettings from '@material-ui/icons/Settings';
import * as React from 'react';
import { RouteProjectParams } from 'Constants/routes';
import cx from 'classnames';
import styles from './NavElements.module.scss';
import { useParams } from 'react-router-dom';
import useProjectNavigation from 'Hooks/useProjectNavigation';
import { useReactiveVar } from '@apollo/client';
import { loadingRuntime, lastRanRuntime, primaryPanel, runningRuntime } from 'Graphql/client/cache';
import usePanel, { PanelType } from 'Graphql/client/hooks/usePanel';
import { USERTOOLS_PANEL_OPTIONS } from 'Pages/Project/panelSettings';
import { PANEL_ID } from 'Graphql/client/models/Panel';
import useRuntime from 'Graphql/client/hooks/useRuntime';
import ReactTooltip from 'react-tooltip';
import TooltipElement from './TooltipElement';

type Props = {
  isOpened: boolean;
};

function NavElements({ isOpened }: Props) {
  const { projectId } = useParams<RouteProjectParams>();
  const { mainRoutes, userToolsRoutes, projectToolsRoutes } = useProjectNavigation(projectId);
  const runtimeLoading = useReactiveVar(loadingRuntime);
  const isLoading = runtimeLoading !== '';

  const runtimeRunning = useReactiveVar(runningRuntime);
  const panelData = useReactiveVar(primaryPanel);
  const runtimeLastRun = useReactiveVar(lastRanRuntime);
  const { startRuntime, pauseRuntime } = useRuntime();

  const { openPanel: openRuntimesList, closePanel: closeRuntimesList } = usePanel(
    PanelType.PRIMARY,
    USERTOOLS_PANEL_OPTIONS,
  );

  function toggleUsertoolsPanel() {
    const shouldOpen = !panelData || panelData.id !== PANEL_ID.RUNTIMES_LIST;

    if (shouldOpen) openRuntimesList();
    else closeRuntimesList();
  }

  function runtimeStart() {
    startRuntime(runtimeLastRun);
  }

  function runtimeStop() {
    pauseRuntime();
  }

  function renderToggleToolsIcon() {
    const Progress = (
      <div className={styles.progressSpinnerContainer}>
        <CircularProgress color="inherit" className={styles.loadingTools} size={12} />{' '}
      </div>
    );
    if (isLoading) return Progress;

    return runtimeRunning ? (
      <TooltipElement cssId="stop" spanText="Stop tools">
        <IconPause className={cx(styles.usertoolsIcon, 'icon-small')} onClick={runtimeStop} data-testid="stopTools" />
      </TooltipElement>
    ) : (
      <TooltipElement cssId="start" spanText="Start tools">
        <IconPlay className={cx(styles.usertoolsIcon, 'icon-small')} onClick={runtimeStart} data-testid="startTools" />
      </TooltipElement>
    );
  }

  return (
    <>
      {mainRoutes.map(({ Icon, label, route }) => (
        <NavButtonLink to={route} key={label}>
          <NavigationButton label={label} Icon={Icon} isNavCollapsed={!isOpened} />
        </NavButtonLink>
      ))}
      <div className={styles.projectTools}>
        {projectToolsRoutes.map(({ Icon, label, route, disabled }) => (
          <NavButtonLink to={route} key={label} disabled={disabled}>
            <NavigationButton label={label} Icon={Icon} isNavCollapsed={!isOpened} />
          </NavButtonLink>
        ))}
        <div
          className={cx(styles.userTools, {
            [styles.started]: runtimeRunning,
            [styles.stopped]: !runtimeRunning,
            [styles.loading]: isLoading,
          })}
        >
          <div className={cx(styles.usertoolsOptions, { [styles.opened]: isOpened })}>
            <AnimateHeight height={isOpened ? 'auto' : 0} duration={300}>
              <div className={styles.userToolLabel}>USER TOOLS</div>
            </AnimateHeight>
            <div
              className={cx(styles.usertoolsSettings, { [styles.opened]: isOpened })}
              data-testid="usertoolsSettings"
            >
              <ReactTooltip
                id="settings"
                effect="solid"
                textColor="white"
                backgroundColor="#888"
                className={styles.toolsTip}
              >
                <span>Show available runtimes</span>
              </ReactTooltip>
              <div data-tip data-for="settings" data-testid="openRuntimeSettings">
                <IconSettings className={cx(styles.usertoolsIcon, 'icon-small')} onClick={toggleUsertoolsPanel} />
              </div>
              {renderToggleToolsIcon()}
            </div>
          </div>
          {userToolsRoutes.map(({ Icon, label, route, disabled }) => (
            <NavButtonLink to={route} key={label} disabled={disabled}>
              <NavigationButton label={label} Icon={Icon} isNavCollapsed={!isOpened} />
            </NavButtonLink>
          ))}
          <div
            className={cx(styles.toggleToolsWrapper, {
              [styles.disabled]: isLoading,
            })}
            data-testid="confirmationModal"
          />
        </div>
      </div>
    </>
  );
}

export default NavElements;
