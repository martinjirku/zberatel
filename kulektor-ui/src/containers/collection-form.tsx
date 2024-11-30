import { useMutation } from "@apollo/client";
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
import { MY_NEW_COLLECTION, MY_COLLECTIONS } from "../queries";

interface Props {
  onSuccess(collectionId: string): void;
}
export default function CollectionForm({ onSuccess }: Props) {
  const {
    register,
    handleSubmit,
    // formState: { errors },
  } = useForm();
  const [addCollection, { loading }] = useMutation(MY_NEW_COLLECTION, {
    refetchQueries: [MY_COLLECTIONS],
  });
  const onSubmit = (data: any) => {
    addCollection({
      variables: {
        input: {
          title: data.title,
          description: data.description,
          type: data.type,
        },
      },
    }).then((d) => onSuccess(d.data?.createMyCollection?.data?.id!));
  };
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
          variant="outline"
          color="error"
          isCircular
          className="flex-grow"
        >
          <Xmark className="h-5 w-5" />
          Cancel
        </Drawer.DismissTrigger>
        {!loading ? (
          <Button className="flex-grow" color="success">
            Create
          </Button>
        ) : (
          <Button className="flex-grow" color="success" disabled>
            <Spinner size="xs" className="mr-2" />
            Create
          </Button>
        )}
      </div>
    </form>
  );
}
