export const removeHttpAndWWW = (url: string) =>
  url.replace(/^(?:https?:\/\/)?(?:www\.)?/i, "");
