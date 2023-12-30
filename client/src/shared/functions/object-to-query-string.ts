export const objectToQueryString = (queryStringData: Record<string, any>) => {
  const queryString = Object.entries(queryStringData)
    .map(([key, value]) => {
      if (Array.isArray(value)) {
        return value.map(
          (v) => `${encodeURIComponent(key)}=${encodeURIComponent(v)}`
        );
      } else {
        return `${encodeURIComponent(key)}=${encodeURIComponent(value)}`;
      }
    })
    .flat(1)
    .join("&");

  return queryString;
};
