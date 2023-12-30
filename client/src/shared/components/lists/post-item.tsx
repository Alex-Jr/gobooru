import { ChatIcon, UpDownIcon } from "@chakra-ui/icons";
import {
  Box,
  Flex,
  HStack,
  Icon,
  Image,
  Link,
  Text,
  Tooltip,
} from "@chakra-ui/react";
import { useState } from "react";
import { FaHeart } from "react-icons/fa";
import { Link as RouterLink } from "react-router-dom";

import { filePathToUrl } from "shared/functions/file_path_to_url";
import { ratingToChakraColor } from "shared/functions/rating-to-chakra-color";
import { RatingEnum } from "shared/types/enums/rating-enum";

interface IPostItem {
  post: {
    id: number;
    thumb_path: string;
    rating: RatingEnum;
    tag_ids: string[];
  };
}

export const PostItem = ({ post }: IPostItem) => {
  const [tagLabel, setTagLabel] = useState<string>();

  return (
    <Box mx={"auto"}>
      <Tooltip
        label={post.tag_ids.slice(0, 20).join(" ")}
        openDelay={300}
        closeOnScroll={true}
        hasArrow={true}
      >
        <Link as={RouterLink} to={`/posts/${post.id}`}>
          <Image
            src={filePathToUrl(post.thumb_path)}
            objectFit="contain"
            border={"1px"}
            borderColor={ratingToChakraColor(post.rating)}
          />
        </Link>
      </Tooltip>

      <Flex justifyContent={"space-evenly"} m={"auto"}>
        <HStack gap={0.5}>
          <Text textDecorationStyle={"unset"}>{0}</Text>
          <Icon as={FaHeart} color={"pink.500"} boxSize={4} />
        </HStack>

        <HStack gap={0.5}>
          <Text>{0}</Text>
          <Icon as={ChatIcon} color={"orange.500"} boxSize={4} />
        </HStack>

        <HStack gap={0.5}>
          <Text>{0}</Text>
          <Icon
            as={UpDownIcon}
            color={0 >= 0 ? "green.500" : "red.500"}
            boxSize={4}
          />
        </HStack>
      </Flex>
    </Box>
  );
};
