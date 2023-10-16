import useLoans from "@/hooks/usePrestamos";
import Login from "@/pages/login";
import { render, screen } from "@testing-library/react";
jest.mock("next/router", () => ({
  push: jest.fn(),
  back: jest.fn(),
  events: {
    on: jest.fn(),
    off: jest.fn(),
  },
  beforePopState: jest.fn(() => null),
  useRouter: () => ({
    push: jest.fn(),
  }),
}));
useLoans;

jest.mock("@/hooks/usePrestamos", () => ({
  setError: false,
  error: false,
}));

describe("Login Test", () => {
  it("Render Componet", () => {
    render(<Login />);
  });
});
