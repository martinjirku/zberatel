import { FC, PropsWithChildren, useEffect, useState } from "react";
import {
  Button,
  Navbar,
  Collapse,
  Typography,
  IconButton,
  Menu,
  Avatar,
} from "@material-tailwind/react";
import { Link, Outlet } from "react-router";
import {
  HeadsetHelp,
  LogOut,
  Menu as MenuIcon,
  MultiplePages,
  SelectFace3d,
  Settings,
  UserCircle,
  Xmark,
} from "iconoir-react";
import { useAuth } from "../auth/AuthState";

const DefaultLayout: FC<PropsWithChildren> = ({ children }) => {
  const [openNav, setOpenNav] = useState(false);
  const { loginWithPopup, isAuthenticated, logout, user } = useAuth();
  useEffect(() => {
    window.addEventListener(
      "resize",
      () => window.innerWidth >= 640 && setOpenNav(false),
    );
  }, []);

  return (
    <div className="flex h-screen flex-col">
      <Navbar className="sticky top-0 mx-auto w-full bg-black dark:bg-surface-dark p-4 border-none rounded-none">
        <div className="flex items-center text-white">
          <Typography
            as={Link}
            to="/"
            type="small"
            className="ml-2 mr-2 block py-1 font-semibold"
          >
            Kulektor
          </Typography>
          <hr className="ml-1 mr-1.5 hidden h-5 w-px border-l border-t-0 border-surface/25 sm:block dark:border-surface" />
          <div className="hidden sm:block">
            <NavList />
          </div>
          <div className="flex-grow"></div>
          {isAuthenticated && (
            <Typography type="small">
              <span className="p-2">Hello, {user?.name}</span>
            </Typography>
          )}
          {isAuthenticated ? (
            <ProfileMenu />
          ) : (
            <Button
              size="sm"
              color="primary"
              onClick={() => loginWithPopup()}
              className="hidden sm:ml-auto sm:inline-block p-2"
            >
              Sign In
            </Button>
          )}
          <IconButton
            size="sm"
            color="secondary"
            onClick={() => setOpenNav(!openNav)}
            className="ml-auto grid sm:hidden"
          >
            {openNav ? (
              <Xmark className="h-4 w-4" />
            ) : (
              <MenuIcon className="h-4 w-4" />
            )}
          </IconButton>
        </div>
        <Collapse open={openNav}>
          <NavList />
          {isAuthenticated ? (
            <Button
              size="sm"
              isFullWidth
              onClick={() => logout()}
              className="mt-4 border-white bg-white text-black hover:border-white hover:bg-white hover:text-black"
            >
              Logout
            </Button>
          ) : (
            <Button
              size="sm"
              isFullWidth
              onClick={() => loginWithPopup()}
              className="mt-4 border-white bg-white text-black hover:border-white hover:bg-white hover:text-black"
            >
              Sign In
            </Button>
          )}
        </Collapse>
      </Navbar>

      <div className="flex flex-1 overflow-hidden">
        {isAuthenticated ? (
          <aside className="bg-gray-100 w-28 sm:w-32 md:w-36 lg:w-56 p-4 overflow-y-auto">
            <nav className="space-y-4">
              <Link to="/my/dashboard" className="block">
                Dashboard
              </Link>
              <Link
                to="/my/profile"
                className="block text-blue-500 font-semibold"
              >
                My Profile
              </Link>
            </nav>
          </aside>
        ) : null}
        {children}
        <Outlet />
      </div>

      {/* Footer */}
      <footer className="bg-gray-200 text-gray-700 text-center py-3">
        © 2024 Kulektor. All rights reserved.
      </footer>
    </div>
  );
};

export default DefaultLayout;

const LINKS = [
  { icon: MultiplePages, title: "Collections", href: "#" },
  { icon: SelectFace3d, title: "Blog", href: "#" },
];

const NavList = () => {
  return (
    <ul className="mt-4 flex flex-col gap-x-3 gap-y-1.5 sm:mt-0 sm:flex-row sm:items-center">
      {LINKS.map(({ icon: Icon, title, href }) => (
        <li key={title}>
          <Typography
            as="a"
            href={href}
            type="small"
            className="flex items-center gap-x-2 p-1 text-white hover:text-white"
          >
            <Icon className="h-4 w-4" />
            {title}
          </Typography>
        </li>
      ))}
    </ul>
  );
};

const ProfileMenu = () => {
  const { user, logout } = useAuth();
  return (
    <Menu>
      <Menu.Trigger
        as={Avatar}
        src={user?.picture}
        alt={user?.name}
        size="sm"
        className="border border-primary p-0.5 lg:ml-auto"
      />
      <Menu.Content>
        <Menu.Item>
          <UserCircle className="mr-2 h-[18px] w-[18px]" /> My Profile
        </Menu.Item>
        <Menu.Item>
          <Settings className="mr-2 h-[18px] w-[18px]" /> Edit Profile
        </Menu.Item>
        <Menu.Item>
          <HeadsetHelp className="mr-2 h-[18px] w-[18px]" /> Support
        </Menu.Item>
        <hr className="!my-1 -mx-1 border-surface" />
        <Menu.Item
          className="text-error hover:bg-error/10 hover:text-error focus:bg-error/10 focus:text-error"
          onClick={() => logout()}
        >
          <LogOut className="mr-2 h-[18px] w-[18px]" />
          Logout
        </Menu.Item>
      </Menu.Content>
    </Menu>
  );
};
