import { useQuery } from "@apollo/client";
import {
  Breadcrumb,
  IconButton,
  Spinner,
  Drawer,
  Tooltip,
} from "@material-tailwind/react";
import { createFileRoute, Link } from "@tanstack/react-router";
import { Cube, Home, NavArrowRight, Plus } from "iconoir-react";
import { PropsWithChildren, useState } from "react";
import { LIST_BLUEPRINT } from "../queries";
import {
  useReactTable,
  getCoreRowModel,
  getSortedRowModel,
  createColumnHelper,
  flexRender,
} from "@tanstack/react-table";
import { BlueprintsListQuery } from "../gql/graphql";
import Table from "../components/table";
import { twMerge } from "tailwind-merge";
import NewBlueprintForm from "../containers/new-blueprint.form";

export const Route = createFileRoute("/blueprints")({
  component: Blueprints,
});

function Blueprints() {
  const [activeDrawer, setActiveDrawer] = useState<"new">();
  const { data, loading } = useQuery(LIST_BLUEPRINT, {
    variables: {
      input: {
        paging: {
          limit: 2,
          offset: 0,
        },
      },
    },
  });
  const table = useReactTable({
    columns,
    data: data?.blueprintsList?.items ?? emptyData,
    getCoreRowModel: getCoreRowModel(),
    getSortedRowModel: getSortedRowModel(),
  });

  if (loading) {
    return (
      <div className="flex flex-col gap-4">
        <Spinner />
      </div>
    );
  }
  return (
    <div className="flex flex-col gap-4 w-full">
      <Row>
        <BlueprintsBreadcrumbs />
      </Row>
      <Row className="flex flex-col gap-1">
        <div className="flex flex-row w-full justify-end">
          <Drawer
            open={activeDrawer === "new"}
            onOpenChange={(open) => {
              if (!open) {
                setActiveDrawer(undefined);
              }
            }}
          >
            <Tooltip>
              <Tooltip.Trigger as={"div"}>
                <IconButton
                  variant="outline"
                  size="xs"
                  onClick={() => setActiveDrawer("new")}
                >
                  <Plus fontSize="1rem" />
                </IconButton>
              </Tooltip.Trigger>
              <Tooltip.Content>
                <Tooltip.Arrow />
                Create Blueprint
              </Tooltip.Content>
            </Tooltip>
            <Drawer.Overlay lockScroll>
              <Drawer.Panel className="w-full sm:w-3/5 md:w-2/5 xl:w-1/5">
                <NewBlueprintForm />
              </Drawer.Panel>
            </Drawer.Overlay>
          </Drawer>
        </div>
        <Table className="w-full">
          <Table.Head>
            {table.getHeaderGroups().map((headerGroup) => {
              return (
                <Table.tr key={headerGroup.id}>
                  {headerGroup.headers.map(
                    (
                      header, // map over the headerGroup headers array
                    ) => (
                      <Table.th
                        key={header.id}
                        colSpan={header.colSpan}
                        className="p-2"
                      >
                        {header.isPlaceholder
                          ? null
                          : flexRender(
                              header.column.columnDef.header,
                              header.getContext(),
                            )}
                      </Table.th>
                    ),
                  )}
                </Table.tr>
              );
            })}
          </Table.Head>
          <Table.Body>
            {table.getRowModel().rows.map((row) => (
              <Table.tr key={row.id}>
                {row.getVisibleCells().map((cell) => (
                  <Table.td key={cell.id} className="p-2">
                    {flexRender(cell.column.columnDef.cell, cell.getContext())}
                  </Table.td>
                ))}
              </Table.tr>
            ))}
          </Table.Body>
        </Table>
      </Row>
    </div>
  );
}

type BlueprintQueryResult =
  BlueprintsListQuery["blueprintsList"]["items"][number];

const columnHelper = createColumnHelper<BlueprintQueryResult>();
const columns = [
  columnHelper.accessor("title", {
    header: "Title",
  }),
  columnHelper.accessor("description", {
    header: "Description",
  }),
  columnHelper.accessor("updatedAt", {
    header: "Updated",
    cell: (v) => new Date(v.getValue()).toLocaleString(),
  }),
];
const emptyData: BlueprintQueryResult[] = [];

const Row = ({
  children,
  className,
}: PropsWithChildren<{ className?: string }>) => {
  return <div className={twMerge("p-5", className)}>{children}</div>;
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
