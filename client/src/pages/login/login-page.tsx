import {
  Button,
  Flex,
  FormControl,
  FormErrorMessage,
  FormLabel,
  Input,
} from "@chakra-ui/react";
import { useContext, useEffect, useState } from "react";
import { SubmitHandler, useForm } from "react-hook-form";
import { useTranslation } from "react-i18next";
import { useNavigate } from "react-router-dom";

import { useLoginMutation } from "services/auth/use-login";
import { PasswordInput } from "shared/components/forms/password-input";
import { UserContext } from "shared/context/userContext";

interface IFormInput {
  name: string;
  password: string;
}

export const LoginPage = () => {
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [isInvalid, setIsInvalid] = useState(false);
  const { setUser } = useContext(UserContext);

  const navigate = useNavigate();
  const { t } = useTranslation();

  const { handleSubmit, register, watch } = useForm<IFormInput>({
    defaultValues: {
      name: "",
      password: "",
    },
  });

  const loginMutation = useLoginMutation();

  useEffect(() => {
    const subscription = watch(() => setIsInvalid(false));
    return () => subscription.unsubscribe();
  }, [watch]);

  const onSubmit: SubmitHandler<IFormInput> = (formData) => {
    setIsSubmitting(true);

    loginMutation.mutate(formData, {
      onSuccess: (user) => {
        setUser(user);
        navigate("/");
      },
      onError: () => {
        setUser(undefined);
        setIsSubmitting(false);
        setIsInvalid(true);
      },
    });
  };

  return (
    <Flex
      as={"form"}
      direction={"column"}
      gap={4}
      m={"auto"}
      onSubmit={handleSubmit(onSubmit)}
      w={"300px"}
    >
      <FormControl isInvalid={isInvalid} isDisabled={isSubmitting}>
        <FormLabel htmlFor={"name"}>{t("glossary.name")!}</FormLabel>

        <Input
          {...register("name")}
          placeholder={t("forms.namePlaceholder")!}
          autoComplete="off"
        />
      </FormControl>

      <FormControl isInvalid={isInvalid} isDisabled={isSubmitting}>
        <FormLabel htmlFor={"password"}>{t("glossary.password")!}</FormLabel>

        <PasswordInput
          {...register("password")}
          placeholder={t("forms.passwordPlaceholder")!}
        />

        <FormErrorMessage>{t("loginPage.failedLoginMsg")!}</FormErrorMessage>
      </FormControl>

      <Button type="submit">Confirm</Button>
    </Flex>
  );
};
