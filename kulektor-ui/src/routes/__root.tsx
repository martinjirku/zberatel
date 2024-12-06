import * as React from "react";
import { createRootRoute } from "@tanstack/react-router";
import DefaultLayout from "../layouts/default-layout";
import { TanStackRouterDevtools } from "@tanstack/router-devtools";

export const Route = createRootRoute({
  component: RootComponent,
});

function RootComponent() {
  return (
    <React.Fragment>
      <DefaultLayout />
      <TanStackRouterDevtools />
    </React.Fragment>
  );
}
