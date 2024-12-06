import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./index.css";
import AuthProvider from "./auth/AuthProvider";
import ApiProvider from "./api/ApiProvider";
import { RouterProvider, createRouter } from "@tanstack/react-router";

import { routeTree } from "./routeTree.gen";
import { useAuth } from "./auth/AuthState";

const router = createRouter({
  routeTree,
  context: { auth: undefined! },
});

declare module "@tanstack/react-router" {
  interface Register {
    router: typeof router;
  }
}

const InnerApp = () => {
  const auth = useAuth();
  return <RouterProvider router={router} context={{ auth }} />;
};

// Render the app
createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <AuthProvider>
      <ApiProvider>
        <InnerApp />
      </ApiProvider>
    </AuthProvider>
  </StrictMode>,
);
