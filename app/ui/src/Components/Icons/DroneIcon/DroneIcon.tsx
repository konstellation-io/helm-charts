import { SvgIcon, SvgIconProps } from '@material-ui/core';
import * as React from 'react';

const DroneIcon = (props: SvgIconProps) => (
  <SvgIcon {...props} viewBox="0 0 2500 2500">
    <path
      clipRule="evenodd"
      d="m503.6 242.3-10.7 10.7 438.1 438.1c-62 96.7-97.6 215.5-97.6 350.6 0 375.6 275.8 625 625 625 130.5 0 250.8-34.8 350.2-98l440.3 440.3c-228.1 301.3-590.7 491-998.8 491-690.3 0-1250.1-542.6-1250.1-1250 0-420.7 198-783.1 503.6-1007.7zm208.4-123.8c162.9-76.2 345.4-118.5 538-118.5 690.3 0 1250 542.6 1250 1250 0 194.6-42.3 376.7-118.1 538.5l-396.2-396.2c62-96.7 97.6-215.5 97.6-350.6 0-375.6-275.8-625-625-625-130.5 0-250.8 34.8-350.2 98zm746.3 1298.2c-209.6 0-375-149.6-375-375s165.4-375 375-375 375 149.6 375 375-165.4 375-375 375z"
      fillRule="evenodd"
    />
  </SvgIcon>
);
export default DroneIcon;
