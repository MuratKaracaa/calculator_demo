import { KeyPad } from "@/calculator/components/keypad/keypad";
import { OutputDisplay } from "@/calculator/components/output-display/outputDisplay";
import { useCalculatorInput } from "@/calculator/hooks/useCalculatorInput";
import type { FC } from "react";
import "@/calculator/calculator.css";

export const Calculator: FC = () => {
  const {
    handleKeyClick,
    outputs,
    disabledItems,
    isKeypadDisabled,
    isLoading,
    error,
  } = useCalculatorInput();
  return (
    <div className="calculator-outer-wrapper">
      <div className="calculator-wrapper">
        <OutputDisplay outputList={outputs} isLoading={isLoading} />
        <KeyPad
          onKeyClick={handleKeyClick}
          disabledItems={disabledItems}
          disabled={isKeypadDisabled}
        />
      </div>
      <span className="text-body calculator-error">{error}</span>
    </div>
  );
};
