import { Button } from 'kwc';
import { D } from '../KGVisualization/KGVisualization';
import IconClose from '@material-ui/icons/Close';
import IconStar from '@material-ui/icons/Star';
import IconUnstar from '@material-ui/icons/StarBorder';
import React from 'react';
import Score from '../KGVisualization/Score';
import cx from 'classnames';
import styles from './ResourceDetails.module.scss';

type Props = {
  resource: D | null;
  onClose: () => void;
  idToFullResource: { [key: string]: any };
};
function ResourceDetails({
  resource: tempResource,
  onClose,
  idToFullResource,
}: Props) {
  const resource: any = idToFullResource[tempResource?.id || ''] || null;
  const starred = tempResource?.starred;

  return (
    <div className={styles.container}>
      <div className={styles.title}>
        <div className={styles.titleText}>Resource Details</div>
        <div className={styles.actions}>
          <div className={styles.starredText}>{starred ? 'Starred' : ''}</div>
          <Button
            Icon={starred ? IconStar : IconUnstar}
            label=""
            className={cx({ [styles.starred]: starred })}
          />
          <Button Icon={IconClose} label="" onClick={onClose} />
        </div>
      </div>
      {resource !== null && (
        <>
          <div
            className={cx(styles.resourceTitleAndTopics, {
              [styles.starred]: starred,
            })}
          >
            <div className={styles.nameAndTopics}>
              <div className={styles.name}>{resource.title}</div>
            </div>
            <div className={styles.score}>
              <Score value={resource.score} />
            </div>
          </div>
          <div className={styles.content}>
            <div className={styles.authors}>
              <div className={styles.sectionTitle}>AUTHORS</div>
              <div className={styles.authorsText}>
                Cavigelli, Lukas; Bernath, Dominic; Magno, Michele; Benini, Luca
              </div>
            </div>
            <div className={styles.type}>{'Paper'}</div>
            <div className={styles.url}>
              <a href="http://www.arxiv-sanity.com/1611.03130" target="_blank">
                http://www.arxiv-sanity.com/1611.03130
              </a>
            </div>
            <div className={styles.topicsG}>
              {resource.topics.length > 0 && (
                <div className={styles.sectionTitle}>TOPICS</div>
              )}
              {resource.topics.map(({ name, relevance }: any) => (
                <div className={styles.topicScore}>
                  <Score value={relevance} inline />
                  <span>{name}</span>
                </div>
              ))}
            </div>
            <div className={styles.abstract}>{resource.abstract}</div>
          </div>
        </>
      )}
    </div>
  );
}

export default ResourceDetails;
