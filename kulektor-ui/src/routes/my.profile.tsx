import { createFileRoute, redirect } from "@tanstack/react-router";

export const Route = createFileRoute("/my/profile")({
  component: MyUser,
  beforeLoad: ({ context, location }) => {
    if (!context.auth.isAuthenticated) {
      throw redirect({
        to: "/",
        search: {
          redirect: location.href,
        },
      });
    }
  },
});

function MyUser() {
  return (
    <main className="flex-1 bg-gray-50 overflow-y-auto p-4">
      <h1 className="text-2xl font-bold mb-4">My Profile</h1>
      <p className="text-gray-700">
        This is section where you can update your user info
      </p>
    </main>
  );
}
