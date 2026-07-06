export enum KeyButtonType {
  EMPTY = "EMPTY",
  NUMBER = "NUMBER",
  OPERATOR = "OPERATOR",
}

export interface KeyButtonProps {
  type: KeyButtonType;
  label: string;
  onClick: () => void;
  disabled?: boolean;
}
