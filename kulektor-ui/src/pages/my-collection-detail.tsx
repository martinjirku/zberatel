import { useQuery } from "@apollo/client";
import {
  Breadcrumb,
  Card,
  Spinner,
  Typography,
} from "@material-tailwind/react";
import { Cube, CursorPointer, Home, NavArrowRight } from "iconoir-react";
import { Link, useParams } from "react-router";
import { MY_COLLECTIONS_DETAIL } from "../queries";

export default function MyCollectionDetail() {
  let { id } = useParams();
  const { data, loading } = useQuery(MY_COLLECTIONS_DETAIL, {
    variables: { input: id },
  });
  if (loading) {
    return (
      <main className="flex flex-col gap-4 flex-1 bg-gray-50 overflow-y-auto p-4">
        <Spinner className="h-32 w-32" />
      </main>
    );
  }
  return (
    <main className="flex flex-col gap-4 flex-1 bg-gray-50 overflow-y-auto p-4">
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
          to="/my/dashboard"
          className="rounded bg-secondary px-2 py-1 text-secondary-foreground hover:bg-primary hover:text-primary-foreground"
        >
          <Cube className="h-4 w-4" />
          Dashboard
        </Breadcrumb.Link>
        <Breadcrumb.Separator>
          <NavArrowRight className="h-4 w-4 stroke-2" />
        </Breadcrumb.Separator>
        <Breadcrumb.Link className="rounded bg-primary px-2 py-1 text-primary-foreground hover:bg-primary hover:text-primary-foreground">
          <CursorPointer className="h-4 w-4 rotate-90" />
          Breadcrumb
        </Breadcrumb.Link>
      </Breadcrumb>
      <Typography type="h3">{data?.myCollectionDetail.title}</Typography>
      <Typography type="p">{data?.myCollectionDetail.description}</Typography>
      <Card variant="solid" title="Collections">
        <Card.Body className="flex flex-row gap-4"></Card.Body>
      </Card>
    </main>
  );
}
