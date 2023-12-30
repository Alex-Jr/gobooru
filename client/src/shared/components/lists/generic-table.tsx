import {
  Table,
  TableContainer,
  Tbody,
  Th,
  Thead,
  ThemingProps,
  Tr,
} from "@chakra-ui/react";
import { ReactElement } from "react";

interface IGenericTableProps<T> {
  headers: string[];
  rows: T[];
  renderRow: (row: T, index: number) => ReactElement;
  size: ThemingProps["size"];
}

export const GenericTable = <T,>({
  headers,
  rows,
  renderRow,
  size,
}: IGenericTableProps<T>) => {
  if (!rows || rows.length === 0) return <></>;

  return (
    <TableContainer flex={1}>
      <Table variant="striped" size={size}>
        <Thead>
          <Tr>
            {headers.map((header) => (
              <Th key={`header-${header}`}>{header}</Th>
            ))}
          </Tr>
        </Thead>
        <Tbody>{rows.map((row, index) => renderRow(row, index))}</Tbody>
      </Table>
    </TableContainer>
  );
};
