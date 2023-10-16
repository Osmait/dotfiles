import { Button, Tooltip } from "@nextui-org/react";
import React from "react";

export const Error = () => {
  return (
    <Tooltip content="Developers love Next.js" color="error">
      <Button flat auto color="error">
        Error
      </Button>
    </Tooltip>
  );
};
