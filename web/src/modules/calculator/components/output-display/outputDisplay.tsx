import type { OutputDisplayProps } from "@/calculator/components/output-display/outputDisplay.types";
import type { FC } from "react";
import "@/calculator/components/output-display/outputDisplay.css";

export const OutputDisplay: FC<OutputDisplayProps> = ({ outputList }) => {
  return (
    <div className="output-display-wrapper">
      {outputList.map((item) => (
        <span>{item.label}</span>
      ))}
    </div>
  );
};
