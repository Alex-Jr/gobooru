import { Button, FormLabel, HStack } from "@chakra-ui/react";

import { RatingEnum } from "shared/types/enums/rating-enum";

interface RatingButtonsProps<Name extends string> {
  name: Name;
  values: { [key in Name]: string[] };
  push: (s: any) => void;
  remove: (n: number) => void;
}

export const RatingButtons = <Name extends string>({
  name,
  values,
  push,
  remove,
}: RatingButtonsProps<Name>) => {
  const ratingsObj: Record<RatingEnum, boolean> = {
    E: values[name].includes(RatingEnum.EXPLICIT),
    Q: values[name].includes(RatingEnum.QUESTIONABLE),
    S: values[name].includes(RatingEnum.SAFE),
  };

  const handleOnClick = (ratings: RatingEnum) => {
    if (ratingsObj[ratings]) {
      remove(values[name].indexOf(ratings));
    } else {
      push(ratings);
    }
  };

  return (
    <HStack justify={"space-between"}>
      <FormLabel m={0} htmlFor={name}>
        Rating
      </FormLabel>
      <HStack>
        <Button
          colorScheme={ratingsObj[RatingEnum.EXPLICIT] ? "red" : "gray"}
          onClick={() => {
            handleOnClick(RatingEnum.EXPLICIT);
          }}
          size={"sm"}
        />
        <Button
          colorScheme={ratingsObj[RatingEnum.QUESTIONABLE] ? "yellow" : "gray"}
          onClick={() => {
            handleOnClick(RatingEnum.QUESTIONABLE);
          }}
          size={"sm"}
        />
        <Button
          colorScheme={ratingsObj[RatingEnum.SAFE] ? "green" : "gray"}
          onClick={() => {
            handleOnClick(RatingEnum.SAFE);
          }}
          size={"sm"}
        />
      </HStack>
    </HStack>
  );
};
