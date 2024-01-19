import { Box, Tooltip } from "@chakra-ui/react";
import { useTranslation } from "react-i18next";

import { usePostsList } from "services/posts/use-posts-list";

export const PostCount = () => {
  const { t } = useTranslation();

  const { count } = usePostsList({
    search: "",
    page: "1",
    page_size: "20",
  });

  return (
    <Tooltip label={t("homepage.countLabel")}>
      <Box fontSize={"4xl"} letterSpacing={"8px"}>
        {count}
      </Box>
    </Tooltip>
  );
};
