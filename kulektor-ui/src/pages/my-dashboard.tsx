import { useState } from "react";
import { useQuery } from "@apollo/client";
import { MY_COLLECTIONS } from "../queries";
import { Button, Card, Typography } from "@material-tailwind/react";
import { Link } from "react-router";

export default function MyDashboard() {
  const [offset, setOffset] = useState(0);
  const { data } = useQuery(MY_COLLECTIONS, {
    variables: { input: { paging: { limit: 4, offset } } },
  });
  const meta = data?.myCollectionsList.meta;
  const hasNext = !!meta?.nextPage;
  const hasPrev = !!meta?.prevPage;
  return (
    <main className="flex-1 bg-gray-50 overflow-y-auto p-4">
      <h1 className="text-2xl font-bold mb-4">My Dashboard</h1>
      <Card variant="solid" title="Collections">
        <Card.Header>
          <Typography type="h6">Collections</Typography>
        </Card.Header>
        <Card.Body className="flex flex-row gap-4">
          {data?.myCollectionsList.items?.map((c) => (
            <Card key={c.id}>
              <Card.Header>
                <Typography type="h5">{c.title}</Typography>
              </Card.Header>
              <Card.Body>
                <Typography type="p">{c.description}</Typography>
              </Card.Body>
              <Card.Footer className="flex flex-row gap-2">
                <Button
                  isFullWidth
                  size="sm"
                  as={Link}
                  to={`/my/collection/${c.id}`}
                >
                  Detail
                </Button>
                <Button isFullWidth size="sm">
                  Delete
                </Button>
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
    </main>
  );
}
