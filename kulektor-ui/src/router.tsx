import { lazy, Suspense } from "react";
import { createBrowserRouter, RouterProvider } from "react-router";

const Home = lazy(() => import("./Home.tsx"));
const AuthLayout = lazy(() => import("./layouts/auth-layout.tsx"));

const router = createBrowserRouter([
  {
    path: "/",
    element: (
      <Suspense fallback="Loading Home...">
        <Home />
      </Suspense>
    ),
    index: true, // Marks this route as the index route
  },
  {
    path: "/",
    element: (
      <Suspense fallback={<div>Loading...</div>}>
        <AuthLayout />
      </Suspense>
    ),
    shouldRevalidate: () => false,
    children: [
      {
        path: "login",
        id: "login",
        element: <div>Login</div>,
      },
      {
        path: "register",
        id: "register",
        element: <div>Register</div>,
      },
    ],
  },
]);

export const Router = () => <RouterProvider router={router} />;
