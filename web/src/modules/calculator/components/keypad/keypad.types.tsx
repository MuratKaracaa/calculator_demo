export interface KeypadProps {
  onKeyClick: (value: string, displayValue: string) => void;
  disabledItems?: Set<string>;
  disabled: boolean;
}
