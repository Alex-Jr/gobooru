import { ChatIcon, UpDownIcon } from "@chakra-ui/icons";
import {
  Flex,
  HStack,
  Icon,
  Image,
  Link,
  Text,
  Tooltip,
  VStack,
} from "@chakra-ui/react";
import { FaHeart } from "react-icons/fa";
import { Link as RouterLink } from "react-router-dom";

import { filePathToUrl } from "shared/functions/file_path_to_url";
import { APIPoolList } from "shared/types/services/pools/APIPoolList";

interface IPostListItem {
  pool: APIPoolList["pools"][0];
}

export const PoolItem = ({ pool }: IPostListItem) => {
  return (
    <VStack w={"100%"} maxW={250}>
      <Text fontSize={"sm"} h={5} textOverflow={"hidden"}>
        {pool.name}
      </Text>

      <Tooltip
        label={pool.description}
        openDelay={300}
        closeOnScroll={true}
        hasArrow={true}
      >
        <Link as={RouterLink} to={`/pools/${pool.id}`}>
          <Image
            src={filePathToUrl(pool.posts[0].thumb_path)}
            h={350}
            objectFit="cover"
            border={"1px"}
            // borderColor={ratingToChakraColor(pool.rating)}
          />
        </Link>
      </Tooltip>

      <Flex justifyContent={"space-evenly"} w={"100%"}>
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
    </VStack>
  );
};
