import {
  OutputState,
  type OutputItemProps,
} from "@/calculator/components/output-display/outputDisplay.types";
import { useCallback, useState } from "react";

export const disabledByLabel: Record<string, Set<string>> = {
  "(": new Set(["+", "x", "/", "^", "%", ")", "submit", "."]),

  ")": new Set(["(", ".", "submit", "s"]),

  "%": new Set(["+", "x", "/", "^", "%", "(", ")", ".", "submit", "s"]),

  "÷": new Set(["+", "x", "/", "^", "%", ")", ".", "submit"]),

  "√": new Set(["+", "x", "/", "^", "%", ")", ".", "submit"]),

  "^": new Set(["+", "x", "/", "^", "%", ")", ".", "submit"]),

  x: new Set(["+", "x", "/", "^", "%", ")", ".", "submit"]),

  "+": new Set(["+", "x", "/", "^", "%", ")", ".", "submit"]),

  "-": new Set(["+", "x", "/", "^", "%", ")", ".", "submit"]),

  ".": new Set([".", "submit", "(", ")", "%", "s", "^", "x", "+", "-"]),
  "0": new Set(["+", ".", "submit", "%", "/", "x", "^", ")", "0"]),
  "1": new Set(["(", "s"]),
  "2": new Set(["(", "s"]),
  "3": new Set(["(", "s"]),
  "4": new Set(["(", "s"]),
  "5": new Set(["(", "s"]),
  "6": new Set(["(", "s"]),
  "7": new Set(["(", "s"]),
  "8": new Set(["(", "s"]),
  "9": new Set(["(", "s"]),
};

const operators = new Set(["+", "-", "x", "÷", "%", "^"]);

export const useCalculatorInput = () => {
  const [outputs, setOutputs] = useState<Array<OutputItemProps>>([
    { label: "0", state: OutputState.DEFAULT, value: "0" },
  ]);
  const [isWritingDecimal, setIsWritingDecimal] = useState(false);
  const disabledItems = disabledByLabel[outputs[outputs.length - 1].label];
  if (isWritingDecimal) {
    disabledItems.add(".");
  }

  const handleKeyClick = useCallback((value: string, displayValue: string) => {
    setOutputs((prev) => {
      const newList = [...prev];

      if (value === "backspace" || value === "clear" || value === "submit") {
        switch (value) {
          case "backspace":
            newList.pop();
            break;
          case "clear":
            return [{ label: "0", state: OutputState.DEFAULT, value: "0" }];
          case "submit":
            break;
        }
      } else {
        if (displayValue === ".") {
          setIsWritingDecimal(true);
        }

        if (operators.has(displayValue)) {
          setIsWritingDecimal(false);
        }
        newList.push({
          label: displayValue,
          state: OutputState.DEFAULT,
          value: value,
        });
      }

      if (newList.length == 0) {
        return [{ label: "0", state: OutputState.DEFAULT, value: "0" }];
      } else if (newList[0].label == "0") {
        newList.shift();
      }

      return newList;
    });
  }, []);

  return { outputs, handleKeyClick, disabledItems };
};
