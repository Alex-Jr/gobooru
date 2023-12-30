import { Box, Checkbox, CheckboxProps, Tooltip } from "@chakra-ui/react";
import { useTranslation } from "react-i18next";

import { RatingEnum } from "shared/types/enums/rating-enum";

interface RatingButtonsProps extends Omit<CheckboxProps, "value"> {
  name: string;
  onChange: (event: any) => void;
  value: string[];
}

export const RatingButtons = ({
  name,
  onChange,
  value: values,
  ...rest
}: RatingButtonsProps) => {
  const { t } = useTranslation();

  const ratings = [
    {
      label: t("glossary.explicit"),
      value: RatingEnum.EXPLICIT,
      colorScheme: "red",
    },
    {
      label: t("glossary.questionable"),
      value: RatingEnum.QUESTIONABLE,
      colorScheme: "orange",
    },
    {
      label: t("glossary.safe"),
      value: RatingEnum.SAFE,
      colorScheme: "green",
    },
  ];

  return (
    <>
      {ratings.map(({ value, colorScheme, label }) => (
        <Tooltip label={label.toLocaleLowerCase()} key={name + value}>
          <Box w={"fit-content"} h={"fit-content"}>
            <Checkbox
              value={value}
              colorScheme={colorScheme}
              isChecked={values.some((v) => v === value)}
              onChange={(e) => {
                onChange(
                  e.target.checked
                    ? [...values, e.target.value]
                    : values.filter((v) => v !== e.target.value)
                );
              }}
              {...rest}
            />
          </Box>
        </Tooltip>
      ))}
    </>
  );
};
