export enum OutputState {
  DEFAULT = "DEFAULT",
  INVALID = "INVALID",
}

export interface OutputItemProps {
  label: string;
  value: string;
  state: OutputState;
}

export interface OutputDisplayProps {
  outputList: Array<OutputItemProps>;
  isLoading?: boolean;
}
