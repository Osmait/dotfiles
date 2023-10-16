import { Button, Container } from "@nextui-org/react";
import React from "react";

type Props = {
  children: JSX.Element;
};

export const Error = ({ children }: Props) => {
  return <Container>{children}</Container>;
};
