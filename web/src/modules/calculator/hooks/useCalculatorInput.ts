import {
  OutputState,
  type OutputItemProps,
} from "@/calculator/components/output-display/outputDisplay.types";
import { useSubmitExpression } from "@/calculator/hooks/useSubmitExpression";
import { useCallback, useRef, useState } from "react";

export const disabledByLabel: Record<string, Array<string>> = {
  "(": ["+", "x", "/", "^", "%", ")", "submit", "."],

  ")": ["(", ".", "s", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", ")"],

  "%": ["+", "x", "/", "^", "%", "(", ")", ".", "submit", "s"],

  "÷": ["+", "x", "/", "^", "%", ")", ".", "submit"],

  "√": ["+", "x", "/", "^", "%", ")", ".", "submit"],

  "^": ["+", "x", "/", "^", "%", ")", ".", "submit"],

  x: ["+", "x", "/", "^", "%", ")", ".", "submit"],

  "+": ["+", "x", "/", "^", "%", ")", ".", "submit"],

  "-": ["+", "x", "/", "^", "%", ")", ".", "submit"],

  ".": [".", "submit", "(", ")", "%", "s", "^", "x", "+", "-"],
  "0": ["+", ".", "submit", "%", "/", "x", "^", ")", "0"],
  "1": ["(", "s"],
  "2": ["(", "s"],
  "3": ["(", "s"],
  "4": ["(", "s"],
  "5": ["(", "s"],
  "6": ["(", "s"],
  "7": ["(", "s"],
  "8": ["(", "s"],
  "9": ["(", "s"],
};

const operators = new Set(["+", "-", "x", "÷", "%", "^"]);

export const useCalculatorInput = () => {
  const { trigger: submitExpression, isLoading } = useSubmitExpression();

  const [outputs, setOutputs] = useState<Array<OutputItemProps>>([
    { label: "0", state: OutputState.DEFAULT, value: "0" },
  ]);

  const isWritingDecimal = useRef(false);

  const disabledItems = new Set(
    disabledByLabel[outputs[outputs.length - 1].label],
  );

  const openedParanthesisCount = useRef(0);

  if (isWritingDecimal.current) {
    disabledItems.add(".");
  }

  if (openedParanthesisCount.current === 0) {
    disabledItems.add(")");
  }

  if (outputs.length > 2) {
    disabledItems.delete("submit");
  }

  const isKeypadDisabled = outputs.length === 15 || isLoading;

  const handleSubmit = async () => {
    const expression = outputs.map((item) => item.value).join("");
    const response = await submitExpression({ expression });
    const stringified = response.result.toFixed(2);
    const newOutPuts = stringified.split("").map((digit) => ({
      label: digit,
      state: OutputState.DEFAULT,
      value: digit,
    }));
    setOutputs(newOutPuts);
  };

  const handleKeyClick = (value: string, displayValue: string) => {
    if (value === "submit") {
      handleSubmit();
    } else {
      setOutputs((prev) => {
        const newList = [...prev];

        if (value === "backspace" || value === "clear" || value === "submit") {
          switch (value) {
            case "backspace":
              const deletedItem = newList.pop();
              if (deletedItem?.value === ")") {
                openedParanthesisCount.current--;
                const firstParanthesisItem = newList.find(
                  (item) => item.label === "(",
                );
                if (firstParanthesisItem) {
                  firstParanthesisItem.state = OutputState.INVALID;
                }
              }
              break;
            case "clear":
              return [{ label: "0", state: OutputState.DEFAULT, value: "0" }];
            case "submit":
              break;
          }
        } else {
          if (displayValue === ".") {
            isWritingDecimal.current = true;
          }

          if (operators.has(displayValue)) {
            isWritingDecimal.current = false;
          }
          if (displayValue === "(") {
            newList.push({
              label: displayValue,
              state: OutputState.INVALID,
              value: value,
            });
            openedParanthesisCount.current++;
          } else {
            if (displayValue === ")") {
              openedParanthesisCount.current--;
              for (let i = newList.length - 1; i >= 0; i--) {
                const item = newList[i];
                if (item.state === OutputState.INVALID) {
                  item.state = OutputState.DEFAULT;
                  break;
                }
                console.log(newList);
              }
            }
            newList.push({
              label: displayValue,
              state: OutputState.DEFAULT,
              value: value,
            });
          }
        }

        if (newList.length == 0) {
          return [{ label: "0", state: OutputState.DEFAULT, value: "0" }];
        } else if (newList[0].label == "0") {
          newList.shift();
        }

        return newList;
      });
    }
  };

  return {
    outputs,
    handleKeyClick,
    disabledItems,
    isKeypadDisabled,
    isLoading,
  };
};
