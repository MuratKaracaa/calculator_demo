import { useState } from "react";

interface CalculationRequest {
  expression: string;
}

interface CalculationResponse {
  result: number;
  error?: string;
}

export const useSubmitExpression = () => {
  const [data, setData] = useState<CalculationResponse | null>(null);
  const [isLoading, setIsLoading] = useState<boolean>(false);

  const trigger = async (
    request: CalculationRequest,
  ): Promise<CalculationResponse> => {
    setIsLoading(true);

    try {
      const response = await fetch(
        "http://localhost:8080/calculator/calculate",
        {
          method: "POST",
          body: JSON.stringify(request),
        },
      );
      const result = await response.json();

      setData(result);

      return result;
    } finally {
      setIsLoading(false);
    }
  };

  return { trigger, data, isLoading };
};
