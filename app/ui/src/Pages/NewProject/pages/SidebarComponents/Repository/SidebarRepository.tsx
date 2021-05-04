import React from 'react';
import styles from '../SidebarComponents.module.scss';

const SidebarRepository = () => (
  <div className={styles.sidebar}>
    <span>In this page you can choose the type of repository you want.</span>
    <span className={styles.line}>We provide two types of repos:</span>
    <ol className={styles.repoList}>
      <li>
        <strong>External repo:</strong> use a version-control system located
        outside the server.
      </li>
      <li>
        <strong>Internal repo:</strong> use a version-control system deployed
        inside the current server.
      </li>
    </ol>
  </div>
);

export default SidebarRepository;
