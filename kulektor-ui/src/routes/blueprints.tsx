import { Breadcrumb } from "@material-tailwind/react";
import { createFileRoute, Link } from "@tanstack/react-router";
import { Cube, Home, NavArrowRight } from "iconoir-react";
import { PropsWithChildren } from "react";

export const Route = createFileRoute("/blueprints")({
  component: Blueprints,
});

function Blueprints() {
  return (
    <div className="flex flex-col gap-4">
      <Row>
        <BlueprintsBreadcrumbs />
      </Row>
    </div>
  );
}

const Row = ({ children }: PropsWithChildren) => {
  return <div className="p-5">{children}</div>;
};

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
        Blueprints
      </Breadcrumb.Link>
    </Breadcrumb>
  );
};
