const wordRegex = /\w+/g;

export const hasNWords = (minWords: number) => {
  return (val: string | undefined) => {
    if (!val) return false;

    const words = val.trim().replaceAll("\n", " ").match(wordRegex);

    if (words === null) {
      return false;
    }

    return words.length >= minWords;
  };
};
