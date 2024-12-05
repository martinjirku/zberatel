import { lazy, Suspense } from "react";
import { createBrowserRouter, RouterProvider } from "react-router";

const Home = lazy(() => import("./Home.tsx"));
// const AuthLayout = lazy(() => import("./layouts/auth-layout.tsx"));
const DefaultLayout = lazy(() => import("./layouts/default-layout.tsx"));
const MyUser = lazy(() => import("./pages/my-profile.tsx"));
const MyDashboard = lazy(() => import("./pages/my-collections.tsx"));
const MyCollectionDetail = lazy(
  () => import("./pages/my-collection-detail.tsx"),
);

const router = createBrowserRouter([
  {
    path: "/",
    element: (
      <Suspense fallback="Loading Home...">
        <DefaultLayout>
          <Home />
        </DefaultLayout>
      </Suspense>
    ),
    index: true, // Marks this route as the index route
  },
  {
    path: "/my",
    element: (
      <Suspense fallback={<div>Loading...</div>}>
        <DefaultLayout />
      </Suspense>
    ),
    children: [
      {
        path: "profile",
        id: "my-profile",
        element: (
          <Suspense fallback={<div>loading User...</div>}>
            <MyUser />
          </Suspense>
        ),
      },
      {
        path: "collections",
        id: "my-dashboard",
        element: (
          <Suspense fallback="loading User">
            <MyDashboard />
          </Suspense>
        ),
      },
      {
        path: "collections/:id",
        id: "my-collections/detail",
        element: (
          <Suspense fallback="loading User">
            <MyCollectionDetail />
          </Suspense>
        ),
      },
    ],
  },
]);

export const Router = () => <RouterProvider router={router} />;
