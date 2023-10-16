import { Button } from "@nextui-org/react";
import React from "react";

export const Error = ({ children }) => {
  return (
    <Button auto color="error">
      {children}
    </Button>
  );
};
