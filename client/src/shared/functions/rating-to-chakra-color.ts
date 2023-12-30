import { RatingEnum } from "shared/types/enums/rating-enum";

export const ratingToChakraColor = (s: RatingEnum): string => {
  switch (s) {
    case RatingEnum.EXPLICIT:
      return "red.500";
    case RatingEnum.QUESTIONABLE:
      return "yellow.500";
    case RatingEnum.SAFE:
      return "green.500";
    default:
      return "gray";
  }
};
