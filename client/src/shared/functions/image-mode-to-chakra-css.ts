export const imageModeToChakraCss = (imageMode: string) => {
  switch (imageMode) {
    case "ORIGINAL":
      return { maxW: "unset", h: "unset" };
    case "FIT_V":
      return { maxW: "unset", h: "85vh" };
    case "FIT_H":
      return { maxW: "100%", h: "auto" };
    case "SAMPLE":
      return { maxW: "480px", h: "auto" };
    default:
      return { maxW: "100%", h: "100%" };
  }
};
