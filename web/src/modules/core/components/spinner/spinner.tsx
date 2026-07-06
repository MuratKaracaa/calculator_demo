import type { FC } from "react";
import { SpinnerSize, type SpinnerProps } from "./spinner.types";
import "@/core/components/spinner/spinner.css";

const sizeMap: Record<SpinnerSize, string> = {
  [SpinnerSize.SMALL]: "spinner-sm",
  [SpinnerSize.MEDIUM]: "spinner-md",
  [SpinnerSize.LARGE]: "spinner-lg",
};

export const Spinner: FC<SpinnerProps> = ({ size = SpinnerSize.MEDIUM }) => {
  return <div className={`spinner ${sizeMap[size]}`} />;
};
