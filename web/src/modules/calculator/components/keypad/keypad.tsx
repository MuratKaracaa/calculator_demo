import { keypadLayout } from "@/calculator/components/keypad/keypad.constants";
import type { KeypadProps } from "@/calculator/components/keypad/keypad.types";
import type { FC } from "react";
import "@/calculator/components/keypad/keypad.css";
import { KeyButton } from "@/calculator/components/key-button/keyButton";

export const KeyPad: FC<KeypadProps> = ({
  onKeyClick,
  disabledItems,
  disabled,
}) => {
  return (
    <div className="keypad-wrapper">
      {keypadLayout.map((row) => (
        <div className="keypad-row">
          {row.map((item) => (
            <KeyButton
              type={item.type}
              label={item.label}
              onClick={() => onKeyClick(item.value, item.displayValue)}
              disabled={
                !item.alwaysActive &&
                (disabled || disabledItems?.has(item.value))
              }
            />
          ))}
        </div>
      ))}
    </div>
  );
};
