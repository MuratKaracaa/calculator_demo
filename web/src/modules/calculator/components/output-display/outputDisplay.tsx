import {
  OutputState,
  type OutputDisplayProps,
} from "@/calculator/components/output-display/outputDisplay.types";
import type { FC } from "react";
import "@/calculator/components/output-display/outputDisplay.css";

export const OutputDisplay: FC<OutputDisplayProps> = ({ outputList }) => {
  return (
    <div className="output-display-wrapper">
      {outputList.map((item) => (
        <span
          className={`text-header-3 ${item.state === OutputState.INVALID ? "output-display-item-invalid" : ""}`}
        >
          {item.label}
        </span>
      ))}
    </div>
  );
};
