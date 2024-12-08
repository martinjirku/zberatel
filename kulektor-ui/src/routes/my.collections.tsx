import { useState } from "react";
import { useMutation, useQuery } from "@apollo/client";
import { DELETE_MY_COLLECTION, MY_COLLECTIONS } from "../queries";
import {
  Button,
  Card,
  IconButton,
  Tooltip,
  Typography,
  Drawer,
  Spinner,
  Breadcrumb,
} from "@material-tailwind/react";
import {
  Cube,
  EditPencil,
  Home,
  NavArrowRight,
  Plus,
  Trash,
} from "iconoir-react";
import CollectionForm from "../containers/collection.form";
import { createFileRoute, Link } from "@tanstack/react-router";

export const Route = createFileRoute("/my/collections")({
  component: MyDashboard,
});

function MyDashboard() {
  const [offset, setOffset] = useState(0);
  const [newCollectionDrawer, setNewCollectionDrawer] = useState(false);
  const [activeCollectionIdDrawer, setActiveCollectionIdDrawer] =
    useState<string>();
  const { data } = useQuery(MY_COLLECTIONS, {
    variables: { input: { paging: { limit: 4, offset } } },
  });
  const meta = data?.myCollectionsList.meta;
  const hasNext = !!meta?.nextPage;
  const hasPrev = !!meta?.prevPage;
  return (
    <main className="flex flex-col flex-1 py-4 bg-gray-50 overflow-y-auto gap-4">
      <div className="px-4">
        <BlueprintsBreadcrumbs />
      </div>
      <div className="px-4">
        <Card variant="solid">
          <Card.Header className="flex flex-row justify-between">
            <Typography type="h6">Collections</Typography>
            <div>
              <Drawer
                open={newCollectionDrawer}
                onOpenChange={setNewCollectionDrawer}
              >
                <Drawer.Trigger onClick={() => setNewCollectionDrawer(true)}>
                  <Tooltip placement="bottom" interactive={false}>
                    <Tooltip.Trigger as={"div"} variant="outline" size="xs">
                      <IconButton variant="outline" size="xs" as="div">
                        <Plus fontSize="sm" />
                      </IconButton>
                    </Tooltip.Trigger>
                    <Tooltip.Content>
                      <Tooltip.Arrow />
                      Add new Collection
                    </Tooltip.Content>
                  </Tooltip>
                </Drawer.Trigger>
                <Drawer.Overlay lockScroll>
                  <Drawer.Panel className="w-full sm:w-3/5 md:w-2/5 xl:w-1/5">
                    <CollectionForm
                      onSuccess={() => setNewCollectionDrawer(false)}
                    />
                  </Drawer.Panel>
                </Drawer.Overlay>
              </Drawer>
            </div>
          </Card.Header>
          <Card.Body className="grid grid-flow-row grid-cols-1 sm:grid-cols-2 lg:grid-cols-4 flex-row gap-4">
            {data?.myCollectionsList.items?.map((c) => (
              <Card key={c.id} className="group flex flex-col">
                <Card.Header className="flex flex-row grow-0 justify-between">
                  <Typography type="h6">{c.title}</Typography>
                  <span className="opacity-0 group-hover:opacity-100 transition-opacity duration-300">
                    <Drawer
                      open={activeCollectionIdDrawer === c.id}
                      onOpenChange={(open) => {
                        if (!open) {
                          setActiveCollectionIdDrawer(undefined);
                        }
                      }}
                    >
                      <Tooltip>
                        <Tooltip.Trigger as={"div"}>
                          <IconButton
                            variant="ghost"
                            size="xs"
                            onClick={() => setActiveCollectionIdDrawer(c.id)}
                          >
                            <EditPencil />
                          </IconButton>
                        </Tooltip.Trigger>
                        <Tooltip.Content>
                          <Tooltip.Arrow />
                          Edit Collection
                        </Tooltip.Content>
                      </Tooltip>
                      <Drawer.Overlay lockScroll>
                        <Drawer.Panel className="w-full sm:w-3/5 md:w-2/5 xl:w-1/5">
                          <CollectionForm
                            id={c.id}
                            onSuccess={() =>
                              setActiveCollectionIdDrawer(undefined)
                            }
                          />
                        </Drawer.Panel>
                      </Drawer.Overlay>
                    </Drawer>
                  </span>
                </Card.Header>
                <Card.Body className="grow">
                  <Typography type="p">{c.description}</Typography>
                </Card.Body>
                <Card.Footer className="grid grid-cols-2 gap-2">
                  <Link to="/my/collections/$id" params={{ id: c.id }}>
                    <Button isFullWidth size="sm" color="primary" as="div">
                      Detail
                    </Button>
                  </Link>
                  <DeleteCollectionButton id={c.id} className="" />
                </Card.Footer>
              </Card>
            ))}
          </Card.Body>
          <Card.Footer className="flex justify-end gap-2">
            {hasPrev ? (
              <Button
                variant="outline"
                onClick={() => setOffset(meta?.prevPage?.offset ?? 0)}
              >
                Prev
              </Button>
            ) : (
              <Button variant="outline" disabled>
                Prev
              </Button>
            )}
            {hasNext ? (
              <Button
                aria-disabled={!hasNext}
                variant="outline"
                onClick={() => setOffset(meta?.nextPage?.offset ?? 0)}
                className={hasNext ? "" : "cursor-default dis"}
              >
                Next
              </Button>
            ) : (
              <Button variant="outline" disabled>
                Prev
              </Button>
            )}
          </Card.Footer>
        </Card>
      </div>
    </main>
  );
}

function DeleteCollectionButton({
  id,
  className,
}: {
  id: string;
  className: string;
}) {
  const [deleteCollection, { loading }] = useMutation(DELETE_MY_COLLECTION, {
    refetchQueries: [MY_COLLECTIONS],
  });

  return (
    <Button
      isFullWidth
      size="sm"
      color="warning"
      variant="outline"
      className={`flex flex-row gap-2 ` + className}
      onClick={() => deleteCollection({ variables: { input: id } })}
      disabled={loading}
    >
      {!loading ? (
        <Trash className="text-xs flex-shrink-0" />
      ) : (
        <Spinner size="xs" className="flex-shrink-0" />
      )}
      <span className="flex-shrink">Delete</span>
    </Button>
  );
}

const BlueprintsBreadcrumbs = () => {
  return (
    <Breadcrumb className="gap-0.5">
      <Breadcrumb.Link
        as={Link}
        to="/"
        className="rounded bg-secondary px-2 py-1 text-secondary-foreground hover:bg-primary hover:text-primary-foreground"
      >
        <Home className="h-4 w-4" />
        Home
      </Breadcrumb.Link>
      <Breadcrumb.Separator>
        <NavArrowRight className="h-4 w-4 stroke-2" />
      </Breadcrumb.Separator>
      <Breadcrumb.Link
        as={Link}
        to="/blueprints"
        className="rounded bg-secondary px-2 py-1 text-secondary-foreground hover:bg-primary hover:text-primary-foreground"
      >
        <Cube className="h-4 w-4" />
        Collections
      </Breadcrumb.Link>
    </Breadcrumb>
  );
};
