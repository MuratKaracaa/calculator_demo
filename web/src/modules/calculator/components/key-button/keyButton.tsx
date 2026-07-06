import {
  KeyButtonType,
  type KeyButtonProps,
} from "@/calculator/components/key-button/keyButton.types";
import type { FC } from "react";
import "@/calculator/components/key-button/keyButton.css";

export const KeyButton: FC<KeyButtonProps> = ({
  label,
  onClick,
  value,
  type,
}) => {
  const handleClick = () => [onClick(value)];
  return (
    <button
      className={
        type === KeyButtonType.EMPTY ? "key-button-empty" : "key-button-wrapper"
      }
      onClick={handleClick}
    >
      <span className="text-header-4 label">{label}</span>
    </button>
  );
};
