import * as React from "react";
import { createRootRouteWithContext } from "@tanstack/react-router";
import DefaultLayout from "../layouts/default-layout";
import { TanStackRouterDevtools } from "@tanstack/router-devtools";
import { AuthState } from "../auth/AuthState";

interface RootContext {
  auth: AuthState;
}
export const Route = createRootRouteWithContext<RootContext>()({
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
