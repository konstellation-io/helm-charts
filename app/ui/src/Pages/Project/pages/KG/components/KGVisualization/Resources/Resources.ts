import { BaseType, Selection } from 'd3-selection';
import { Quadtree, quadtree } from 'd3-quadtree';

import { D } from '../KGVisualization';
import { DComplete } from './../../../KGUtils';
import styles from './Resources.module.scss';

export const RESOURCE_R = 4;
const RESOURCE_STROKE = 1;
const FONT_SIZE = 20;

const x = (d: DComplete) => d.x;
const y = (d: DComplete) => d.y;

export let lastSection: string | undefined;

const favoriteResources = [1, 5, 12, 13, 43, 76, 128, 654, 765, 734, 812];

const COLORS = {
  DEFAULT: 'rgba(12, 52, 72, 1)',
  DEFAULT_HIGHLIGHT: '#33FFFF',
  STARRED: '#CC7B55',
  STARRED_HIGHLIGHT: '#fc915f',
};

export default class Resources {
  container: Selection<SVGGElement, unknown, null, undefined>;
  data: DComplete[] = [];
  resourceR: number = RESOURCE_R;
  resourceStroke: number = RESOURCE_STROKE;
  fontSize: number = FONT_SIZE;
  context: CanvasRenderingContext2D | null;
  canvas: Selection<HTMLCanvasElement, unknown, null, undefined>;
  clearCanvas: () => void;
  onShowTooltip: (e: MouseEvent, d: DComplete) => void;
  onHideTooltip: (element: BaseType) => void;
  onResourceSelection: (d: D) => void;
  center: { x: number; y: number };
  qt: Quadtree<DComplete>;
  moving: boolean;
  onHoverResource: (d: DComplete | null) => void;

  constructor(
    onShowTooltip: (e: MouseEvent, d: DComplete) => void,
    onHideTooltip: (element: BaseType) => void,
    container: Selection<SVGGElement, unknown, null, undefined>,
    onResourceSelection: (d: D) => void,
    context: CanvasRenderingContext2D | null,
    clearCanvas: () => void,
    center: { x: number; y: number },
    canvas: Selection<HTMLCanvasElement, unknown, null, undefined>,
    onHoverResource: (d: DComplete | null) => void
  ) {
    this.onShowTooltip = onShowTooltip;
    this.onHideTooltip = onHideTooltip;
    this.onResourceSelection = onResourceSelection;
    this.container = container.select('g');
    this.context = context;
    this.canvas = canvas;
    this.clearCanvas = clearCanvas;
    this.center = center;
    this.qt = quadtree<DComplete>().x(x).y(y);
    this.moving = false;
    this.onHoverResource = onHoverResource;
  }

  drawCircles = (hover?: string | undefined) => {
    const {
      clearCanvas,
      context,
      data,
      center: { x, y },
    } = this;

    if (context === null) return;

    context.globalCompositeOperation = 'screen';
    // context.fillStyle = 'rgba(12, 52, 72, 0.8)';
    // context.globalCompositeOperation = 'lighter';

    let lastElement: DComplete | undefined;

    clearCanvas();
    data.forEach((d, i) => {
      let r = d.outsideMax ? RESOURCE_R * 0.7 : RESOURCE_R;
      let fillStyle = COLORS.DEFAULT;

      if (hover && d.name === hover) {
        lastElement = d;
      }

      if (favoriteResources.includes(i)) {
        fillStyle = COLORS.STARRED;
      }

      context.fillStyle = fillStyle;

      context.beginPath();
      context.moveTo(x + d.x, y + d.y);
      context.arc(x + d.x, y + d.y, r, 0, 2 * Math.PI);
      context.fill();
      context.closePath();
    });

    if (lastElement) {
      let fillStyle = COLORS.DEFAULT_HIGHLIGHT;

      if (favoriteResources.includes(data.indexOf(lastElement))) {
        fillStyle = COLORS.STARRED_HIGHLIGHT;
      }

      context.globalCompositeOperation = 'source-over';
      context.shadowBlur = 10;
      context.shadowColor = fillStyle;

      context.beginPath();
      context.moveTo(x + lastElement.x, y + lastElement.y);
      context.arc(
        x + lastElement.x,
        y + lastElement.y,
        RESOURCE_R * 1.7,
        0,
        2 * Math.PI
      );
      context.lineWidth = 0.5;
      context.strokeStyle = fillStyle;
      context.fillStyle = fillStyle;
      context.stroke();
      context.closePath();

      context.beginPath();
      context.moveTo(x + lastElement.x, y + lastElement.y);
      context.arc(
        x + lastElement.x,
        y + lastElement.y,
        RESOURCE_R,
        0,
        2 * Math.PI
      );
      context.fill();
      context.closePath();

      context.shadowBlur = 0;
    }
  };

  init = (
    container: Selection<SVGGElement, unknown, null, undefined>,
    data: DComplete[]
  ) => {
    this.data = data;
    this.container = container;

    const { drawCircles } = this;

    container.append('g').classed(styles.resourcesWrapper, true);

    this.qt.addAll(data);
    drawCircles();

    this.canvas.on('mousemove', this.onMouseMove);
    this.canvas.on('mousedown', this.onMouseDown);
    this.canvas.on('mouseup', this.onMouseUp);
    // this.canvas.on('click', this.onMouseClick);
  };

  onMouseMove = (e: any) => {
    const hovered = this.qt.find(
      e.offsetX - this.center.x,
      e.offsetY - this.center.y,
      50
    );

    this.moving = true;

    lastSection = hovered?.category;

    if (hovered) {
      this.onHoverResource(hovered);
      //   this.onShowTooltip(e, hovered);
    } else {
      this.onHoverResource(null);
      //   this.onHideTooltip(e);
    }

    this.drawCircles(hovered?.name);
  };

  onMouseClick = (e: any) => {
    const resource = this.qt.find(
      e.offsetX - this.center.x,
      e.offsetY - this.center.y,
      50
    );

    if (resource) {
      this.onResourceSelection(resource);
    }
  };

  onMouseDown = (e: any) => {
    this.moving = false;
  };

  onMouseUp = (e: any) => {
    console.log('moving', this.moving);

    if (!this.moving) {
      const resource = this.qt.find(
        e.offsetX - this.center.x,
        e.offsetY - this.center.y,
        50
      );

      if (resource) {
        this.onResourceSelection(resource);
      }
    }
  };

  performUpdate = (data: DComplete[]) => {
    this.qt.removeAll(this.data);
    this.data = data;
    this.qt.addAll(data);

    this.drawCircles();
  };

  highlightResource = (resourceName: string | null) => {
    this.drawCircles(resourceName || '');
  };
}
