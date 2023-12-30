import { Box, Flex, HStack, Link, Text, Tooltip } from "@chakra-ui/react";
import { Link as RouterLink } from "react-router-dom";

import { ImageModeEnum } from "shared/types/enums/image-mode-enum";
import { APIPost } from "shared/types/services/posts/APIPost";

import { FilePreview } from "./file-preview";

interface PostRelationsProps {
  relations?: APIPost["post"]["relations"];
}

const PostItem = ({
  label,
  relation,
}: {
  label: string;
  relation: APIPost["post"]["relations"][0];
}) => {
  return (
    <Tooltip label={label}>
      <Link as={RouterLink} to={`/posts/${relation.other_post.id}`}>
        <Flex w={125} h={125}>
          <FilePreview
            filePath={relation.other_post.thumb_path}
            imageMode={ImageModeEnum.FIT_H}
          />
        </Flex>
      </Link>
    </Tooltip>
  );
};

export const PostRelations = ({ relations }: PostRelationsProps) => {
  if (!relations || relations.length === 0) return <></>;

  const alreadyRender: number[] = [];

  const parent = relations.find((r) => r.type === "PARENT");

  return (
    <Box p={4} bgColor={"gray.700"}>
      <Text>Relations</Text>
      <HStack>
        {parent && (
          <PostItem relation={parent} label={`Post: ${parent.other_post.id}`} />
        )}

        {relations.map((relation) => {
          //remove any duplicated
          if (alreadyRender.includes(relation.other_post.id)) return <></>;
          alreadyRender.push(relation.other_post.id);

          return (
            <PostItem
              key={`relation-${relation.other_post.id}`}
              label={`Post: ${relation.other_post.id}. Similarity: ${
                relation.similarity / 100
              }%`}
              relation={relation}
            />
          );
        })}
      </HStack>
    </Box>
  );
};
