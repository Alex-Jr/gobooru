import { ChevronDownIcon, HamburgerIcon } from "@chakra-ui/icons";
import {
  Box,
  Button,
  Flex,
  HStack,
  IconButton,
  Link,
  Menu,
  MenuButton,
  MenuDivider,
  MenuItem,
  MenuList,
  Text,
  VStack,
  useColorModeValue,
} from "@chakra-ui/react";
import { useQueryClient } from "@tanstack/react-query";
import { useContext } from "react";
import { useTranslation } from "react-i18next";
import { FaHome } from "react-icons/fa";
import { Link as RouterLink } from "react-router-dom";

import { useLogout } from "services/auth/use-logout";
import { UserContext } from "shared/context/userContext";
import { APIUser } from "shared/types/services/user/APIUser";

const navItems = ["posts", "pools", "tags"];

const UserMenu = ({
  user,
  onLogoutClick,
}: {
  user: APIUser;
  onLogoutClick: () => void;
}) => {
  const { t } = useTranslation();

  return (
    <Menu>
      <MenuButton
        as={Button}
        rounded={"full"}
        cursor={"pointer"}
        bgColor={"green.400"}
        _hover={{
          bg: "green.200",
        }}
        minW={0}
      >
        <HStack>
          {/* <Avatar
            size={"sm"}
            src={
              user.image && `${BASE_URL}/${user.image.replace("./public", "")}`
            }
          /> */}
          <VStack
            display={{ base: "none", md: "flex" }}
            alignItems="flex-start"
            spacing="1px"
            ml="2"
          >
            <Text fontSize="sm">{user.name}</Text>
            <Text fontSize="xs" color="gray.600">
              ADMIN
            </Text>
          </VStack>
          <Box display={{ base: "none", md: "flex" }}>
            <ChevronDownIcon />
          </Box>
        </HStack>
      </MenuButton>
      <MenuList>
        <MenuItem>
          <Link as={RouterLink} to="/posts/new">
            {t("loggedMenu.newPost")}
          </Link>
        </MenuItem>
        <MenuItem>
          <Link as={RouterLink} to="/posts/new/batch">
            {t("loggedMenu.newPostBatch")}
          </Link>
        </MenuItem>
        {/* <MenuItem>
          <Link as={RouterLink} to="/pool/new">
            Nova pool
          </Link>
        </MenuItem> */}
        <MenuDivider />

        <MenuItem isDisabled>
          <Link pointerEvents={"none"} as={RouterLink}>
            {t("loggedMenu.myAccount")}
          </Link>
        </MenuItem>
        <MenuDivider />

        <MenuItem onClick={onLogoutClick}>{t("loggedMenu.logout")}</MenuItem>
      </MenuList>
    </Menu>
  );
};

export const NavBar = () => {
  const { user, setUser } = useContext(UserContext);
  const logoutMutation = useLogout();
  const queryClient = useQueryClient();

  return (
    <Flex
      as={"nav"}
      bgColor={useColorModeValue("lightgray", "gray.700")}
      justify={"space-between"}
      align={"center"}
      px={4}
      py={2}
    >
      {/* DesktopNav  */}
      <HStack gap={6} display={{ sm: "none", md: "flex" }}>
        <IconButton
          as={RouterLink}
          icon={<FaHome />}
          aria-label="home"
          variant={"ghost"}
          to={"/"}
        />

        {navItems.map((item) => (
          <Link key={`nav-desktop-${item}`} as={RouterLink} to={`/${item}`}>
            {item}
          </Link>
        ))}
      </HStack>

      {/* MobileNav  */}
      <Menu>
        <MenuButton
          as={IconButton}
          aria-label="Options"
          icon={<HamburgerIcon />}
          variant="outline"
          display={{ sm: "unset", md: "none" }}
        />
        <MenuList>
          {navItems.map((item) => (
            <MenuItem
              key={`nav-mobile-${item}`}
              as={RouterLink}
              to={`/${item}`}
            >
              {item}
            </MenuItem>
          ))}
        </MenuList>
      </Menu>

      {user ? (
        <UserMenu
          user={user}
          onLogoutClick={() => {
            //! clears cache because futures users can have different permissions
            queryClient.invalidateQueries();
            logoutMutation.mutate();
            setUser(undefined);
          }}
        />
      ) : (
        <HStack>
          {/* <ColorModeSwitcher /> */}
          {/* <Button as={RouterLink} to={"/signin"} variant={"ghost"}>
            Cadastro
          </Button> */}
          <Button as={RouterLink} to={"/login"} colorScheme={"orange"}>
            Entrar
          </Button>
        </HStack>
      )}
    </Flex>
  );
};
