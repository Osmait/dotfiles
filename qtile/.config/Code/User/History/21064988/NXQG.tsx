import React from "react";
import { Page, Document, StyleSheet } from "@react-pdf/renderer";
import { Table, Text } from "@nextui-org/react";
import dayjs from "dayjs";
// Create styles
const styles = StyleSheet.create({
  page: {
    flexDirection: "row",
    backgroundColor: "#E4E4E4",
  },
  section: {
    margin: 10,
    padding: 10,
    flexGrow: 1,
  },
});

export const PdfTable = ({ amount, interest, date, tiempo }: any) => {
  const pagos = Array(tiempo).fill(0);
  let now = dayjs(date);

  for (let i = 0; i < pagos.length; i++) {
    pagos[i] = (amount * interest) / tiempo;
  }
  let amountInterest = amount * interest;

  return (
    <Document>
      <Page size="A4" style={styles.page}>
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
              pagos.map((pago: number, i) => (
                <Table.Row key={i}>
                  <Table.Cell>
                    {
                      now
                        .add(1 + i, "month")
                        .toISOString()
                        .split("T")[0]
                    }
                  </Table.Cell>
                  <Table.Cell>{pago.toFixed(2)}</Table.Cell>
                  <Table.Cell>{(amountInterest -= pago).toFixed(2)}</Table.Cell>
                </Table.Row>
              ))}
          </Table.Body>
        </Table>
        <Text h3> Total: ${amount * interest}</Text>
        <Text h3> Ganancias: ${amount * interest - amount}</Text>
      </Page>
    </Document>
  );
};
