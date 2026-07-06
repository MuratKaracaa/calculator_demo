import {
  KeyButtonType,
  type KeyButtonProps,
} from "@/calculator/components/key-button/keyButton.types";

export const keypadLayout: Omit<KeyButtonProps, "onClick">[][] = [
  [
    { type: KeyButtonType.OPERATOR, label: "C", value: "clear" },
    { type: KeyButtonType.OPERATOR, label: "⌫", value: "backspace" },
    { type: KeyButtonType.OPERATOR, label: "%", value: "%" },
    { type: KeyButtonType.OPERATOR, label: "(", value: "(" },
    { type: KeyButtonType.OPERATOR, label: ")", value: ")" },
  ],
  [
    { type: KeyButtonType.NUMBER, label: "7", value: "7" },
    { type: KeyButtonType.NUMBER, label: "8", value: "8" },
    { type: KeyButtonType.NUMBER, label: "9", value: "9" },
    { type: KeyButtonType.OPERATOR, label: "÷", value: "÷" },
    { type: KeyButtonType.OPERATOR, label: "x²", value: "s" },
  ],
  [
    { type: KeyButtonType.NUMBER, label: "4", value: "4" },
    { type: KeyButtonType.NUMBER, label: "5", value: "5" },
    { type: KeyButtonType.NUMBER, label: "6", value: "6" },
    { type: KeyButtonType.OPERATOR, label: "×", value: "×" },
    { type: KeyButtonType.OPERATOR, label: "^", value: "^" },
  ],
  [
    { type: KeyButtonType.NUMBER, label: "1", value: "1" },
    { type: KeyButtonType.NUMBER, label: "2", value: "2" },
    { type: KeyButtonType.NUMBER, label: "3", value: "3" },
    { type: KeyButtonType.OPERATOR, label: "-", value: "-" },
    { type: KeyButtonType.OPERATOR, label: "+", value: "+" },
  ],
  [
    { type: KeyButtonType.EMPTY, label: "", value: "" },
    { type: KeyButtonType.NUMBER, label: "0", value: "0" },
    { type: KeyButtonType.EMPTY, label: "", value: "" },
    { type: KeyButtonType.NUMBER, label: ".", value: "." },
    { type: KeyButtonType.OPERATOR, label: "=", value: "equals" },
  ],
];
