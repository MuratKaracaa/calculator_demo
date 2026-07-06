import {
  OutputState,
  type OutputDisplayProps,
} from "@/calculator/components/output-display/outputDisplay.types";
import type { FC } from "react";
import "@/calculator/components/output-display/outputDisplay.css";
import { SpinnerSize } from "@/core/components/spinner/spinner.types";
import { Spinner } from "@/core/components/spinner/spinner";

export const OutputDisplay: FC<OutputDisplayProps> = ({
  outputList,
  isLoading,
}) => {
  return (
    <div className="output-display-wrapper">
      {isLoading ? (
        <Spinner size={SpinnerSize.LARGE} />
      ) : (
        outputList.map((item) => (
          <span
            className={`text-header-3 ${item.state === OutputState.INVALID ? "output-display-item-invalid" : ""}`}
          >
            {item.label}
          </span>
        ))
      )}
    </div>
  );
};
