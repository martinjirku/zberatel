import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { Router } from "./router";
import "./index.css";
import AuthProvider from "./auth/AuthProvider";
import ApiProvider from "./api/ApiProvider";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <AuthProvider>
      <ApiProvider>
        <Router />
      </ApiProvider>
    </AuthProvider>
  </StrictMode>,
);
