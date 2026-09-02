import type { SVGProps } from 'react'

// A dozen 16px stroke icons, hand-rolled. An icon *font* or an icon package
// would be several megabytes of dependency for shapes that are four path
// commands each, and the design system asks for one accent colour and thin
// lines — so these inherit `currentColor` and a 1.5 stroke throughout.

export type IconProps = SVGProps<SVGSVGElement> & { size?: number }

function Icon({ size = 16, children, ...rest }: IconProps) {
  return (
    <svg
      width={size}
      height={size}
      viewBox='0 0 16 16'
      fill='none'
      stroke='currentColor'
      strokeWidth={1.5}
      strokeLinecap='round'
      strokeLinejoin='round'
      aria-hidden='true'
      focusable='false'
      {...rest}
    >
      {children}
    </svg>
  )
}

export const SearchIcon = (p: IconProps) => (
  <Icon {...p}>
    <circle cx='7' cy='7' r='4.5' />
    <path d='M10.5 10.5 14 14' />
  </Icon>
)

export const CopyIcon = (p: IconProps) => (
  <Icon {...p}>
    <rect x='5.5' y='5.5' width='8' height='8' rx='1.5' />
    <path d='M10.5 3.5h-7a1 1 0 0 0-1 1v7' />
  </Icon>
)

export const CheckIcon = (p: IconProps) => (
  <Icon {...p}>
    <path d='m3 8.5 3.2 3.2L13 4.8' />
  </Icon>
)

export const ExternalIcon = (p: IconProps) => (
  <Icon {...p}>
    <path d='M9 3h4v4' />
    <path d='M13 3 7.5 8.5' />
    <path d='M12 10v2.5a1 1 0 0 1-1 1H3.5a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1H6' />
  </Icon>
)

export const ChevronDownIcon = (p: IconProps) => (
  <Icon {...p}>
    <path d='m4 6 4 4 4-4' />
  </Icon>
)

export const ChevronLeftIcon = (p: IconProps) => (
  <Icon {...p}>
    <path d='m10 3-5 5 5 5' />
  </Icon>
)

export const ChevronRightIcon = (p: IconProps) => (
  <Icon {...p}>
    <path d='m6 3 5 5-5 5' />
  </Icon>
)

export const ArrowUpIcon = (p: IconProps) => (
  <Icon {...p}>
    <path d='M8 13V3' />
    <path d='m4 7 4-4 4 4' />
  </Icon>
)

export const ArrowDownIcon = (p: IconProps) => (
  <Icon {...p}>
    <path d='M8 3v10' />
    <path d='m4 9 4 4 4-4' />
  </Icon>
)

export const CloseIcon = (p: IconProps) => (
  <Icon {...p}>
    <path d='m4 4 8 8M12 4l-8 8' />
  </Icon>
)

export const MenuIcon = (p: IconProps) => (
  <Icon {...p}>
    <path d='M2.5 4.5h11M2.5 8h11M2.5 11.5h11' />
  </Icon>
)

export const InfoIcon = (p: IconProps) => (
  <Icon {...p}>
    <circle cx='8' cy='8' r='6' />
    <path d='M8 7.5v3.5M8 5.2v.6' />
  </Icon>
)

export const WarningIcon = (p: IconProps) => (
  <Icon {...p}>
    <path d='M8 2.6 14.2 13H1.8z' />
    <path d='M8 6.5v3M8 11.2v.5' />
  </Icon>
)

export const DotIcon = ({ size = 8, ...rest }: IconProps) => (
  <svg width={size} height={size} viewBox='0 0 8 8' aria-hidden='true' focusable='false' {...rest}>
    <circle cx='4' cy='4' r='4' fill='currentColor' />
  </svg>
)
