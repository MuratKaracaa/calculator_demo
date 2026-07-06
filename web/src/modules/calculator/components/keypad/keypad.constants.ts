import {
  KeyButtonType,
  type KeyButtonProps,
} from "@/calculator/components/key-button/keyButton.types";

export const keypadLayout: (Omit<KeyButtonProps, "onClick"> & {
  displayValue: string;
  value: string;
  alwaysActive?: boolean;
})[][] = [
  [
    {
      type: KeyButtonType.OPERATOR,
      label: "C",
      displayValue: "C",
      value: "clear",
      alwaysActive: true,
    },
    {
      type: KeyButtonType.OPERATOR,
      label: "⌫",
      displayValue: "⌫",
      value: "backspace",
      alwaysActive: true,
    },
    { type: KeyButtonType.OPERATOR, label: "%", displayValue: "%", value: "%" },
    { type: KeyButtonType.OPERATOR, label: "(", displayValue: "(", value: "(" },
    { type: KeyButtonType.OPERATOR, label: ")", displayValue: ")", value: ")" },
  ],
  [
    { type: KeyButtonType.NUMBER, label: "7", displayValue: "7", value: "7" },
    { type: KeyButtonType.NUMBER, label: "8", displayValue: "8", value: "8" },
    { type: KeyButtonType.NUMBER, label: "9", displayValue: "9", value: "9" },
    { type: KeyButtonType.OPERATOR, label: "÷", displayValue: "÷", value: "/" },
    { type: KeyButtonType.OPERATOR, label: "√", displayValue: "√", value: "s" },
  ],
  [
    { type: KeyButtonType.NUMBER, label: "4", displayValue: "4", value: "4" },
    { type: KeyButtonType.NUMBER, label: "5", displayValue: "5", value: "5" },
    { type: KeyButtonType.NUMBER, label: "6", displayValue: "6", value: "6" },
    { type: KeyButtonType.OPERATOR, label: "x", displayValue: "x", value: "x" },
    { type: KeyButtonType.OPERATOR, label: "^", displayValue: "^", value: "^" },
  ],
  [
    { type: KeyButtonType.NUMBER, label: "1", displayValue: "1", value: "1" },
    { type: KeyButtonType.NUMBER, label: "2", displayValue: "2", value: "2" },
    { type: KeyButtonType.NUMBER, label: "3", displayValue: "3", value: "3" },
    { type: KeyButtonType.OPERATOR, label: "-", displayValue: "-", value: "-" },
    { type: KeyButtonType.OPERATOR, label: "+", displayValue: "+", value: "+" },
  ],
  [
    { type: KeyButtonType.EMPTY, label: "", displayValue: "", value: "" },
    { type: KeyButtonType.NUMBER, label: "0", displayValue: "0", value: "0" },
    { type: KeyButtonType.EMPTY, label: "", displayValue: "", value: "" },
    { type: KeyButtonType.NUMBER, label: ".", displayValue: ".", value: "." },
    {
      type: KeyButtonType.OPERATOR,
      label: "=",
      displayValue: "=",
      value: "submit",
    },
  ],
];
