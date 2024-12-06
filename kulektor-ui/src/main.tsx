import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./index.css";
import AuthProvider from "./auth/AuthProvider";
import ApiProvider from "./api/ApiProvider";
import { RouterProvider, createRouter } from "@tanstack/react-router";

import { routeTree } from "./routeTree.gen";

const router = createRouter({ routeTree });

declare module "@tanstack/react-router" {
  interface Register {
    router: typeof router;
  }
}

// Render the app
createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <AuthProvider>
      <ApiProvider>
        <RouterProvider router={router} />
      </ApiProvider>
    </AuthProvider>
  </StrictMode>,
);
