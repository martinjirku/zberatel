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
} from "@material-tailwind/react";
import { Link, Outlet } from "react-router";
import { EditPencil, Plus, Trash } from "iconoir-react";
import CollectionForm from "../containers/collection-form";

export default function MyDashboard() {
  const [offset, setOffset] = useState(0);
  const [newCollectionDrawer, setNewCollectionDrawer] = useState(false);
  const { data } = useQuery(MY_COLLECTIONS, {
    variables: { input: { paging: { limit: 4, offset } } },
  });
  const meta = data?.myCollectionsList.meta;
  const hasNext = !!meta?.nextPage;
  const hasPrev = !!meta?.prevPage;
  return (
    <main className="flex-1 bg-gray-50 overflow-y-auto p-4">
      <h1 className="text-2xl font-bold mb-4">My Dashboard</h1>
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
        <Card.Body className="flex flex-row gap-4">
          {data?.myCollectionsList.items?.map((c) => (
            <Card key={c.id} className="group">
              <Card.Header className="flex flex-row justify-between">
                <Typography type="h6">{c.title}</Typography>
                <span className="opacity-0 group-hover:opacity-100 transition-opacity duration-300">
                  <Tooltip>
                    <Tooltip.Trigger as={"div"}>
                      <IconButton
                        as={Link}
                        to={`my/collection/${c.id}/edit`}
                        variant="ghost"
                        size="xs"
                      >
                        <EditPencil />
                      </IconButton>
                    </Tooltip.Trigger>
                    <Tooltip.Content>
                      <Tooltip.Arrow />
                      Edit Collection
                    </Tooltip.Content>
                  </Tooltip>
                </span>
              </Card.Header>
              <Card.Body>
                <Typography type="p">{c.description}</Typography>
              </Card.Body>
              <Card.Footer className="flex flex-row gap-2">
                <Button
                  isFullWidth
                  size="sm"
                  color="primary"
                  as={Link}
                  to={`/my/collection/${c.id}`}
                >
                  Detail
                </Button>
                <DeleteCollectionButton id={c.id} />
              </Card.Footer>
            </Card>
          ))}
        </Card.Body>
        <Card.Footer className="flex justify-end gap-4">
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
      <Outlet />
    </main>
  );
}

function DeleteCollectionButton({ id }: { id: string }) {
  const [deleteCollection, { loading }] = useMutation(DELETE_MY_COLLECTION, {
    refetchQueries: [MY_COLLECTIONS],
  });

  return (
    <Button
      isFullWidth
      size="sm"
      color="warning"
      variant="outline"
      className="flex flex-row gap-2"
      onClick={() => deleteCollection({ variables: { input: id } })}
      disabled={loading}
    >
      {!loading ? <Trash className="text-xs" /> : <Spinner size="xs" />}
      Delete
    </Button>
  );
}
