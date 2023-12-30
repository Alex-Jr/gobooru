export const parseQueryString = (
  data: Record<string, undefined | string | string[]>,
  removeEmpty = true
) => {
  const qs = new URLSearchParams();

  Object.entries(data).forEach(([key, value]) => {
    if (!value) return;

    if (removeEmpty && value.length === 0) return;

    if (Array.isArray(value)) {
      value.forEach((v) => {
        qs.append(key, v);
      });
    } else {
      qs.append(key, value);
    }
  });

  return qs;
};
