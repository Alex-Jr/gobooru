import { useSearchParams } from "react-router-dom";

export const useQueryString = <S extends string, A extends string>({
  asString,
  asArray,
}: {
  asString?: S[];
  asArray?: A[];
}): {
  asString: { [K in S]: string | undefined };
  asArray: { [K in A]: string[] | undefined };
  searchParams: URLSearchParams;
  setSearchParams: ReturnType<typeof useSearchParams>[1];
} => {
  const asStringParsed = {} as {
    [K in S]: string | undefined;
  };

  const asArrayParsed = {} as {
    [K in A]: string[] | undefined;
  };

  const [searchParams, setSearchParams] = useSearchParams();

  asString?.forEach((param) => {
    asStringParsed[param] = searchParams.get(param) ?? undefined;
  });

  asArray?.forEach((param) => {
    const paramData = searchParams.getAll(param);
    asArrayParsed[param] = paramData.length > 0 ? paramData : undefined;
  });

  return {
    asString: asStringParsed,
    asArray: asArrayParsed,
    searchParams,
    setSearchParams,
  };
};
