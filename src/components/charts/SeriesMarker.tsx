import type { SeriesMarkerShape } from './chartTheme';

type SeriesMarkerIconProps = {
  color: string;
  shape: SeriesMarkerShape;
};

type SeriesMarkerDotProps = SeriesMarkerIconProps & {
  cx?: number;
  cy?: number;
  size?: number;
};

function marker(shape: SeriesMarkerShape, color: string, size: number) {
  const center = size / 2;
  const radius = size * 0.28;

  if (shape === 'diamond') {
    return (
      <polygon
        fill={color}
        points={`${center},${center - radius} ${center + radius},${center} ${center},${center + radius} ${center - radius},${center}`}
      />
    );
  }

  if (shape === 'square') {
    return (
      <rect
        fill={color}
        height={radius * 1.75}
        width={radius * 1.75}
        x={center - radius * 0.875}
        y={center - radius * 0.875}
      />
    );
  }

  if (shape === 'triangle') {
    return (
      <polygon
        fill={color}
        points={`${center},${center - radius} ${center + radius},${center + radius} ${center - radius},${center + radius}`}
      />
    );
  }

  return <circle cx={center} cy={center} fill={color} r={radius} />;
}

export function SeriesMarkerIcon({ color, shape }: SeriesMarkerIconProps) {
  return (
    <svg aria-hidden="true" height="14" viewBox="0 0 14 14" width="14">
      {marker(shape, color, 14)}
    </svg>
  );
}

export function SeriesMarkerDot({
  color,
  cx = 0,
  cy = 0,
  shape,
  size = 14,
}: SeriesMarkerDotProps) {
  return (
    <g transform={`translate(${cx - size / 2} ${cy - size / 2})`}>
      {marker(shape, color, size)}
    </g>
  );
}
