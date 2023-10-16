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
      <Page size="A4" style={styles.page}></Page>
    </Document>
  );
};
