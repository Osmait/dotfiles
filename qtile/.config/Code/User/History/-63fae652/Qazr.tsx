import { Table } from "@nextui-org/react";
import dayjs from "dayjs";
import React from "react";

export const InterestTable = ({ amount }: any) => {
  const pagos = Array(6).fill(0);
  let now = dayjs(new Date());

  for (let i = 0; i < pagos.length; i++) {
    pagos[i] = (amount * 1.2) / 6;
  }
  const amountInterest = amount * 1.2;
  console.log(amountInterest);
  console.log(now.toISOString());

  return (
    <Table
      aria-label="Example static collection table"
      css={{
        height: "auto",
        minWidth: "100%",
      }}
      selectionMode="single"
    >
      <Table.Header>
        <Table.Column>Fecha</Table.Column>
        <Table.Column>Pagos</Table.Column>
        <Table.Column>Balance</Table.Column>
      </Table.Header>
      <Table.Body>
        {pagos &&
          pagos.map((pago: any, i) => (
            <Table.Row key={i}>
              <Table.Cell>
                {
                  now
                    .add(1 + i, "month")
                    .toISOString()
                    .split("T")[0]
                }
              </Table.Cell>
              <Table.Cell>{pago}</Table.Cell>
              <Table.Cell>{amount}</Table.Cell>
            </Table.Row>
          ))}
      </Table.Body>
    </Table>
  );
};
