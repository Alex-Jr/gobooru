export const isFileType = (types: string[]) => {
  return (val: File | undefined) => {
    if (!val) return false;

    return types.includes(val.type);
  };
};
