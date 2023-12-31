import { AddIcon, ExternalLinkIcon, MinusIcon } from "@chakra-ui/icons";
import {
  Button,
  Divider,
  Flex,
  HStack,
  IconButton,
  Link,
  ListItem,
  SimpleGrid,
  Text,
  UnorderedList,
  VStack,
  useColorModeValue,
} from "@chakra-ui/react";
import { format } from "date-fns";
import { Link as RouterLink } from "react-router-dom";

import { useTagCategoryList } from "services/tag-categories/use-tag-category";
import { ImageModeSelect } from "shared/components/buttons/image-mode-select";
import { formatFileSize } from "shared/functions/format_file_size";
import { removeHttpAndWWW } from "shared/functions/remove_http_and_www";
import { tagGroupOrder } from "shared/functions/tag-group-order";
import { ImageModeEnum } from "shared/types/enums/image-mode-enum";
import { APIPost } from "shared/types/services/posts/APIPost";

export function Sidebar({
  post,
  imageMode,
  setImageMode,
  onEditClick,
  onDeleteClick,
}: {
  post: APIPost["post"];
  imageMode: ImageModeEnum;
  setImageMode: (x: ImageModeEnum) => void;
  onEditClick: () => void;
  onDeleteClick: () => void;
}) {
  const { tagCategoriesObj } = useTagCategoryList();

  const tagsCategory = post.tags.reduce((acc, cur) => {
    if (acc[cur.category_id]) acc[cur.category_id].push(cur);
    else acc[cur.category_id] = [cur];

    return acc;
  }, {} as Record<string, APIPost["post"]["tags"]>);

  return (
    <Flex
      direction={"column"}
      gap={2}
      h={"fit-content"}
      p={4}
      w={{ base: "100%", lg: 275 }}
      minW={275}
      bgColor={useColorModeValue("lightgray", "gray.700")}
      mx={"auto"}
    >
      <Text>Tags</Text>

      {Object.entries(tagsCategory)
        .sort(
          ([aGroup], [bGroup]) => tagGroupOrder[aGroup] - tagGroupOrder[bGroup]
        )
        .map(([group, tags]) => (
          <VStack key={group} gap={1} align={"flex-start"}>
            <Text color={"white"} fontWeight={"bold"} userSelect={"none"}>
              {group}
            </Text>

            {tags.map(({ id: tagId }) => (
              <HStack key={tagId}>
                <IconButton
                  size={"xxs"}
                  aria-label={`Add ${tagId} to search`}
                  icon={<AddIcon />}
                  as={RouterLink}
                  to={`/posts?search=${tagId}`}
                />
                <IconButton
                  size={"xxs"}
                  aria-label={`Remove ${tagId} from search`}
                  icon={<MinusIcon />}
                  as={RouterLink}
                  to={`/posts?search=-${tagId}`}
                />

                <Link
                  fontSize={"xs"}
                  as={RouterLink}
                  to={`/posts?search=${tagId}`}
                  color={tagCategoriesObj[group] || "unset"}
                  wordBreak={"break-all"}
                >
                  {tagId}
                </Link>
              </HStack>
            ))}
          </VStack>
        ))}

      <Divider />
      <Text>Image size</Text>
      <ImageModeSelect value={imageMode} onChange={setImageMode} />

      <Divider />
      <Text>Sources</Text>
      {post.sources.length === 0 ? (
        <Text>None available</Text>
      ) : (
        <UnorderedList>
          {post.sources.map((source, index) => {
            const parsedSource = removeHttpAndWWW(source);

            return (
              <ListItem key={"source" + index}>
                <Link
                  href={source}
                  isExternal
                  display={"flex"}
                  gap={2}
                  alignItems={"center"}
                  wordBreak={"break-all"}
                >
                  {parsedSource}
                  <ExternalLinkIcon />
                </Link>
              </ListItem>
            );
          })}
        </UnorderedList>
      )}

      <Divider />
      <Text>Information</Text>
      <VStack align={"start"}>
        {/* <Link as={RouterLink} to={`/user/${post.userId}`}>
          Author: {post.userId}
        </Link> */}

        <Text>Rating: {post.rating}</Text>

        <Text>File Size: {formatFileSize(post.file_size)}</Text>

        <Text>
          Published: {format(new Date(post.created_at), "yyyy-MM-dd HH:mm:ss")}
        </Text>

        <Text wordBreak={"break-all"}> MD5: {post.md5}</Text>
      </VStack>

      <Divider />
      <Text>Actions</Text>
      <SimpleGrid minChildWidth={"100px"} gap={1}>
        <Button size={"sm"} colorScheme="green">
          Favorite
        </Button>
        <Button size={"sm"} colorScheme="yellow" onClick={onEditClick}>
          Edit
        </Button>
        <Button size={"sm"} colorScheme="red" onClick={onDeleteClick}>
          Delete
        </Button>
      </SimpleGrid>
    </Flex>
  );
}
