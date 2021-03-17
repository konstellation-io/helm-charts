import { useCallback, useEffect, useRef, useState } from 'react';
import { useLazyQuery } from '@apollo/client';
import {
  GetQualityProjectDesc,
  GetQualityProjectDescVariables,
} from 'Graphql/queries/types/GetQualityProjectDesc';
import { loader } from 'graphql.macro';

const GetQualityProjectDescQuery = loader(
  'Graphql/queries/getQualityProjectDesc.graphql'
);

type Options = {
  skipFirstRun?: boolean;
  debounceTime?: number;
  minWordsNumber?: number;
};

function useQualityDescription(
  description: string,
  {
    skipFirstRun = true,
    debounceTime = 2000,
    minWordsNumber = 50,
  }: Options = {}
) {
  const [descriptionScore, setDescriptionScore] = useState(0);

  const [getQualityProjectDesc] = useLazyQuery<
    GetQualityProjectDesc,
    GetQualityProjectDescVariables
  >(GetQualityProjectDescQuery, {
    onCompleted: (data) => {
      setDescriptionScore(data.qualityProjectDesc.quality || 0);
    },
  });

  const fetchDescriptionScore = useCallback(() => {
    const isLengthAcceptable = description.split(' ').length >= minWordsNumber;

    if (!isLengthAcceptable) setDescriptionScore(0);
    else getQualityProjectDesc({ variables: { description } });
  }, [description, minWordsNumber, getQualityProjectDesc]);

  useEffect(() => {
    if (!skipFirstRun) fetchDescriptionScore();
    // We want to run this only on hook instantiation.
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  const firstRun = useRef(true);
  useEffect(() => {
    if (firstRun.current) {
      firstRun.current = false;
      return;
    }

    const scoreTimeoutId = setTimeout(fetchDescriptionScore, debounceTime);

    return () => clearTimeout(scoreTimeoutId);
  }, [description, debounceTime, fetchDescriptionScore]);

  return {
    descriptionScore,
    fetchDescriptionScore,
  };
}

export default useQualityDescription;
