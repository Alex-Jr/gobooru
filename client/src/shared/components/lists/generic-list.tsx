import { Flex, Grid, Spinner } from "@chakra-ui/react";
import { ReactElement } from "react";

interface IGenericGrid<T> {
  w: string;
  items?: T[];
  renderItem: (item: T, index?: number) => ReactElement;
}

export const GenericGrid = <T,>({ items, renderItem, w }: IGenericGrid<T>) => {
  if (!items)
    return (
      <Flex w={"100%"}>
        <Spinner m={"auto"} />
      </Flex>
    );

  return (
    <Grid
      gap={4}
      flex={1}
      gridTemplateColumns={`repeat(auto-fill, ${w})`}
      placeItems={{ sm: "center", lg: "unset" }}
      justifyContent={{ sm: "center", lg: "unset" }}
    >
      {items.map((item, index) => renderItem(item, index))}
    </Grid>
  );
};
