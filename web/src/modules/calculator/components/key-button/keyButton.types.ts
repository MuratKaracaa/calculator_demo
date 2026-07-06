export enum KeyButtonType {
  EMPTY = "EMPTY",
  NUMBER = "NUMBER",
  OPERATOR = "OPERATOR",
}

export interface KeyButtonProps {
  type: KeyButtonType;
  label: string;
  value: string;
  onClick: (value: string) => void;
}
