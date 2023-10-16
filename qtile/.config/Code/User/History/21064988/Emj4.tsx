import React from "react";
import { Page, Text, View, Document, StyleSheet } from "@react-pdf/renderer";
import dayjs from "dayjs";
// Create styles
const styles = StyleSheet.create({
  table: {
    display: "table",
    width: "auto",
    borderStyle: "solid",
    borderWidth: 1,
    borderRightWidth: 0,
    borderBottomWidth: 0,
  },
  tableRow: {
    margin: "auto",
    flexDirection: "row",
  },
  tableCol: {
    width: "25%",
    borderStyle: "solid",
    borderWidth: 1,
    borderLeftWidth: 0,
    borderTopWidth: 0,
  },
  tableCell: {
    margin: "auto",
    marginTop: 5,
    fontSize: 10,
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
      <Page style={styles.body}>
        <View style={styles.table}>
          <View style={styles.tableRow}>
            <View style={styles.tableCol}>
              <Text style={styles.tableCell}>Fecha</Text>
            </View>
            <View style={styles.tableCol}>
              <Text style={styles.tableCell}>Pagos</Text>
            </View>
            <View style={styles.tableCol}>
              <Text style={styles.tableCell}>Balance</Text>
            </View>
          </View>
        </View>
        {/*  */}
        <View style={styles.tableRow}>
          <View style={styles.tableCol}>
            <Text style={styles.tableCell}>React-PDF</Text>
          </View>
          <View style={styles.tableCol}>
            <Text style={styles.tableCell}>3 User </Text>
          </View>
          <View style={styles.tableCol}>
            <Text style={styles.tableCell}>2019-02-20 - 2020-02-19</Text>
          </View>
          <View style={styles.tableCol}>
            <Text style={styles.tableCell}>5€</Text>
          </View>
        </View>
        {/*  */}
        {pagos &&
          pagos.map((pago: number, i) => (
            <View style={styles.tableRow} key={i}>
              <View style={styles.tableCol}>
                <Text style={styles.tableCell}>
                  {
                    now
                      .add(1 + i, "month")
                      .toISOString()
                      .split("T")[0]
                  }
                </Text>
              </View>
              <View style={styles.tableCol}>
                <Text style={styles.tableCell}>{pago.toFixed(2)}</Text>
              </View>
              <View style={styles.tableCol}>
                <Text style={styles.tableCell}>
                  {(amountInterest -= pago).toFixed(2)}
                </Text>
              </View>
            </View>
          ))}
      </Page>
    </Document>
  );
};
