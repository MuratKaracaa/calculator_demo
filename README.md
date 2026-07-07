# Full Stack Calculator Demo (Murat Karaca)

## Prerequisites

- Go 1.25.0+ (tested with 1.26.4)
- Node.js 22.12+ (tested with Node 24.16.0)

## Setup instructions (Backend)

- `cd server`
- `go mod download`
- Run directly with `go run ./cmd/api`
- or build with `go build -o ./bin/app ./cmd/api`, for Windows `go build -o ./bin/app.exe ./cmd/api`
- After building run with `./bin/app` or `./bin/app.exe`

## Setup instructions (Frontend)

- `cd web`
- `npm install` or `yarn install`
- Run directly with `npm run dev` or `yarn dev`

## Environment variables

- There is no environment configuration. Both projects should be run locally.

## Supported Operations

- Addition (+), Subtraction (-), Multiplication (x), Division (/), Exponentiation(^), Square Root(√), Percentage(%)

## API

- Backend exposes only one endpoint for all operations "/calculator/calculate". It accepts a json object with a single field named "expression" of type string and returns an object with field "result" or "error".

- The "expression" field accepts a string containing digits and characters for the operations.

## API Example

- POST /calculator/calculate

- Request

```json
{
  "expression": "s(2^3+4^2)%+100"
}
```

- Response

```json
{
  "result": 100.04898979485566
}
```

- Request

```json
{
  "expression": "5/0"
}
```

- Response

```json
{
  "error": "division by zero at position 1"
}
```

## Design

- The frontend is responsible only for maintaining an interactive ui for the users and directing them to create a sensible mathematical expression, and displaying the result of the evaluation.

- The API accepts mathematical expressions instead of separate operation endpoints. This keeps the interface simple and allows the backend implementation to evolve without changing the frontend contract.

- For parsing and computations, the "Shunting Yard" algorithm by Edsger Dijkstra was followed. The service contains an implementation of the two-phased algorithm and test cases.

- The API initially checks for valid characters. Grammar validation is performed before parsing to improve reliability, testability, and prevent invalid expressions from reaching the evaluator.

## Operations

- The Addition operation is represented with the `+` character and it is a binary operation, requiring before and after elements

- The Subtraction operation is represented with the `-` character and it is a binary operation although '-' character is internally utilized for unary minus evaluations as well.

- The Multiplication operation is represented with the `x` character and it is a binary operation.

- The Division operation is represented with the `/` character and it is a binary operation.

- The Exponentiation operation is represented with the `^` character and it is a binary operation.

- The square root operation is represented by the `s` character internally. It is implemented as a unary operation and supports both direct numbers and expressions that evaluate to a number.

- The percentage operation is represented with the `%` character and it is designed to follow any number or any computable expression.

## Testing

### Backend

Unit tests cover the calculator engine and expression parser.

- `cd server`
- `go test ./...`

```
ok      server/internal/calculator      0.390s  coverage: 91.6% of statements
```

### Frontend

No automated frontend tests were added. The frontend was manually tested for UI interactions, expression input handling, and API communication.

## AI Usage

AI tools were used during development for:

- Generating test cases
- Simulating the execution of the algorithm for comparisons and debugging
- Reviewing implementation details
- Generating a small typography scale and reviewing UI styling decisions

Example prompts

- For an implementation of the "Shunting Yard" algorithm, generate expressions for basic and advanced arithmetic operations

- Given this input, display the step by step evolution of the stacks and the queues while parsing the expression

- Review this Shunting Yard implementation. Unary minus is supported, but operator precedence is incorrect. Identify possible issues.
