import {
  KeyButtonType,
  type KeyButtonProps,
} from "@/calculator/components/key-button/keyButton.types";
import type { FC } from "react";
import "@/calculator/components/key-button/keyButton.css";

export const KeyButton: FC<KeyButtonProps> = ({
  label,
  onClick,
  type,
  disabled,
}) => {
  return (
    <button
      disabled={disabled}
      className={
        type === KeyButtonType.EMPTY ? "key-button-empty" : "key-button-wrapper"
      }
      onClick={onClick}
    >
      <span className="text-header-4 label">{label}</span>
    </button>
  );
};
