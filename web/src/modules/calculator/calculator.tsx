import { KeyPad } from "@/calculator/components/keypad/keypad";
import { OutputDisplay } from "@/calculator/components/output-display/outputDisplay";
import { useCalculatorInput } from "@/calculator/hooks/useCalculatorInput";
import type { FC } from "react";
import "@/calculator/calculator.css";

export const Calculator: FC = () => {
  const { handleKeyClick, outputs, disabledItems } = useCalculatorInput();
  return (
    <div className="calculator-wrapper">
      <OutputDisplay outputList={outputs} />
      <KeyPad onKeyClick={handleKeyClick} disabledItems={disabledItems} />
    </div>
  );
};
