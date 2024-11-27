import { Suspense } from "react";
import { Link, Outlet, useMatches } from "react-router";
import { Card, Tabs } from "@material-tailwind/react";

export default function AuthLayout() {
  const matches = useMatches();
  const value = matches[1].id ?? "login";
  return (
    <div className="flex items-center justify-center min-h-screen bg-gray-900 text-gray-200">
      <Card className="w-full max-w-md shadow-md ">
        <Tabs orientation="horizontal" defaultValue={value}>
          <Tabs.List value={value} className="w-full">
            <Tabs.Trigger value="login" className="w-full" as={Link} to="login">
              Login
            </Tabs.Trigger>
            <Tabs.Trigger
              value="register"
              className="w-full"
              as={Link}
              to="register"
            >
              Register
            </Tabs.Trigger>
          </Tabs.List>
          <Tabs.Panel value={value}>
            <Suspense fallback={<div>Loading...</div>}>
              <Outlet />
            </Suspense>
          </Tabs.Panel>
        </Tabs>
      </Card>
    </div>
  );
}
