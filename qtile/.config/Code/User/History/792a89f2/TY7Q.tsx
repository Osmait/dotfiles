import { Button, Card, Container } from "@nextui-org/react";
import React from "react";

type Props = {
  children: JSX.Element;
};

export const Error = ({ children }: Props) => {
  return (
    <Container>
      <Card css={{ $$cardColor: "$colors$primary" }}>
        <Card.Body>{children}</Card.Body>
      </Card>
    </Container>
  );
};
