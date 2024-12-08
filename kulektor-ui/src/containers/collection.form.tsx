import { useMutation, useQuery } from "@apollo/client";
import {
  Button,
  Drawer,
  IconButton,
  Input,
  Spinner,
  Textarea,
  Typography,
} from "@material-tailwind/react";
import { Xmark } from "iconoir-react";
import { useForm } from "react-hook-form";
import {
  MY_NEW_COLLECTION,
  MY_COLLECTIONS,
  MY_COLLECTIONS_DETAIL,
  UPDATE_MY_COLLECTION,
} from "../queries";
import { CollectionField } from "../gql/graphql";

interface Props {
  id?: string;
  onSuccess(collectionId: string): void;
}
export default function CollectionForm({ id, onSuccess }: Props) {
  const {
    register,
    handleSubmit,
    setValue,
    // formState: { errors },
  } = useForm({});
  const { loading } = useQuery(MY_COLLECTIONS_DETAIL, {
    variables: { input: id },
    skip: !id,
    fetchPolicy: "network-only",
    onCompleted(data) {
      setValue("title", data.myCollectionDetail.title);
      setValue("description", data.myCollectionDetail.description);
      setValue("type", data.myCollectionDetail.type);
    },
  });
  const [addCollection, { loading: addingCollection }] = useMutation(
    MY_NEW_COLLECTION,
    {
      refetchQueries: [MY_COLLECTIONS],
    },
  );
  const [updateCollection, { loading: editingCollection }] = useMutation(
    UPDATE_MY_COLLECTION,
    {
      refetchQueries: [MY_COLLECTIONS],
    },
  );
  const onSubmit = (data: any) => {
    if (!id) {
      addCollection({
        variables: {
          input: {
            title: data.title,
            description: data.description,
            type: data.type,
          },
        },
      }).then((d) => onSuccess(d.data?.createMyCollection?.data?.id!));
    } else {
      updateCollection({
        variables: {
          input: {
            id,
            collection: {
              title: data.title,
              description: data.description,
              type: data.type,
            },
            fieldsToUpdate: [
              CollectionField.Title,
              CollectionField.Description,
              CollectionField.Type,
            ],
          },
        },
      }).then((d) => onSuccess(d.data?.updateMyCollection?.data?.id!));
    }
  };
  if (loading) {
    return (
      <div className="w-full h-full flex flex-col justify-center align-middle">
        <Spinner size="xl" />
      </div>
    );
  }
  const mutationLoading = editingCollection || addingCollection;
  const mutationBtnLabel = id ? "Udate" : "Create";
  return (
    <form
      onSubmit={handleSubmit(onSubmit)}
      className="flex flex-col gap-4 h-full"
    >
      <div className="flex items-center justify-between gap-4 flex-grow-0 flex-shrink-0">
        <Typography type="h6">New Collection</Typography>
        <Drawer.DismissTrigger
          as={IconButton}
          size="sm"
          type="button"
          variant="ghost"
          className="absolute right-2 top-2"
          isCircular
        >
          <Xmark className="h-5 w-5" />
        </Drawer.DismissTrigger>
      </div>
      <div className="flex-grow flex-shrink-0 flex flex-col gap-4">
        <div className="w-full space-y-1">
          <Typography
            as="label"
            htmlFor="title"
            type="small"
            color="default"
            className="font-semibold"
          >
            Title
          </Typography>
          <Input {...register("title")} />
        </div>
        <div className="w-full space-y-1">
          <Typography
            as="label"
            htmlFor="description"
            type="small"
            color="default"
            className="font-semibold"
          >
            Description
          </Typography>
          <Textarea {...register("description")} />
        </div>
        <div className="w-full space-y-1">
          <Typography
            as="label"
            htmlFor="type"
            type="small"
            color="default"
            className="font-semibold"
          >
            Type
          </Typography>
          <Input
            placeholder="Pop-Heads, Stamsp, Coins,..."
            {...register("type")}
          />
        </div>
      </div>
      <div className="flex flex-row gap-4 flex-grow-0 flex-shrink-0">
        <Drawer.DismissTrigger
          as={Button}
          size="sm"
          type="button"
          variant="outline"
          color="error"
          className="flex-grow"
        >
          <Xmark className="h-5 w-5" />
          Cancel
        </Drawer.DismissTrigger>
        {!mutationLoading ? (
          <Button type="submit" className="flex-grow" color="success">
            {mutationBtnLabel}
          </Button>
        ) : (
          <Button type="submit" className="flex-grow" color="success" disabled>
            <Spinner size="xs" className="mr-2" />
            {mutationBtnLabel}
          </Button>
        )}
      </div>
    </form>
  );
}
